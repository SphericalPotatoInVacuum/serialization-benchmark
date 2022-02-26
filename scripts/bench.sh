#!/bin/sh
go test -bench=.Ser/ -cpu=1 -benchtime=100x -timeout 0 | tee serialization_results.txt 1>&2
go test -bench=.Deser/ -cpu=1 -benchtime=100x -timeout 0 | tee deserialization_results.txt 1>&2
go test -bench=.Data/ -cpu=1 -benchtime=1x -timeout 0 | tee data_results.txt 1>&2
go run stat.go
