package main

import (
	"fmt"
	"strings"

	"github.com/lywonglokyin/advent_of_code_2019/intcode"
	"github.com/lywonglokyin/advent_of_code_2019/utils"
)

func main() {
	filePath := "./day2"
	lines := utils.ReadFile(filePath)
	if len(lines) != 1 {
		panic("Should only have 1 line...")
	}

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
	commands[1] = 12
	commands[2] = 2
	interpreter := intcode.NewIntcodeInterpreter(commands)
	interpreter.Execute()
	return interpreter.At(0)
}

func part2(commands []int64) int64 {
	var target int64 = 19690720
	found := false
	var targetNoun, targetVerb int64
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			commandsCopy := make([]int64, len(commands))
			copy(commandsCopy, commands)
			commandsCopy[1] = int64(noun)
			commandsCopy[2] = int64(verb)

			interpreter := intcode.NewIntcodeInterpreter(commandsCopy)
			interpreter.Execute()
			output := interpreter.At(0)
			if output == target {
				found = true
				targetNoun = int64(noun)
				targetVerb = int64(verb)
				break
			}
		}
	}
	if !found {
		panic("No found solution!")
	}
	return 100*targetNoun + targetVerb
}
