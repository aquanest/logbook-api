# Check current directory
function validate_current_dir(){
  local APP_NAME=$1
  local CURRENT_DIR=$(basename "$PWD")
  if [ "${CURRENT_DIR}" != "${APP_NAME}" ]; then
    echo -e "\nInvalid working directory. Please move to '${APP_NAME}'.\n"
    exit 1
  fi
}

# Check curl exists
function validate_curl_exists(){
  if ! curl --version > /dev/null ; then
    echo -e "\nMissing curl. Please download from https://curl.se/\n"
    exit 1
  fi
}

# Check docker exists
function validate_docker_exists(){
  if ! docker --version > /dev/null ; then
    echo -e "\nMissing docker. Please download from https://www.docker.com/\n"
    exit 1
  fi
}
