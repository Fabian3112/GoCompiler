package main

import (
	"strings"
	"testing"
)

func Test_parseAndRun(t *testing.T) {
	type args struct {
		prog string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{"print(1234)"}, "1234"},
		{"simple while", args{"i:=0;while(i<10){i=i+1;print(i)}"}, "12345678910"},
		{"simple while 2", args{"i:=0;while(i<10){i=i+1;print(i)}"}, "12345678910"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.TrimSpace(parseAndRun(tt.args.prog)); got != tt.want {
				t.Errorf("parseAndRun() = %v, want %v", got, tt.want)
			}
		})
	}
}
