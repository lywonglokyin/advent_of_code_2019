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
		{[]int64{1002, 4, 3, 4, 33}, []int64{1002, 4, 3, 4, 99}},
	}
	for _, fixture := range fixtures {
		interpreter := intcode.NewIntcodeInterpreter(fixture.initial)
		interpreter.Execute()
		reflect.DeepEqual(fixture.expected, interpreter.Intcodes)
	}
}
