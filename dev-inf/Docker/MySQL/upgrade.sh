#!/bin/bash

# Exit on error.
set -e

# MySQL Stop.
docker stop mysql
docker rm mysql

# Docker Image Update.
docker pull mysql

# MySQL Start.
docker compose up -d
