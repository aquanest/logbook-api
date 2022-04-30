#!/bin/bash

# Prepare variables
ENV=$1
APP_NAME=logbook-api
LOCAL_BUILD_DIR=.

# Load libraries
source ./scripts/lib/build.sh
source ./scripts/lib/deploy.sh
source ./scripts/lib/help.sh
source ./scripts/lib/validation.sh

# Validate in advance
validate_current_dir $APP_NAME
validate_curl_exists
validate_docker_exists

# Build container and start it on Docker Engine
if build_on_local $APP_NAME $LOCAL_BUILD_DIR; then
  deploy_to_local $APP_NAME $LOCAL_BUILD_DIR
fi
