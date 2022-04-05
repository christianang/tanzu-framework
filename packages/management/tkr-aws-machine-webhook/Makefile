# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

configure-package: ## Configure package before creating the package

reset-package: ## Reset configured package

build-and-push: export PACKAGE_NAME = tkr-aws-machine-webhook
build-and-push: export PACKAGE_BUNDLES = $(PACKAGE_NAME)

ifeq ($(strip $(PACKAGE_SUB_VERSION)),)
	PACKAGE_TAG ?= $(PACAKGE_VERSION)
else
	PACKAGE_TAG ?= $(PACKAGE_VERSION)_$(PACKAGE_SUB_VERSION)
endif

.PHONY: build-and-push
build-and-push: configure-package build-and-push-image reset-package
	imgpkg push -b ${OCI_REGISTRY}/${PACKAGE_NAME}:${PACKAGE_TAG} -f bundle

.PHONY: build-and-push-image
build-and-push-image: configure-package kbld-build-push-image reset-package

.PHONY: kbld-build-push-image
kbld-build-push-image:
	mkdir -p ./bundle/.imgpkg && \
	../../../hack/tools/bin/ytt -f kbld-config.yaml -f kbld-config-overlay.yaml -v registry="${OCI_REGISTRY}" --data-value-yaml tags="${IMAGE_TAGS}" \
		| ../../../hack/tools/bin/kbld --imgpkg-lock-output ./bundle/.imgpkg/images.yml -f bundle -f-
