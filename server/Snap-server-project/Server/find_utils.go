package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

// find the name of the searched snap
func getSearchedSnapName(url string) string {
	splittedString := strings.Split(url, "=")
	return splittedString[1]
}

// get the content of the .json files if possible and return the struct with infos
func getFindJsonFile(dir os.FileInfo) snapDetails {
	path := getPathToSnaps() + dir.Name()
	files, err := ioutil.ReadDir(path)
	checkError(err)

	var details snapDetails
	for _, file := range files {
		if strings.Contains(file.Name(), "find") {
			fileBytes, err := ioutil.ReadFile(path + "/" + file.Name())
			checkError(err)

			err = json.Unmarshal(fileBytes, &details)
			checkError(err)

			return details
		}
	}
	return details
}
