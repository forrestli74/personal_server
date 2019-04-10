FROM golang:latest
ARG repo_slug=lijiaqigreat/personal_server
ENV BUILD_DIR=$GOPATH/src/github.com/$repo_slug
RUN mkdir -p $BUILD_DIR
RUN ln -s $BUILD_DIR /build
WORKDIR $BUILD_DIR

ADD install.sh .
RUN bash install.sh
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/twitchtv/twirp/protoc-gen-twirp

ADD . .
RUN make proto
RUN go get ./...
RUN go build -o main .
CMD ["go", "test"]

# FROM alpine
# RUN adduser -S -D -H -h /app appuser
# USER appuser
# COPY --from=builder /build/main /app/
# COPY --from=builder /build/server.* /app/
# COPY --from=builder /build/protobuf/command.pb.go /app/
# WORKDIR /app
# CMD [""]
