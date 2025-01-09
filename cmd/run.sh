#!/bin/bash

cd "$(dirname "$0")" || { echo "Failed to change directory"; exit 1; }
./build.sh || { echo "Build failed"; exit 1; }
clear
cd ../out || { echo "Failed to change directory to 'out'"; exit 1; }
./laskb_server || { echo "Failed to start the server"; exit 1; }
