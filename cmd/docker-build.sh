#!/bin/bash

cd "$(dirname "$0")" || { echo "Failed to change directory"; exit 1; }
cd .. || { echo "Failed to change directory"; exit 1; }

docker build -t laskb-server .