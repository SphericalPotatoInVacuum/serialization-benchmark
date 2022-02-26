package main

import (
	"fmt"
	"os"
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

type TestCase struct {
	name       string
	serializer serializers.Serializer
}

var TestCases []TestCase

func init() {
	TestCases = []TestCase{
		{name: "Native(gob)", serializer: sgob.NewSerializer()},
		{name: "XML", serializer: sxml.NewSerializer()},
		{name: "Json", serializer: sjson.NewSerializer()},
		{name: "MsgPack", serializer: smsgpack.NewSerializer()},
		{name: "Yaml", serializer: syaml.NewSerializer()},
		{name: "Avro", serializer: savro.NewSerializer()},
	}
}

func TestSerializers(t *testing.T) {
	for _, testCase := range TestCases {
		name, serializer := testCase.name, testCase.serializer
		t.Run(name, func(t *testing.T) {
			var other_struct data.C
			byte_data, err := serializer.Serialize(data.SampleC)
			if err != nil {
				t.Fatal(err)
			}
			os.WriteFile(fmt.Sprintf("serialization_results/%s", typeof(serializer)), byte_data, 'w')
			err = serializer.Deserialize(byte_data, &other_struct)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(other_struct, data.SampleC) {
				t.Fatalf("Value changed after serialization cycle")
			}
		})
	}
	t.Run("Protobuf", func(t *testing.T) {
		serializer := sproto.NewSerializer()
		other_struct := data.SampleProtoC
		byte_data, err := serializer.Serialize(data.SampleProtoC)
		if err != nil {
			t.Fatal(err)
		}
		err = serializer.Deserialize(byte_data, &other_struct)
		if err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(&other_struct, &data.SampleProtoC) {
			t.Fatalf("Value changed after serialization cycle")
		}
	})
}

func BenchmarkSerializationTime(b *testing.B) {
	for _, testCase := range TestCases {
		name, serializer := testCase.name, testCase.serializer
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := serializer.Serialize(data.SampleC)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
	b.Run("Protobuf", func(b *testing.B) {
		serializer := sproto.NewSerializer()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := serializer.Serialize(data.SampleProtoC)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkDataSize(b *testing.B) {
	for _, testCase := range TestCases {
		name, serializer := testCase.name, testCase.serializer
		b.Run(name, func(b *testing.B) {
			var data_size uint64
			data, err := serializer.Serialize(data.SampleC)
			if err != nil {
				b.Fatal(err)
			}
			data_size += uint64(len(data))
			b.ReportMetric(0, "ns/op")
			b.ReportMetric(float64(data_size), "data_size")
		})
	}
	b.Run("Protobuf", func(b *testing.B) {
		serializer := sproto.NewSerializer()
		var data_size uint64
		data, err := serializer.Serialize(data.SampleProtoC)
		if err != nil {
			b.Fatal(err)
		}
		data_size += uint64(len(data))
		b.ReportMetric(0, "ns/op")
		b.ReportMetric(float64(data_size), "data_size")
	})
}
func BenchmarkDeserializationTime(b *testing.B) {
	for _, testCase := range TestCases {
		name, serializer := testCase.name, testCase.serializer
		other_struct := data.SampleC
		b.Run(name, func(b *testing.B) {
			data_bytes, err := serializer.Serialize(data.SampleC)
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
	b.Run("Protobuf", func(b *testing.B) {
		serializer := sproto.NewSerializer()
		other_struct := data.SampleProtoC
		data_bytes, err := serializer.Serialize(data.SampleProtoC)
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
