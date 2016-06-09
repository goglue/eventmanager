package eventmanager

import "testing"

type subscriber struct {
	t *testing.T
}
func (s *subscriber) Update(u interface{}) {
	actual, ok := u.(int)
	if !ok {
		panic("something wrong")
	}

	if actual != expected {
		s.t.Fail()
	}
}

var expected int = 10

// Testing the dispatch method
func TestDispatcher_Dispatch(t *testing.T) {
	d := NewDispatcher()
	subscribers := make([]Subscriber, 0)
	subscribers = append(subscribers, &subscriber{t})

	d.Dispatch("something", expected, subscribers)
}
