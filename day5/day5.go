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
	fakestdin, err := utils.NewFakeStdin("1\n")
	defer fakestdin.Close()
	if err != nil {
		panic(err)
	}
	interpreter := intcode.NewIntcodeInterpreter(commands)
	interpreter.Execute()
	return interpreter.At(0)
}

func part2(commands []int64) string {
	fakestdin, err := utils.NewFakeStdin("5\n")
	if err != nil {
		panic(err)
	}
	defer fakestdin.Close()
	fakestdout, err := utils.NewFakeStdout()
	if err != nil {
		panic(err)
	}
	interpreter := intcode.NewIntcodeInterpreter(commands)
	interpreter.Execute()
	return fakestdout.ReadandClose()
}
