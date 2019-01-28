package types

type EventType uint8

const (
	EventBlockCommitted    EventType = iota // 0, block submit successfully
	EventBlockCommitFailed                  // 1, block submit failed
	EventBlockVerifyFailed                  // 2, block verified failed
	EventBlockExisted                       // 3. block has exist
	EventConsensusFailed                    // 4. to consensus failed
	EventBlockWritten                       // 5. block has been written
	EventBlockWriteFailed                   // 6. block write failed
	EventTxVerifySucceeded                  // 7. tx has been verified successfully
	EventTxVerifyFailed                     // 8. tx was verified failed
	EventMasterChange                       // 9. change master
	EventOnline                             // 10. node online
	EventBlockWithoutTxs                    // 11. block without any txs

	//P2P Event
	EventRemovePeer
	EventAddPeer
	EventBroadCastMsg
	EventRecvNewMsg
)

type EventFunc func(v interface{})

type Subscriber chan interface{}

type EventCenter interface {

	// subscribe specified eventType with eventFunc
	Subscribe(eventType EventType, eventFunc EventFunc) Subscriber

	// unsubscribe specified eventType and subscriber
	UnSubscribe(eventType EventType, subscriber Subscriber) (err error)

	// notify subscriber of eventType
	Notify(eventType EventType, value interface{}) (err error)

	// notify specified eventFunc
	NotifySubscriber(eventFunc EventFunc, value interface{})

	// notify subscriber traversing all events
	NotifyAll() (errs []error)

	// unsubscribe all event
	UnSubscribeAll()
}
