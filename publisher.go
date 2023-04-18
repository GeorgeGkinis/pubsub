package pubsub

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type PublisherIF interface {
	TypeChecker
	Name() string
	GetSubscriptions() ([]Topic, error)
}

type TypeChecker interface {
	CheckType(topicName string, checkMsg interface{}) (bool, error)
}

type Publisher struct {
	name          string
	subscriptions Subscriptions
}

func (p Publisher) CheckType(topicName TopicName, checkMsg interface{}) (typeOk bool, err error) {

	t := reflect.TypeOf(checkMsg).Name()
	topic := tm.Topic(topicName)
	if topic == nil {
		err = fmt.Errorf("non-existing topic %s", topicName)
		return
	}
	_, typeOk = tm.Topic(topicName).Types()[t]
	if typeOk {
		log.Debugf("topic %s allows Type %T", topicName, checkMsg)
	} else {
		log.Debugf("topic %s does not allow Type %T", topicName, checkMsg)
	}
	return

}

func (p Publisher) Name() string {
	//TODO implement me
	panic("implement me")
}

func (p Publisher) GetSubscriptions() ([]Topic, error) {
	//TODO implement me
	panic("implement me")
}
