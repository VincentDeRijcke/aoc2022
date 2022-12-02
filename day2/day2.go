package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var scores1 = map[string]int{
	"A X": 1 + 3, // Rock vs Rock
	"A Y": 2 + 6, // Rock vs Paper🏆
	"A Z": 3 + 0, // Rock🏆 vs Scissor
	"B X": 1 + 0, // Paper🏆 vs Rock
	"B Y": 2 + 3, // Paper vs Paper
	"B Z": 3 + 6, // Paper vs Scissor🏆
	"C X": 1 + 6, // Scissor vs Rock🏆
	"C Y": 2 + 0, // Scissor🏆 vs Paper
	"C Z": 3 + 3, // Scissor vs Scissor
}
var scores2 = map[string]int{
	"A X": 0 + 3, // Lose vs Rock (Scissor)
	"A Y": 3 + 1, // Draw vs Rock (Rock)
	"A Z": 6 + 2, // Win vs Rock  (Paper)
	"B X": 0 + 1, // Lose vs Paper (Rock)
	"B Y": 3 + 2, // Draw vs Paper (Paper)
	"B Z": 6 + 3, // Win vs Paper (Scissor)
	"C X": 0 + 2, // Lose vs Scissor (Paper)
	"C Y": 3 + 3, // Draw vs Scissor (Scissor)
	"C Z": 6 + 1, // Win vs Scissor (Rock)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func scoreGame(game io.Reader) (int, int) {
	var scanner = bufio.NewScanner(game)

	part1Score := 0
	part2Score := 0
	for scanner.Scan() {
		round := strings.TrimSpace(scanner.Text())
		if round != "" {
			part1Score += scores1[round]
			part2Score += scores2[round]
		}
	}

	return part1Score, part2Score
}

func main() {
	fmt.Println("Reading input")
	var content, err = os.ReadFile("./day2/input.txt")
	check(err)
	part1, part2 := scoreGame(bytes.NewReader(content))

	fmt.Println("Total Score Part 1:", part1)
	fmt.Println("Total Score Part 2:", part2)
}
