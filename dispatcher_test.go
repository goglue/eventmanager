package eventmanager

import (
	"bytes"
	"encoding/gob"
	"testing"
)

type subscriber struct {
	t *testing.T
}

func (s *subscriber) Update(u []byte) {
	buff := bytes.NewBuffer(u)
	decoder := gob.NewDecoder(buff)

	actual := new(testEventPayload)
	decoder.Decode(actual)

	if actual.X != expected.X {
		s.t.Fail()
	}
}

type testEventPayload struct {
	X int
}

var expected *testEventPayload

// Testing the dispatch method
func TestDispatcher_Dispatch(t *testing.T) {
	d := NewDispatcher()
	subscribers := make([]Subscriber, 0)
	subscribers = append(subscribers, &subscriber{t})

	expected = &testEventPayload{10}
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	enc.Encode(expected)

	d.Dispatch("something", buff.Bytes(), subscribers)
}
