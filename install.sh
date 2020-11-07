# install protoc
set -e
PROTOC_VERSION=3.7.1
PROTOC_ZIP=protoc-$PROTOC_VERSION-linux-x86_64.zip
curl -OL https://github.com/google/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_ZIP
unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
rm -f $PROTOC_ZIP

# TODO
# ln -s ~/git/personal_server ~/go/src/github.com/lijiaqigreat/personal_server
