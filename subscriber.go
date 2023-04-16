package pubsub

type Subscriber interface {
	Listener
	Processor
	Name() string
	//	Sub(topicName TopicName) error
	Channel() chan interface{}
	GetSubscriptions() ([]TopicStr, error)
	//	GetTopic(topicName string) (TopicStr, error)
}

type SubscriberStr struct {
	ch chan interface{}
}

type Listener interface {
	Listen()
}

type Processor interface {
	Process(msg ...interface{})
}
