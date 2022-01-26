#!/bin/bash

# Prepare variables
ENV=$1
APP_NAME=logbook-api
GCP_PROJECT=$GCP_PROJECT
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

# Build container and start it on GCP
if [ "${ENV}" = "gcp" ]; then
  validate_gcloud_exists
  validate_gcloud_access_token_exists
  validate_gcloud_project_configured ${GCP_PROJECT}

  if build_on_gcp $APP_NAME $LOCAL_BUILD_DIR $GCP_PROJECT; then
    deploy_to_gcp $APP_NAME $LOCAL_BUILD_DIR $GCP_PROJECT
  fi
  exit
fi

# Build container and start it on Docker Engine
if build_on_local $APP_NAME $LOCAL_BUILD_DIR; then
  deploy_to_local $APP_NAME $LOCAL_BUILD_DIR
fi
