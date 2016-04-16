package eventmanager

import (
	"time"
)

type DispatcherContract interface {
	Dispatch(interface{}, Subscriber)
}

type Storage interface {
	Subscribers(string) *[]Subscriber
	Attach(string, Subscriber)
	DeAttach(string, Subscriber)
}

type Recorder interface {
	SnapShot(eventName string, eventPayload interface{}, on time.Time)
}

type Subscriber interface {
	Update(interface{})
}
