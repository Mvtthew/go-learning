package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {

	file, _ := os.Open("file.txt")

	a := bufio.NewScanner(file)

	for a.Scan() {

		line := strings.Split(a.Text(), " ")

		var sum int

		for _, item := range line {
			value, _ := strconv.Atoi(item)
			sum += value
		}

		println(sum)

	}

}
