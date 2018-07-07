package typedconf

func newDecoders(createInstance createInstanceFunc) Decoders {
	return &decoders{
		decoders:       make(map[string]func() interface{}),
		createInstance: createInstance,
	}
}

type createInstanceFunc func(tObjs typedObjects) Instance

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
