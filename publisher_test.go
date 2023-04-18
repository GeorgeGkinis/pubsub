package pubsub

import (
	"reflect"
	"testing"
)

func TestPublisher_CheckType(t *testing.T) {
	type fields struct {
		name          string
		subscriptions Subscriptions
	}
	type args struct {
		topicName TopicName
		checkMsg  interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantTypeOk bool
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Publisher{
				name:          tt.fields.name,
				subscriptions: tt.fields.subscriptions,
			}
			gotTypeOk, err := p.CheckType(tt.args.topicName, tt.args.checkMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTypeOk != tt.wantTypeOk {
				t.Errorf("CheckType() gotTypeOk = %v, want %v", gotTypeOk, tt.wantTypeOk)
			}
		})
	}
}

func TestPublisher_GetSubscriptions(t *testing.T) {
	type fields struct {
		name          string
		subscriptions Subscriptions
	}
	tests := []struct {
		name   string
		fields fields
		want   Subscriptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Publisher{
				name:          tt.fields.name,
				subscriptions: tt.fields.subscriptions,
			}
			if got := p.GetSubscriptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscriptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublisher_Name(t *testing.T) {
	type fields struct {
		name          string
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
			p := Publisher{
				name:          tt.fields.name,
				subscriptions: tt.fields.subscriptions,
			}
			if got := p.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
