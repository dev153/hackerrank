package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readline(reader *bufio.Reader) (string, bool) {
	var aByte byte
	var bytesArray []byte
	var err error
	eof := false
	line := ""

	for {
		aByte, err = reader.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				eof = true
				break
			}
		} else {
			bytesArray = append(bytesArray, aByte)
			if aByte == 10 {
				break
			}
		}
	}
	line = string(bytesArray)
	line = strings.Trim(line, "\n")
	return line, eof
}

func main() {
	var line string
	reader := bufio.NewReader(os.Stdin)
	// Read s and t
	line, _ = readline(reader)
	var values []string
	values = strings.Split(line, " ")
	s, _ := strconv.Atoi(values[0])
	t, _ := strconv.Atoi(values[1])

	// Read a and b
	values = make([]string, 0)
	line, _ = readline(reader)
	values = strings.Split(line, " ")
	a, _ := strconv.Atoi(values[0])
	b, _ := strconv.Atoi(values[1])

	// Read m and n
	values = make([]string, 0)

	line, _ = readline(reader)
	values = strings.Split(line, " ")
	m, _ := strconv.Atoi(values[0])
	n, _ := strconv.Atoi(values[1])

	// Read apple distances
	values = make([]string, 0)
	line, _ = readline(reader)
	values = strings.Split(line, " ")
	var applesDistances []int
	for i := 0; i < m; i++ {
		tmp, _ := strconv.Atoi(values[i])
		applesDistances = append(applesDistances, tmp)
	}

	// Read orange distances
	values = make([]string, 0)
	line, _ = readline(reader)
	values = strings.Split(line, " ")
	var orangeDistances []int
	for i := 0; i < n; i++ {
		tmp, _ := strconv.Atoi(values[i])
		orangeDistances = append(orangeDistances, tmp)
	}

	// Find number of apples that will fall in Sam's house.
	applesCount := 0
	for i := 0; i < len(applesDistances); i++ {
		tmp := a + applesDistances[i]
		if tmp >= s && tmp <= t {
			applesCount++
		}
	}
	// Find number of oranges that will fall in Sam's house.
	orangesCount := 0
	for i := 0; i < len(orangeDistances); i++ {
		tmp := b + orangeDistances[i]
		if tmp >= s && tmp <= t {
			orangesCount++
		}
	}
	fmt.Println(applesCount)
	fmt.Println(orangesCount)
}
