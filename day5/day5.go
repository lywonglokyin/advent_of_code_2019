package main

import (
	"fmt"
	"strings"

	"github.com/lywonglokyin/advent_of_code_2019/intcode"
	"github.com/lywonglokyin/advent_of_code_2019/utils"
)

func main() {
	filePath := "./day5"
	lines := utils.ReadFile(filePath)

	commandsText := strings.Split(lines[0], ",")
	commands, err := utils.SliceAtoi(commandsText)
	if err != nil {
		panic(err)
	}

	commandsPart1 := make([]int64, len(commands))

	copy(commandsPart1, commands)
	commandsPart2 := commands

	fmt.Println(part1(commandsPart1))
	fmt.Println(part2(commandsPart2))
}

func part1(commands []int64) int64 {
	inputCh := make(chan int64, 10)
	inputCh <- 1
	interpreter := intcode.NewIntcodeInterpreter(commands, inputCh, 10)
	go interpreter.Execute()
	var last int64
	for i := range interpreter.OutputCh() {
		last = i
	}
	return last
}

func part2(commands []int64) int64 {
	inputCh := make(chan int64, 10)
	inputCh <- 5
	interpreter := intcode.NewIntcodeInterpreter(commands, inputCh, 10)
	interpreter.Execute()
	return <-interpreter.OutputCh()
}
