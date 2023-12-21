#!/bin/bash

# Exit on error.
set -e

# Push an Image to the Docker Registry.
IMAGE_ID="1234567" # This ID is known after Docker Image is built.
REGISTRY="localhost:5000"
IMAGE_NAME="my_image"
docker tag $IMAGE_ID $REGISTRY/$IMAGE_NAME
docker push $REGISTRY/$IMAGE_NAME

# Source:
# https://docs.docker.com/registry/deploying/
