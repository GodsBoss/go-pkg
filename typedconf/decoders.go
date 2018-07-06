package typedconf

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func NewDecoders() Decoders {
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

type Instance interface {
	Value() interface{}
}

type instance struct {
	decoders *decoders
	value    interface{}
}

func (inst *instance) UnmarshalJSON(data []byte) error {
	detect := &jsonTypeDetect{}
	err := json.Unmarshal(data, detect)
	if err != nil {
		return err
	}
	create, ok := inst.decoders.decoders[detect.Type]
	if !ok {
		return fmt.Errorf("unknown type %s", detect.Type)
	}
	obj := create()
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
	create, ok := inst.decoders.decoders[typeKey]
	if !ok {
		return fmt.Errorf("unknown type %s", typeKey)
	}
	obj := create()
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
