package logkeeper

import (
	"reflect"
	"testing"
)

func TestGetCloudWatch(t *testing.T) {
	tests := []struct {
		name string
		want LogKeeper
	}{
		{
			name: "success",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCloudWatch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCloudWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
