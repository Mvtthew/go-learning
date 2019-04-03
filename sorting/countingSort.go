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

	max := 0
	for i := 0; i < len(tab); i++ {
		if tab[i] > max {
			max = tab[i]
		}
	}

	tab = sort(tab, max)


	fmt.Println("After")
	fmt.Println(tab)

}

func sort(tab [10]int, max int) [10]int  {

	hist := make([]int, max + 1)

	for i := 0; i < len(hist); i++ {
		hist[i] = 0
	}

	for i := 0; i < len(tab); i++ {
		hist[tab[i]]++
	}

	outIndex := 0
	for i := 0; i < len(hist); i++ {
		for j := 0; j < hist[i]; j++ {
			tab[outIndex] = i
			outIndex++
		}
	}

	return tab

}