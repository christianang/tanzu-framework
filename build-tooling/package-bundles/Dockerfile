FROM ubuntu:20.04

ENV GOVERSION 1.17.7
ENV PATH="${PATH}:/usr/local/go/bin"

RUN apt-get update && \
    apt-get install -y make git wget curl tar zip && \
    rm -rf /var/lib/apt/lists/*

RUN wget -q https://go.dev/dl/go${GOVERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz

RUN curl -sSL https://get.docker.com/ | sh

RUN wget -O- https://carvel.dev/install.sh | bash

COPY build-tooling/package-bundles/build-and-publish-package-bundle.sh /build-and-publish-package-bundle.sh

ENTRYPOINT ["/build-and-publish-package-bundle.sh"]
