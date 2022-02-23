package smsgpack

import (
	"github.com/vmihailenco/msgpack/v5"

	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers"
)

type MsgpackSerializer struct{ serializers.Serializer }

func NewSerializer() *MsgpackSerializer {
	return &MsgpackSerializer{}
}

func (s *MsgpackSerializer) Serialize(input interface{}) ([]byte, error) {
	return msgpack.Marshal(input)
}

func (s *MsgpackSerializer) Deserialize(input []byte, output interface{}) error {
	return msgpack.Unmarshal(input, output)
}
