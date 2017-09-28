package unique

// Key represents a unique key. It cannot be implemented by other packages.
type Key interface {
	// Equals checks wether this key equals another key.
	Equals(Key) bool

	// internal makes it impossible to implement Key from outside of this package.
	internal()
}

// NewKey creates a new, unique key.
func NewKey() Key {
	n := 0
	return &key{
		p: &n,
	}
}

type key struct {
	p *int
}

func (k *key) internal() {}

func (k *key) Equals(other Key) bool {
	return k == other
}
