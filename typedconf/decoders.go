package typedconf

import "fmt"

func newDecoders(createInstance createInstanceFunc) Decoders {
	return &decoders{
		decoders:       make(map[string]func() interface{}),
		createInstance: createInstance,
	}
}

type createInstanceFunc func(tObjs unmarshaler) Instance

type Decoders interface {
	Register(key string, create func() interface{})
	Instance() Instance
}

type decoders struct {
	decoders       map[string]func() interface{}
	createInstance createInstanceFunc
}

func (dec *decoders) Register(key string, create func() interface{}) {
	dec.decoders[key] = create
}

func (dec *decoders) Instance() Instance {
	return dec.createInstance(dec)
}

func (dec *decoders) create(objectType string) (interface{}, bool) {
	if cr, ok := dec.decoders[objectType]; ok {
		return cr(), true
	}
	return nil, false
}

func (dec *decoders) unmarshal(detect func() (string, error), concreteUnmarshal func(obj interface{}) error, set func(interface{})) error {
	desiredType, err := detect()
	if err != nil {
		return err
	}
	dest, ok := dec.create(desiredType)
	if !ok {
		return fmt.Errorf("unknown type %s", desiredType)
	}
	err = concreteUnmarshal(dest)
	if err != nil {
		return err
	}
	set(dest)
	return nil
}

type unmarshaler interface {
	unmarshal(detect func() (string, error), concreteUnmarshal func(obj interface{}) error, set func(interface{})) error
}

type typedObjects interface {
	create(objectType string) (interface{}, bool)
}

type Instance interface {
	Value() interface{}
}

type instance struct {
	value interface{}
}

func (inst instance) Value() interface{} {
	return inst.value
}

func (inst *instance) setValue(value interface{}) {
	inst.value = value
}
