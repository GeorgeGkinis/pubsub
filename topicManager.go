package pubsub

var tm *TopicManager

func init() {
	tm = new(TopicManager)
}

type topics map[TopicName]*Topic

type TopicsManagerConfig struct {
	autoCreate bool
}

type TopicManager struct {
	topics
	TopicsManagerConfig
}

func NewTopicManager() (t *TopicManager) {
	t.topics = make(topics, 0)
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
