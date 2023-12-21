#!/bin/bash

# Exit on error.
set -e

# Docker Registry Start.
docker run -d -p 5000:5000 --restart=always --name registry registry:2

# Source:
# https://docs.docker.com/registry/deploying/
