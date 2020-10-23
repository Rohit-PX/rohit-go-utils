package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Pallinder/go-randomdata"
	"github.com/portworx/spawn/buildinfo"
	"github.com/portworx/spawn/instance"
	"github.com/portworx/spawn/log"
	"github.com/portworx/spawn/pkg/stringutils"
	"github.com/portworx/spawn/provisioners"
	"github.com/portworx/spawn/pxinfra"
	"github.com/spf13/cobra"
)

var (
	stackName          string
	vmCount            int
	vmClusters         int
	esxiHost           string
	esxiLogin          string
	esxiPassword       string
	vmCPU              int
	vmMemory           int
	vmOS               string
	defaultProvisioner string

	// Additional node pools
	disableAdditionalPools bool

	// HBA
	vmHBACount int

	// PX disks
	vmDiskCount  int
	vmDiskSizePX string

	// Special disks
	vmDiskDockerEnabled   bool
	vmDiskJournalEnabled  bool
	vmDiskMetadataEnabled bool
	vmDiskCacheEnabled    bool
	vmDiskCacheSize       int

	// IKS
	iksK8sVersion string

	// kops
	kopsKeyId              string
	kopsKeySecret          string
	kopsRegion             string
	kopsHostedZoneId       string
	kopsK8sVersion         string
	kopsAvailabilityZones  string
	kopsVpc                string
	kopsSubnets            string
	kopsUtilitySubnets     string
	kopsSecurityGroups     string
	kopsUserGroup          string
	kopsImage              string
	kopsUser               string
	kopsKeyPropagationTime int
	kopsMasterVolSize      int
	kopsNodeVolSize        int

	// GKE
	gkeName            string
	gkeZone            string
	gkeAdditionalZones string
	gkeNodeCount       int
	gkeMultizone       bool

	// AKS
	aksK8sVersion string
	aksLocation   string
	aksVMSize     string
	aksVMSetType  string
)

const defaultStackName = "dev"

