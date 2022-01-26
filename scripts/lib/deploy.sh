# Build a container from local image to boot on Cloud Run
function deploy_to_gcp() {
  local APP_NAME=$1
  local LOCAL_BUILD_DIR=$2
  local GCP_PROJECT=$3
  local GCR_LOCATION=asia.gcr.io
  local GCR_REGION=asia-northeast1
  local CONTAINER_IMAGE_TAG=latest

  # Build and run a container from image
  gcloud run deploy ${APP_NAME} \
    --image ${GCR_LOCATION}/${GCP_PROJECT}/${APP_NAME}:${CONTAINER_IMAGE_TAG} \
    --project ${GCP_PROJECT} \
    --region ${GCR_REGION} \
    --update-env-vars \
      LOGBOOK_ATMOS_API=$LOGBOOK_ATMOS_API,LOGBOOK_ATMOS_TOKEN=$LOGBOOK_ATMOS_TOKEN \
    --no-traffic \
    --allow-unauthenticated

  # Show more help
  show_gcloud_run_help ${APP_NAME} ${GCR_REGION}
}

# Build a container from local image to boot on local
function deploy_to_local() {
  local APP_NAME=$1
  local LOCAL_BUILD_DIR=$2
  local LOCAL_FORWARDING_PORT=80
  local APP_PORT=8080

  # Build a container from image
  DOCKER_CONTAINER_ID=$(docker container create --publish ${LOCAL_FORWARDING_PORT}:${APP_PORT} "${APP_NAME}")

  # # Boot the container
  docker container start "${DOCKER_CONTAINER_ID}"

  # Show more help
  show_docker_container_help ${DOCKER_CONTAINER_ID}
}
