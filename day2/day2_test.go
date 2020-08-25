package main

import (
	"reflect"
	"testing"

	"github.com/lywonglokyin/advent_of_code_2019/intcode"
)

type Fixture struct {
	initial  []int64
	expected []int64
}

func TestIntepreter(t *testing.T) {
	fixtures := []Fixture{
		{[]int64{1, 0, 0, 0, 99}, []int64{2, 0, 0, 0, 99}},
		{[]int64{2, 3, 0, 0, 99}, []int64{2, 3, 0, 6, 99}},
		{[]int64{2, 4, 4, 5, 99, 0}, []int64{2, 4, 4, 5, 99, 9801}},
		{[]int64{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int64{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, fixture := range fixtures {
		interpreter := intcode.NewIntcodeInterpreter(fixture.initial)
		interpreter.Execute()
		reflect.DeepEqual(fixture.expected, interpreter.Intcodes)
	}
}
