#!/bin/bash

set -euo pipefail

# TODO: For local dev usage can we have a path that skips login if they are already auth'd?
docker login --username "${REGISTRY_USERNAME}" --password "${REGISTRY_PASSWORD}" "${REGISTRY_SERVER}"

cd /tanzu-framework

make -C ./hack/tools ytt kbld imgpkg
make package-bundle
OCI_REGISTRY="${REGISTRY_SERVER}" make push-package-bundle
