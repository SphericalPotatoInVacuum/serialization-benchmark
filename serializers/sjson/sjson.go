package sjson

import (
	"encoding/json"

	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers"
)

type JsonSerializer struct{ serializers.Serializer }

func NewSerializer() *JsonSerializer {
	return &JsonSerializer{}
}

func (s *JsonSerializer) Serialize(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

func (s *JsonSerializer) Deserialize(input []byte, output interface{}) error {
	return json.Unmarshal(input, output)
}
