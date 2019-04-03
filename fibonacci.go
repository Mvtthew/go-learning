package sorting

import "fmt"

func main() {

	i := 0

	fmt.Println("Which element of fibonacci sequence?")
	fmt.Scanf("%d", &i)

	fmt.Println(fibonacci(i))

}

func fibonacci(n int)int {
	if n < 3 {
		return 1
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}
