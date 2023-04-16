package pubsub

type Node interface {
	Publisher
	Subscriber
}
