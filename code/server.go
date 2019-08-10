package main

import (
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Run a particular script
func runScript(script string) {
	cmd := exec.Command("bash", dir+"/scripts/"+script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go cmd.Run()
}

// Turn the server off after five minutes of no players
func autokillServer() {
	// Number of minutes server has sat empty
	emptyMinutes := 0
	for {
		// Check server once every minute
		time.Sleep(time.Minute)
		status, players, _ := checkServer()

		// Server is on and empty
		if status == "Server Online!" && players == "0" {

			// Count up to ten minutes
			if emptyMinutes < 10 {
				emptyMinutes += 1
			} else {
				
				// Shut the server down
				runScript("stop.sh")
				running = false
			}
		} else {
			
			// Reset to 0 if the server's down or players on
			emptyMinutes = 0
		}
	}
}

// Return information about the current state of the minecraft server
func checkServer() (connectStatus, currentPlayers, maxPlayers string) {
	
	// Default values
	connectStatus = "Server Offline"
	currentPlayers = ""
	maxPlayers = ""

	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:25565")
	if errCheck(err) {
		connectStatus = "Could not contact Server"
		return
	}

	// Ask the server for info
	conn.Write([]byte("\xFE\x01"))
	raw_data := make([]byte, 512)
	_, err = conn.Read(raw_data)
	defer conn.Close()
	if errCheck(err) {
		connectStatus = "Problem talking with Server"
		return
	}
	
	// Make sure the server actually replied
	if nil == raw_data || len(raw_data) == 0 {
		connectStatus = "Empty response received from Server"
		return
	}

	// Successful query for data from server
	data := strings.Split(string(raw_data[:]), "\x00\x00\x00")
	connectStatus = "Server Online!"
	currentPlayers = data[4]
	maxPlayers = data[5]
	return

	/*
		fmt.Fprint(w, "Version:", data[2])
		fmt.Fprint(w, "Motd:", data[3])
		fmt.Fprint(w, "Current Players:", data[4])
		fmt.Fprint(w, "Max Players:", data[5])
	*/
}
