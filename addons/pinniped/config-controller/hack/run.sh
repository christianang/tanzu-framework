#!/bin/bash

# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

MANIFESTS_DIR="./manifests"
RBAC_FILE="${MANIFESTS_DIR}/role.yaml"
DEPLOYMENT_FILE="${MANIFESTS_DIR}/deployment.yaml"
MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Always run from config-controller directory for reproducibility
cd "${MY_DIR}/.."

tag="dev"
# tag="$(uuidgen)" # Uncomment to create random image every time
image="harbor-repo.vmware.com/tkgiam/$(whoami)/pinniped-config-controller-manager:$tag"

./hack/generate.sh
kubectl apply -f "${RBAC_FILE}"

docker build -t "$image" .
docker push "$image"
ytt --data-value "image=$image" -f "${DEPLOYMENT_FILE}" | kbld -f - | kapp deploy -a pinniped-config-controller-manager -f - -y
kubectl logs -n pinniped deploy/pinniped-config-controller-manager -f
