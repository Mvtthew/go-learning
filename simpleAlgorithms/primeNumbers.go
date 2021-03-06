package main

import (
	"fmt"
	"math"
)

func main() {

	i := 1

	for {

		i = nextPrimary(i + 1)

		fmt.Println("Next prime is: ", i, "(click / hold enter for next prime)")
		fmt.Scanln()

	}

}

func nextPrimary(n int) int {

	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			n++
			i = 2
		}
	}

	return n

}
