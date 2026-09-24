#!/bin/bash

# Exit on error.
set -e

# MySQL Stop.
docker stop postgresql
docker rm postgresql

# Docker Image Update.
# PostgreSQL 18 was released on 25.09.2025.
# End-of-Life (EOL) Date: 14.11.2030.
docker pull postgres:18

# MySQL Start.
docker compose up -d
