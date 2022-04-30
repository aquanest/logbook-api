# Build and upload new image into Docker Registry on local
function build_on_local() {
  local APP_NAME=$1
  local LOCAL_BUILD_DIR=$2
  local CONTAINER_IMAGE_TAG=latest

  # Build new image
  docker build -t "${APP_NAME}:${CONTAINER_IMAGE_TAG}" "${LOCAL_BUILD_DIR}"
}
