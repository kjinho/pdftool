package utils

import (
	"testing"
)

func TestGenerateFmtString(t *testing.T) {
	type args struct {
		prefix    string
		separator string
		padding   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"MRDK", "-", 20},
			"MRDK-%020d",
		},
		{
			"2",
			args{"", "", 10},
			"%010d",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateFmtString(tt.args.prefix, tt.args.separator, tt.args.padding); got != tt.want {
				t.Errorf("GenerateFmtString() = %v, want %v", got, tt.want)
			}
		})
	}
}
