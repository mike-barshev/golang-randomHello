package main

import (
	"fmt"
	"math/rand"
	"time"
)

func privet(p chan string) {
	fmt.Println(<-p)
}

func main() {
	p := make(chan string)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10)
	go privet(p)
	time.Sleep(time.Duration(r) * time.Second)
	p <- "Привет"
	time.Sleep(time.Duration(r) * time.Second)
	close(p)

}
