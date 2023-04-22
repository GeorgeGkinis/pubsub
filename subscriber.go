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

type Handlers map[string]HandlerFunc

type Subscriptions []*Topic

type Subscriber struct {
	name          string
	listening     bool
	ch            chan interface{}
	handlers      Handlers
	subscriptions Subscriptions
}

func NewSubscriber(name string, handlers Handlers, subscriptions ...*Topic) *Subscriber {
	s := Subscriber{
		name:          name,
		listening:     false,
		ch:            make(chan interface{}, 0),
		handlers:      make(Handlers, 0),
		subscriptions: make(Subscriptions, 0),
	}
	for k, v := range handlers {
		s.handlers[k] = v
	}
	for _, v := range subscriptions {
		s.subscriptions = append(s.subscriptions, v)
	}

	return &s
}

func (h Subscriber) Listen() {
	if !h.listening {
		h.listening = true
		go func() {
			for msg := range h.ch {
				log.Debugf("Received message of type %T", msg)
				handler, ok := h.handlers[reflect.TypeOf(msg).Name()]
				if !ok {
					log.Errorf("Subscriber %s has no handler for message type %T", h.name, msg)
				}
				//log.Errorf("%s", handler)
				if err := handler(msg); err != nil {
					log.Errorf("error handling message: %v of type %T on handler: %s", msg, msg, h.Name())
				}
			}
		}()
	}
}

func (h Subscriber) AddHandler(typeOf interface{}, handler HandlerFunc) {

	if typeOf == nil || handler == nil {
		log.Errorf("Required: typeOf and handler. Provided: typeOf: %v", typeOf)
		return
	}
	h.handlers[reflect.TypeOf(typeOf).Name()] = handler
	log.Debugf("Added handler for type %s, %v for Subscriber %s", reflect.TypeOf(typeOf), &handler, h.name)

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
