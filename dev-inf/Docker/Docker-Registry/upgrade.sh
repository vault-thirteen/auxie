#!/bin/bash

# Exit on error.
set -e

# Portainer Stop.
docker stop registry
docker rm registry

# Docker Image Update.
docker pull registry:2

# Start.
docker run -d -p 5000:5000 --restart=always --name registry registry:2

# Source:
# https://docs.portainer.io/start/upgrade/docker
