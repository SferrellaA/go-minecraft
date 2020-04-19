package main

import (
	"io/ioutil"
	"path"
)

type configs struct {
	serverName      string
	port            string
	autoKill        bool
	portListen      bool
	autoKillTimeout int
	scriptsDir      string
	backupsDir      string
	webpageDir      string
}

func readConfigs() (configs, error) {
	cfg := configs{
		serverName:      "Alex's Minecraft Server",
		port:            "8080",
		autoKill:        true,
		portListen:      true,
		autoKillTimeout: 5,
		scriptsDir:      "./scripts",
		backupsDir:      "../backups",
		webpageDir:      "../web",
	}
	return cfg, nil
}

type webAssets struct {
	minecraftPNG []byte
	mainCSS      []byte
	indexHTML    []byte
	actionsHTML  []byte
}

func loadAssets() (w webAssets, err error) {
	w.minecraftPNG, err = ioutil.ReadFile(path.Join(cfg.webpageDir, "minecraft.png"))
	if nil != err {
		return
	}

	w.mainCSS, err = ioutil.ReadFile(path.Join(cfg.webpageDir, "main.css"))
	if nil != err {
		return
	}

	w.indexHTML, err = ioutil.ReadFile(path.Join(cfg.webpageDir, "index.html"))
	if nil != err {
		return
	}

	w.actionsHTML, err = ioutil.ReadFile(path.Join(cfg.webpageDir, "actions.html"))
	if nil != err {
		return
	}

	return
}
