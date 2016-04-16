package eventmanager

import (
	"github.com/op/go-logging"
)

func Instance(storage Storage, dispatcher DispatcherContract, recorder Recorder) *EventManager {
	return &EventManager{
		storage,
		dispatcher,
		recorder,
	}
}

func StorageInstance() *Memory {
	return &Memory{
		make(map[string][]Subscriber),
	}
}

func DispatcherInstance() *Dispatcher {
	return &Dispatcher{}
}

func RecorderInstance(writer *logging.Logger) *LoggerRecorder {
	return &LoggerRecorder{
		writer,
	}
}
