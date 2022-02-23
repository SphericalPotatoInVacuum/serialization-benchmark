package sgob

import (
	"bytes"
	"encoding/gob"
)

type GobSerializer struct {
	buffer  *bytes.Buffer
	encoder gob.Encoder
	decoder gob.Decoder
}

func NewSerializer() *GobSerializer {
	var buffer bytes.Buffer
	return &GobSerializer{buffer: &buffer, encoder: *gob.NewEncoder(&buffer), decoder: *gob.NewDecoder(&buffer)}
}

func (s *GobSerializer) Serialize(input interface{}) ([]byte, error) {
	s.buffer.Reset()
	err := s.encoder.Encode(input)
	return s.buffer.Bytes(), err
}

func (s *GobSerializer) Deserialize(input []byte, output interface{}) error {
	s.buffer.Reset()
	s.buffer.Write(input)
	err := s.decoder.Decode(output)
	return err
}
