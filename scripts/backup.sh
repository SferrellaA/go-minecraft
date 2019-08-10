#!/bin/bash

# Move into the server directory
stuff_to_backup=(
	"banned-ips.json"\
	"banned-players.json"\
	"ops.json"\
	"server.properties"\
	"usercache.json"\
	"whitelist.json"\
	"world"\
)

# Generate the name of the backup file
name_of_backup=`date '+%d-%B-%y--%H%M'`.zip
zip -r backups/$name_of_backup ${stuff_to_backup[@]/#/server\/}

# Remove old backups
oldest_backup_prefix=`date +%d-%B-%y -d "7 days ago"`
for i in backups/$oldest_backup_prefix--*.zip; do
	if [ -f $i ]; then rm $i; fi
done