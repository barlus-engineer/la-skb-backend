#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
BLUE='\033[0;34m'
RESET='\033[0m'

cd "$(dirname "$0")" || exit 1
echo -e "${CYAN}Starting build process...${RESET}"

time_counter() {
    local pid=$!
    local start_time=$(date +%s%3N)
    
    while [ "$(ps a | awk '{print $1}' | grep "$pid")" ]; do
        local current_time=$(date +%s%3N)
        local elapsed=$((current_time - start_time))
        
        local seconds=$((elapsed / 1000))
        local milliseconds=$((elapsed % 1000))
        local seconds_remaining=$((seconds % 60))
        
        printf "%02d:%02d\r" $seconds_remaining $milliseconds
        sleep 0.01
    done
    printf "${GREEN}Build finished in %02d:%02ds${RESET}\n" $seconds_remaining $milliseconds
}

echo -e "${CYAN}Compiling the Go project...${RESET}"
{
    go build -o laskb_server app/main.go
} & time_counter

if [ ! -f laskb_server ]; then
    echo -e "${RED}Build failed!${RESET}"
    exit 1
fi
echo -e "${GREEN}Build successful!${RESET}"

echo -e "${CYAN}Setting up the project files...${RESET}"
{
    cd .. || exit 1
    [ -d "out" ] && rm -r out
    mkdir out
    mv cmd/laskb_server out/laskb_server
    cp -r locale out/locale
}

echo -e "${CYAN}Checking for .env file...${RESET}"

if [ -f ".env" ]; then
    echo -e "${CYAN}.env file found. Copying to out directory...${RESET}"
    cp .env out/
else
    echo -e "${YELLOW}.env file not found. Skipping copy...${RESET}"
fi

if [ $? -eq 0 ]; then
    echo -e "${GREEN}Setup complete! Files are in the 'out' folder.${RESET}"
else
    echo -e "${RED}Setup failed!${RESET}"
    exit 1
fi