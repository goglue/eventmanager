package eventmanager

// Memory storage structure
type Memory struct {
	storage map[string][]Subscriber
}

// Attaches a subscriber for an event
func (m *Memory) Attach(eventName string, subscriber Subscriber) {
	subscribers := m.storage[eventName]
	m.storage[eventName] = append(subscribers, subscriber)
}

// Retrieve all subscribers of a given event
func (m *Memory) Subscribers(eventName string) []Subscriber {
	sr := m.storage[eventName]
	return sr
}

// De attaches a subscriber
func (m *Memory) DeAttach(eventName string, subscriber Subscriber) {
	var key int
	for k, v := range m.storage[eventName] {
		if subscriber == v {
			key = k
			break
		}
	}
	m.storage[eventName] = append(m.storage[eventName][:key], m.storage[eventName][1+key:]...)
}

// Factory method for the storage structure
func NewMemoryStorage() *Memory {
	m := make(map[string][]Subscriber)

	return &Memory{
		m,
	}
}
