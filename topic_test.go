package pubsub

import (
	"reflect"
	"testing"
)

var (
	p1    = NewPublisher("p1")
	p2    = NewPublisher("p2")
	ct1   = "type1"
	ct2   = 2
	t1    = reflect.TypeOf(ct1)
	t2    = reflect.TypeOf(ct2)
	types = NewTypes(ct1, ct2)
)

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
				types: *types,
			},
		}, wantT: &Topic{
			name: "TestNewTopic2Types",
			cfg: TopicConfig{
				typeSafe: true,
				types: Types{
					"string": reflect.TypeOf("type1"),
					"int":    reflect.TypeOf(2),
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := NewTopic(tt.args.name, tt.args.cfg, tt.args.pubs...); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("NewTopic() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func TestTopic_AddPub(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	type args struct {
		pub *Publisher
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
			if err := t.AddPub(tt.args.pub); (err != nil) != tt.wantErr {
				t1.Errorf("AddPub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopic_AddSub(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	type args struct {
		sub *Subscriber
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
			if err := t.AddSub(tt.args.sub); (err != nil) != tt.wantErr {
				t1.Errorf("AddSub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTopic_IsTypeSafe(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
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
			if got := t.IsTypeSafe(); got != tt.want {
				t1.Errorf("IsTypeSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopic_Name(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   TopicName
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
			if got := t.Name(); got != tt.want {
				t1.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopic_Pub(t1 *testing.T) {
	type fields struct {
		name        TopicName
		subscribers Subscribers
		publishers  Publishers
		cfg         TopicConfig
	}
	type args struct {
		pub Publisher
		msg []interface{}
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
			if err := t.Pub(tt.args.pub, tt.args.msg...); (err != nil) != tt.wantErr {
				t1.Errorf("Pub() error = %v, wantErr %v", err, tt.wantErr)
			}
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
