# Copyright 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

configure-package: ## Configure package before creating the package

reset-package: ## Reset configured package

PHONY: build-and-push-image
build-and-push-image:
	cd ../../../pkg/v1/sdk/capabilities && make docker-build-and-publish # build the controller image

build-and-push: export PACKAGE_REPOSITORY = standalone
build-and-push: export PACKAGE_NAME = capabilities

PHONY: build-and-push
build-and-push: build-and-push-image
	cd ../../../ && \
	make package-bundle && \
	make push-package-bundle

