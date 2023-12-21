#!/bin/bash

# Exit on error.
set -e

# Docker Registry Stop.
docker container stop registry

# Source:
# https://docs.docker.com/registry/deploying/
