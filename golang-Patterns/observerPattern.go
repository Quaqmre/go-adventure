package main

import "fmt"

type Event struct {
	id string
}

type observer struct {
	name string
}

type EventObserver interface {
	OnNotify(event Event)
}

func NewEventObserver(name string) EventObserver {
	return &observer{name}
}

func (o *observer) OnNotify(event Event) {
	t := fmt.Sprintf("%s receved %s event", o.name, event.id)
	fmt.Println(t)
}

type EventNotifier interface {
	Add(obs EventObserver)
	Remove(obs EventObserver)
	Notify(event Event)
}

type eventNotifier struct {
	observers []EventObserver
}

func NewEventNotifier() *eventNotifier {
	return &eventNotifier{}
}

func (e *eventNotifier) Add(eo EventObserver) {
	e.observers = append(e.observers, eo)
}
func (e *eventNotifier) Remove(eo EventObserver) {
	for index, item := range e.observers {
		if item == eo {
			e.observers = append(e.observers[:index], e.observers[index+1:]...)
		}
	}
}

func (e *eventNotifier) Notify(event Event) {
	for _, item := range e.observers {
		item.OnNotify(event)
	}
}

func main() {
	t := NewEventNotifier()
	t.Add(NewEventObserver("ali"))
	t.Add(NewEventObserver("akif"))
	t.Add(NewEventObserver("ayşe"))
	t.Notify(Event{"koşun"})

}
