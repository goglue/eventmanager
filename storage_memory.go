package eventmanager

// Memory storage structure
type Memory struct {
	Storage map[string][]Subscriber
}

// Attaches a subscriber for an event
func (m *Memory) Attach(eventName string, subscriber Subscriber) {
	subscribers := m.Storage[eventName]
	m.Storage[eventName] = append(subscribers, subscriber)
}

// Retrieve all subscribers of a given event
func (m *Memory) Subscribers(eventName string) *[]Subscriber {
	sr := m.Storage[eventName]
	return &sr
}

// De attaches a subscriber
func (m *Memory) DeAttach(eventName string, subscriber Subscriber) {
	var key int
	for k, v := range m.Storage[eventName] {
		if subscriber == v {
			key = k
			break
		}
	}
	m.Storage[eventName] = append(m.Storage[eventName][:key], m.Storage[eventName][1+key:]...)
}
