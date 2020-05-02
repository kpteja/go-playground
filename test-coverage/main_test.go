package main

import (
	"testing"
)

func TestSize(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{-1},
			want: "negative",
		},
		{
			name: "Test Case 2",
			args: args{0},
			want: "zero",
		},
		{
			name: "Test Case 3",
			args: args{9},
			want: "small",
		},
		{
			name: "Test Case 4",
			args: args{99},
			want: "big",
		},
		{
			name: "Test Case 5",
			args: args{999},
			want: "huge",
		},
		{
			name: "Test Case 6",
			args: args{1001},
			want: "enormous",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Size(tt.args.a); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test Case 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
