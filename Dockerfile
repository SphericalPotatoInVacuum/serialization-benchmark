FROM golang:1.17

RUN apt-get update
RUN apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install golang.org/x/perf/cmd/benchstat@latest

WORKDIR /build
COPY . ./
RUN go get -u -v -f all
RUN protoc -I=serializers/sproto --go_out=serializers/sproto schema.proto

ENTRYPOINT go test -bench=. -cpu=1 -benchtime=100x -count=10 -timeout 0 | tee results.txt && benchstat results.txt
