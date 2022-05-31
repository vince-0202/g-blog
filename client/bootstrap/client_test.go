package bootstrap

import (
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	type args struct {
		options ClientOptions
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{name: "new server with port",
			args: args{
				options: ClientOptions{Port: 1234},
			},
			want: &Server{Port: 1234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewServer(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
