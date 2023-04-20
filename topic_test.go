package pubsub

import (
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
	"time"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

var (
	p1 = NewPublisher("p1")
	p2 = NewPublisher("p2")
	s1 = NewSubscriber("s1", nil, nil)
	s2 = NewSubscriber("s2", nil, nil)
	s3 = NewSubscriber("s3", nil, nil)

	ct1 = "type1"
	ct2 = 2

	types = NewTypes(ct1, ct2)
)

func simpleConsoleStringHandler(msg interface{}) (err error) {
	log.Errorf("simpleConsoleStringHandler received message: %v", msg)
	return
}

func simpleConsoleIntHandler(msg interface{}) (err error) {
	log.Errorf("simpleConsoleIntHandler received message: %v", msg)
	return
}

func TestNewTopic(t *testing.T) {

	type args struct {
		name  TopicName
		cfg   TopicConfig
		pubs  []*Publisher
		types []interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantT *Topic
	}{
		{name: "Create Topic with name \"TestNewTopicName\"", args: args{
			name:  "TestNewTopicName",
			cfg:   TopicConfig{},
			pubs:  make([]*Publisher, 0),
			types: make([]interface{}, 0),
		}, wantT: &Topic{
			name:        "TestNewTopicName",
			subscribers: make(Subscribers, 0),
			publishers:  make(Publishers, 0),
			cfg:         TopicConfig{},
		}},
		{name: "Create Topic with no (nil) config", args: args{
			name:  "TestNewTopicNoConfig",
			pubs:  make([]*Publisher, 0),
			types: make([]interface{}, 0),
		}, wantT: &Topic{
			name:        "TestNewTopicNoConfig",
			subscribers: make(Subscribers, 0),
			publishers:  make(Publishers, 0),
			cfg: TopicConfig{
				types:              nil,
				typeSafe:           false,
				allowSetTypes:      false,
				allowSetTypeSafe:   false,
				allowSetName:       false,
				allowOverride:      false,
				allowAddPub:        false,
				allowAllPublishers: false,
			},
		}},
		{name: "Create Topic with 2 publishers", args: args{
			name:  "TestNewTopic2Publishers",
			pubs:  []*Publisher{p1, p2},
			types: make([]interface{}, 0),
		}, wantT: &Topic{
			name:        "TestNewTopic2Publishers",
			subscribers: make(Subscribers, 0),
			publishers: Publishers{
				"p1": &Publisher{
					name:          "p1",
					subscriptions: nil,
				},
				"p2": &Publisher{
					name:          "p2",
					subscriptions: nil,
				}},
			cfg: TopicConfig{},
		}},
		{name: "Create Topic with 2 Types", args: args{
			name: "TestNewTopic2Types",
			cfg: TopicConfig{
				types: *types},
			pubs:  nil,
			types: nil,
		}, wantT: &Topic{
			name:        "TestNewTopic2Types",
			subscribers: make(Subscribers, 0),
			publishers:  make(Publishers, 0),
			cfg: TopicConfig{
				types: Types{
					"string": reflect.TypeOf("type1"),
					"int":    reflect.TypeOf(2),
				},
				typeSafe: true},
		}},
		{name: "Create Topic no name", args: args{
			name:  "",
			cfg:   TopicConfig{},
			pubs:  nil,
			types: nil,
		}, wantT: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, err := NewTopic(tt.args.name, tt.args.cfg, tt.args.pubs...)
			log.Errorf("Error: %v", err)
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("NewTopic()\nGot : %#v\nWant: %#v\n", gotT, tt.wantT)
			}
		})
	}
}

