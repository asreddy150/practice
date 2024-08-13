package main

import (
	"fmt"
)

var even, odd = 0, 1

//var wg = sync.WaitGroup{}

func printEven(ch1, ch2 chan int) {
	defer wg.Done()
	for range ch2 {
		if even > 100 {
			break
		}
		fmt.Println(even)
		ch1 <- even
		even = even + 2
	}
	close(ch1)
}
func printOdd(ch1, ch2 chan int) {
	defer wg.Done()
	ch2 <- odd
	for range ch1 {
		if odd > 100 {
			break
		}
		fmt.Println(odd)
		ch2 <- odd
		odd = odd + 2
	}
	close(ch2)
}

// func main() {
// 	ch1, ch2 := make(chan int), make(chan int)
// 	wg.Add(2)
// 	go printEven(ch1, ch2)
// 	go printOdd(ch1, ch2)
// 	wg.Wait()
// }
