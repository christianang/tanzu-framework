#!/bin/bash

set -euo pipefail

docker login --username "${REGISTRY_USERNAME}" --password "${REGISTRY_PASSWORD}" "${REGISTRY_SERVER}"

cd /tanzu-framework

make local-registry

OCI_REGISTRY="${OCI_REGISTRY}" \
  PACKAGE_NAME="${PACKAGE_REPOSITORY}" \
  PACKAGE_REPOSITORY="${PACKAGE_REPOSITORY}" \
  PACKAGE_VALUES_FILE="${PACKAGE_VALUES_FILE}" \
  PACKAGE_BUNDLES="${PACKAGE_REPOSITORY}" \
  LOCAL_REGISTRY_URL="$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' registry):5000" \
  make package-repo-bundle push-package-repo-bundle
