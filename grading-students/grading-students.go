package main

import "fmt"
import "strconv"

func roundedGrade(grade int) int {
	if grade < 38 {
		return grade
	}

	i := grade % 5
	nextMultipleOfFive := grade + (5 - i)
	diff := nextMultipleOfFive - grade

	if diff < 3 {
		return nextMultipleOfFive
	}

	return grade
}

func main() {
	var line string
	var lastError error
	fmt.Scanf("%s", &line)
	var n int
	var m int
	n, lastError = strconv.Atoi(line)
	grades := make([]int, n)
	if lastError == nil {
		for i := 0; i < n; i++ {
			fmt.Scanf("%s", &line)
			m, lastError = strconv.Atoi(line)
			if lastError == nil {
				grades[i] = m
			}
		}
	}
	for _, value := range grades {
		fmt.Println(roundedGrade(value))
	}
}
