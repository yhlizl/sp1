package library

import (
	"reflect"
	"testing"
)

func TestCompareConfigList(t *testing.T) {
	type args struct {
		root     string
		filelist []string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]map[string]interface{}
		want1 map[string]bool
	}{
		// TODO: Add test cases.
		{"", args{"/Users/royale/go/src/sp1/filesystem/", []string{"/Users/royale/go/src/sp1/filesystem/F18AF18bP1P5_asp1_sp1_2_sp1.ini", "/Users/royale/go/src/sp1/filesystem/F18AF18bP1P5_esp1_sp1_2_sp1.ini"}}, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CompareConfigList(tt.args.root, tt.args.filelist)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareConfigList() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CompareConfigList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
