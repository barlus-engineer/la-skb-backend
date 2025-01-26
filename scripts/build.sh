#!/bin/bash

# Set path
cd "$(dirname "$0")" || exit 1
cd ..

# Define the Docker image name and tag
IMAGE_NAME="laskb-server-api"
IMAGE_TAG="latest"

# Build the Docker image
echo "Building Docker image ${IMAGE_NAME}:${IMAGE_TAG}..."
docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Docker image ${IMAGE_NAME}:${IMAGE_TAG} built successfully."
else
    echo "Failed to build Docker image ${IMAGE_NAME}:${IMAGE_TAG}."
    exit 1
fi