# Benchmark of serialization formats using Golang

![docker pulls](https://img.shields.io/docker/pulls/sphericalpotatoinvacuum/serialization-benchmark)![build status](https://img.shields.io/github/workflow/status/sphericalpotatoinvacuum/serialization-benchmark/ci)

## Tested packages

- [encoding/gob](https://pkg.go.dev/encoding/gob)
- [encoding/xml](https://pkg.go.dev/encoding/xml)
- [encoding/json](https://pkg.go.dev/encoding/json)
- [github.com/vmihailenco/msgpack/v5](https://msgpack.uptrace.dev/)
- [yaml.v2](https://pkg.go.dev/gopkg.in/yaml.v2)
- [github.com/hamba/avro](https://pkg.go.dev/github.com/hamba/avro@v1.6.6)
- [google.golang.org/protobuf/proto](https://pkg.go.dev/google.golang.org/protobuf/proto)

## Results

<pre>
    FORMAT    | SERIALIZATION TIME | DESERIALIZATION TIME |   DATA SIZE    
--------------|--------------------|----------------------|----------------
  Native(gob) |       2.59ms,   1% |         2.93ms,   1% |  1.56MB,  14%  
          XML |      63.85ms,  25% |       378.59ms, 100% | 11.34MB, 100%  
         Json |       9.01ms,   4% |        50.51ms,  13% |  3.22MB,  28%  
      MsgPack |      10.00ms,   4% |        16.37ms,   4% |  1.69MB,  15%  
         Yaml |     256.17ms, 100% |       264.18ms,  70% |  5.39MB,  48%  
         Avro |       2.88ms,   1% |         5.05ms,   1% |  1.53MB,  14%  
     Protobuf |       2.01ms,   1% |         3.36ms,   1% |  1.53MB,  13%  
</pre>

## Running the benchmarks

```console
go get -u -t
./scripts/bench.sh > results.md
```

`results.md` now contains a table that you see above in the [results](#results)
section. For how to update that table see the next section. For your convenience
there is a [docker container](https://hub.docker.com/r/sphericalpotatoinvacuum/serialization-benchmark)
with all the needed packages that does the same thing:
```console
docker run --rm sphericalpotatoinvacuum/serialization-benchmark > results.md
```

## Updating the results table in README.md

To update the results you need to save the output of `bench.sh` script. Let's
assume we saved it to the `results.md` file. In that case to update the README.md
you would need to run:

```console
./scripts/update_readme.sh results.md
```

The first and only argument of that script is the name of the file with new results.

## Data used for benchmarking

Here is the proto schema of the data:

```protobuf
message C {
  repeated double SomeDoubleData = 1;

  message B {
    message A {
      string SomeString = 1;
      double SomeDouble = 2;
      repeated int32 SomeInt32Data = 3;
    }

    repeated A SomeAData = 1;
    string SomeString = 2;
    string OtherString = 3;
    repeated float SomeFloatData = 4;
  }

  map<string, B> SomeMap = 2;
}
```

Every `repeated` field is filled with n ∈ [50; 100) entries. `SomeMap` contains 
50 elements. We used `uuid` urn representation for strings for consistency.
