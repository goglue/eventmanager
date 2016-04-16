package eventmanager

import (
	"bytes"
	"testing"
)

func TestEventManager(t *testing.T) {
	var buffer bytes.Buffer
	em := Instance(StorageInstance(), DispatcherInstance(), RecorderInstance(&buffer))
	em.Attach("onSomething", &SubscriberConcrete{t})
	em.Dispatch("onSomething", &EventConcrete{39})
	log := buffer.String()
	println(log)
}

type SubscriberConcrete struct {
	T *testing.T
}

func (s *SubscriberConcrete) Update(event interface{}) {
	ec, ok := event.(EventConcrete)
	if ok {
		if !check(ec) {
			s.T.Errorf("Expected 39 got something else")
		}
	} else {
		s.T.Errorf("Not OK type")
	}
}

func check(e EventConcrete) bool {
	if e.ID() != 39 {
		return false
	}

	return true
}

type EventConcrete struct {
	Id int64
}

func (e *EventConcrete) ID() int64 {
	return e.Id
}
