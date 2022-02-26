FROM golang:1.17

RUN apt-get update
RUN apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install golang.org/x/perf/cmd/benchstat@latest

WORKDIR /go/src/github.com/SphericalPotatoInVacuum/serialization-benchmark
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY data data/
COPY serializers serializers/
COPY scripts scripts/
RUN protoc -I=serializers/sproto --go_out=serializers/sproto schema.proto
COPY serialization_test.go stat.go bench.sh ./

CMD ["/bin/bash", "-c", "./scripts/bench.sh"]
