#!/bin/bash

# Set path
cd "$(dirname "$0")" || exit 1
cd ..

docker-compose -p laskb-server-api up -d