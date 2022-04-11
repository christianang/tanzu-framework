# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

PACKAGE_NAME := tkr-vsphere-machine-webhook
PACKAGE_VERSION := dev
IMG_DEFAULT_TAG := dev
IMG_VERSION_OVERRIDE ?= $(IMG_DEFAULT_TAG)

ifeq ($(strip $(PACKAGE_SUB_VERSION)),)
	PACKAGE_TAG ?= $(PACKAGE_VERSION)
else
	PACKAGE_TAG ?= $(PACKAGE_VERSION)_$(PACKAGE_SUB_VERSION)
endif

YTT := ../../../hack/tools/bin/ytt
KBLD := ../../../hack/tools/bin/kbld
IMGPKG := ../../../hack/tools/bin/imgpkg

configure-package: ## Configure package before creating the package

reset-package: ## Reset configured package

.PHONY: build-and-publish
build-and-publish: configure-package build-and-publish-images reset-package
	$(IMGPKG) push -b $(OCI_REGISTRY)/$(PACKAGE_NAME):$(PACKAGE_TAG) -f bundle

bundle/.imgpkg:
	mkdir -p ./bundle/.imgpkg/

.PHONY: clean-imagelock
clean-imagelock:
	rm -f ./bundle/.imgpkg/images.yml

.PHONY: build-and-publish-images
build-and-publish-images: bundle/.imgpkg clean-imagelock
	$(YTT) -f kbld-config-template.yaml \
		-v registry="$(OCI_REGISTRY)" \
		-v image_tag="$(IMG_VERSION_OVERRIDE)" | \
	$(KBLD) --imgpkg-lock-output ./bundle/.imgpkg/images.yml -f bundle -f- > /dev/null
