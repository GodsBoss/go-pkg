package typedconf

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

func (inst instance) Value() interface{} {
	return inst.value
}
