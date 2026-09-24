#!/bin/bash

# Exit on error.
set -e

# Registry Stop.
docker stop registry
docker rm registry

# Docker Image Update.
docker pull registry:3

# Registry Start.
docker compose up -d
