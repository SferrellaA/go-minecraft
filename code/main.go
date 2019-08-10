package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var (
	dir     = ""    // Path of the golang binary
	running = false // Server is off by default
)

// Return true if err is an error
func errCheck(err error) bool {
	if nil != err {
		return true
	}
	return false
}

// Exit the program if err is an error
func errFail(err error) {
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	
	// Find where we're running from
	d, err := os.Executable()
	if nil != err {
		panic(err)
	}
	dir = filepath.Dir(d)

	// Turn server off if no one on
	go autokillServer()

	// Handle the admin page
	configure_server()
	err = http.ListenAndServe(":80", nil)
	errFail(err)
}
