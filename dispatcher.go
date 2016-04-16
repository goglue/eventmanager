package eventmanager

type Dispatcher struct {
}

func (d *Dispatcher) Dispatch(event interface{}, subscriber Subscriber) {
	subscriber.Update(event)
}
