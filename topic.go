package pubsub

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type TopicName string

type Types map[string]reflect.Type

func NewTypes(types ...interface{}) *Types {
	t := make(Types, 0)
	for _, v := range types {
		t[reflect.TypeOf(v).Name()] = reflect.TypeOf(v)
	}
	return &t
}

type Subscribers map[string]SubscriberIF

type Publishers map[string]*Publisher

type TopicConfig struct {
	types              Types
	typeSafe           bool
	allowSetTypes      bool
	allowSetTypeSafe   bool
	allowSetName       bool
	allowOverride      bool
	allowAddPub        bool
	allowAllPublishers bool
}

type Topic struct {
	name        TopicName
	subscribers Subscribers
	publishers  Publishers
	cfg         TopicConfig
}

func NewTopic(name TopicName, cfg TopicConfig, pubs ...*Publisher) (topic *Topic, err error) {

	if name == "" {
		err = fmt.Errorf("Cannot create Topic without name.")
		return nil, err
	}

	t := new(Topic)
	t.cfg = cfg
	t.name = name
	t.subscribers = make(Subscribers, 0)
	t.publishers = make(Publishers, 0)

	if pubs != nil {
		for _, v := range pubs {
			t.publishers[v.Name()] = v
		}
	}

	if cfg.types != nil && len(cfg.types) > 0 {
		t.cfg.typeSafe = true
	}
	log.Debugf("Created Topic %v", t)
	return t, err
}

func (t *Topic) Pub(pub Publisher, msg ...interface{}) (err error) {
	if _, ok := t.publishers[pub.Name()]; !ok && t.cfg.allowAllPublishers {
		err = fmt.Errorf("publisher %s is not whitelisted for topic %s and allowAllPublishers is false", pub.Name(), t.name)
		return
	}
	for _, s := range t.subscribers {
		for _, m := range msg {
			(s).Channel() <- m
		}

	}
	return
}

func (t *Topic) Name() TopicName {
	return t.name
}

func (t *Topic) SetName(name TopicName) (err error) {
	if !t.cfg.allowSetName {
		err = fmt.Errorf("allow.SetName is false")
		return
	}
	log.Debugf("Changing topic name from %s to %s", t.name, name)
	t.name = name
	return
}

func (t *Topic) Subscribers() (s Subscribers) {
	for k, v := range t.subscribers {
		s[k] = v
	}
	return
}

func (t *Topic) AddSub(sub SubscriberIF) (err error) {
	if _, ok := t.subscribers[(sub).Name()]; ok {
		if !t.cfg.allowOverride {
			err = fmt.Errorf("subscriber %s already exists and allowOverwrite is false", sub.Name())
			return
		}
	}
	t.subscribers[(sub).Name()] = sub
	return
}

func (t *Topic) Publishers() (p Publishers) {
	for k, v := range t.publishers {
		p[k] = v
	}
	return
}

func (t *Topic) AddPub(pub *Publisher) (err error) {
	if !t.cfg.allowAddPub {
		err = fmt.Errorf("AddPub not allowed for topic %s", t.name)
		return
	}
	if _, ok := t.publishers[(*pub).Name()]; ok {
		if !t.cfg.allowOverride {
			err = fmt.Errorf("publisher %s already exists and allowOverride is false", (*pub).Name())
			return
		}
	}
	t.publishers[(*pub).Name()] = pub
	return
}

func (t *Topic) Types() (ty Types) {
	for k, v := range t.cfg.types {
		ty[k] = v
	}
	return
}

func (t *Topic) SetTypes(types map[string]reflect.Type) (err error) {
	if !t.cfg.allowSetTypes {
		err = fmt.Errorf("allow.SetTypes is false")
		return
	}
	t.cfg.types = types
	return
}

func (t *Topic) IsTypeSafe() bool {
	return t.cfg.typeSafe
}

func (t *Topic) SetTypeSafe(typeSafe bool) (err error) {
	if !t.cfg.allowSetTypeSafe {
		err = fmt.Errorf("allow.SetTypeSafe is false")
		return
	}
	t.cfg.typeSafe = typeSafe
	return
}
