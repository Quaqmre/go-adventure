package main

import (
	"fmt"
	"time"
)

//UserConnected is a event and it just have eventId
type UserConnected struct {
	id int
}

//UserConnectedListener is much listenerss and  must implament HandleUserConnected metot
type UserConnectedListener interface {
	HandleUserConnected(*UserConnected)
}

//UserConnectHandler is a handler state it have event Id and ListenerList
type UserConnectHandler struct {
	event     *UserConnected
	listeners []UserConnectedListener
}

//handler is a handler listeners implemented function and it will executed by dispatcher
func (handler *UserConnectHandler) handle() {
	for _, listener := range handler.listeners {
		listener.HandleUserConnected(handler.event)
	}
}

type eventHandler interface {
	handle()
}

//Dispacher is a top level struct and pass all event to all listerners
type Dispacher struct {
	EventQueue             chan eventHandler
	userConnectedListeners []UserConnectedListener
}

//RunEventLoop is a infinite loop for get all event on one way
func (dispatcher *Dispacher) RunEventLoop() {
	for {
		select {
		case handler := <-dispatcher.EventQueue:
			handler.handle()
		}
	}
}

//Sender orcastrate hole events
type Sender struct {
	name string
}

//HandleUserConnected handled sender struct and doin job unique
func (sender *Sender) HandleUserConnected(event *UserConnected) {
	fmt.Println(event.id, "is connected succesfully from ", sender.name)
}

type SenderDiff struct {
	name string
}

func (senderdiff *SenderDiff) HandleUserConnected(event *UserConnected) {
	fmt.Println("omw I cant get any more from", senderdiff.name)
}

func main() {
	dispach := &Dispacher{make(chan eventHandler, 100), []UserConnectedListener{}}
	sender1 := &Sender{"sender1"}
	sender2 := &SenderDiff{"senderdiff"}
	dispach.userConnectedListeners = append(dispach.userConnectedListeners, sender1)
	dispach.userConnectedListeners = append(dispach.userConnectedListeners, sender2)
	go dispach.RunEventLoop()

	userCHandler := &UserConnectHandler{
		event:     &UserConnected{12},
		listeners: dispach.userConnectedListeners,
	}

	dispach.EventQueue <- userCHandler
	time.Sleep(3 * time.Second)
}