var vmsCmd = &cobra.Command{
	Use:   "vms",
	Short: "create vms",
	Long: `Create one or more VMs. Examples:

# Can be called with no arguments in a BuildInfo job
spawn vms

# Common args
spawn vms --name example --count 1 --os centos/7

# Specify an ESXi host
# To see all options run: spawn esxi list
spawn vms --esxi-host 70.0.0.124

# To create AKS cluster
spawn vms --os aks --name aks-test --count 3 --aks-k8s-version 1.12.8 --aks-location eastus
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.EnableLogging(verboseLogging)

		existingInstances, _ := instance.Load()
		newStackName := getStackName()
		fixMemory()

		spec := make(provisioners.VMSpecs, 0)
		for i := 0; i < vmCount; i++ {

			// Boot disk
			disks := make([]provisioners.DiskSpec, 0)
			disks = append(disks, provisioners.DiskSpec{
				Type:       "os",
				SizeGB:     128,
				MountPoint: "/",
			})

			// Docker disk
			if vmDiskDockerEnabled {
				disks = append(disks, provisioners.DiskSpec{
					Type:   "docker",
					SizeGB: 32,
				})
			}

			// Journal disk
			if vmDiskJournalEnabled {
				disks = append(disks, provisioners.DiskSpec{
					Type:   "journal",
					SizeGB: 25,
				})
			}

			// Metadata disk
			if vmDiskMetadataEnabled {
				disks = append(disks, provisioners.DiskSpec{
					Type:   "metadata",
					SizeGB: 64,
				})
			}

			// Cache disk
			if vmDiskCacheEnabled {
				disks = append(disks, provisioners.DiskSpec{
					Type:   "cache",
					SizeGB: vmDiskCacheSize,
				})
			}

			// PX disks
			disks = addPxDisks(disks)

			// Add HBA, convert count to int list
			if vmHBACount > 1 {
				if vmHBACount > 4 { // 4 HBA maximum per VM is allowed
					log.Fatal("Failed to provision. Error: You have selected to have %s HBAs, maximum number of HBAs allowed per VM is 4. Exiting.", vmHBACount)
				}
				log.Warn("You have selected a custom HBA count number, be aware drives will be distributed between %d HBA(s) 1:1 ratio, all extra drives will be provisioned on last HBA", vmHBACount)
			}
			var hbaList []int
			for hba := 0; hba <= vmHBACount-1; hba++ {
				hbaList = append(hbaList, hba)
			}

			spec = append(spec, provisioners.VMSpec{
				Name: fmt.Sprintf("%s-%d",
					newStackName,
					i),
				OS:       vmOS,
				CPU:      vmCPU,
				MemGB:    vmMemory,
				Disks:    disks,
				HBACount: vmHBACount,
				HBAList:  hbaList,
			})
		}

		provOptions := map[string]string{
			"esxi-host":                     esxiHost,
			"iks-k8s-version":               iksK8sVersion,
			"disk-docker":                   fmt.Sprintf("%t", vmDiskDockerEnabled),
			"disk-journal":                  fmt.Sprintf("%t", vmDiskJournalEnabled),
			"disk-metadata":                 fmt.Sprintf("%t", vmDiskMetadataEnabled),
			"disk-cache":                    fmt.Sprintf("%t", vmDiskCacheEnabled),
			"kops-key-id":                   kopsKeyId,
			"kops-key-secret":               kopsKeySecret,
			"kops-region":                   kopsRegion,
			"kops-hosted-zone-id":           kopsHostedZoneId,
			"kops-cluster-name":             newStackName,
			"kops-az":                       kopsAvailabilityZones,
			"kops-k8s-version":              kopsK8sVersion,
			"kops-vpc":                      kopsVpc,
			"kops-subnets":                  kopsSubnets,
			"kops-utility-subnets":          kopsUtilitySubnets,
			"kops-security-groups":          kopsSecurityGroups,
			"kops-user-group":               kopsUserGroup,
			"kops-image":                    kopsImage,
			"kops-user":                     kopsUser,
			"kops-key-propagation":          fmt.Sprintf("%d", kopsKeyPropagationTime),
			"kops-master-vol-size":          fmt.Sprintf("%d", kopsMasterVolSize),
			"kops-node-vol-size":            fmt.Sprintf("%d", kopsNodeVolSize),
			"gke-name":                      gkeName,
			"gke-zone":                      gkeZone,
			"gke-additional-zones":          gkeAdditionalZones,
			"gke-node-count-per-zone":       fmt.Sprintf("%d", vmCount),
			"gke-multizone":                 fmt.Sprintf("%t", gkeMultizone),
			"aks-cluster-name":              newStackName,
			"aks-worker-count":              fmt.Sprintf("%d", vmCount),
			"aks-k8s-version":               aksK8sVersion,
			"aks-location":                  aksLocation,
			"aks-vm-size":                   aksVMSize,
			"aks-vm-set-type":               aksVMSetType,
			"disable-additional-node-pools": fmt.Sprintf("%t", disableAdditionalPools),
		}

		var provisioner provisioners.Provisioner
		var ok bool

		if defaultProvisioner != "" {
			provisioner, ok = ProvisionerMap[defaultProvisioner]
			if !ok {
				log.Fatal("Unable to find a provisioner for %s. Exiting.", defaultProvisioner)
			}
		}

		// TODO: This should go away when we make --provisioner mandatory
		if defaultProvisioner == "" {
			log.Warn("DEPRECATED: in future, you must specify the --provisioner flag")
			if esxiHost != "" {
				provisioner = ProvisionerMap["esxiclone"]
			} else {
				provisioner, ok = VMOStoProvisioners[vmOS]
				if !ok {
					log.Fatal("Unable to find a provisioner for OS %s. Exiting.", vmOS)
				}
			}
		}

		instances, err := provisioner.Create(spec, provOptions)
		if err != nil {
			log.Fatal("Failed to provision. Error: %s. Exiting.", err)
		}

		// For GKE, treat all the instances as part of same cluster.
		if defaultProvisioner == "gke" || provisioner.SupportedOSs()[0] == "gke" {
			vmClusters = len(instances)
		}
		instances = append(instances, existingInstances...)
		instances.Save()
		instances.WriteClusterInfo(vmClusters)
	},
}

func addPxDisks(disks []provisioners.DiskSpec) []provisioners.DiskSpec {
	sizeList := make([]int, 0) // List of sizes

	diskSizeList := strings.Split(vmDiskSizePX, ",") // Trim white spaces and commas from string

	if len(diskSizeList) > 1 { // Incase of multiple sizes, convert to int and add to sizeList
		for _, ds := range diskSizeList {
			d, _ := strconv.Atoi(ds)
			sizeList = append(sizeList, d)
		}
	} else { // If single value size, convert to int and add to sizeList
		singleSize, _ := strconv.Atoi(diskSizeList[0])
		sizeList = append(sizeList, singleSize)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizeList))) // Sort sizes from big to small
	log.Debug("Requested PX disk size(s): %v", sizeList)

	for d := 0; d < vmDiskCount; d++ {
		if len(sizeList) > 1 && vmDiskCount > 1 { // Incase of multiple disks or sizes, assign to each disk count
			if len(sizeList) <= d { // If there is less sizes than disks, take last size
				log.Warn("There are more PX disks than PX disk sizes requested, assigning smallest size")
				disks = append(disks, provisioners.DiskSpec{
					Type:   "px",
					SizeGB: sizeList[len(sizeList)-1],
				})
			} else { // Assign sizes to disks, as long as there is enough number of sizes
				disks = append(disks, provisioners.DiskSpec{
					Type:   "px",
					SizeGB: sizeList[d],
				})
			}
		} else { // Incase of 1 disk size, use it for all disks
			disks = append(disks, provisioners.DiskSpec{
				Type:   "px",
				SizeGB: sizeList[len(sizeList)-1],
			})
		}
	}

	return disks
}

func getStackName() string {
	if stackName != defaultStackName {
		// Someone explicitly overrode the name
		stackName = fmt.Sprintf("%s-%s-%s",
			stackName,
			randomdata.Adjective(),
			randomdata.Noun(),
		)
		log.Info("Using given stack name %s", stackName)

	} else if buildinfo.OnJenkins() {
		bi := buildinfo.GetBuildInfo()
		stackName = fmt.Sprintf("%s-%s", bi.Job, bi.BuildNumber)
		log.Info("Using Jenkins to set job name to %s", stackName)

	} else {
		// Someone called with default name
		stackName = fmt.Sprintf("%s-%s-%s",
			stackName,
			randomdata.Adjective(),
			randomdata.Noun(),
		)
		log.Info("Using default stack name %s", stackName)
	}

	// Stack name ends up as hostname, so sanitize
	return stringutils.SanitizeName(stackName)
}

func fixMemory() {
	// Normalize RAM value
	if vmMemory > 1024 {
		log.Warn("RAM value should be in GB")
		vmMemory = vmMemory / 1024
	}
}

func init() {
	rootCmd.AddCommand(vmsCmd)

	vmsCmd.Flags().StringVar(&stackName, "name", defaultStackName, "stack name, if on Jenkins, filled automatically")
	provisionerNames := make([]string, 0)
	for provisionerName, _ := range ProvisionerMap {
		provisionerNames = append(provisionerNames, provisionerName)
	}
	// TODO: in the future, this flag should be made mandatory
	vmsCmd.Flags().StringVar(&defaultProvisioner, "provisioner", "", fmt.Sprintf("Provisioner to use: %s", strings.Join(provisionerNames, ", ")))

	vmsCmd.Flags().IntVar(&vmCount, "count", 4, "number of VMs")
	vmsCmd.Flags().IntVar(&vmCPU, "cpu", 2, "number of CPU cores per VM")
	vmsCmd.Flags().IntVar(&vmHBACount, "hba-count", 1, "number of HBAs (4 maximum) per VM, disks will be distributed accordingly")
	vmsCmd.Flags().IntVar(&vmDiskCount, "disk-count-px", 3, "number of disks in addition to root per VM (0 to disable)")
	vmsCmd.Flags().StringVar(&vmDiskSizePX, "disk-size-px", "128", "size of PX disks in GB, example for multiple disks sizes: 50,100,128,25")
	vmsCmd.Flags().IntVar(&vmMemory, "memory", 8, "VM memory in GB")

	vmsCmd.Flags().BoolVar(&vmDiskDockerEnabled, "disk-docker", true, "enable docker disk (32GB)")
	vmsCmd.Flags().BoolVar(&vmDiskJournalEnabled, "disk-journal", false, "enable journal disk (25GB)")
	vmsCmd.Flags().BoolVar(&vmDiskMetadataEnabled, "disk-metadata", false, "enable metadata disk (64GB)")
	vmsCmd.Flags().BoolVar(&vmDiskCacheEnabled, "disk-cache", false, "enable cache disk")
	vmsCmd.Flags().IntVar(&vmDiskCacheSize, "disk-cache-size", 128, "size of cache disk")

	vmsCmd.Flags().StringVar(&vmOS, "os", "centos/7-base", "")

	vmsCmd.Flags().StringVar(&esxiHost, "esxi-host", "", "ESXi name / IP address")
	vmsCmd.Flags().StringVar(&iksK8sVersion, "iks-k8s-version", "1.10.8", "supported: 1.9.10, 1.10.8, 1.11.3")

	vmsCmd.Flags().StringVar(&kopsKeyId, "kops-key-id", "", "KOPS access key id")
	vmsCmd.Flags().StringVar(&kopsKeySecret, "kops-key-secret", "", "KOPS access key secret")
	vmsCmd.Flags().StringVar(&kopsRegion, "kops-region", pxinfra.DefaultRegion, "KOPS region")
	vmsCmd.Flags().StringVar(&kopsHostedZoneId, "kops-hosted-zone-id", pxinfra.DefaultHostedZoneId, "KOPS hosted zone ID")
	vmsCmd.Flags().StringVar(&kopsK8sVersion, "kops-k8s-version", "1.11.3", "KOPS K8s supported: 1.9.10, 1.10.8, 1.11.3")
	vmsCmd.Flags().StringVar(&kopsAvailabilityZones, "kops-az", pxinfra.DefaultAvailabilityZone, "KOPS ASG availability zones")
	vmsCmd.Flags().StringVar(&kopsVpc, "kops-vpc", pxinfra.DefaultVPC, "KOPS vpc")
	vmsCmd.Flags().StringVar(&kopsSubnets, "kops-subnets", "", "KOPS subnets, must be odd number")
	vmsCmd.Flags().StringVar(&kopsUtilitySubnets, "kops-utility-subnets", "", "KOPS utility subnets, must be odd number")
	vmsCmd.Flags().StringVar(&kopsSecurityGroups, "kops-security-groups", pxinfra.DefaultSecurityGroup, "KOPS additional security groups")
	vmsCmd.Flags().StringVar(&kopsUserGroup, "kops-user-group", pxinfra.DefaultUserGroup, "KOPS IAM user group that has the required roles")
	vmsCmd.Flags().StringVar(&kopsImage, "kops-image", "", "KOPS AMI image")
	vmsCmd.Flags().StringVar(&kopsUser, "kops-user", "", "KOPS AMI image ssh username")
	vmsCmd.Flags().IntVar(&kopsKeyPropagationTime, "kops-key-propagation", pxinfra.DefaultKeyPropagationTime, "KOPS waiting time for access key propagation")
	vmsCmd.Flags().IntVar(&kopsMasterVolSize, "kops-master-vol-size", 0, "KOPS master root volume size in GB")
	vmsCmd.Flags().IntVar(&kopsNodeVolSize, "kops-node-vol-size", 0, "KOPS node root volume size in GB")

	vmsCmd.Flags().StringVar(&gkeName, "gke-name", "", "GKE cluster name")
	vmsCmd.Flags().StringVar(&gkeZone, "gke-zone", "us-central1-a", "GKE cluster zone")
	vmsCmd.Flags().StringVar(&gkeAdditionalZones, "gke-additional-zones", "us-central1-a,us-central1-b,us-central1-c", "GKE cluster additiona zones (in case of multizone setup)")
	vmsCmd.Flags().BoolVar(&gkeMultizone, "gke-multizone", true, "Whether to setup a multi-zone GKE cluster")

	vmsCmd.Flags().StringVar(&aksK8sVersion, "aks-k8s-version", "1.12.8", "Default: 1.12.8. To get currently supported AKS version, run `az aks get-versions --location <location> --output table`")
	vmsCmd.Flags().StringVar(&aksLocation, "aks-location", "eastus", "Default: eastus")
	vmsCmd.Flags().StringVar(&aksVMSize, "aks-vm-size", "Standard_DS3_v2", "Azure VM size for AKS worker nodes. Default: Standard_DS2_v2 (4 vCPU, 14 GB RAM)")
	vmsCmd.Flags().StringVar(&aksVMSetType, "aks-vm-set-type", "VirtualMachineScaleSets", "Azure agent pool vm set type.VirtualMachineScaleSets or AvailabilitySet. Default: VirtualMachineScaleSets")
	vmsCmd.Flags().BoolVar(&disableAdditionalPools, "disable-additional-node-pools", false, "Disable additional node pools (to run torpedo), by default we add additional pools ")

	vmsCmd.Flags().IntVar(&vmClusters, "groupvms-cluster", 4, "Group given VMs to one cluster")
}
