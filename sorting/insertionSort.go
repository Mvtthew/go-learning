package main

import (
	"math/rand"
	"time"
	"fmt"
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

		minIndex := i

		for j := i; j < 10; j ++ {

			if tab[j] < tab[minIndex] {
				minIndex = j
			}

		}

		tmp := tab[i]
		tab[i] = tab[minIndex]
		tab[minIndex] = tmp

	}

	fmt.Println("After")
	fmt.Println(tab)

}
