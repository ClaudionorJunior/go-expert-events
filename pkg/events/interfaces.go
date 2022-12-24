package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	HandleEvent(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	RegisterHandler(eventName string, handler EventHandlerInterface) error
	DispatchEvent(event EventInterface) error
	RemoveHandler(eventName string, handler EventHandlerInterface) error
	HasHandler(eventName string, handler EventHandlerInterface) bool
	ClearHandler(eventName string) error
}
