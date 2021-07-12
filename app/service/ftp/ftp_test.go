package ftp

import (
	"testing"

	"github.com/dutchcoders/goftp"
)

func TestGetFromFTP(t *testing.T) {
	c := ConnnectFTP("192.168.0.4", "23", "Royale", "hanshans")
	type args struct {
		c          *goftp.FTP
		remotepath string
		localpath  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		//TODO: Add test cases.
		{
			"test1",
			args{c, "loader/sp1_1/sp1.ini", "/Users/royale/Documents/test.txt"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetFromFTP(tt.args.c, tt.args.remotepath, tt.args.localpath); (err != nil) != tt.wantErr {
				t.Errorf("GetFromFTP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
