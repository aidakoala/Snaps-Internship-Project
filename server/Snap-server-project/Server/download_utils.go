package main

import (
	"io/ioutil"
	"os"
	"strings"
)

// get the content of the .json files if possible and return the struct with infos
func getSnapFile(dir os.FileInfo) []byte {
	var err error
	path := getPathToSnaps() + dir.Name()
	files, err := ioutil.ReadDir(path)
	checkError(err)

	var fileBytes []byte
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".snap") {
			writeDebugInfo(file.Name() + "\n")
			fileBytes, err = ioutil.ReadFile(path + "/" + file.Name())
			checkError(err)
			break
		}
	}
	return fileBytes
}

// get the content of the .json files if possible and return the struct with infos
func getSnapFileLocation(dir string) os.FileInfo {
	var err error
	path := getPathToSnaps() + dir
	files, err := ioutil.ReadDir(path)
	checkError(err)
	var f os.FileInfo

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".snap") {
			//writeDebugInfo(file.Name() + "\n")
			f = file
			break
		}
	}
	return f
}
