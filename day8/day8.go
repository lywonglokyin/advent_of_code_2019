package main

import (
	"fmt"
	"strings"

	"github.com/lywonglokyin/advent_of_code_2019/utils"
)

func main() {
	filePath := "./day8"
	lines := utils.ReadFile(filePath)

	digits := lines[0]

	width := 25
	height := 6

	fmt.Println(part1(digits, width, height))

	image := part2(digits, width, height)
	rowCount := len(image) / width
	for i := 0; i < rowCount; i++ {
		startPos := i * width
		endPos := (i + 1) * width
		for _, pixel := range image[startPos:endPos] {
			if pixel == '0' {
				fmt.Print(" ")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println("")
	}
}

type layer struct {
	digits    string
	zeroCount int
}

func readDigits(digits string, width int, height int) []*layer {
	layerCount := len(digits) / width / height

	layers := make([]*layer, layerCount)

	for i := 0; i < layerCount; i++ {
		startPos := i * (width * height)
		endPos := startPos + (width * height) //exclusive bound

		layerDigits := digits[startPos:endPos]
		zeroCount := strings.Count(layerDigits, "0")
		layers[i] = &layer{
			digits:    digits[startPos:endPos],
			zeroCount: zeroCount,
		}
	}
	return layers
}

func part1(digits string, width int, height int) int {
	layers := readDigits(digits, width, height)

	minZeroCount := layers[0].zeroCount
	minIndex := 0
	for i, layer := range layers {
		if layer.zeroCount < minZeroCount {
			minZeroCount = layer.zeroCount
			minIndex = i
		}
	}

	minLayerOneCount := strings.Count(layers[minIndex].digits, "1")
	minLayerTwoCount := strings.Count(layers[minIndex].digits, "2")
	return minLayerOneCount * minLayerTwoCount
}

func part2(digits string, width int, height int) string {
	layers := readDigits(digits, width, height)
	pixelCount := width * height

	finalImage := make([]byte, pixelCount)

	for i := 0; i < pixelCount; i++ {
		color := parseColor(layers, i)
		finalImage[i] = color
	}
	return string(finalImage[:])
}

func parseColor(layers []*layer, pos int) byte {
	for _, layer := range layers {
		digit := layer.digits[pos] // This implicitly assumes the digits are of bytes only (no unicode shits)
		switch digit {
		case '0':
			return '0'
		case '1':
			return '1'
		case '2':
			continue
		default:
			panic("Unknown pixel value!")
		}
	}
	panic("No color found!")
}
