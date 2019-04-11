FROM golang:1 AS builder
ARG repo_slug=lijiaqigreat/personal_server
ENV BUILD_DIR=$GOPATH/src/github.com/$repo_slug
RUN mkdir -p $BUILD_DIR
RUN ln -s $BUILD_DIR /build
WORKDIR $BUILD_DIR

ADD install.sh .
RUN sh install.sh
ADD generate_cert.sh .
RUN sh generate_cert.sh
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/twitchtv/twirp/protoc-gen-twirp
ADD . .
RUN make proto
RUN go get ./...
RUN go build -o main .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
COPY --from=builder /build/server.* /app/
WORKDIR /app
CMD ["./main"]
