#!/bin/bash

# Exit on error.
set -e

docker inspect mysql | grep MYSQL_VERSION
