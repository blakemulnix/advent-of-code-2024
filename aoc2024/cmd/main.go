package main

import (
	"aoc2024/internal/challenges/day1"
	"aoc2024/internal/format"
	"aoc2024/internal/timing"
)

func main() {
	format.PrintHeader()

	days := []struct {
		part    int
		day     int
		solver  func() any
	}{
		{1, 1, func() any { return day1.Day1Part1() }},
		{1, 2, func() any { return day1.Day1Part2() }},
	}

	for _, d := range days {
		result := timing.Track(d.solver)
		format.PrintAnswer(d.day, d.part, result.Result, timing.FormatDuration(result.Duration))
	}

	format.PrintFooter()
}
