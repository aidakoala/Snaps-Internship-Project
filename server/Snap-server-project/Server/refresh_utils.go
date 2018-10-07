package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

// get the content of the .json files if possible and return the struct with infos
func getRefreshJsonFile(dir os.FileInfo) storeSnap {
	path := getPathToSnaps() + dir.Name()
	files, err := ioutil.ReadDir(path)
	checkError(err)

	var details storeSnap
	for _, file := range files {
		if strings.Contains(file.Name(), "refresh") {
			fileBytes, err := ioutil.ReadFile(path + "/" + file.Name())
			checkError(err)

			err = json.Unmarshal(fileBytes, &details)
			checkError(err)

			return details
		}
	}
	return details
}
