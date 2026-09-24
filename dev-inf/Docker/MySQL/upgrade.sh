#!/bin/bash

# Exit on error.
set -e

# MySQL Stop.
docker stop mysql
docker rm mysql

# Docker Image Update.
# MySQL 8.4 LTS was released on 30.04.2026.
# Support till April 2032.
docker pull mysql:8.4

# MySQL Start.
docker compose up -d
