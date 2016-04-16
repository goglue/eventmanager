package eventmanager

import "time"

type EventManager struct {
	Storage    Storage
	Dispatcher DispatcherContract
	Recorder   Recorder
}

// Attaches a subscriber to an event
func (e *EventManager) Attach(eventName string, sub interface {
	Update(interface{})
}) {
	e.Storage.Attach(eventName, sub)
}

// Dispatches the event across all the subscribers
func (e *EventManager) Dispatch(eventName string, eventState interface{}) {
	subscribers := e.Storage.Subscribers(eventName)

	for _, subscriber := range *subscribers {
		go e.Dispatcher.Dispatch(eventState, subscriber)
	}

	e.Recorder.SnapShot(eventName, eventState, time.Now())
}

// De attaches a subscriber from an event
func (e *EventManager) DeAttach(eventName string, subscriber Subscriber) {
	e.Storage.DeAttach(eventName, subscriber)
}