func TestTopic_AddPub(t1 *testing.T) {
	type args struct {
		pub *Publisher
	}
	tests := []struct {
		name    string
		fields  Topic
		args    args
		wantErr bool
	}{
		{name: "Topic no Publishers and allowAddPub true", fields: Topic{
			name:       "allowAddPub true",
			publishers: make(Publishers, 0),
			cfg: TopicConfig{
				allowAddPub: true,
			},
		}, args: struct{ pub *Publisher }{pub: p1}, wantErr: false},
		{name: "Topic no Publishers and allowAddPub false", fields: Topic{
			name:       "allowAddPub false",
			publishers: make(Publishers, 0),
			cfg: TopicConfig{
				allowAddPub: false,
			},
		}, args: struct{ pub *Publisher }{pub: p1}, wantErr: true},

		{name: "Topic with Publishers and allowAddPub true", fields: Topic{
			name: "allowAddPub true",
			publishers: Publishers{
				"p1": &Publisher{
					name: "p1",
				}},
			cfg: TopicConfig{
				allowAddPub: true,
			},
		}, args: struct{ pub *Publisher }{pub: p2}, wantErr: false},
		{name: "Topic with Publishers and allowAddPub false", fields: Topic{
			name: "allowAddPub false",
			publishers: Publishers{
				"p1": &Publisher{
					name: "p1",
				}},
			cfg: TopicConfig{
				allowAddPub: false,
			},
		}, args: struct{ pub *Publisher }{pub: p1}, wantErr: true},
		{name: "Add Publisher with same name to topic and allowOverwrite true", fields: Topic{
			name: "allowOverwrite false",
			publishers: Publishers{
				"p1": &Publisher{
					name: "p1",
				}},
			cfg: TopicConfig{
				allowAddPub:   true,
				allowOverride: true,
			},
		}, args: struct{ pub *Publisher }{pub: p1}, wantErr: false},
		{name: "Add Publisher with same name to topic and allowOverwrite false", fields: Topic{
			name: "allowOverwrite false",
			publishers: Publishers{
				"p1": &Publisher{
					name: "p1",
				}},
			cfg: TopicConfig{
				allowAddPub:   true,
				allowOverride: false,
			},
		}, args: struct{ pub *Publisher }{pub: p1}, wantErr: true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.fields.name,
				subscribers: tt.fields.subscribers,
				publishers:  tt.fields.publishers,
				cfg:         tt.fields.cfg,
			}
			err := t.AddPub(tt.args.pub)
			log.Errorf("Error: %v", err)
			if (err != nil) != tt.wantErr {
				t1.Errorf("AddPub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopic_AddSub(t1 *testing.T) {
	type args struct {
		sub SubscriberIF
	}
	tests := []struct {
		name    string
		topic   Topic
		args    args
		wantErr bool
	}{
		{name: "Topic no Subs and allowOverride true", topic: Topic{
			name:        "Topic no Subs and allowOverride true",
			subscribers: make(Subscribers, 0),
			cfg:         TopicConfig{allowOverride: true},
		}, args: args{
			s1,
		}, wantErr: false},
		{name: "Topic no Subs and allowOverride false", topic: Topic{
			name:        "Topic no Subs and allowOverride false",
			subscribers: make(Subscribers, 0),
			cfg:         TopicConfig{allowOverride: false},
		}, args: args{s1}, wantErr: false},
		{name: "Add Subscriber with same name to topic allowOverride true", topic: Topic{
			name: "Topic same Sub and allowOverride true",
			subscribers: Subscribers{
				"s1": s1,
			},
			cfg: TopicConfig{allowOverride: true},
		}, args: args{s1}, wantErr: false},
		{name: "Add Subscriber with same name to topic and allowOverride false", topic: Topic{
			name: "Topic same Sub and allowOverride false",
			subscribers: Subscribers{
				"s1": s1,
			},
			cfg: TopicConfig{allowOverride: false},
		}, args: args{s1}, wantErr: true},
		{name: "Add Subscriber with other name to topic allowOverride true", topic: Topic{
			name: "Topic same Sub and allowOverride true",
			subscribers: Subscribers{
				"s1": s1,
			},
			cfg: TopicConfig{allowOverride: true},
		}, args: args{s2}, wantErr: false},
		{name: "Add Subscriber with other name to topic and allowOverride false", topic: Topic{
			name: "Topic same Sub and allowOverride false",
			subscribers: Subscribers{
				"s1": s1,
			},
			cfg: TopicConfig{allowOverride: false},
		}, args: args{s2}, wantErr: false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.topic.name,
				subscribers: tt.topic.subscribers,
				publishers:  tt.topic.publishers,
				cfg:         tt.topic.cfg,
			}
			err := t.AddSub(tt.args.sub)
			log.Errorf("Error: %v", err)
			if (err != nil) != tt.wantErr {
				t1.Errorf("AddSub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopic_IsTypeSafe(t1 *testing.T) {

	tests := []struct {
		name  string
		topic Topic
		want  bool
	}{
		{name: "Topic isTypeSafe true", topic: Topic{
			name:        "Topic isTypeSafe true",
			subscribers: nil,
			publishers:  nil,
			cfg: TopicConfig{
				typeSafe: true,
			},
		}, want: true},
		{name: "Topic isTypeSafe false", topic: Topic{
			name:        "Topic isTypeSafe false",
			subscribers: nil,
			publishers:  nil,
			cfg: TopicConfig{
				typeSafe: false,
			},
		}, want: false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.topic.name,
				subscribers: tt.topic.subscribers,
				publishers:  tt.topic.publishers,
				cfg:         tt.topic.cfg,
			}
			if got := t.IsTypeSafe(); got != tt.want {
				t1.Errorf("IsTypeSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopic_Name(t1 *testing.T) {
	tests := []struct {
		name  string
		topic Topic
		want  TopicName
	}{
		{name: "Topic with name", topic: Topic{
			name: "Topic with name",
		}, want: "Topic with name"},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.topic.name,
				subscribers: tt.topic.subscribers,
				publishers:  tt.topic.publishers,
				cfg:         tt.topic.cfg,
			}
			if got := t.Name(); got != tt.want {
				t1.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopic_Pub(t1 *testing.T) {

	type args struct {
		pub Publisher
		msg []interface{}
	}
	type handler struct {
		typeof     interface{}
		handlefunc func(msg interface{}) (err error)
	}
	tests := []struct {
		name     string
		topic    Topic
		handlers []handler
		args     args
		wantErr  bool
	}{
		{name: "Publisher exists, allowAllPublishers true, 2 messages sent", topic: Topic{
			name:        "Publisher exists and allowAllPublishers true",
			subscribers: Subscribers{"s3": s3},
			publishers:  Publishers{"p1": p1},
			cfg:         TopicConfig{},
		},
			handlers: []handler{
				{
					typeof:     "",
					handlefunc: simpleConsoleStringHandler,
				},
				{
					typeof:     42,
					handlefunc: simpleConsoleIntHandler,
				}},
			args: args{
				pub: *p1,
				msg: []interface{}{"Message: Publisher exists and allowAllPublishers true", 42},
			}, wantErr: false},
		//{name: "Publisher does not exist and allowAllPublishers true", topic: , args: , wantErr: },
		//{name: "Publisher does not exist and allowAllPublishers false", topic: , args: , wantErr: },

	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.topic.name,
				subscribers: tt.topic.subscribers,
				publishers:  tt.topic.publishers,
				cfg:         tt.topic.cfg,
			}
			log.SetLevel(log.DebugLevel)
			for _, v := range tt.handlers {
				s3.AddHandler(v.typeof, v.handlefunc)
			}

			log.Debugf("Handlers registered: %v", s3.handlers)
			s3.Listen()
			if err := t.Pub(tt.args.pub, tt.args.msg...); (err != nil) != tt.wantErr {
				t1.Errorf("Pub() error = %v, wantErr %v", err, tt.wantErr)
			}
			time.Sleep(100 * time.Millisecond)
		})
	}
}

func TestTopic_Publishers(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	tests := []struct {
		name   string
		fields fields
		wantP  Publishers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.fields.name,
				subscribers: tt.fields.subscribers,
				publishers:  tt.fields.publishers,
				cfg:         tt.fields.cfg,
			}
			if gotP := t.Publishers(); !reflect.DeepEqual(gotP, tt.wantP) {
				t1.Errorf("Publishers() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestTopic_SetName(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	type args struct {
		name TopicName
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.fields.name,
				subscribers: tt.fields.subscribers,
				publishers:  tt.fields.publishers,
				cfg:         tt.fields.cfg,
			}
			if err := t.SetName(tt.args.name); (err != nil) != tt.wantErr {
				t1.Errorf("SetName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopic_SetTypeSafe(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	type args struct {
		typeSafe bool
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.fields.name,
				subscribers: tt.fields.subscribers,
				publishers:  tt.fields.publishers,
				cfg:         tt.fields.cfg,
			}
			if err := t.SetTypeSafe(tt.args.typeSafe); (err != nil) != tt.wantErr {
				t1.Errorf("SetTypeSafe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopic_SetTypes(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	type args struct {
		types map[string]reflect.Type
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.fields.name,
				subscribers: tt.fields.subscribers,
				publishers:  tt.fields.publishers,
				cfg:         tt.fields.cfg,
			}
			if err := t.SetTypes(tt.args.types); (err != nil) != tt.wantErr {
				t1.Errorf("SetTypes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopic_Subscribers(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	tests := []struct {
		name   string
		fields fields
		wantS  Subscribers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.fields.name,
				subscribers: tt.fields.subscribers,
				publishers:  tt.fields.publishers,
				cfg:         tt.fields.cfg,
			}
			if gotS := t.Subscribers(); !reflect.DeepEqual(gotS, tt.wantS) {
				t1.Errorf("Subscribers() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestTopic_Types(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	tests := []struct {
		name   string
		fields fields
		wantTy Types
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.fields.name,
				subscribers: tt.fields.subscribers,
				publishers:  tt.fields.publishers,
				cfg:         tt.fields.cfg,
			}
			if gotTy := t.Types(); !reflect.DeepEqual(gotTy, tt.wantTy) {
				t1.Errorf("Types() = %v, want %v", gotTy, tt.wantTy)
			}
		})
	}
}
