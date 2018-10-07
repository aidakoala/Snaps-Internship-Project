package main

import (
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

// get the path to the directory where the snaps are saved at
func getPathToSnaps() string {
	usr, err := user.Current()
	checkError(err)
	return usr.HomeDir + "/snaps-sources/"
}

// directories that represent snaps
func getAvailableSnaps(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	checkError(err)
	return files
}

// parse the URL params func if we need it someday
func parseURL(url string, n int) string {
	paramStr := strings.Split(url, "?")
	writeDebugInfo("paramStr = " + paramStr[1] + "\n")

	params := strings.Split(paramStr[1], "&")
	writeDebugInfo("params = " + params[n] + "\n")

	return params[n]
}

func getSnapDir(snapName string) os.FileInfo {
	var f os.FileInfo
	snapDirectories := getAvailableSnaps(getPathToSnaps())
	// writeDebugInfo("DIRS:\n")
	for _, file := range snapDirectories {
		// writeDebugInfo("FILE NAME: " + file.Name() + "\n")
		if file.Mode().IsDir() && strings.Contains(file.Name(), snapName) {
			f = file
			break
		}
	}
	return f
}

// func getYamlFile(snapName string) string {
// 	dirs, err := getSnapDir(snapName)
// }
