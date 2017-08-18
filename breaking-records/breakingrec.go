package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var highestScoreBreaks int
	var lowestScoreBreaks int
	var numberOfGames int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberOfGames, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	numberOfGamesLine := scanner.Text()
	scoresStr := strings.Split(numberOfGamesLine, " ")
	var scores []int
	for _, value := range scoresStr {
		score, err := strconv.Atoi(value)
		if err == nil {
			scores = append(scores, score)
		}
	}

	highestScore := scores[0]
	lowestScore := scores[0]

	for i := 1; i < numberOfGames; i++ {
		if scores[i] > highestScore {
			highestScoreBreaks++
			highestScore = scores[i]
		}
		if scores[i] < lowestScore {
			lowestScoreBreaks++
			lowestScore = scores[i]
		}
	}
	fmt.Println(highestScoreBreaks, lowestScoreBreaks)
}
