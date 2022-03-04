#!/bin/bash

cd "/tanzu-framework/$COMPONENT_PATH"

docker login --username "$REGISTRY_USERNAME" --password "$REGISTRY_PASSWORD" "$REGISTRY_SERVER"
OCI_REGISTRY="$OCI_REGISTRY" make docker-build-and-publish
