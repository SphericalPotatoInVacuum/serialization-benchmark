package savro

import (
	"log"
	"os"

	"github.com/hamba/avro"

	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers"
)

type AvroSerializer struct {
	serializers.Serializer
	schema avro.Schema
}

func NewSerializer() *AvroSerializer {
	schema_data, err := os.ReadFile("serializers/savro/schema.json")
	if err != nil {
		log.Fatalf("could not read schema.json: %v", err)
	}
	schema, err := avro.Parse(string(schema_data))
	if err != nil {
		log.Fatalf("error parsing avro schema: %v\n", err)
	}
	return &AvroSerializer{
		schema: schema,
	}
}

func (s *AvroSerializer) Serialize(input interface{}) ([]byte, error) {
	return avro.Marshal(s.schema, input)
}

func (s *AvroSerializer) Deserialize(input []byte, output interface{}) error {
	return avro.Unmarshal(s.schema, input, output)
}
