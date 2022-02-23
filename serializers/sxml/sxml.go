package sxml

import (
	"encoding/xml"
)

type XmlSerializer struct {
}

func NewSerializer() *XmlSerializer {
	return &XmlSerializer{}
}

func (s *XmlSerializer) Serialize(input interface{}) ([]byte, error) {
	return xml.Marshal(input)
}

func (s *XmlSerializer) Deserialize(input []byte, output interface{}) error {
	return xml.Unmarshal(input, output)
}
