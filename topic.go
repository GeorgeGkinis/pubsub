package pubsub

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type TopicName string

type Types map[string]reflect.Type

func NewTypes(types ...interface{}) Types {
	t := make(Types, 0)
	for _, v := range types {
		t[reflect.TypeOf(v).Name()] = reflect.TypeOf(v)
	}
	return t
}

type Subscribers map[string]*Subscriber

type Publishers map[string]*Publisher

type TopicConfig struct {
	Types              Types
	TypeSafe           bool
	AllowSetTypes      bool
	AllowSetTypeSafe   bool
	AllowSetName       bool
	AllowOverride      bool
	AllowAddPub        bool
	AllowAllPublishers bool
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
	t.cfg.Types = make(Types, 0)
	t.name = name
	t.subscribers = make(Subscribers, 0)
	t.publishers = make(Publishers, 0)

	if pubs != nil {
		for _, v := range pubs {
			if v != nil {
				t.publishers[v.Name()] = v
			}
		}
	}

	if cfg.Types != nil && len(cfg.Types) > 0 {
		t.cfg.TypeSafe = true
		t.cfg.Types = cfg.Types
	}
	log.Debugf("Created Topic %v", t)

	err = TM.RegisterTopic(t)

	return t, err
}

func (t *Topic) Pub(pub *Publisher, msg ...interface{}) (err error) {
	if _, ok := t.publishers[pub.Name()]; ok == false && t.cfg.AllowAllPublishers == false {
		err = fmt.Errorf("publisher %s is not whitelisted for topic: \"%s\" and AllowAllPublishers is false", pub.Name(), t.name)
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
	if !t.cfg.AllowSetName {
		err = fmt.Errorf("allow.SetName is false")
		return
	}
	if name == "" {
		err = fmt.Errorf("tried to set name to empty string for Topic %s", t.name)
	}
	log.Debugf("Changing topic name from %s to %s", t.name, name)
	t.name = name
	return
}

func (t *Topic) Subscribers() (s Subscribers) {
	return t.subscribers
}

func (t *Topic) AddSub(sub *Subscriber) (err error) {
	if _, ok := t.subscribers[sub.Name()]; ok {
		if !t.cfg.AllowOverride {
			err = fmt.Errorf("subscriber %s already exists and allowOverwrite is false", (sub).Name())
			return
		}
	}
	t.subscribers[sub.Name()] = sub
	return
}

func (t *Topic) Publishers() (p Publishers) {
	return t.publishers
}

func (t *Topic) AddPub(pub *Publisher) (err error) {
	if !t.cfg.AllowAddPub {
		err = fmt.Errorf("AddPub not allowed for topic %s", t.name)
		return
	}
	if _, ok := t.publishers[pub.Name()]; ok {
		if !t.cfg.AllowOverride {
			err = fmt.Errorf("publisher %s already exists and AllowOverride is false", pub.Name())
			return
		}
	}
	t.publishers[pub.Name()] = pub
	return
}

func (t *Topic) Types() (ty Types) {
	return t.cfg.Types
}

func (t *Topic) SetTypes(types ...interface{}) (err error) {
	if !t.cfg.AllowSetTypes {
		err = fmt.Errorf("allow.SetTypes is false for Topic %s", t.name)
		return
	}
	if len(types) == 0 {
		err = fmt.Errorf("no Types received for Topic: %s", t.name)
		return
	}
	for _, v := range types {
		t.cfg.Types[reflect.TypeOf(v).Name()] = reflect.TypeOf(v)
	}
	return
}

func (t *Topic) IsTypeSafe() bool {
	return t.cfg.TypeSafe
}

func (t *Topic) SetTypeSafe(typeSafe bool) (err error) {
	if !t.cfg.AllowSetTypeSafe {
		err = fmt.Errorf("AllowSetTypeSafe is false for Topic %s", t.name)
		return
	}
	t.cfg.TypeSafe = typeSafe
	return
}
