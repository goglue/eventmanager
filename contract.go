package eventmanager

import (
	"time"
)

// Dispatcher interface
type DispatcherContract interface {
	Dispatch(eventName string, eventState []byte, subscribers []Subscriber)
	GoDispatch(eventName string, eventState []byte, subscribers []Subscriber)
}

// The storage interface
type Storage interface {
	Subscribers(string) []Subscriber
	Attach(string, Subscriber)
	DeAttach(string, Subscriber)
}

// The recorder interface
type Recorder interface {
	SnapShot(eventName string, eventPayload []byte, on time.Time)
}

// The subscriber interface
type Subscriber interface {
	Update([]byte)
}
