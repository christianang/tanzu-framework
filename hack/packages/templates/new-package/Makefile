# Copyright 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

configure-package: ## Configure package before creating the package

reset-package: ## Reset configured package

build-and-push: export PACKAGE_REPOSITORY = TEMPLATE_REPO
build-and-push: export PACKAGE_NAME = TEMPLATE_PACKAGE_NAME
build-and-push: export PACKAGE_BUNDLES = $(PACKAGE_NAME)

.PHONY: build-and-push
build-and-push:
	cd ../../../ && \
	make package-bundle && \
	make push-package-bundles
