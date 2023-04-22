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
		name TopicName
		cfg  TopicConfig
		pubs []*Publisher
	}
	tests := []struct {
		name  string
		args  args
		wantT *Topic
	}{
		{name: "Create Topic with name \"TestNewTopicName\"", args: args{
			name: "TestNewTopicName",
			cfg:  TopicConfig{},
		}, wantT: &Topic{
			name:        "TestNewTopicName",
			subscribers: make(Subscribers, 0),
			publishers:  make(Publishers, 0),
			cfg: TopicConfig{
				types: make(Types, 0),
			},
		}},
		{name: "Create Topic with (nil) config", args: args{
			name: "TestNewTopicNoConfig",
			pubs: make([]*Publisher, 0),
		}, wantT: &Topic{
			name:        "TestNewTopicNoConfig",
			subscribers: make(Subscribers, 0),
			publishers:  make(Publishers, 0),
			cfg: TopicConfig{
				types: make(Types, 0),
			},
		}},
		{name: "Create Topic with 2 publishers", args: args{
			name: "TestNewTopic2Publishers",
			pubs: []*Publisher{p1, p2},
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
			cfg: TopicConfig{
				types: make(Types, 0),
			},
		}},
		{name: "Create Topic with nil publishers", args: args{
			name: "Create Topic with nil publishers",
			pubs: nil,
		}, wantT: &Topic{
			name:        "Create Topic with nil publishers",
			subscribers: make(Subscribers, 0),
			publishers:  make(Publishers, 0),
			cfg: TopicConfig{
				types: make(Types, 0),
			},
		}},
		{name: "Create Topic with 2 Types", args: args{
			name: "TestNewTopic2Types",
			cfg: TopicConfig{
				types: Types{
					"string": reflect.TypeOf("type1"),
					"int":    reflect.TypeOf(2),
				},
			},
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
			name: "",
			cfg:  TopicConfig{},
			pubs: nil,
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
	type tp struct {
		name        TopicName
		cfg         TopicConfig
		subscribers []*Subscribers
		publishers  []*Publisher
	}
	tests := []struct {
		name     string
		topic    tp
		handlers []handler
		args     args
		wantErr  bool
	}{
		{name: "Publisher exists, allowAllPublishers true, 2 messages sent", topic: tp{
			name:       "Publisher exists and allowAllPublishers true",
			publishers: []*Publisher{p1},
			cfg: TopicConfig{
				allowAllPublishers: true,
			},
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
				msg: []interface{}{"Message: Publisher exists and allowAllPublishers false", 42},
			}, wantErr: false},
		{name: "Publisher exists, allowAllPublishers false, 2 messages sent", topic: tp{
			name:       "Publisher exists and allowAllPublishers false",
			publishers: []*Publisher{p1},
			cfg: TopicConfig{
				allowAllPublishers: false,
			},
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
				msg: []interface{}{"Message: Publisher exists and allowAllPublishers false", 42},
			}, wantErr: false},
		{
			name: "Publisher not exists, allowAllPublishers false, 2 messages sent",
			topic: tp{
				name:       "Publisher not exists, allowAllPublishers false, 2 messages sent",
				publishers: nil,
				cfg: TopicConfig{
					allowAllPublishers: false,
				},
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
			}, wantErr: true},
		{
			name: "Publisher not exists, allowAllPublishers true, 2 messages sent",
			topic: tp{
				name:       "Publisher not exists, allowAllPublishers true, 2 messages sent",
				publishers: nil,
				cfg: TopicConfig{
					allowAllPublishers: true,
				},
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
			}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			log.SetLevel(log.DebugLevel)

			t, err := NewTopic(tt.topic.name, tt.topic.cfg, tt.topic.publishers...)
			if err != nil {
				log.Error(err)
			}

			s3 := NewSubscriber("s3", nil, nil)
			if err := t.AddSub(s3); err != nil {
				log.Error(err)
			}
			for _, v := range tt.handlers {
				s3.AddHandler(v.typeof, v.handlefunc)
			}
			log.Debugf("Handlers registered: %v", s3.handlers)
			s3.Listen()
			err = t.Pub(tt.args.pub, tt.args.msg...)
			log.Error(err)
			if err != nil != tt.wantErr {
				t1.Errorf("Pub() error = %v, wantErr %v", err, tt.wantErr)
			}
			time.Sleep(100 * time.Millisecond)
		})
	}
}

func TestTopic_Publishers(t1 *testing.T) {
	type tp struct {
		name        TopicName
		cfg         TopicConfig
		subscribers []*Subscribers
		publishers  []*Publisher
	}
	tests := []struct {
		name  string
		topic tp
		wantP Publishers
	}{
		{name: "No Publishers",
			topic: tp{
				name:       "No Publishers",
				publishers: nil,
			}, wantP: make(Publishers, 0)},
		{name: "One Publisher",
			topic: tp{
				name:       "No Publishers",
				publishers: []*Publisher{p1},
			}, wantP: Publishers{"p1": p1}},
		{name: "Two Publishers",
			topic: tp{
				name:       "No Publishers",
				publishers: []*Publisher{p1, p2},
			}, wantP: Publishers{"p1": p1, "p2": p2}},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {

			log.SetLevel(log.DebugLevel)

			t, err := NewTopic(tt.topic.name, tt.topic.cfg, tt.topic.publishers...)
			if err != nil {
				log.Error(err)
			}

			if gotP := t.Publishers(); !reflect.DeepEqual(gotP, tt.wantP) {
				t1.Errorf("Publishers() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestTopic_SetName(t1 *testing.T) {
	type tp struct {
		name        TopicName
		cfg         TopicConfig
		subscribers []*Subscribers
		publishers  []*Publisher
	}
	type args struct {
		name TopicName
	}
	tests := []struct {
		name    string
		topic   tp
		args    args
		wantErr bool
	}{
		{name: "Set Name, allowSetName true",
			topic: tp{
				name: "Name Before",
				cfg: TopicConfig{
					allowSetName: true,
				},
			}, args: args{name: "Name After"},
			wantErr: false},
		{name: "Set Name to empty, allowSetName true",
			topic: tp{
				name: "Name Before",
				cfg: TopicConfig{
					allowSetName: true,
				},
			}, args: args{name: ""},
			wantErr: true},
		{name: "Set Name , allowSetName false",
			topic: tp{
				name: "Name Before",
				cfg: TopicConfig{
					allowSetName: false,
				},
			}, args: args{name: "Name After"},
			wantErr: true},
		{name: "Set Name to empty, allowSetName false",
			topic: tp{
				name: "Name Before",
				cfg: TopicConfig{
					allowSetName: false,
				},
			}, args: args{name: ""},
			wantErr: true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			log.SetLevel(log.DebugLevel)

			t, err := NewTopic(tt.topic.name, tt.topic.cfg, tt.topic.publishers...)
			if err != nil {
				log.Error(err)
			}
			err = t.SetName(tt.args.name)
			log.Error(err)
			if (err != nil) != tt.wantErr {
				t1.Errorf("SetName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopic_SetTypeSafe(t1 *testing.T) {
	type args struct {
		typeSafe bool
	}
	tests := []struct {
		name    string
		topic   Topic
		args    args
		wantErr bool
	}{
		{name: "SetTypeSafe, allowSetTypeSafe true",
			topic: Topic{
				name: "SetTypeSafe, allowSetTypeSafe true",
				cfg: TopicConfig{
					allowSetTypeSafe: true,
				},
			}, args: args{typeSafe: true}, wantErr: false},
		{name: "SetTypeSafe, allowSetTypeSafe false",
			topic: Topic{
				name: "SetTypeSafe, allowSetTypeSafe false",
				cfg: TopicConfig{
					allowSetTypeSafe: false,
				},
			}, args: args{typeSafe: true}, wantErr: true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Topic{
				name:        tt.topic.name,
				subscribers: tt.topic.subscribers,
				publishers:  tt.topic.publishers,
				cfg:         tt.topic.cfg,
			}
			err := t.SetTypeSafe(tt.args.typeSafe)
			if err != nil {
				log.Error(err)
			}
			if (err != nil) != tt.wantErr {
				t1.Errorf("SetTypeSafe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopic_SetTypes(t1 *testing.T) {
	type args struct {
		types []interface{}
	}
	type tp struct {
		name        TopicName
		cfg         TopicConfig
		subscribers []*Subscribers
		publishers  []*Publisher
	}
	tests := []struct {
		name    string
		topic   tp
		args    args
		wantErr bool
	}{
		{name: "SetTypes and allowsSetTypes true",
			topic: tp{
				name: "SetTypes and allowsSetTypes true",
				cfg: TopicConfig{
					allowSetTypes: true,
				},
			}, args: args{types: []interface{}{
				"42",
				42,
				Topic{},
			},
			}, wantErr: false},
		{name: "SetTypes and allowSetTypes false",
			topic: tp{
				name: "SetTypes and allowsSetTypes false",
				cfg: TopicConfig{
					allowSetTypes: false,
				},
			}, args: args{types: []interface{}{
				"42",
				42,
				Topic{},
			},
			}, wantErr: true},
		{name: "SetTypes with empty args and allowsSetTypes true",
			topic: tp{
				name: "SetTypes and allowsSetTypes true",
				cfg: TopicConfig{
					allowSetTypes: true,
				},
			}, args: args{types: []interface{}{}}, wantErr: true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t, _ := NewTopic(tt.topic.name, tt.topic.cfg)
			err := t.SetTypes(tt.args.types...)
			if err != nil {
				log.Error(err)
			}
			log.Debugf("Types: %s", t.cfg.types)

			tmpTypes := func(types ...interface{}) (t Types) {
				t = make(Types, 0)
				for _, v := range types {
					t[reflect.TypeOf(v).Name()] = reflect.TypeOf(v)
				}
				return
			}
			tmpt := tmpTypes(tt.args.types...)
			b := reflect.DeepEqual(tmpt, t.cfg.types)
			log.Debugf("DeepEqual: %v,Wanted %s, got %s", b, tmpt, t.cfg.types)

			if (err != nil) != tt.wantErr {
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
		{name: "Topic no Subscribers",
			fields: fields{
				name:        "Topic no Subscribers",
				subscribers: make(Subscribers, 0),
			}, wantS: make(Subscribers, 0)},
		{name: "Topic one Subscriber",
			fields: fields{
				name:        "Topic one Subscribers",
				subscribers: Subscribers{"s1": s1},
			}, wantS: Subscribers{"s1": s1}},
		{name: "Topic two Subscriber",
			fields: fields{
				name:        "Topic two Subscribers",
				subscribers: Subscribers{"s1": s1, "s2": s2},
			}, wantS: Subscribers{"s1": s1, "s2": s2}},
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
		{name: "Topic no Types",
			fields: fields{
				name: "Topic no Types",
				cfg: TopicConfig{
					types: make(Types, 0),
				},
			}, wantTy: make(Types, 0)},
		{name: "Topic one Type",
			fields: fields{
				name: "Topic one Type",
				cfg: TopicConfig{
					types: NewTypes(Topic{}),
				},
			}, wantTy: Types{
				"Topic": reflect.TypeOf(Topic{}),
			},
		},
		{name: "Topic three Types",
			fields: fields{
				name: "Topic three Types",
				cfg: TopicConfig{
					types: NewTypes("", 41, Topic{}),
				},
			}, wantTy: Types{
				"Topic":  reflect.TypeOf(Topic{}),
				"int":    reflect.TypeOf(42),
				"string": reflect.TypeOf(""),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t, err := NewTopic(tt.fields.name, tt.fields.cfg)
			log.Error(err)
			if gotTy := t.Types(); !reflect.DeepEqual(gotTy, tt.wantTy) {
				t1.Errorf("Types() = %v, want %v", gotTy, tt.wantTy)
			}
		})
	}
}
