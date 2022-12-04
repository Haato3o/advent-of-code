package main

import (
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int64
	End   int64
}

func (r Range) Intersects(other Range) bool {
	return (r.Start >= other.Start && r.Start <= other.End) || (r.End >= other.Start && r.End <= other.End)
}

func toInt(str string) int64 {
	result, _ := strconv.ParseInt(str, 10, 64)
	return result
}

func toElfsRange(sections []string) (Range, Range) {
	section1 := strings.Split(sections[0], "-")
	section2 := strings.Split(sections[1], "-")

	return Range{
			Start: toInt(section1[0]),
			End:   toInt(section1[1]),
		}, Range{
			Start: toInt(section2[0]),
			End:   toInt(section2[1]),
		}
}

func main() {
	rawInput, _ := os.ReadFile("input")
	input := string(rawInput)
	lines := strings.Split(input, "\n")

	intersections := 0
	for _, line := range lines {
		sections := strings.Split(line, ",")
		elf1, elf2 := toElfsRange(sections)

		if elf1.Intersects(elf2) || elf2.Intersects(elf1) {
			intersections++
		}
	}

	println(intersections)
}
