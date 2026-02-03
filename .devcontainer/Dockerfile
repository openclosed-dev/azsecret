FROM golang:1.25.6-trixie

ARG USERNAME
ARG USER_UID
ARG USER_GID=$USER_UID

RUN groupadd -g ${USER_GID} ${USERNAME} \
  && useradd -u ${USER_UID} -g ${USER_GID} -m -s /bin/bash ${USERNAME}

RUN apt-get update \
  && export DEBIAN_FRONTEND=noninteractive \
  && apt-get install -y --no-install-recommends git openssh-client

RUN mkdir /workspaces && chown ${USERNAME}:${USERNAME} /workspaces

USER ${USERNAME}

RUN go install github.com/go-delve/delve/cmd/dlv@latest \
  && go install golang.org/x/tools/gopls@latest
