package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var line string
	var numOfQueries int
	var results []string
	fmt.Scanf("%s", &line)
	numOfQueries, lastError := strconv.Atoi(line)
	if lastError == nil {
		for i := 0; i < numOfQueries; i++ {
			var output string
			var catAPos float64
			var catBPos float64
			var mousePos float64
			fmt.Scanf("%f %f %f", &catAPos, &catBPos, &mousePos)
			var catAMouseDiff = math.Abs((catAPos - mousePos))
			var catBMouseDiff = math.Abs((catBPos - mousePos))
			if catAMouseDiff == catBMouseDiff {
				output = "Mouse C"
			} else {
				output += "Cat "
				if catAMouseDiff < catBMouseDiff {
					output += "A"
				} else {
					output += "B"
				}
			}
			results = append(results, output)
		}
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
