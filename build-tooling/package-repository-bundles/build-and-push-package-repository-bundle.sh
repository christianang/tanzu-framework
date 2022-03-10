#!/bin/bash

set -euo pipefail

docker login --username "${REGISTRY_USERNAME}" --password "${REGISTRY_PASSWORD}" "${REGISTRY_SERVER}"

make -C /tanzu-framework/hack/tools yq ytt kbld imgpkg

cd /tanzu-framework

OCI_REGISTRY="${OCI_REGISTRY}" \
  PACKAGE_NAME="${PACKAGE_REPOSITORY}" \
  PACKAGE_REPOSITORY="${PACKAGE_REPOSITORY}" \
  PACKAGE_VALUES_FILE="${PACKAGE_VALUES_FILE}" \
  make package-repo-bundle push-package-repo-bundle
