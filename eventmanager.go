package eventmanager

import "time"

// Event manager structure
type EventManager struct {
	Storage    Storage
	Dispatcher DispatcherContract
	Recorder   Recorder
}

// Attaches a subscriber to an event
func (e *EventManager) Attach(eventName string, sub Subscriber) {
	e.Storage.Attach(eventName, sub)
}

// Dispatches the event across all the subscribers
func (e *EventManager) Dispatch(eventName string, eventState interface{}) {
	subscribers := e.Storage.Subscribers(eventName)
	e.Dispatcher.Dispatch(eventName, eventState, subscribers)
	record(e.Recorder, eventName, eventState)
}

// Dispatches the event across all the subscribers
func (e *EventManager) GoDispatch(eventName string, eventState interface{}) {
	subscribers := e.Storage.Subscribers(eventName)
	e.Dispatcher.GoDispatch(eventName, eventState, subscribers)
	record(e.Recorder, eventName, eventState)
}

// De attaches a subscriber from an event
func (e *EventManager) DeAttach(eventName string, subscriber Subscriber) {
	e.Storage.DeAttach(eventName, subscriber)
}

// Record the current event
func record(r Recorder, eventName string, eventState interface{}) {
	if nil != r {
		r.SnapShot(eventName, eventState, time.Now())
	}
}

// Factory method for the event manager
func NewEventManager(storage Storage, dispatcher DispatcherContract, recorder Recorder) *EventManager {
	return &EventManager{
		storage,
		dispatcher,
		recorder,
	}
}
