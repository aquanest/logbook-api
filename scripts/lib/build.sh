# Build and upload new image into Docker Registry on local
function build_on_local() {
  local APP_NAME=$1
  local LOCAL_BUILD_DIR=$2
  local CONTAINER_IMAGE_TAG=latest

  # Build new image
  docker build -t "${APP_NAME}:${CONTAINER_IMAGE_TAG}" "${LOCAL_BUILD_DIR}"
}

# Build and upload new image into Cloud Build Registry on GCP
function build_on_gcp() {
  local APP_NAME=$1
  local LOCAL_BUILD_DIR=$2
  local GCP_PROJECT=$3
  local GCR_LOCATION=asia.gcr.io
  local CONTAINER_IMAGE_TAG=latest

  # Build and push new image
  gcloud builds submit ${LOCAL_BUILD_DIR} \
    --tag ${GCR_LOCATION}/${GCP_PROJECT}/${APP_NAME}:${CONTAINER_IMAGE_TAG} \
    --project ${GCP_PROJECT}
}
