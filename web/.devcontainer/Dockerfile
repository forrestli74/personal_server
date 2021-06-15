ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-buster
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Setup user
# RUN adduser $USERNAME -s /bin/sh -D -u $USER_UID $USER_GID
# RUN mkdir -p /etc/sudoers.d
# RUN echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME
# RUN chmod 0440 /etc/sudoers.d/$USERNAME

# Install packages and Go language server
# RUN apk add -q --update --progress --no-cache git sudo openssh-client bash unzip
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

RUN apt-get update && apt-get install -y unzip


ARG PROTOC_VERSION=3.7.1
ARG PROTOC_ZIP=protoc-$PROTOC_VERSION-linux-x86_64.zip
ADD https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/$PROTOC_ZIP $PROTOC_ZIP
RUN unzip -o $PROTOC_ZIP -d / bin/protoc
RUN rm $PROTOC_ZIP



# Setup shell
USER root
