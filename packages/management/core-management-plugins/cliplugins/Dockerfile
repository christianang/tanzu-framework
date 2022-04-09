# Copyright 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

FROM ubuntu:20.04 as builder

ARG DEBIAN_FRONTEND=noninteractive

ENV GOVERSION 1.17.8
ENV PATH="${PATH}:/usr/local/go/bin"

RUN apt-get update && \
    apt-get install -y make git wget curl tar zip gnupg lsb-release ca-certificates

RUN wget -q https://go.dev/dl/go${GOVERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz

COPY . /tanzu-framework

ARG OS
ARG ARCH
ARG PLUGIN_NAME

RUN cd ./tanzu-framework && make build-cli-local-$OS-$ARCH-$PLUGIN_NAME

FROM scratch
ARG OS
ARG ARCH
ARG PLUGIN_NAME

COPY --from=builder /tanzu-framework/artifacts/$OS/$ARCH/cli/$PLUGIN_NAME/**/tanzu-* .
