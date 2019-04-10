FROM golang:latest as builder
RUN mkdir /build 
WORKDIR /build 
ADD generate_cert.sh /build/
RUN bash generate_cert.sh
ADD . /build/
RUN go get -d ./...
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
COPY --from=builder /build/server.cert /app/
WORKDIR /app
CMD ["cat server.cert"]
