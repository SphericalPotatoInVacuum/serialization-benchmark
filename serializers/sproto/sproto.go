package sproto

import (
	"errors"

	"github.com/SphericalPotatoInVacuum/serialization-benchmark/serializers"
	"google.golang.org/protobuf/proto"
)

type ProtoSerializer struct{ serializers.Serializer }

func NewSerializer() *ProtoSerializer {
	return &ProtoSerializer{}
}

func (s *ProtoSerializer) serializeProto(input C) ([]byte, error) {
	return proto.Marshal(&input)
}

func (s *ProtoSerializer) deserializeProto(input []byte, output *C) error {
	return proto.Unmarshal(input, output)
}

func (s *ProtoSerializer) Serialize(input interface{}) ([]byte, error) {
	obj, ok := input.(C)
	if ok {
		return s.serializeProto(obj)
	} else {
		return nil, errors.New("this is a placeholder function")
	}
}

func (s *ProtoSerializer) Deserialize(input []byte, output interface{}) error {
	obj, ok := output.(*C)
	if ok {
		return s.deserializeProto(input, obj)
	} else {
		return errors.New("this is a placeholder function")
	}
}
