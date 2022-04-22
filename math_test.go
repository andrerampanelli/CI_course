package main

import "testing"

func TetSum(t *testing.T) {
	total := sum(15, 15)

	if total != 30 {
		t.Errorf("Expected %d; Received %d", 30, total)
	}
}
