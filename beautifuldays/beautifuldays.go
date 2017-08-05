package main

import "fmt"
import "strconv"

func intAbs(num int) int {
	if num >= 0 {
		return num
	}
	return (-1 * num)
}

func intReversed(num int) int {
	var numStr = strconv.Itoa(num)
	buffer := make([]byte, len(numStr))
	j := 0
	for i := len(numStr) - 1; i >= 0; i-- {
		buffer[j] = numStr[i]
		j++
	}
	reversedNumStr := string(buffer)
	reversedNum, _ := strconv.Atoi(reversedNumStr)
	return reversedNum
}

func isBeautifulDay(num int, special int) bool {
	var reversedNum = intReversed(num)
	var absOfDiff = intAbs(num - reversedNum)
	if absOfDiff%special == 0 {
		return true
	}
	return false
}

func main() {
	var begin int
	var end int
	var special int
	fmt.Scanf("%d %d %d", &begin, &end, &special)
	var beautifulDaysCount int
	for i := begin; i <= end; i++ {
		if isBeautifulDay(i, special) == true {
			beautifulDaysCount++
		}
	}
	fmt.Println(beautifulDaysCount)
}
