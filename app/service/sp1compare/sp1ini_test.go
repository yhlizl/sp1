package sp1compare

import (
	"reflect"
	"testing"
)

func TestCompareini(t *testing.T) {
	tests := []struct {
		name string
		want map[string]map[string]map[string]string
	}{
		// TODO: Add test cases.
		{"", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compareini(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compareini() = %v, want %v", got, tt.want)
			}
		})
	}
}
