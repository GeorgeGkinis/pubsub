package pubsub

import (
	"reflect"
	"testing"
)

func TestNewTopicManager(t *testing.T) {
	tests := []struct {
		name  string
		wantT *TopicManager
	}{
		{name: "Test constructor", wantT: NewTopicManager()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := NewTopicManager(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("NewTopicManager() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func TestTopicManager_Topic(t *testing.T) {
	type fields struct {
		topics              topics
		TopicsManagerConfig TopicsManagerConfig
	}
	type args struct {
		n TopicName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantT  *Topic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TopicManager{
				topics:              tt.fields.topics,
				TopicsManagerConfig: tt.fields.TopicsManagerConfig,
			}
			if gotT := tm.Topic(tt.args.n); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("Topic() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func TestNewTopicManager1(t *testing.T) {
	tests := []struct {
		name  string
		wantT *TopicManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := NewTopicManager(); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("NewTopicManager() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func TestTopicManager_RegisterTopic(t *testing.T) {
	type fields struct {
		topics              topics
		TopicsManagerConfig TopicsManagerConfig
	}
	type args struct {
		topic *Topic
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TopicManager{
				topics:              tt.fields.topics,
				TopicsManagerConfig: tt.fields.TopicsManagerConfig,
			}
			if err := tm.RegisterTopic(tt.args.topic); (err != nil) != tt.wantErr {
				t.Errorf("RegisterTopic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopicManager_Topic1(t *testing.T) {
	type fields struct {
		topics              topics
		TopicsManagerConfig TopicsManagerConfig
	}
	type args struct {
		n TopicName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantT  *Topic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TopicManager{
				topics:              tt.fields.topics,
				TopicsManagerConfig: tt.fields.TopicsManagerConfig,
			}
			if gotT := tm.Topic(tt.args.n); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("Topic() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func TestTopicManager_Topics(t *testing.T) {
	type fields struct {
		topics              topics
		TopicsManagerConfig TopicsManagerConfig
	}
	type args struct {
		topicNames []TopicName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantT  []*Topic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TopicManager{
				topics:              tt.fields.topics,
				TopicsManagerConfig: tt.fields.TopicsManagerConfig,
			}
			if gotT := tm.Topics(tt.args.topicNames); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("Topics() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
