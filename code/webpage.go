package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Return a 404 message 
func return404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 :("))
}

// Let users control the minecraft server
func server(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("web/server.html")
	if errCheck(err) {
		return404(w, r)
		return
	}
	fmt.Fprint(w, string(html))
}

// Start the Minecraft server
func server_start(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://"+r.Host+"/server", 302)
	runScript("start.sh")
	running = true
}

// Stop the Minecraft server
func server_stop(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://"+r.Host+"/server", 302)
	runScript("stop.sh")
	running = false
}

// Reboot the Minecraft server
func server_reboot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://"+r.Host+"/server", 302)
	runScript("reboot.sh")
}

// Show the current status of the server
func status(w http.ResponseWriter, r *http.Request) {
	
	// Query the minecraft server for it's current status
	connectMessage, currentPlayers, maxPlayers := checkServer()
	
	// Craft text to show on status tab of admin page
	content := connectMessage;
	if connectMessage == "Server Online!" {
		content += "&#10;"
		content += currentPlayers + "/" + maxPlayers + " players online"
	}
	
	// Weave in and present html
	html, err := ioutil.ReadFile(dir+"/web/status.html")
	if errCheck(err) {
		return404(w, r)
		return
	}
	content = strings.Replace(string(html), "<!--Content Goes Here-->", content, 1)
	fmt.Fprint(w, content)
}

// Set up the web server
func configure_server() {
	http.Handle("/", http.FileServer(http.Dir(dir+"/web")))
	http.HandleFunc("/server", server)
	http.HandleFunc("/server/start", server_start)
	http.HandleFunc("/server/stop", server_stop)
	http.HandleFunc("/server/reboot", server_reboot)
	http.HandleFunc("/status", status)
	http.HandleFunc("/fake", return404)
	http.Handle("/backups/", http.StripPrefix("/backups/", http.FileServer(http.Dir(dir+"/backups"))))
}