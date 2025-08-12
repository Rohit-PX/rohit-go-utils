package main

/*RunPoolServiceDisruptiveTests covers the cases where pool deletion should not happen
1. Deleting a pool without entering into maintenance mode
2. Deleting a pool having volume
3. Deleting a pool with resync happening because of ha-increase
*/
func (pDeleteManager *PdeleteManager) RunPoolServiceDisruptiveTests() error {
	utils.Debug.Println("Starting RunPoolServiceDisruptiveTests")
	var err error
	var poolDeleteNode *node_pkg.Node
	job := pDeleteManager.Job
	test := pDeleteManager.Test
	instances := job.Stacks[0].Provider.Instances

	if len(instances) < 3 {
		test.Skip(fmt.Sprintf("Only %d instances in the cluster, skip the test", len(instances)))
	}

	storageNodes := GetXStorageNodesRandomly(3)
	numPools := 0
	for _, storageNode := range storageNodes {
		numPools, err = storageNode.GetNumOfPools()
		if err != nil {
			return err
		}
		// Trying to get 1 node for doing pool delete operation and another node for ha-update
		if numPools > 1 {
			poolDeleteNode = storageNode
			break
		}
	}
	if poolDeleteNode == nil {
		return fmt.Errorf("No storage node has more than 1 pool for running RunPoolServiceDisruptiveTests")
	}
}
