#!/bin/bash
date +"[%d/%m/%Y - %T]  Starting!"
pushd server &> /dev/null
java -Xmx1024M -Xms1024M -jar server.jar nogui &>> ../log.txt
popd &> /dev/null