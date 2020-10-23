package main

import (
	"fmt"
	"io/ioutil"
	"path"
)

func readDir(specDir string) error {
	appList, err := ioutil.ReadDir(specDir)
	for _, file := range appList {
		if file.IsDir() && file.Name() == "mysql" {
			cloudProviderDirList, err := ioutil.ReadDir(path.Join(specDir, file.Name()))
			fmt.Printf("ERR-1: %v\n", err)
			fmt.Printf("RK=> Subdirect in %v: %d\n", specDir, len(cloudProviderDirList))
		}
	}
	return err
}

func main() {
	specID := "/root/git/go/src/github.com/rohit-go-utils/specs/"
	err := readDir(specID)
	fmt.Printf("ERR-2: %v\n", err)
}
