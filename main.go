package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go func(c chan int) {
			c <- i
		}(ch)
	}
	sum(ch)

}

func sum(ch chan int) {
	sum := 0
	for val := range ch {
		sum += val
	}
	close(ch)

	fmt.Printf("Sum: %d\n", sum)
}
