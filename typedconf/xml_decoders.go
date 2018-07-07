package typedconf

import (
	"encoding/xml"
	"fmt"
)

func NewXMLDecoders() Decoders {
	return newDecoders(
		func(tObjs unmarshaler) Instance {
			return &xmlInstance{
				decoders: tObjs,
			}
		},
	)
}

type xmlInstance struct {
	instance
	decoders unmarshaler
}

func (inst *xmlInstance) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	return inst.decoders.unmarshal(
		func() (string, error) {
			typeKey := ""
			for _, attr := range start.Attr {
				if attr.Name.Local == "type" {
					typeKey = attr.Value
					break
				}
			}
			if typeKey == "" {
				return "", fmt.Errorf("no type found")
			}
			return typeKey, nil
		},
		func(obj interface{}) error {
			return decoder.DecodeElement(obj, &start)
		},
		inst.setValue,
	)
}
