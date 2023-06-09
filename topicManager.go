package pubsub

import "fmt"

var TM = NewTopicManager()

func init() {

}

type topics map[TopicName]*Topic

type TopicsManagerConfig struct {
	autoCreate bool
}

type TopicManager struct {
	topics
	TopicsManagerConfig
}

func NewTopicManager() *TopicManager {
	t := &TopicManager{
		topics:              make(topics, 0),
		TopicsManagerConfig: TopicsManagerConfig{},
	}
	return t
}

func (tm *TopicManager) Topic(n TopicName) (t *Topic) {
	if tm.autoCreate {
		t, _ := NewTopic(n, TopicConfig{})
		tm.topics[n] = t
	} else {
		return nil
	}
	return tm.topics[n]
}

func (tm *TopicManager) Topics(topicNames []TopicName) (t []*Topic) {
	for _, tn := range topicNames {
		if topic, ok := tm.topics[tn]; ok {
			t = append(t, topic)
		}
	}
	return t
}

func (tm *TopicManager) RegisterTopic(topic *Topic) (err error) {
	if _, ok := tm.topics[topic.Name()]; !ok {
		tm.topics[topic.Name()] = topic
	} else {
		err = fmt.Errorf("topic with name %s already exists", topic.Name())
	}

	return
}
