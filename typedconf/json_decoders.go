package typedconf

import (
	"encoding/json"
	"fmt"
)

func NewJSONDecoders() Decoders {
	return newDecoders(
		func(tObjs typedObjects) Instance {
			return &jsonInstance{
				decoders: tObjs,
			}
		},
	)
}

type jsonInstance struct {
	instance
	decoders typedObjects
}

func (inst *jsonInstance) UnmarshalJSON(data []byte) error {
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
