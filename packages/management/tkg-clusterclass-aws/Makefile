# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

PACKAGE_NAME := tkg-clusterclass-aws
PACKAGE_VERSION := dev
IMG_DEFAULT_TAG := dev
IMG_VERSION_OVERRIDE ?= $(IMG_DEFAULT_TAG)

ifeq ($(strip $(PACKAGE_SUB_VERSION)),)
	PACKAGE_TAG ?= $(PACAKGE_VERSION)
else
	PACKAGE_TAG ?= $(PACKAGE_VERSION)_$(PACKAGE_SUB_VERSION)
endif

KBLD := ../../../hack/tools/bin/kbld

configure-package: ## Configure package before creating the package

reset-package: ## Reset configured package

.PHONY: build-and-publish
build-and-publish: configure-package build-and-publish-images reset-package
	imgpkg push -b $(OCI_REGISTRY)/$(PACKAGE_NAME):$(PACKAGE_TAG) -f bundle

bundle/.imgpkg:
	mkdir -p ./bundle/.imgpkg/

.PHONY: clean-imagelock
clean-imagelock:
	rm -f ./bundle/.imgpkg/images.yml

.PHONY: build-and-publish-images
build-and-publish-images: bundle/.imgpkg clean-imagelock
	$(KBLD) --imgpkg-lock-output ./bundle/.imgpkg/images.yml -f bundle > /dev/null
