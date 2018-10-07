package main

import (
	"io"
	"os"
)

// get debug file name
func getDebugFileName() string {
	dir, err := os.Getwd()
	checkError(err)
	return dir + "/debug.log"
}

// func to write debug info to a file
func writeDebugInfo(s string) {
	f, _ := os.OpenFile(getDebugFileName(), os.O_APPEND|os.O_RDWR, 0666)
	io.WriteString(f, s+"\n")
}
