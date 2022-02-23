package syaml

import (
	"gopkg.in/yaml.v2"

	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers"
)

type YamlSerializer struct{ serializers.Serializer }

func NewSerializer() *YamlSerializer {
	return &YamlSerializer{}
}

func (s *YamlSerializer) Serialize(input interface{}) ([]byte, error) {
	return yaml.Marshal(input)
}

func (s *YamlSerializer) Deserialize(input []byte, output interface{}) error {
	return yaml.Unmarshal(input, output)
}
