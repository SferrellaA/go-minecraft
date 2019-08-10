#!/bin/bash

# Build the webserver binary
build_binary () {
	if [ -d "code" ]; then
		go build -o handler ./code/*
		sudo setcap 'cap_net_bind_service=+ep' ./handler
	else
		echo "Cannot find code to build server with!"
		exit 1
	fi
}

# Check if a folder and it's files exist
check_folder () {
	if [ -d $1 ]; then
		for f in ${@:2}; do
			if ! [ -f "$1/$f" ]; then
				echo "Could not find \"$1/$f\"!"
				exit 1
			fi
		done
	else
		echo "Could not find folder \"$1\"!"
		exit 1
	fi
}

# Start the server
run_server () {
	touch server.log
	./handler &>> server.log &
	disown
}

# If the binary does not exist, build it
if ! [ -f "handler" ]; then
	build_binary
fi

# Check that important folders exist
check_folder "backups"
check_folder "server" "server.jar" "eula.txt"
check_folder "web" "index.html" "main.css" "minecraft.png" "server.html" "status.html"
check_folder "scripts" "backup.sh" "reboot.sh" "start.sh" "stop.sh"

# Run the server
run_server