package typedconf

import (
	"encoding/xml"
	"fmt"
)

func NewXMLDecoders() Decoders {
	return newDecoders(
		func(tObjs typedObjects) Instance {
			return &xmlInstance{
				decoders: tObjs,
			}
		},
	)
}

type xmlInstance struct {
	decoders typedObjects
	value    interface{}
}

func (inst xmlInstance) Value() interface{} {
	return inst.value
}

func (inst *xmlInstance) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	typeKey := ""
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			typeKey = attr.Value
			break
		}
	}
	if typeKey == "" {
		return fmt.Errorf("no type found")
	}
	obj, ok := inst.decoders.create(typeKey)
	if !ok {
		return fmt.Errorf("unknown type %s", typeKey)
	}
	err := decoder.DecodeElement(obj, &start)
	if err != nil {
		return err
	}
	decoder.Skip()
	inst.value = obj
	return nil
}
