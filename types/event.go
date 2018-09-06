package types

type EventType uint8

const (
	EventBlockCommitted EventType = iota
	EventBlockCommitFailed
	EventBlockVerifyFailed
)

type EventFunc func(v interface{})

type Subscriber chan interface{}

type EventCenter interface {

	// subscriber subscribe specified eventType with eventFunc
	Subscribe(eventType EventType, eventFunc EventFunc) Subscriber

	// subscriber unsubscribe specified eventType
	UnSubscribe(eventType EventType, subscriber Subscriber) (err error)

	// notify subscriber of eventType
	Notify(eventType EventType, value interface{}) (err error)

	// notify specified eventFunc
	NotifySubscriber(eventFunc EventFunc, value interface{})

	// notify subscriber traversing all events
	NotifyAll() (errs []error)

	// unsubscrible all event
	UnSubscribeAll()
}

var GlobalEventCenter EventCenter
