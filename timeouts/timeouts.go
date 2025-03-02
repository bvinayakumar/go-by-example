package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	}
}
