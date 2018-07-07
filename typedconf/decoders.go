package typedconf

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func newDecoders() Decoders {
	return &decoders{
		decoders: make(map[string]func() interface{}),
	}
}

type Decoders interface {
	Register(key string, create func() interface{})
	Instance() Instance
}

type decoders struct {
	decoders map[string]func() interface{}
}

func (dec *decoders) Register(key string, create func() interface{}) {
	dec.decoders[key] = create
}

func (dec *decoders) Instance() Instance {
	return &instance{
		decoders: dec,
	}
}

func (dec *decoders) create(objectType string) (interface{}, bool) {
	if cr, ok := dec.decoders[objectType]; ok {
		return cr(), true
	}
	return nil, false
}

type typedObjects interface {
	create(objectType string) (interface{}, bool)
}

type Instance interface {
	Value() interface{}
}

type instance struct {
	decoders typedObjects
	value    interface{}
}

func (inst *instance) UnmarshalJSON(data []byte) error {
	detect := &jsonTypeDetect{}
	err := json.Unmarshal(data, detect)
	if err != nil {
		return err
	}
	obj, ok := inst.decoders.create(detect.Type)
	if !ok {
		return fmt.Errorf("unknown type %s", detect.Type)
	}
	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	inst.value = obj
	return nil
}

type jsonTypeDetect struct {
	Type string `json:"type"`
}

func (inst *instance) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
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

func (inst instance) Value() interface{} {
	return inst.value
}
