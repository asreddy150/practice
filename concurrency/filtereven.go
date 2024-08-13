package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var mu = sync.Mutex{}

func firstFunc(ch1, ch2 chan int) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		ch2 <- i
	}
	close(ch2)
	for num := range ch1 {
		fmt.Println(num)
	}
}
func secondFunc(ch1, ch2 chan int) {
	defer wg.Done()
	for i := range ch2 {
		if i%2 == 0 {
			ch1 <- i
		}
	}
	close(ch1)
}

func newFirstFunc(ch1, ch2 chan int) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		ch2 <- i
	}
	close(ch2)
	even := <-ch1
	fmt.Println(even)
}
func newSecondFunc(ch1, ch2 chan int) {
	defer wg.Done()
	i := <-ch2
	if i%2 == 0 {
		ch1 <- i
	}
	close(ch1)
}

// func main() {
// 	ch1, ch2 := make(chan int), make(chan int, 20)
// 	wg.Add(2)
// 	go firstFunc(ch1, ch2)
// 	go secondFunc(ch1, ch2)
// 	wg.Wait()
// }
