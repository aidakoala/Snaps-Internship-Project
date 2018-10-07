package main

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "golang.org/x/crypto/sha3"

	"github.com/snapcore/snapd/osutil"
)

// EncodeDigest encodes the digest from hash algorithm to be put in an assertion header.
func EncodeDigest(hash crypto.Hash, hashDigest []byte) (string, error) {
	algo := ""
	switch hash {
	case crypto.SHA512:
		algo = "sha512"
	case crypto.SHA3_384:
		algo = "sha3-384"
	default:
		return "", fmt.Errorf("unsupported hash")
	}
	if len(hashDigest) != hash.Size() {
		return "", fmt.Errorf("hash digest by %s should be %d bytes", algo, hash.Size())
	}
	return base64.RawURLEncoding.EncodeToString(hashDigest), nil
}

// SnapFileSHA3_384 computes the SHA3-384 digest of the given snap file.
// It also returns its size.
func SnapFileSHA3_384(snapPath string) (digest string, size uint64, err error) {
	sha3_384Dgst, size, err := osutil.FileDigest(snapPath, crypto.SHA3_384)
	if err != nil {
		return "", 0, fmt.Errorf("cannot compute snap %q digest: %v", snapPath, err)
	}

	sha3_384, err := EncodeDigest(crypto.SHA3_384, sha3_384Dgst)
	if err != nil {
		return "", 0, fmt.Errorf("cannot encode snap %q digest: %v", snapPath, err)
	}
	return sha3_384, size, nil
}

func findSnaps(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		writeDebugInfo(r.URL.String())
		snapName := getSearchedSnapName(parseURL(r.URL.String(), 2))
		writeDebugInfo("SNAP NAME IS:" + snapName + "\n")

		var structToBeSent snapsNall
		var packs snapPacks
		var URLs topLevelLinks
		var selfLink linkTypes

		dir := getSnapDir(snapName)
		snapDetails := getFindJsonFile(dir)
		writeDebugInfo("JSON INFO:\n" + snapDetails.Title + "\n")
		packs.ClickIndex_Package = append(packs.ClickIndex_Package, snapDetails)

		structToBeSent.Embedded = packs
		URLs.Href = r.URL.String()
		selfLink.SelfLink = URLs
		structToBeSent.Links = selfLink

		response, err := json.Marshal(structToBeSent)
		checkError(err)

		w.Header().Set("Content-Type", "application/hal+json")
		w.Write(response)
	}
}

func refresh(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b, e1 := ioutil.ReadAll(r.Body)
		checkError(e1)
		writeDebugInfo(string(b) + "\n\n")

		data := refreshDetails{}
		e2 := json.Unmarshal(b, &data)
		checkError(e2)

		// the action requested by the client
		action := data.Actions[0]
		writeDebugInfo(action.Action + "\n" + action.Name + "\n")
		snapName := action.Name

		// send json response
		var structToBeSend snapActionResultList
		var res result

		res.Channel = action.Channel
		res.InstanceKEY = action.InstanceKEY
		res.Name = action.Name
		res.Res = action.Action

		dir := getSnapDir(snapName)
		snapDetails := getRefreshJsonFile(dir)
		res.Snap = snapDetails
		// get yaml file

		res.SnapID = snapDetails.SnapID
		structToBeSend.Results = append(structToBeSend.Results, res)

		response, err := json.Marshal(structToBeSend)
		checkError(err)

		w.Header().Set("Content-Type", "application/hal+json")
		w.Write(response)
	}
}

func downloadSnap(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		writeDebugInfo(r.URL.String())

		// get snap name to be downloaded
		snapFullName := r.URL.Path[len("/api/v1/snaps/download/"):]
		writeDebugInfo(snapFullName + "\n")
		tokens := strings.Split(snapFullName, "_")
		snapName := tokens[0]
		writeDebugInfo(snapName + "\n")
		dir := getSnapDir(snapName)
		snapBinary := getSnapFile(dir)

		w.Write(snapBinary)
	}
}

