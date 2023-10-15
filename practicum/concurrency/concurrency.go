package main

import (
	"fmt"
	"time"
)

const LESSON_NUMBER int = 12

func sum(arr []int, c chan int) {
	result := 0
	for _, val := range arr {
		result += val
	}
	c <- result
}

func in(c chan<- string, msg string) {
	c <- msg
}

func out(c <-chan string) {
	msg := <-c
	fmt.Println(msg)
}

func waitForMe(c chan int, length int) {
	for i := 1; i <= length; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
}

func PrintWhatILearned() {
	c := make(chan int)

	go sum([]int{1, 2, 3, 4, 5}, c)
	go sum([]int{2, 3, 4, 5, 6}, c)

	result1, result2 := <-c, <-c
	fmt.Println(result1 + result2)

	messenger := make(chan string, 1)
	in(messenger, "Hello World!")
	out(messenger)

	waiter := make(chan int, 1)

	go waitForMe(waiter, 1)

	select {
	case num := <-waiter:
		fmt.Println(num)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}

	queue := make(chan int, 3)
	go waitForMe(queue, 3)

	for value := range queue {
		fmt.Println(value)
	}

	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(5 * time.Second)

	go func() {
		for tick := range ticker.C {
			fmt.Println(tick)
		}
	}()

	<-timer.C
	ticker.Stop()
}

func main() {
	fmt.Printf("\n%d. Concurrency.\n", LESSON_NUMBER)

	PrintWhatILearned()
}
