package pubsub

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type PublisherIF interface {
	TypeChecker
	Name() string
	GetSubscriptions() Subscriptions
}

type TypeChecker interface {
	CheckType(topicName TopicName, checkMsg interface{}) (bool, error)
}

type Publisher struct {
	name          string
	subscriptions Subscriptions
}

func NewPublisher(name string) *Publisher {
	return &Publisher{name: name}
}

func (p *Publisher) CheckType(topicName TopicName, checkMsg interface{}) (typeOk bool, err error) {

	t := reflect.TypeOf(checkMsg).Name()
	topic := TM.Topic(topicName)
	if topic == nil {
		err = fmt.Errorf("non-existing topic %s", topicName)
		return
	}
	_, typeOk = TM.Topic(topicName).Types()[t]
	if typeOk {
		log.Debugf("topic %s allows Type %T", topicName, checkMsg)
	} else {
		log.Debugf("topic %s does not allow Type %T", topicName, checkMsg)
	}
	return

}

func (p *Publisher) Name() string {
	return p.name
}

func (p *Publisher) GetSubscriptions() Subscriptions {
	return p.subscriptions
}
