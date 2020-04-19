package main

import (
	"net"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

// Run a particular script
func runScript(script string) {
	cmd := exec.Command("bash", path.Join(cfg.scriptsDir, script))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go cmd.Run()
}

// Turn the server off after five minutes of no players
func autokillServer(timeout int) {

	// Number of minutes server has sat empty
	emptyMinutes := 0
	for {
		// Check server once every minute
		time.Sleep(time.Minute)
		serverStatus, playerCount, _ := checkServer()

		// Server is on and empty
		if serverStatus == "Server Online!" && playerCount == "0" {

			// Count up to ten minutes
			if emptyMinutes < timeout {
				emptyMinutes++
			} else {

				// Shut the server down
				runScript("stop.sh")
			}
		} else {

			// Reset to 0 if the server's down or players on
			emptyMinutes = 0
		}
	}

}

// Return information about the current state of the minecraft server
func checkServer() (serverStatus, playerCount, maxPlayers string) {

	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:25565")
	if nil != err {
		serverStatus = "Could not contact Server"
		return
	}

	// Ask the server for info
	conn.Write([]byte("\xFE\x01"))
	rawData := make([]byte, 512)
	_, err = conn.Read(rawData)
	defer conn.Close()
	if nil != err {
		serverStatus = "Problem talking with Server"
		return
	}

	// Make sure the server actually replied
	if nil == rawData || len(rawData) == 0 {
		serverStatus = "Empty response received from Server"
		return
	}

	// Successful query for data from server
	data := strings.Split(string(rawData[:]), "\x00\x00\x00")
	serverStatus = "Server Online!"
	// s.ServerVersion = data[2]
	// s.Motd = data[3]
	playerCount = data[4]
	maxPlayers = data[5]
	return
}

// genHTML generates HTML so the mess isn't elsewhere
func genHTML(body string) []byte {
	htmlString := "<head>"
	htmlString += "<meta http-equiv=\"Cache-Control\" content=\"no-cache, no-store, must-revalidate\">"
	htmlString += "<meta http-equiv=\"Pragma\" content=\"no-cache\">"
	htmlString += "<meta http-equiv=\"Expires\" content=\"0\">"
	htmlString += "<link rel=\"stylesheet\" href=\"/web/main.css\" type=\"text/css\">"
	htmlString += "<link rel=\"stylesheet\" media=\"screen\" href=\"https://fontlibrary.org/face/minecraftia\" type=\"text/css\"/>"
	htmlString += "</head>"
	htmlString += "<body style=\"font-family: 'MinecraftiaRegular';\">"
	htmlString += body
	htmlString += "</body>"
	return []byte(htmlString)
}
