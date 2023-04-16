package pubsub

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type TopicName string

type Types map[string]reflect.Type

type Subscribers map[string]Subscriber

type TopicConfig struct {
	types            Types
	typeSafe         bool
	allowSetTypes    bool
	allowSetTypeSafe bool
	allowSetName     bool
	allowOverride    bool
}

type TopicStr struct {
	name        TopicName
	subscribers Subscribers
	cfg         TopicConfig
}

func NewTopic(name TopicName, cfg TopicConfig, types ...interface{}) (t *TopicStr) {
	t.cfg = cfg
	t.name = name
	t.cfg.types = make(Types, 0)
	t.subscribers = make(Subscribers, 0)

	if len(types) > 0 {
		t.cfg.typeSafe = true
		for _, v := range types {
			tName := reflect.TypeOf(v).Name()
			tType := reflect.TypeOf(v)
			t.cfg.types[tName] = tType
		}
	}
	log.Debugf("Created Topic %v", t)
	return
}

func (t *TopicStr) Pub(msg ...interface{}) (err error) {
	for _, s := range t.subscribers {
		s.Channel() <- msg
	}
	return
}

func (t *TopicStr) Name() TopicName {
	return t.name
}

func (t *TopicStr) SetName(name TopicName) (err error) {
	if !t.cfg.allowSetName {
		err = fmt.Errorf("allow.SetName is false")
		return
	}
	log.Debugf("Changing topic name from %s to %s", t.name, name)
	t.name = name
	return
}

func (t *TopicStr) AddSub(sub Subscriber) (err error) {
	if _, ok := t.subscribers[sub.Name()]; ok {
		if !t.cfg.allowOverride {
			err = fmt.Errorf("subscriber %s already exists and allowOverwrite is false", sub.Name())
			return
		}
	}
	t.subscribers[sub.Name()] = sub
	return
}

func (t *TopicStr) Subscribers() (s Subscribers) {
	for k, v := range t.subscribers {
		s[k] = v
	}
	return
}

func (t *TopicStr) Types() (ty Types) {
	for k, v := range t.cfg.types {
		ty[k] = v
	}
	return
}

func (t *TopicStr) SetTypes(types map[string]reflect.Type) (err error) {
	if !t.cfg.allowSetTypes {
		err = fmt.Errorf("allow.SetTypes is false")
		return
	}
	t.cfg.types = types
	return
}

func (t *TopicStr) IsTypeSafe() bool {
	return t.cfg.typeSafe
}

func (t *TopicStr) SetTypeSafe(typeSafe bool) (err error) {
	if !t.cfg.allowSetTypeSafe {
		err = fmt.Errorf("allow.SetTypeSafe is false")
		return
	}
	t.cfg.typeSafe = typeSafe
	return
}
