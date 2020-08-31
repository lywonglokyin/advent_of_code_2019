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
	commandsPart2 := commands

	fmt.Println(part1(commandsPart1))
	fmt.Println(part2(commandsPart2))
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

func part2(commands []int64) int64 {
	max := int64(math.Inf(-1))
	perms := utils.Permutations(5)
	for _, perm := range perms {

		thruttleCh := make(chan int64) // channel for storing final thruttle

		commandCopies := make([][]int64, 5)
		for i := 0; i < 5; i++ {
			commandCopies[i] = make([]int64, len(commands))
			copy(commandCopies[i], commands)
		}

		ampAInputCh := make(chan int64, 10)
		ampA := intcode.NewIntcodeInterpreter(commandCopies[0], ampAInputCh, 10)
		ampB := intcode.NewIntcodeInterpreter(commandCopies[1], ampA.OutputCh(), 10)
		ampC := intcode.NewIntcodeInterpreter(commandCopies[2], ampB.OutputCh(), 10)
		ampD := intcode.NewIntcodeInterpreter(commandCopies[3], ampC.OutputCh(), 10)
		ampE := intcode.NewIntcodeInterpreter(commandCopies[4], ampD.OutputCh(), 10)

		go func() {
			var last int64 = 0
			for output := range ampE.OutputCh() {
				last = output
				ampAInputCh <- last
			}
			thruttleCh <- last
		}()

		ampA.InputCh() <- int64(perm[0] + 5)
		ampB.InputCh() <- int64(perm[1] + 5)
		ampC.InputCh() <- int64(perm[2] + 5)
		ampD.InputCh() <- int64(perm[3] + 5)
		ampE.InputCh() <- int64(perm[4] + 5)

		ampA.InputCh() <- 0

		go ampA.Execute()
		go ampB.Execute()
		go ampC.Execute()
		go ampD.Execute()
		go ampE.Execute()

		thruttle := <-thruttleCh
		if thruttle > max {
			max = thruttle
		}
	}
	return max
}
