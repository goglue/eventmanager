# Event manager library for GOLANG

[![Build Status](https://travis-ci.org/goglue/event-manager.svg?branch=master)](https://travis-ci.org/goglue/event-manager)

This library implements an event manager for common usages

### How to use

The subscriber should implement the following interface

```go
type Subscriber int

func (s Subscriber) Update(eventState interface{}) {
    println("I have recieved the event:")
    println(fmt.Sprintf("#%v", eventState))
}
```

Attaching a subscriber

```go
memory := NewMemoryStorage()
dispatcher := NewDispatcher()
var sub Subscriber = 1

eventManager := NewEventManager(memory, dispatcher, nil)

eventManager.Attach("randomEventName", sub)

eventManager.Dispatch("randomEventName", "any value can be dispatched")
```

There is also another method to dispatch the event for the subscribers in go routines:

```go
eventManager.GoDispatch("randomEventName", "any value can be dispatched")
```