package main

import (
	"os"
	"strings"
)

const (
	ROCK     int = 1
	PAPER    int = 2
	SCISSORS int = 3
)

var opponent = map[string]int{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var player = map[string]int{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

var rules = map[int]int{
	ROCK:     SCISSORS,
	SCISSORS: PAPER,
	PAPER:    ROCK,
}

func main() {
	inputRaw, _ := os.ReadFile("input")
	input := string(inputRaw)
	cases := strings.Split(input, "\n")

	var points = 0

	for _, caseRaw := range cases {
		parsedCase := strings.Split(caseRaw, " ")
		opponentPlay := opponent[parsedCase[0]]
		localPlay := player[parsedCase[1]]

		if opponentPlay == localPlay {
			points += 3
		} else if rules[localPlay] == opponentPlay {
			points += 6
		}

		points += localPlay
	}

	println(points)
}
