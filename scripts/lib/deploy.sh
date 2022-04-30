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
