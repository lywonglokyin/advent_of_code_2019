package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lywonglokyin/advent_of_code_2019/intcode"

	"github.com/lywonglokyin/advent_of_code_2019/utils"
)

func main() {
	filePath := "./day7"
	lines := utils.ReadFile(filePath)

	commandsText := strings.Split(lines[0], ",")
	commands, err := utils.SliceAtoi(commandsText)
	if err != nil {
		panic(err)
	}

	commandsPart1 := make([]int64, len(commands))

	copy(commandsPart1, commands)
	//commandsPart2 := commands

	fmt.Println(part1(commandsPart1))
}

func part1(commands []int64) int64 {
	max := int64(math.Inf(-1))
	perms := utils.Permutations(5)
	for _, perm := range perms {
		var previousOutput int64 = 0
		for _, val := range perm {
			commandsCopy := make([]int64, len(commands))
			copy(commandsCopy, commands)

			inputCh := make(chan int64, 10)
			intepreter := intcode.NewIntcodeInterpreter(commandsCopy, inputCh, 10)
			inputCh <- int64(val)
			inputCh <- previousOutput
			intepreter.Execute()
			previousOutput = <-intepreter.OutputCh()
		}
		if previousOutput > max {
			max = previousOutput
		}
	}
	return max
}
