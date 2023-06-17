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
	p1    = NewPublisher("p1")
	p2    = NewPublisher("p2")
	s1, _ = NewSubscriber("s1", nil, Subscriptions{})
	s2, _ = NewSubscriber("s2", nil, Subscriptions{})
)

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
				Types: make(Types, 0),
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
				Types: make(Types, 0),
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
					subscriptions: Subscriptions{},
				},
				"p2": &Publisher{
					name:          "p2",
					subscriptions: Subscriptions{},
				}},
			cfg: TopicConfig{
				Types: make(Types, 0),
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
				Types: make(Types, 0),
			},
		}},
		{name: "Create Topic with 2 Types", args: args{
			name: "TestNewTopic2Types",
			cfg: TopicConfig{
				Types: Types{
					"string": reflect.TypeOf("type1"),
					"int":    reflect.TypeOf(2),
				},
			},
		}, wantT: &Topic{
			name:        "TestNewTopic2Types",
			subscribers: make(Subscribers, 0),
			publishers:  make(Publishers, 0),
			cfg: TopicConfig{
				Types: Types{
					"string": reflect.TypeOf("type1"),
					"int":    reflect.TypeOf(2),
				},
				TypeSafe: true},
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
		{name: "Topic no Publishers and AllowAddPub true", fields: Topic{
			name:       "AllowAddPub true",
			publishers: make(Publishers, 0),
			cfg: TopicConfig{
				AllowAddPub: true,
			},
		}, args: struct{ pub *Publisher }{pub: p1}, wantErr: false},
		{name: "Topic no Publishers and AllowAddPub false", fields: Topic{
			name:       "AllowAddPub false",
			publishers: make(Publishers, 0),
			cfg: TopicConfig{
				AllowAddPub: false,
			},
		}, args: struct{ pub *Publisher }{pub: p1}, wantErr: true},

		{name: "Topic with Publishers and AllowAddPub true", fields: Topic{
			name: "AllowAddPub true",
			publishers: Publishers{
				"p1": &Publisher{
					name: "p1",
				}},
			cfg: TopicConfig{
				AllowAddPub: true,
			},
		}, args: struct{ pub *Publisher }{pub: p2}, wantErr: false},
		{name: "Topic with Publishers and AllowAddPub false", fields: Topic{
			name: "AllowAddPub false",
			publishers: Publishers{
				"p1": &Publisher{
					name: "p1",
				}},
			cfg: TopicConfig{
				AllowAddPub: false,
			},
		}, args: struct{ pub *Publisher }{pub: p1}, wantErr: true},
		{name: "Add Publisher with same name to topic and allowOverwrite true", fields: Topic{
			name: "allowOverwrite false",
			publishers: Publishers{
				"p1": &Publisher{
					name: "p1",
				}},
			cfg: TopicConfig{
				AllowAddPub:   true,
				AllowOverride: true,
			},
		}, args: struct{ pub *Publisher }{pub: p1}, wantErr: false},
		{name: "Add Publisher with same name to topic and allowOverwrite false", fields: Topic{
			name: "allowOverwrite false",
			publishers: Publishers{
				"p1": &Publisher{
					name: "p1",
				}},
			cfg: TopicConfig{
				AllowAddPub:   true,
				AllowOverride: false,
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
		sub *Subscriber
	}
	tests := []struct {
		name    string
		topic   Topic
		args    args
		wantErr bool
	}{
		{name: "Topic no Subs and AllowOverride true", topic: Topic{
			name:        "Topic no Subs and AllowOverride true",
			subscribers: make(Subscribers, 0),
			cfg:         TopicConfig{AllowOverride: true},
		}, args: args{
			s1,
		}, wantErr: false},
		{name: "Topic no Subs and AllowOverride false", topic: Topic{
			name:        "Topic no Subs and AllowOverride false",
			subscribers: make(Subscribers, 0),
			cfg:         TopicConfig{AllowOverride: false},
		}, args: args{s1}, wantErr: false},
		{name: "Add Subscriber with same name to topic AllowOverride true", topic: Topic{
			name: "Topic same Sub and AllowOverride true",
			subscribers: Subscribers{
				"s1": s1,
			},
			cfg: TopicConfig{AllowOverride: true},
		}, args: args{s1}, wantErr: false},
		{name: "Add Subscriber with same name to topic and AllowOverride false", topic: Topic{
			name: "Topic same Sub and AllowOverride false",
			subscribers: Subscribers{
				"s1": s1,
			},
			cfg: TopicConfig{AllowOverride: false},
		}, args: args{s1}, wantErr: true},
		{name: "Add Subscriber with other name to topic AllowOverride true", topic: Topic{
			name: "Topic same Sub and AllowOverride true",
			subscribers: Subscribers{
				"s1": s1,
			},
			cfg: TopicConfig{AllowOverride: true},
		}, args: args{s2}, wantErr: false},
		{name: "Add Subscriber with other name to topic and AllowOverride false", topic: Topic{
			name: "Topic same Sub and AllowOverride false",
			subscribers: Subscribers{
				"s1": s1,
			},
			cfg: TopicConfig{AllowOverride: false},
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
				TypeSafe: true,
			},
		}, want: true},
		{name: "Topic isTypeSafe false", topic: Topic{
			name:        "Topic isTypeSafe false",
			subscribers: nil,
			publishers:  nil,
			cfg: TopicConfig{
				TypeSafe: false,
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
		handlefunc HandlerFunc
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
		{name: "Publisher exists, AllowAllPublishers true, 2 messages sent", topic: tp{
			name:       "Publisher exists and AllowAllPublishers true",
			publishers: []*Publisher{p1},
			cfg: TopicConfig{
				AllowAllPublishers: true,
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
				msg: []interface{}{"Message: Publisher exists and AllowAllPublishers false", 42},
			}, wantErr: false},
		{name: "Publisher exists, AllowAllPublishers false, 2 messages sent", topic: tp{
			name:       "Publisher exists and AllowAllPublishers false",
			publishers: []*Publisher{p1},
			cfg: TopicConfig{
				AllowAllPublishers: false,
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
				msg: []interface{}{"Message: Publisher exists and AllowAllPublishers false", 42},
			}, wantErr: false},
		{
			name: "Publisher not exists, AllowAllPublishers false, 2 messages sent",
			topic: tp{
				name:       "Publisher not exists, AllowAllPublishers false, 2 messages sent",
				publishers: nil,
				cfg: TopicConfig{
					AllowAllPublishers: false,
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
				msg: []interface{}{"Message: Publisher exists and AllowAllPublishers true", 42},
			}, wantErr: true},
		{
			name: "Publisher not exists, AllowAllPublishers true, 3 messages sent",
			topic: tp{
				name:       "Publisher not exists, AllowAllPublishers true, 3 messages sent",
				publishers: nil,
				cfg: TopicConfig{
					AllowAllPublishers: true,
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
				},
				{
					typeof:     "any",
					handlefunc: simpleConsoleAnyHandler,
				}},
			args: args{
				pub: *p1,
				msg: []interface{}{"Message: Publisher exists and AllowAllPublishers true", 42, 0.2},
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

			s3, err := NewSubscriber("s3", nil, Subscriptions{})
			err = t.AddSub(s3)
			if err != nil {
				log.Error(err)
			}
			for _, v := range tt.handlers {
				_ = s3.AddHandler(v.typeof, &v.handlefunc)
			}
			log.Debugf("Handlers registered: %v", s3.handlers)
			s3.Listen()
			err = t.Pub(&tt.args.pub, tt.args.msg...)
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
		{name: "Set Name, AllowSetName true",
			topic: tp{
				name: "Name Before",
				cfg: TopicConfig{
					AllowSetName: true,
				},
			}, args: args{name: "Name After"},
			wantErr: false},
		{name: "Set Name to empty, AllowSetName true",
			topic: tp{
				name: "Name Before",
				cfg: TopicConfig{
					AllowSetName: true,
				},
			}, args: args{name: ""},
			wantErr: true},
		{name: "Set Name , AllowSetName false",
			topic: tp{
				name: "Name Before",
				cfg: TopicConfig{
					AllowSetName: false,
				},
			}, args: args{name: "Name After"},
			wantErr: true},
		{name: "Set Name to empty, AllowSetName false",
			topic: tp{
				name: "Name Before",
				cfg: TopicConfig{
					AllowSetName: false,
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
		{name: "SetTypeSafe, AllowSetTypeSafe true",
			topic: Topic{
				name: "SetTypeSafe, AllowSetTypeSafe true",
				cfg: TopicConfig{
					AllowSetTypeSafe: true,
				},
			}, args: args{typeSafe: true}, wantErr: false},
		{name: "SetTypeSafe, AllowSetTypeSafe false",
			topic: Topic{
				name: "SetTypeSafe, AllowSetTypeSafe false",
				cfg: TopicConfig{
					AllowSetTypeSafe: false,
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
					AllowSetTypes: true,
				},
			}, args: args{types: []interface{}{
				"42",
				42,
				Topic{},
			},
			}, wantErr: false},
		{name: "SetTypes and AllowSetTypes false",
			topic: tp{
				name: "SetTypes and allowsSetTypes false",
				cfg: TopicConfig{
					AllowSetTypes: false,
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
					AllowSetTypes: true,
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
			log.Debugf("Types: %s", t.cfg.Types)

			tmpTypes := func(types ...interface{}) (t Types) {
				t = make(Types, 0)
				for _, v := range types {
					t[reflect.TypeOf(v).Name()] = reflect.TypeOf(v)
				}
				return
			}
			tmpt := tmpTypes(tt.args.types...)
			b := reflect.DeepEqual(tmpt, t.cfg.Types)
			log.Debugf("DeepEqual: %v,Wanted %s, got %s", b, tmpt, t.cfg.Types)

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
					Types: make(Types, 0),
				},
			}, wantTy: make(Types, 0)},
		{name: "Topic one Type",
			fields: fields{
				name: "Topic one Type",
				cfg: TopicConfig{
					Types: NewTypes(Topic{}),
				},
			}, wantTy: Types{
				"Topic": reflect.TypeOf(Topic{}),
			},
		},
		{name: "Topic three Types",
			fields: fields{
				name: "Topic three Types",
				cfg: TopicConfig{
					Types: NewTypes("", 41, Topic{}),
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
