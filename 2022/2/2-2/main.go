package main

import (
	"os"
	"strings"
)

const (
	ROCK     int = 1
	PAPER    int = 2
	SCISSORS int = 3

	OP_DRAW int = 1
	OP_WIN  int = 2
	OP_LOSE int = 3
)

var opponent = map[string]int{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var commands = map[string]int{
	"X": OP_LOSE,
	"Y": OP_DRAW,
	"Z": OP_WIN,
}

var rules = map[int]int{
	ROCK:     SCISSORS,
	SCISSORS: PAPER,
	PAPER:    ROCK,
}

func findWeakness(hand int) int {
	for i := 0; i < 2; i++ {
		hand = rules[hand]
	}
	return hand
}

func main() {
	inputRaw, _ := os.ReadFile("input")
	input := string(inputRaw)
	cases := strings.Split(input, "\n")

	var points = 0

	for _, caseRaw := range cases {
		parsedCase := strings.Split(caseRaw, " ")
		opponentPlay := opponent[parsedCase[0]]
		localPlay := commands[parsedCase[1]]

		switch localPlay {
		case OP_LOSE:
			points += rules[opponentPlay]
			break
		case OP_DRAW:
			points += 3 + opponentPlay
			break
		case OP_WIN:
			points += 6 + findWeakness(opponentPlay)
			break
		}
	}

	println(points)
}
