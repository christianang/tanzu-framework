#!/bin/bash

set -euo pipefail

docker login --username "${REGISTRY_USERNAME}" --password "${REGISTRY_PASSWORD}" "${REGISTRY_SERVER}"

make -C /tanzu-framework/hack/tools ytt kbld imgpkg

cd /tanzu-framework/${PACKAGE_PATH}

OCI_REGISTRY="$OCI_REGISTRY" make build-and-push # all packages are expected to have a build-and-push make target
