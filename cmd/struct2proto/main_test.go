package main

import "testing"

func Test_getgoal(t *testing.T) {
	type args struct {
		last  float32
		start float32
		end   float32
		goal  float32
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{4, 8, 8, 5}},
		{"", args{8, 16, 16, 9}},
		{"", args{252, 252, 252, 256}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getgoal(tt.args.last, tt.args.start, tt.args.end, tt.args.goal)
		})
	}
}
