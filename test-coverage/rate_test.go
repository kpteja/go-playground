package main

import "testing"

func TestRate(t *testing.T) {
	type args struct {
		rate int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{0},
			want: "Worst",
		},
		{
			name: "Test Case 2",
			args: args{5},
			want: "Superb",
		},
		{
			name: "Test Case 3",
			args: args{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rate(tt.args.rate); got != tt.want {
				t.Errorf("Rate() = %v, want %v", got, tt.want)
			}
		})
	}
}
