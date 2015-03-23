package EventManager;

import (
	"testing"
)

var (
	eventName = "eventName"
	secondEventName = "secondEventName"
)

func TestWhenTriggeredAnEventShouldExecuteCallback(t *testing.T) {
	Init()

	Listen(eventName, func(event Event) {
		if event.Name() != eventName {
			t.Error("event name doesnt match")
		}
	})

	anEvent := Event{name: eventName}
	listener := Listener()
	listener <- anEvent
}

func TestWhenTriggeredTwoEventsShouldPassCorrectEventToCallback(t *testing.T) {
	Init()

	Listen(eventName, func(event Event) {
		if event.Name() != eventName {
			t.Error("event name doesnt match")
		}
	})

	Listen(secondEventName, func(event Event) {
		if event.Name() != secondEventName {
			t.Error("event name doesnt match")
		}
	})

	anEvent := Event{name: eventName}
	aSecondEvent := Event{name: secondEventName}
	listener := Listener()
	listener <- aSecondEvent
	listener <- anEvent
}
