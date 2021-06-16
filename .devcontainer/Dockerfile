ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-buster
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

# Setup user
# RUN adduser $USERNAME -s /bin/sh -D -u $USER_UID $USER_GID
# RUN mkdir -p /etc/sudoers.d
# RUN echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME
# RUN chmod 0440 /etc/sudoers.d/$USERNAME

# Install packages and Go language server
# RUN apk add -q --update --progress --no-cache git sudo openssh-client bash unzip


RUN wget -qO- https://deb.nodesource.com/setup_16.x | bash

# install protoc, node
RUN apt-get update && apt-get install -y protobuf-compiler nodejs 

# install protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# install protoc-gen-grpc-web
ADD https://github.com/grpc/grpc-web/releases/download/1.2.1/protoc-gen-grpc-web-1.2.1-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web
RUN chmod +x /usr/local/bin/protoc-gen-grpc-web

# install node
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash


# Setup shell
# USER vscode
