# Event manager library for GO

[![Build Status](https://travis-ci.org/goglue/event-manager.svg?branch=master)](https://travis-ci.org/goglue/event-manager)

This library implements an event manager for common usages

### How to use

#### Interfaces
- The subscriber interface

```go
type Subscriber interface {
    Update(interface{})
}

// Example for a subscriber
type Subscriber int

func (s Subscriber) Update(eventState interface{}) {
    // Do whatever with the event-state you've received
}
```

- Recorder: After dispatching an event a single snapshot will be taken for th

```go
type Recorder interface{
    Snapshot(eventName string, eventState interface{}, on time.Time)
}
```

#### Example

```go
memory := NewMemoryStorage()
dispatcher := NewDispatcher()

var subscriber Subscriber = 1

eventManager := NewEventManager(memory, dispatcher, nil)
eventManager.Attach("randomEventName", subscriber)
eventManager.Dispatch("randomEventName", "any value can be dispatched")
```

There is also another method to dispatch the event for the subscribers in go routines:

```go
eventManager.GoDispatch("randomEventName", "any value can be dispatched")
```