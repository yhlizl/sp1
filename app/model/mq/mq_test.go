package mq

import (
	"testing"
)

func TestToggleMQ(t *testing.T) {

	type args struct {
		item string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{"sp1_ini"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ToggleMQ(tt.args.item)
		})
	}
}

func TestGetStatus(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
		{"", args{"sp1_ini"}, "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetStatus(tt.args.item)
			if got != tt.want {
				t.Errorf("GetStatus() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