func assertionsRevision(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Printf("url2 assertion:" + r.URL.Path + "\n")

		var f os.FileInfo
		var sha384Dgst string
		f = getSnapFileLocation("opensshd")
		fmt.Printf("Sha FIle " + f.Name() + "\n")
		sha384Dgst, _, _ = SnapFileSHA3_384("/home/heri/snaps-sources/opensshd/" + f.Name())
		fmt.Printf("Sha digest: " + sha384Dgst + "\n")

		w.Header().Set("Content-Type", "application/x.ubuntu.assertion")

		w.Write([]byte("type: snap-revision\n" + "authority-id: keysight\n" + "revision: 14\n" + "snap-sha3-384: " + sha384Dgst + "\n" + "developer-id: keysight\n" + "snap-id: RmBXKl6HO6YOC2DE4G2q1JzWImC04EUy\n" + "snap-revision: 1\n" + "snap-size: " + strconv.FormatInt(f.Size(), 10) + "\n" + "timestamp: 2018-09-12T10:23:51+00:00\n" + "sign-key-sha3-384: bCEJWw5FZgRZDFuZhPAdPgo1lZgljyAPlEOcVD3UCWzldtB4qc8yMtKoMgejdXEY\n\nAcLBXAQAAQoABgUCW5rPmAAKCRCgMssgMDM9MQT9D/9c7jFxt++LOoYfLe73fot8uvnwmsH5m3JK\nUBSRIFzJoiUX7B0jdEidsgJtCF5KFIWLW1nR238IRGL2kj2dqkC3vHJ7yoVIu92dHlJlreOQk5eA\n9UNrV4Y1EQExBhnuZ4O+U4fK1HCXgAYH0IfUZ48+Pd2t/xEhGN7n4VOs7wijLRADRBE8HJLM4w8E\nN0QNknnILWC/pF24CXqBlnnmXWcvWHE3xQXJFZTmJLRXoocbXRpICJfrysFIFuPV+mzuWA/vJwxP\nIzES2GJxX9lTJbBgfpCeZt3Adp2WgIHTafuIz2gG7ZmOrvijMOFNoIQHhu7ni2qz3hx3206Vv7nN\npHP8JjKlVAdOSGJu8ZB2qO2ICxdM3Yl0zHIFE6L218BNvdF6uGgdKCQ2hLTwdXPwpjND/t8uP8ib\nqmhbqLya4/zFddVeWyas1qNnyYgcRAwA06nkIYZzruUoSoQlISJ8Q6outnCrUmtHF/deWuyxJc3L\nL23iyeFJQCaWdzyUdQQbuutEB6SDxKlLnoNF8BHF0Nj1Lr4LDFpIO2A+QPqmj3UjZgxY89g3Yixm\nzZo80rbgudTTS4QinWh/Nn7h2l/kLUs2nwFFvnlBGwxPVEc0Z0HK1DdIWYzFFUsrT3Nw6xfVFLbN\nnfoqxaEfI/eTdQL5r/UP3fTOZpHmxCtQYfL2qyQSuw==\n"))
	}
}

func assertionsDeclaration(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/x.ubuntu.assertion3")
		var writtenBytes int
		writtenBytes , err := w.Write([]byte("type: snap-declaration\n" + "authority-id: keysight\n" + "series: 16\n" + "snap-id: RmBXKl6HO6YOC2DE4G2q1JzWImC04EUy\n" + "publisher-id: keysight\nsnap-name: opensshd\n" + "timestamp: 2018-09-12T10:23:51+00:00\n" + "sign-key-sha3-384: bCEJWw5FZgRZDFuZhPAdPgo1lZgljyAPlEOcVD3UCWzldtB4qc8yMtKoMgejdXEY\n\nAcLBXAQAAQoABgUCW5l4HgAKCRCgMssgMDM9MeVED/wMF2WoIYj+OuKfpfUVUHDD+k5N9DglyjqR\n+kNlszG8drw5zqU1Xz1Xtdw5ozs6kTmen8/LvYEvWbzkoaZZswkn95wmaEShvK40aJbMRCjsfmk1\ntpKz90WB3j7YpweCrq0REVXbpOvcJHOSMtcXhBMl0aOC722zBwyy/gfs34wvpWw9ny53gvu+vW8S\nMKWVOKKAW4u/f1eKECEJPIB0Hs0yHWFntNHHkb0IjPxhWZPVf+yoWRd5Y/C02VG13heMPX5Bpf/S\n/irlAT7jWhl6W+yfFnKs0sEIW6tJQaDNn3Z8XWNYQi6QDhUNORTBoyeHSzaFZxQUVA8AZUWJFfmW\nsV3HXNT+oLgQXRiZAXTiCmLq8Hw6KxjxH1sBTALdppVyR5gt4+Z4MHhKGUCDns3xQu7yrrUrJGqv\nvPbZcD4sPQxv3zctbvNeK9Gt8DDmpgCanc/xQjaOfOzzl5m7ep8/LDtvtektNaTRTRrTsIkwCPlQ\nUUDt54zFjqbe+JcXlt0URq8v3SeqlVKkaRGtrzNMyvf7ieFfm50JmrYjparUtoAu+pJ/vwJDT+5e\nKrjClSN8mSFUscpZe1n9dLaoEwkNplrRHV7z8963ve0J9xVgkfq90yd/lRXVgM7VhVEtsObxYU2r\nan1yBo8I2YEg0ZTQZLJkkKHiTITNbt4BT9tpe7yWhA==\n"))
	fmt.Printf("WriitenBytes = " +  strconv.FormatInt(int64(writtenBytes),10) + "\n")
	checkError(err)
	}
}

func main() {
	// debug flag
	// when set, all debug info will be displayed in a .log file
	debug := flag.Bool("debug", false, "")
	flag.Parse()
	fmt.Printf("debug flag = %t\n", *debug)

	path := getDebugFileName()
	// delete previous .log file
	_ = os.Remove(path)

	if *debug {
		// create file to write debug info
		_, err := os.Create(path)
		checkError(err)
	}

	// FIND
	http.HandleFunc("/api/v1/snaps/search", findSnaps)

	// REFRESH + INSTALL
	http.HandleFunc("/v2/snaps/refresh", refresh)

	// DOWNLOAD A .snap FILE
	http.HandleFunc("/api/v1/snaps/download/", downloadSnap)

	// SERVE ASSERTIONS FOR A CERTAIN FILE
	http.HandleFunc("/api/v1/snaps/assertions/snap-revision/", assertionsRevision)
	http.HandleFunc("/api/v1/snaps/assertions/snap-declaration/", assertionsDeclaration)

	// err := http.ListenAndServeTLS("localhost:9090", "server.crt", "server.key", nil)
	err := http.ListenAndServe("0.0.0.0:9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
