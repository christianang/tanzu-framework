# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

configure-package: ## Configure package before creating the package

reset-package: ## Reset configured package

build-and-push: export PACKAGE_REPOSITORY = management
build-and-push: export PACKAGE_NAME = addons-manager-v2
build-and-push: export PACKAGE_BUNDLES = $(PACKAGE_NAME)

.PHONY: build-and-push
build-and-push: build-and-push-image
	cd ../../../ && \
	make package-bundle && \
	make push-package-bundles


.PHONY: build-and-push-image
build-and-push-image:
	cd ../../../addons && OCI_REGISTRY=${OCI_REGISTRY} make docker-build-and-publish # build the controller image
