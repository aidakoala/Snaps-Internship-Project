package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)
import "io/ioutil"

// prints first n lines from a string
func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func main() {
	// wordPtr := flag.String("find", "", "snap name to be searched")
	// request catre server http
	resp, err := http.Get("http://localhost:9091/search/")
	checkError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("get: %s\n", string(body))

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
