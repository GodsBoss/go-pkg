package unique_test

import (
	"github.com/GodsBoss/go-pkg/unique"

	"testing"
)

func TestUniqueKeyEqualsItself(t *testing.T) {
	key := unique.NewKey()

	if !key.Equals(key) {
		t.Errorf("Expected key to equal itself")
	}
}

func TestUniqueKeyDoesNotEqualOtherKey(t *testing.T) {
	aKey := unique.NewKey()
	otherKey := unique.NewKey()

	if aKey.Equals(otherKey) {
		t.Errorf("Expected key not to equal a different key")
	}
}
