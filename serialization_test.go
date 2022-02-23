package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/SphericalPotatoInVacuum/serialization-benchmark/data"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/sgob"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/sjson"
	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers/sxml"
)

func typeof(v interface{}) string {
	str := fmt.Sprintf("%v", reflect.TypeOf(v))
	dotIdx := strings.LastIndex(str, ".")
	return str[dotIdx+1:]
}

func TestSerializers(t *testing.T) {
	serializers := []serializers.Serializer{
		sjson.NewSerializer(),
		sgob.NewSerializer(),
		sxml.NewSerializer(),
	}

	for _, serializer := range serializers {
		t.Run(typeof(serializer), func(t *testing.T) {
			byte_data, err := serializer.Serialize(data.SampleC)
			if err != nil {
				t.Fatal(err)
			}
			var OtherC data.C
			err = serializer.Deserialize(byte_data, &OtherC)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(OtherC, data.SampleC) {
				t.Fatalf("Value changed after serialization cycle")
			}
		})
	}
}

func BenchmarkSerializers(b *testing.B) {
	serializers := []serializers.Serializer{
		sjson.NewSerializer(),
		sgob.NewSerializer(),
		sxml.NewSerializer(),
	}
	b.ResetTimer()
	for _, serializer := range serializers {
		b.Run(typeof(serializer), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				data, err := serializer.Serialize(data.SampleC)
				if err != nil {
					b.Fatal(err)
				}
				b.ReportMetric(float64(len(data)), "data_size")
			}
		})
	}
}
