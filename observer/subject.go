package observer

type Subject interface {
	AddObserver(o Observer)
	RemoveObserver(id string)
	Notify(msg string)
}

type ConcreteSubject struct {
	observers map[string]Observer
}

func NewConcreteSubject() ConcreteSubject {
	return ConcreteSubject{
		observers: make(map[string]Observer),
	}
}

func (s ConcreteSubject) AddObserver(o Observer) {
	s.observers[o.Id()] = o
}

func (s ConcreteSubject) RemoveObserver(id string) {
	delete(s.observers, id)
}

func (s ConcreteSubject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}
