package main

import (
	"strconv"
	"fmt"
)

func main() {

	var x string
	x = "97+3*54-2*-"

	var tab [30]int

	k := 0

	for i := 0; i < len(x); i++ {

		number, err := strconv.Atoi(string(x[i]))

		if err != nil { // nie numer

			b := tab[k-1]
			a := tab[k-2]
			k -= 2

			switch string(x[i]) {

			case "+":
				tab[k] = a + b
				break

			case "-":
				tab[k] = a - b
				break

			case "*":
				tab[k] = a * b
				break

			}

		} else {
			// number
			tab[k] = number
		}

		k++

	}

	fmt.Println(x)
	fmt.Println(tab[0])

}
