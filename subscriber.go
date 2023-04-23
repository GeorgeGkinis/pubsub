package pubsub

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type SubscriberIF interface {
	Listen()
	Name() string
	AddHandler(interface{}, *HandlerFunc) error
	Sub(topicName *Topic) error
	Channel() chan interface{}
	GetSubscriptions() []*Topic
}

type HandlerFunc func(msg interface{}) (err error)

type HandleType string

type Handlers map[string]*HandlerFunc

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
		_ = s.AddHandler(k, v)
	}
	if subscriptions != nil {
		for _, v := range subscriptions {
			s.subscriptions = append(s.subscriptions, v)
		}
	}

	return &s
}

func (s *Subscriber) Listen() {
	if !s.listening {
		s.listening = true
		go func() {
			for msg := range s.ch {
				log.Debugf("Received message of type %T", msg)
				handler, ok := s.handlers[reflect.TypeOf(msg).Name()]
				if !ok {
					log.Errorf("Subscriber %s has no handler for message type %T", s.name, msg)
				}
				if err := (*handler)(msg); err != nil {
					log.Errorf("error handling message: %v of type %T on handler: %s", msg, msg, s.Name())
				}
			}
		}()
	}
}

func (s *Subscriber) AddHandler(typeOf interface{}, handler *HandlerFunc) (err error) {

	if typeOf == nil || handler == nil {
		err = fmt.Errorf("Required: typeOf and handler. Provided: typeOf: %v", typeOf)
		return
	}
	s.handlers[reflect.TypeOf(typeOf).Name()] = handler
	log.Debugf("Added handler for type %s, %v for Subscriber %s", reflect.TypeOf(typeOf), &handler, s.name)

	//TODO: Check if overwriting existing handler
	return
}

func (s *Subscriber) Name() string {
	return s.name
}

func (s *Subscriber) Channel() chan interface{} {
	return s.ch
}

func (s *Subscriber) GetSubscriptions() []*Topic {
	return s.subscriptions
}

func (s *Subscriber) Sub(topic *Topic) (err error) {
	s.subscriptions = append(s.subscriptions, topic)
	err = topic.AddSub(s)
	return err
}
