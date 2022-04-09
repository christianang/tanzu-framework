# Copyright 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

NUL = /dev/null
ifeq ($(GOHOSTOS),windows)
	NUL = NUL
endif
BUILD_VERSION ?= $(shell git describe --tags --abbrev=0 2>$(NUL))

ifeq ($(strip $(BUILD_VERSION)),)
BUILD_VERSION = dev
endif

PACKAGE_NAME := core-management-plugins
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

configure-package: ## Configure package before creating the package
	sed -e "s/\VERSION/${BUILD_VERSION}/" values.template.yaml > bundle/config/zz_generated_values.yaml

reset-package: ## Reset configured package
	rm bundle/config/zz_generated_values.yaml | true

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
	$(YTT) -f kbld-config-template.yaml \
		-v registry="$(OCI_REGISTRY)" \
		-v image_tag="$(IMG_VERSION_OVERRIDE)" | \
	$(KBLD) --build-concurrency 1 --imgpkg-lock-output ./bundle/.imgpkg/images.yml -f bundle -f- > /dev/null
