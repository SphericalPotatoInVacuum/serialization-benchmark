package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/SphericalPotatoInVacuum/serialization-benchmark/data"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/savro"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/sgob"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/sjson"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/smsgpack"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/sproto"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/sxml"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/syaml"
	"google.golang.org/protobuf/proto"
)

func typeof(v interface{}) string {
	str := fmt.Sprintf("%v", reflect.TypeOf(v))
	dotIdx := strings.LastIndex(str, ".")
	return str[dotIdx+1:]
}

var Serializers []serializers.Serializer

func init() {
	Serializers = []serializers.Serializer{
		sgob.NewSerializer(),
		sxml.NewSerializer(),
		sjson.NewSerializer(),
		smsgpack.NewSerializer(),
		syaml.NewSerializer(),
		sproto.NewSerializer(),
		savro.NewSerializer(),
	}
}

func TestSerializers(t *testing.T) {
	for _, serializer := range Serializers {
		t.Run(typeof(serializer), func(t *testing.T) {
			var byte_data []byte
			var err error
			switch serializer.(type) {
			case *sproto.ProtoSerializer:
				byte_data, err = serializer.Serialize(data.SampleProtoC)
			default:
				byte_data, err = serializer.Serialize(data.SampleC)
			}
			if err != nil {
				t.Fatal(err)
			}
			switch serializer.(type) {
			case *sproto.ProtoSerializer:
				var OtherC sproto.C
				err = serializer.Deserialize(byte_data, &OtherC)
				if err != nil {
					t.Fatal(err)
				}
				if !proto.Equal(&OtherC, &data.SampleProtoC) {
					t.Fatalf("Value changed after serialization cycle")
				}
			default:
				var OtherC data.C
				err = serializer.Deserialize(byte_data, &OtherC)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(OtherC, data.SampleC) {
					t.Fatalf("Value changed after serialization cycle")
				}
			}
		})
	}
}

func BenchmarkSerializers(b *testing.B) {
	for _, serializer := range Serializers {
		var data_struct interface{}
		switch serializer.(type) {
		case *sproto.ProtoSerializer:
			data_struct = data.SampleProtoC
		default:
			data_struct = data.SampleC
		}
		b.Run(typeof(serializer), func(b *testing.B) {
			var data_size uint64
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				data, err := serializer.Serialize(data_struct)
				if err != nil {
					b.Fatal(err)
				}
				data_size += uint64(len(data))
			}
			b.ReportMetric(float64(data_size)/float64(b.N), "data_size")
		})
	}
}

func BenchmarkDeserializers(b *testing.B) {
	for _, serializer := range Serializers {
		switch serializer.(type) {
		case *sproto.ProtoSerializer:
			data_struct := data.SampleProtoC
			other_struct := sproto.C{}
			b.Run(typeof(serializer), func(b *testing.B) {
				data_bytes, err := serializer.Serialize(data_struct)
				if err != nil {
					b.Fatal(err)
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					err := serializer.Deserialize(data_bytes, &other_struct)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		default:
			data_struct := data.SampleC
			other_struct := data.C{}
			b.Run(typeof(serializer), func(b *testing.B) {
				data_bytes, err := serializer.Serialize(data_struct)
				if err != nil {
					b.Fatal(err)
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					err := serializer.Deserialize(data_bytes, &other_struct)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}
