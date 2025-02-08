FROM ubuntu:focal

RUN apt update && apt install -y build-essential wget

ARG GO_VERSION=1.23.6

RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm -f go${GO_VERSION}.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

RUN mkdir -p /repo
WORKDIR /repo

ENTRYPOINT ["make"]
CMD ["build"]
