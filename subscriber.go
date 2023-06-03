package pubsub

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"runtime"
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

func NewSubscriber(name string, handlers Handlers, subscriptions ...*Topic) (s *Subscriber, err error) {
	s = &Subscriber{
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
			err = v.AddSub(s)
			if s == nil {
				return s, fmt.Errorf("cannot subscribe Subscriber %v to Topic %v", s.name, v.Name())
			}
		}
	}

	if s == nil {
		return s, fmt.Errorf("could not create Subscriber")
	}

	return s, err
}

func (s *Subscriber) Listen() {
	if !s.listening {
		s.listening = true
		go func() {
			for msg := range s.ch {
				log.Debugf("Received message of type %T", msg)
				handler, ok := s.handlers[reflect.TypeOf(msg).Name()]
				if !ok {
					log.Debugf("Subscriber %s has no handler for message type %T, checking existence of handler for \"any\" type.", s.name, msg)
					handler, ok = s.handlers["any"]
					if !ok {
						log.Errorf("Subscriber %s has no handler for message type %T, and no handler for \"any\" type.", s.name, msg)
						return
					}
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
	if typeOf == "any" {
		s.handlers["any"] = handler
		log.Debugf("Added handler for type %s, %v for Subscriber %s", "any", runtime.FuncForPC(reflect.ValueOf(*handler).Pointer()).Name(), s.name)
	} else {
		s.handlers[reflect.TypeOf(typeOf).Name()] = handler
		log.Debugf("Added handler for type %s, %v for Subscriber %s", reflect.TypeOf(typeOf), &handler, s.name)
	}

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
