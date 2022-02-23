FROM golang:1.17
WORKDIR /build
COPY . ./
RUN go get -u -v -f all

ENTRYPOINT go test && go test -bench=. -cpu=1 -benchtime=1000x
