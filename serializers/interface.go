package serializers

type Serializer interface {
	Serialize(input interface{}) ([]byte, error)
	Deserialize(input []byte, output interface{}) error
}
