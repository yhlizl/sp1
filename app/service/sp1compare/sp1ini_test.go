package sp1compare

import (
	"reflect"
	"testing"
)

func Test_downloadConfig(t *testing.T) {
	type args struct {
		fab   []string
		phase []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"", args{[]string{"F18A", "F18B"}, []string{"P1", "P5"}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := downloadConfig(tt.args.fab, tt.args.phase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("downloadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
