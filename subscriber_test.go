package pubsub

import (
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

var (
	stringTopic, _             = NewTopic("stringTopic", TopicConfig{})
	fa             HandlerFunc = simpleConsoleIntHandler
)

func simpleConsoleStringHandler(msg interface{}) (err error) {
	log.Debugf("simpleConsoleStringHandler received message: %v", msg)
	return
}

func simpleConsoleIntHandler(msg interface{}) (err error) {
	log.Debugf("simpleConsoleIntHandler received message: %v", msg)
	return
}

func TestNewSubscriber(t *testing.T) {

	type args struct {
		name          string
		handlers      Handlers
		subscriptions []*Topic
	}
	tests := []struct {
		name string
		args args
		want *Subscriber
	}{
		{name: "Subscriber no Topic",
			args: args{
				name:          "Subscriber no Topic",
				handlers:      Handlers{"string": &fa},
				subscriptions: []*Topic{},
			}, want: &Subscriber{
				name:          "Subscriber no Topic",
				handlers:      Handlers{"string": &fa},
				subscriptions: []*Topic{},
			}},
		{name: "Subscriber One Topic",
			args: args{
				name:          "Subscriber One Topic",
				handlers:      Handlers{"string": &fa},
				subscriptions: []*Topic{stringTopic},
			}, want: &Subscriber{
				name:          "Subscriber One Topic",
				handlers:      Handlers{"string": &fa},
				subscriptions: []*Topic{stringTopic},
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewSubscriber(tt.args.name, tt.args.handlers, tt.args.subscriptions...)
			tt.want.ch = got.ch
			equal := reflect.DeepEqual(got, tt.want)
			if !equal {
				t.Errorf("Equal: %v,NewSubscriber()\n got: %v\nwant: %v", equal, got, tt.want)
			}
		})
	}
}
func TestSubscriber_AddHandler(t *testing.T) {

	type args struct {
		typeOf  interface{}
		handler *HandlerFunc
	}
	tests := []struct {
		name    string
		fields  *Subscriber
		args    args
		wantErr bool
	}{
		{name: "Add Handler no error",
			fields: NewSubscriber("Sub1", nil),
			args: args{
				typeOf:  "string",
				handler: &fa,
			}, wantErr: false},
		{name: "Add Handler with error",
			fields: NewSubscriber("Sub1", nil),
			args: args{
				typeOf:  nil,
				handler: nil,
			}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			if err := s.AddHandler(tt.args.typeOf, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("AddHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubscriber_Listen(t *testing.T) {
	type fields struct {
		name          string
		listening     bool
		ch            chan interface{}
		handlers      Handlers
		subscriptions Subscriptions
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subscriber{
				name:          tt.fields.name,
				listening:     tt.fields.listening,
				ch:            tt.fields.ch,
				handlers:      tt.fields.handlers,
				subscriptions: tt.fields.subscriptions,
			}
			s.Listen()
		})
	}
}

func TestSubscriber_Sub(t *testing.T) {
	type fields struct {
		name          string
		listening     bool
		ch            chan interface{}
		handlers      Handlers
		subscriptions Subscriptions
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
			s := &Subscriber{
				name:          tt.fields.name,
				listening:     tt.fields.listening,
				ch:            tt.fields.ch,
				handlers:      tt.fields.handlers,
				subscriptions: tt.fields.subscriptions,
			}
			if err := s.Sub(tt.args.topic); (err != nil) != tt.wantErr {
				t.Errorf("Sub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
