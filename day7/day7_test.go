package main

import (
	"testing"
)

type Fixture struct {
	initial  []int64
	expected int64
}

func TestIntepreter(t *testing.T) {
	fixtures := []Fixture{
		{[]int64{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, 43210},
		{[]int64{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
			101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, 54321},
	}
	for _, fixture := range fixtures {
		max := recuriveMaxThruttle(0, fixture.initial, 1)
		if max != fixture.expected {
			t.Fatal("Incorrect max thruttle")
		}
	}
}
