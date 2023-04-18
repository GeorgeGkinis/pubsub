package pubsub

import (
	log "github.com/sirupsen/logrus"
	"reflect"
)

type SubscriberIF interface {
	Listen()
	Name() string
	Sub(topicName TopicName) error
	Channel() chan interface{}
	GetSubscriptions() []*Topic
}

type HandlerFunc func(msg interface{}) (err error)

type Handlers map[reflect.Type]HandlerFunc

type Subscriptions []*Topic

type Subscriber struct {
	name          string
	ch            chan interface{}
	handlers      map[reflect.Type]HandlerFunc
	subscriptions Subscriptions
}

func NewSubscriber(name string, handlers Handlers, subscriptions ...*Topic) *Subscriber {
	return &Subscriber{name: name, handlers: handlers, subscriptions: subscriptions}
}

func (h Subscriber) Listen() {
	go func() {
		for msg := range h.ch {
			if err := h.handlers[reflect.TypeOf(msg)](msg); err != nil {
				log.Errorf("error handling message: %v of type %T on handler: %s", msg, msg, h.Name())
			}
		}
	}()
}

func (h Subscriber) AddHandlers(handlers Handlers) {
	for k, v := range handlers {
		h.handlers[k] = v
	}
	//TODO: Check if overwriting existing handler
}

func (h Subscriber) Name() string {
	return h.name
}

func (h Subscriber) Channel() chan interface{} {
	return h.ch
}

func (h Subscriber) GetSubscriptions() []*Topic {
	return h.subscriptions
}

func (h Subscriber) Sub(topic TopicName) (err error) {
	t := tm.Topic(topic)
	h.subscriptions = append(h.subscriptions, t)
	err = t.AddSub(&h)
	return err
}
