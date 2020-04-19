package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	cfg configs
	web webAssets
)

// Exit the program if err is an error
func errFail(err error) {
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	var err error

	// Gather settings
	s := gin.Default()
	cfg, err = readConfigs()
	errFail(err)
	web, err = loadAssets()
	errFail(err)

	// Begin sub-processes
	if cfg.autoKill {
		go autokillServer(cfg.autoKillTimeout)
	}

	// Establish routes
	s.GET("/", mainPage)
	s.GET("/web/:asset", webAsset)
	s.GET("/status", serverStatus)
	s.GET("/server", listActions)
	s.GET("/server/:serverAction", serverAction)
	s.GET("/backup", listBackups)
	s.GET("/backup/:version", getBackup)

	// Start the server
	s.Run("localhost:" + cfg.port)

	// Turn server off if no one on
	// go autokillServer()
}
