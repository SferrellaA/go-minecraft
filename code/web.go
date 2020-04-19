package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func mainPage(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html", web.indexHTML)
}

func webAsset(ctx *gin.Context) {
	switch ctx.Param("asset") {
	case "main.css":
		ctx.Data(http.StatusOK, "text/css", web.mainCSS)
	case "minecraft.png":
		ctx.Data(http.StatusOK, "image/png", web.minecraftPNG)
	default:
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}

// serverStatus returns a json of the current server status
func serverStatus(ctx *gin.Context) {
	// Query the minecraft server for it's current status
	serverStatus, playerCount, maxPlayers := checkServer()

	// Generate HTML
	bodyHTML := "<div style=\"color: white;text-align: center;font-weight: bold;white-space: pre-wrap;\">"
	bodyHTML += serverStatus
	if serverStatus == "Server Online!" { // Check there's no errors
		bodyHTML += "<br>" + playerCount + "/" + maxPlayers
	}
	bodyHTML += "</div>"

	// Present status to client
	ctx.Data(http.StatusOK, "text/html", genHTML(bodyHTML))
}

// listActions lists available actions for server management
func listActions(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html", web.actionsHTML)
}

// serverAction runs a given script
func serverAction(ctx *gin.Context) {
	switch ctx.Param("serverAction") {
	case "start":
		runScript("start.sh")
	case "stop":
		runScript("stop.sh")
	case "reboot":
		runScript("reboot.sh")
	case "backup":
		runScript("backup.sh")
	default:
		ctx.String(http.StatusOK, "unknown action")
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/server")
}

// listBackups shows a list of files in backup folder
func listBackups(ctx *gin.Context) {
	// Get the list of files
	files, err := ioutil.ReadDir(cfg.backupsDir)
	errFail(err)

	// Generate HTML
	bodyHTML := ""
	for _, f := range files {
		bodyHTML += "<a href=\"/backup/" + f.Name() + "\">" + f.Name() + "</a><br>"
	}

	// Serve the html list
	ctx.Data(http.StatusOK, "text/html", genHTML(bodyHTML))
}

// getBackup serves a given file from the backup folder
func getBackup(ctx *gin.Context) {
	target := filepath.Join(cfg.backupsDir, ctx.Param("version"))
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+ctx.Param("version"))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(target)
}
