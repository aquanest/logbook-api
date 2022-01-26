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

# Guide gcloud run commands
function show_gcloud_run_help() {
  local APP_NAME=$1
  local GCR_REGION=$2

  echo -e "\nNext Step:"
  echo -e "\n* Describe the detail about a service"
  echo -e "gcloud run services describe ${APP_NAME} --region ${GCR_REGION}"
  echo -e "\n* List available revisions for a service"
  echo -e "gcloud run revisions list --service ${APP_NAME} --region ${GCR_REGION}"
  echo -e "\n* Assign 100% of traffic to the latest revision <<CAUTION: UPDATES WILL BE APPLIED TO CLIENTS!>>"
  echo -e "gcloud run services update-traffic ${APP_NAME} --region ${GCR_REGION} --to-latest"
  echo -e "\n* Describe the detail about a revision (Note: {REVISION_NANE} can be get by 'gcloud run revisions list')"
  echo -e "gcloud run revisions describe --region ${GCR_REGION} {REVISION_NAME}"
  echo -e "\n* Rollback to a revision (Note: {REVISION_NANE} can be get by 'gcloud run revisions list')"
  echo -e "gcloud run services update-traffic ${APP_NAME} --region ${GCR_REGION} --to-revisions {REVISION_NAME}=100"
  echo
}
