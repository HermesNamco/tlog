package tlog

import "testing"

func TestInfof(t *testing.T) {
	type args struct {
		template string
		args     string
	}
	arg := make([]string, 0)
	arg = append(arg, "World")

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Normal", args{"Hello %v", "world"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.template, tt.args)
		})
	}
}
