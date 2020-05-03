package main

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestMakeDBConn(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test Case 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeDBConn()
		})
	}
}
