# Copyright 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

include ../../../common.mk

configure-package: ## Configure package before creating the package
	sed -e "s/\VERSION/${BUILD_VERSION}/" values.template.yaml > bundle/config/zz_generated_values.yaml

reset-package: ## Reset configured package
	rm bundle/config/zz_generated_values.yaml | true

build-and-push: export PACKAGE_REPOSITORY = management
build-and-push: export PACKAGE_NAME = core-management-plugins
build-and-push: export PACKAGE_BUNDLES = $(PACKAGE_NAME)

.PHONY: build-and-push
build-and-push: build-and-push-image
	cd ../../../ && \
	make package-bundle && \
	make push-package-bundles

.PHONY: build-and-push-image
build-and-push-image:
	cd ../../../ && \
	make build-cli-plugins-for-core-management-package && \
	cd cliplugins && \
	STANDALONE_PLUGINS="" ADMIN_PLUGINS="" CONTEXTAWARE_PLUGINS="cluster kubernetes-release feature" OCI_REGISTRY=${OCI_REGISTRY} make docker-build-and-publish # build the controller image
