package eventmanager

import "testing"

type test int

func (t test) Update([]byte) {}

// Testing the attach method
func TestMemory_Attach(t *testing.T) {
	memory := NewMemoryStorage()
	var attachable test = 1

	memory.Attach("test", attachable)

	if len(memory.storage["test"]) != 1 {
		t.Fail()
	}

	if memory.storage["test"][0] != attachable {
		t.Fail()
	}
}

// Testing the de-attach method
func TestMemory_DeAttach(t *testing.T) {
	memory := NewMemoryStorage()
	var attachable test = 1

	memory.Attach("test", attachable)
	memory.DeAttach("test", attachable)
	if len(memory.storage["test"]) != 0 {
		t.Fail()
	}
}
