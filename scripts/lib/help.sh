# Guide docker container commands
function show_docker_container_help() {
  local DOCKER_CONTAINER_ID=$1

  echo -e "\nNext Step:"
  echo -e "\n* Stop container"
  echo -e "docker container stop ${DOCKER_CONTAINER_ID}"
  echo -e "\n* Start container"
  echo -e "docker container start ${DOCKER_CONTAINER_ID}"
  echo -e "\n* Show logs in container"
  echo -e "docker container logs ${DOCKER_CONTAINER_ID}"
  echo -e "\n* Show processes in container"
  echo -e "docker container top ${DOCKER_CONTAINER_ID}"
  echo -e "\n* Show statistics of container"
  echo -e "docker container stats ${DOCKER_CONTAINER_ID} --no-stream"
  echo -e "\n* Enter shell in container"
  echo -e "docker container exec -it ${DOCKER_CONTAINER_ID} /bin/sh"
  echo
}
