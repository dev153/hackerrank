package main

import "fmt"

func intAbs(num int) int {
	if num >= 0 {
		return num
	}
	return (-1 * num)
}

func main() {
	fmt.Println("Magic Square...")
}
