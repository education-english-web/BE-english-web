package eventfactory

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/education-english-web/BE-english-web/pkg/notifier/event"
)

func TestNewEventFactory(t *testing.T) {
	type args struct {
		events []event.Event
	}

	tests := []struct {
		name string
		args args
		want EventFactory
	}{
		{
			name: "success",
			args: args{
				events: []event.Event{},
			},
			want: &eventFactory{
				mEventByName: map[event.Name]event.Event{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventFactory(tt.args.events...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEventConfig(t *testing.T) {
	tests := []struct {
		name string
		want EventMapping
	}{
		{
			name: "success",
			want: globalConfig,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEventConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadEventConfig(t *testing.T) {
	t.Parallel()

	t.Run("read config file failed", func(t *testing.T) {
		backupConfig := globalConfig
		defer func() {
			globalConfig = backupConfig
		}()

		filePath := ""

		err := errors.New("open : no such file or directory")
		wantErr := fmt.Errorf("error while reading config file: %w", err)

		gotErr := LoadEventConfig(filePath)
		if gotErr == nil || gotErr.Error() != wantErr.Error() {
			t.Errorf("TestLoadEventConfig error mismatch:\ngot: %v\nwant: %v", gotErr, wantErr)

			return
		}
	})

	t.Run("unmarshal data failed", func(t *testing.T) {
		backupConfig := globalConfig
		defer func() {
			globalConfig = backupConfig
		}()

		filePath := "../../../assets/box.pdf"

		err := errors.New("yaml: found unexpected non-alphabetical character")
		wantErr := fmt.Errorf("error while unmarshaling data: %w", err)

		gotErr := LoadEventConfig(filePath)
		if gotErr == nil || gotErr.Error() != wantErr.Error() {
			t.Errorf("TestLoadEventConfig error mismatch:\ngot: %v\nwant: %v", gotErr, wantErr)

			return
		}
	})

	t.Run("success", func(t *testing.T) {
		backupConfig := globalConfig

		defer func() {
			globalConfig = backupConfig
		}()

		filePath := "../../../assets/event_config.yaml"

		gotErr := LoadEventConfig(filePath)
		if gotErr != nil {
			t.Errorf("TestLoadEventConfig error mismatch:\ngot: %s\nwant: nil", gotErr)
		}
	})
}

func Test_eventFactory_GetEventByName(t *testing.T) {
	type fields struct {
		mEventByName map[event.Name]event.Event
	}

	type args struct {
		name event.Name
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   event.Event
	}{
		{
			name: "success",
			fields: fields{
				mEventByName: map[event.Name]event.Event{
					event.EventNameInternalUserLogin: event.NewEventInternalUserLogin("", "", nil),
				},
			},
			args: args{
				name: event.EventNameInternalUserLogin,
			},
			want: event.NewEventInternalUserLogin("", "", nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventFactory{
				mEventByName: tt.fields.mEventByName,
			}

			if got := e.GetEventByName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventFactory.GetEventByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
