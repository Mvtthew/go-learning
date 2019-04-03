package sorting

import (
	"time"
	"fmt"
)

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {

		for {
			chan1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}

	}()

	go func() {

		for {
			chan2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}

	}()

	for {

		select {
			case message1 := <- chan1:
				fmt.Println(message1)
			case message2 := <- chan2:
				fmt.Println(message2)
		}

	}

}
