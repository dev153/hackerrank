package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// file, _ := os.Open("birthday-cake-candles.txt")
	// defer file.Close()
	// reader := bufio.NewReader(file)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	var numberOfCandles int
	if err != nil {
		log.Fatal(err)
	} else {
		numberOfCandles, err = strconv.Atoi(strings.Trim(str, "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}

	str, err = reader.ReadString('\n')
	candlesHeightsTokens := strings.Split(str, " ")
	var candlesHeights []int
	for _, candleHeight := range candlesHeightsTokens {
		convHeight, _ := strconv.Atoi(candleHeight)
		candlesHeights = append(candlesHeights, convHeight)
	}

	max := candlesHeights[0]
	count := 1

	for i := 1; i < numberOfCandles; i++ {
		if candlesHeights[i] > max {
			max = candlesHeights[i]
			count = 1
		} else if candlesHeights[i] == max {
			count++
		}
	}
	fmt.Println(count)
}
