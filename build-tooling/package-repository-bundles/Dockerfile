FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive

ENV DOCKER_CE_VERSION 5:20.10.14~3-0~ubuntu-focal
ENV DOCKER_CE_CLI_VERSION 5:20.10.14~3-0~ubuntu-focal
ENV CONTAINERD_VERSION 1.5.11-1
ENV GOVERSION 1.17.7
ENV PATH="${PATH}:/usr/local/go/bin"

RUN apt-get update && \
    apt-get install -y make git wget curl tar zip gnupg lsb-release ca-certificates

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

RUN echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

RUN apt-get update && \
    apt-get install -y \
        docker-ce=${DOCKER_CE_VERSION} \
        docker-ce-cli=${DOCKER_CE_CLI_VERSION} \
        containerd.io=${CONTAINERD_VERSION} && \
    rm -rf /var/lib/apt/lists/*

RUN wget -q https://go.dev/dl/go${GOVERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz

RUN curl -sSL https://get.docker.com/ | sh

COPY build-tooling/package-repository-bundles/build-and-push-package-repository-bundle.sh /build-and-push-package-repository-bundle.sh

ENTRYPOINT ["/build-and-push-package-repository-bundle.sh"]
