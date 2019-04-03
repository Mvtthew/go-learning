package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	var tab [10]int

	for i := 0; i < 10; i++ {
		tab[i] = random.Intn(10)
	}

	fmt.Println("Before")
	fmt.Println(tab)

	for i := 0; i < 10; i++ {

		for j := 1; j < 10-i; j ++ {

			if tab[j-1] > tab[j] {
				tmp := tab[j-1]
				tab[j-1] = tab[j]
				tab[j] = tmp
			}

		}

	}

	fmt.Println("After")
	fmt.Println(tab)

}
