#!/bin/bash

# Exit on error.
set -e

# Portainer Stop.
docker stop portainer
docker rm portainer

# Docker Image Update.
docker pull portainer/portainer-ce

# Portainer Start.
docker run -d -p 8000:8000 -p 9443:9443 --name=portainer --restart=always \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v portainer_data:/data portainer/portainer-ce:latest

# Source:
# https://docs.portainer.io/start/upgrade/docker
