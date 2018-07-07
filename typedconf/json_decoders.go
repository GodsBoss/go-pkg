package typedconf

import (
	"encoding/json"
)

func NewJSONDecoders() Decoders {
	return newDecoders(
		func(tObjs unmarshaler) Instance {
			return &jsonInstance{
				decoders: tObjs,
			}
		},
	)
}

type jsonInstance struct {
	instance
	decoders unmarshaler
}

func (inst *jsonInstance) UnmarshalJSON(data []byte) error {
	obj, err := inst.decoders.unmarshal(
		func() (string, error) {
			detect := &jsonTypeDetect{}
			err := json.Unmarshal(data, detect)
			if err != nil {
				return "", err
			}
			return detect.Type, nil
		},
		func(obj interface{}) error {
			return json.Unmarshal(data, obj)
		},
	)
	if err != nil {
		return err
	}
	inst.value = obj
	return nil
}

type jsonTypeDetect struct {
	Type string `json:"type"`
}
