package copyguard

// Guard can be used as a guard against copying a struct after using it. Its usage is simple:
//
//  // This struct should not be copied after first use.
//  type SomeStruct struct {
//      guard copyguard.Guard
//  }
//
//  // GuardedFunction is a function which is not safe to call from a copied struct if the
//  // original struct has called it before the copy was made.
//  func (s *SomeStruct) GuardedFunction() {
//      s.guard.Check(s, "Panic message (may be empty)")
//      // Do something here.
//  }
type Guard struct {
	owner interface{}
}

// Check assigns the guard to an owner if that had not already happened. Otherwise,
// the current owner is checked against the owner stored by the guard. If they
// are different, Check panics.
func (guard *Guard) Check(currentOwner interface{}, message string) {
	// First use.
	if guard.owner == nil {
		guard.owner = currentOwner
		return
	}

	// Stored owner is not current owner. Copy after use. Panic!
	if guard.owner != currentOwner {
		if message == "" {
			message = "Copied after use!"
		}
		panic(message)
	}
}
