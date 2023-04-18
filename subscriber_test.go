package pubsub

import (
	"reflect"
	"testing"
)

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSubscriber(tt.args.name, tt.args.handlers, tt.args.subscriptions...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubscriber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriber_AddHandlers(t *testing.T) {
	type fields struct {
		name          string
		ch            chan interface{}
		handlers      map[reflect.Type]HandlerFunc
		subscriptions Subscriptions
	}
	type args struct {
		handlers Handlers
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Subscriber{
				name:          tt.fields.name,
				ch:            tt.fields.ch,
				handlers:      tt.fields.handlers,
				subscriptions: tt.fields.subscriptions,
			}
			h.AddHandlers(tt.args.handlers)
		})
	}
}

func TestSubscriber_Channel(t *testing.T) {
	type fields struct {
		name          string
		ch            chan interface{}
		handlers      map[reflect.Type]HandlerFunc
		subscriptions Subscriptions
	}
	tests := []struct {
		name   string
		fields fields
		want   chan interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Subscriber{
				name:          tt.fields.name,
				ch:            tt.fields.ch,
				handlers:      tt.fields.handlers,
				subscriptions: tt.fields.subscriptions,
			}
			if got := h.Channel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Channel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriber_GetSubscriptions(t *testing.T) {
	type fields struct {
		name          string
		ch            chan interface{}
		handlers      map[reflect.Type]HandlerFunc
		subscriptions Subscriptions
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Topic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Subscriber{
				name:          tt.fields.name,
				ch:            tt.fields.ch,
				handlers:      tt.fields.handlers,
				subscriptions: tt.fields.subscriptions,
			}
			if got := h.GetSubscriptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscriptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriber_Listen(t *testing.T) {
	type fields struct {
		name          string
		ch            chan interface{}
		handlers      map[reflect.Type]HandlerFunc
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
			h := Subscriber{
				name:          tt.fields.name,
				ch:            tt.fields.ch,
				handlers:      tt.fields.handlers,
				subscriptions: tt.fields.subscriptions,
			}
			h.Listen()
		})
	}
}

func TestSubscriber_Name(t *testing.T) {
	type fields struct {
		name          string
		ch            chan interface{}
		handlers      map[reflect.Type]HandlerFunc
		subscriptions Subscriptions
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Subscriber{
				name:          tt.fields.name,
				ch:            tt.fields.ch,
				handlers:      tt.fields.handlers,
				subscriptions: tt.fields.subscriptions,
			}
			if got := h.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriber_Sub(t *testing.T) {
	type fields struct {
		name          string
		ch            chan interface{}
		handlers      map[reflect.Type]HandlerFunc
		subscriptions Subscriptions
	}
	type args struct {
		topic TopicName
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
			h := Subscriber{
				name:          tt.fields.name,
				ch:            tt.fields.ch,
				handlers:      tt.fields.handlers,
				subscriptions: tt.fields.subscriptions,
			}
			if err := h.Sub(tt.args.topic); (err != nil) != tt.wantErr {
				t.Errorf("Sub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
