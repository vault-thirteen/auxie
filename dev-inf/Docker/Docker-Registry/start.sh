#!/bin/bash

# Exit on error.
set -e

# Docker Registry Start.
docker container start registry

# Source:
# https://docs.docker.com/registry/deploying/
