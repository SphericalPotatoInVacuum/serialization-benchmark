FROM golang:1.17

RUN apt-get update
RUN apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

WORKDIR /build
COPY . ./
RUN go get -u -v -f all
RUN protoc -I=serializers/sproto --go_out=serializers/sproto schema.proto

ENTRYPOINT go test && go test -bench=. -cpu=1 -benchtime=1000x
