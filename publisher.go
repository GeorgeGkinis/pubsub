package pubsub

type Publisher interface {
	TypeChecker
}

type TypeChecker interface {
	CheckType(topicName string, checkMsg interface{}) (bool, error)
}
