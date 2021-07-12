package library

import (
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	p := make(map[string]interface{})
	p["test"] = true
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{{
		"",
		args{"/Users/royale/Documents/test.toml"},
		p,
	},
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadConfig(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
