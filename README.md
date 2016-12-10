# Event manager library for GO

[![Build Status](https://travis-ci.org/goglue/eventmanager.svg?branch=master)](https://travis-ci.org/goglue/eventmanager)

This library implements an event manager for common usages

### How to use

#### Interfaces
- The subscriber interface

```go
type Subscriber interface {
    Update([]byte)
}

// Example for a subscriber
type Subscriber int

func (s Subscriber) Update(eventState []byte) {
    // Do whatever with the event-state you've received
}
```

- Recorder: After dispatching an event a single snapshot will be taken for th

```go
type Recorder interface{
    Snapshot(eventName string, eventState []byte, on time.Time)
}
```

#### Example

```go
import "github.com/goglue/eventmanager"

func main(){
        memory := eventmanager.NewMemoryStorage()
        dispatcher := eventmanager.NewDispatcher()
        
        var subscriber Subscriber = 1
        
        evManager := eventmanager.NewEventManager(memory, dispatcher, nil)
        evManager.Attach("randomEventName", subscriber)
        evManager.Dispatch("randomEventName", "any value can be dispatched")
}
```

There is also another method to dispatch the event for the subscribers in go routines:

```go
evManager.GoDispatch("randomEventName", "any value can be dispatched")
```
