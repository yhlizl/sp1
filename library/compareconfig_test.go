package library

import (
	"reflect"
	"testing"
)

func TestCompareConfig(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]map[string]interface{}
		want1 map[string]bool
	}{
		// TODO: Add test cases.
		{"", args{"/Users/royale/Documents/compare"}, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CompareConfig(tt.args.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareConfig() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CompareConfig() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
