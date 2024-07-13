#!/bin/bash

# Exit on error.
set -e

# MySQL Stop.
docker stop postgresql
docker rm postgresql

# Docker Image Update.
docker pull postgres:16.2-alpine

# MySQL Start.
docker compose up -d
