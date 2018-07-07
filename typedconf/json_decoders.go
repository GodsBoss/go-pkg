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
	return inst.decoders.unmarshal(
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
		inst.setValue,
	)
}

type jsonTypeDetect struct {
	Type string `json:"type"`
}
