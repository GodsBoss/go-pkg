package typedconf

import (
	"encoding/json"
	"fmt"
)

type Decode func(data []byte, dest interface{}) error

func NewDecoders(decode Decode) Decoders {
	return &decoders{
		decode:   decode,
		decoders: make(map[string]func() interface{}),
	}
}

type Decoders interface {
	Register(key string, create func() interface{})
	Instance() Instance
}

type decoders struct {
	decode   Decode
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

func (inst instance) Value() interface{} {
	return inst.value
}
