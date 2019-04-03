package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var number string
	var choice int

	fmt.Println("Type number")
	fmt.Scanf("%s", &number)

	fmt.Println("1 - to decimal")
	fmt.Println("2 - to binary")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		fmt.Println("Output:", toDecimal(number))
		break
	case 2:
		fmt.Println("Output:", toBinary(number))
		break
	default:
		fmt.Println("Bad input, terminating...")
		break
	}

}

func toBinary(number string) string {

	num, _ := strconv.Atoi(number)
	numCheck := num
	var converted string

	i := 0

	for {

		if numCheck > int(math.Pow(float64(2), float64(i))) {
			numCheck -= int(math.Pow(float64(2), float64(i)))
			i++

		} else {

			for {
				if i < 0 {
					break
				}

				if num >= int(math.Pow(float64(2), float64(i))) && num > 0 {

					num -= int(math.Pow(float64(2), float64(i)))
					converted += "1"

				} else  {

					converted += "0"

				}

				i--
			}

			break
		}
	}

	return converted

}

func toDecimal(number string) string {

	var sum int

	for i := len(number) - 1; i >= 0; i-- {
		switch string(number[i]) {
		case "1":
			sum += int(math.Pow(float64(2), float64(len(number)-1-i)))
			break
		}
	}

	return strconv.Itoa(sum)
}
