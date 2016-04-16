package eventmanager

import (
	"github.com/op/go-logging"
	"time"
)

type LoggerRecorder struct {
	Log *logging.Logger
}

func (r *LoggerRecorder) SnapShot(eventName string, eventState interface{}, on time.Time) {
	r.Log.Infof("EVENT[%s][%s]#%v", eventName, on.Format(time.RFC3339Nano), eventState)
}
