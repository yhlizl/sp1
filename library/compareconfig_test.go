package library

import (
	"reflect"
	"testing"

	"github.com/gogf/gf/frame/g"
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

func TestCompareConfig(t *testing.T) {
	aa := g.Map{}
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
		want g.Map
	}{
		// TODO: Add test cases.
		{"", args{"/Users/royale/Documents/compare"}, aa},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareConfig(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
