FROM golang:1.25.6-trixie

ARG USERNAME=vscode

RUN groupadd -g 1000 ${USERNAME} \
  && useradd -u 1000 -g 1000 -m -s /bin/bash ${USERNAME}

RUN apt-get update \
  && export DEBIAN_FRONTEND=noninteractive \
  && apt-get install -y --no-install-recommends git

USER ${USERNAME}
