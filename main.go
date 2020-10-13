package main

import (
	"fmt"
	"math/rand"
	"time"
)

func privet() {
	fmt.Println("Привет")
}

func outRandTen(p chan bool) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)
	p <- true
}

func main() {
	p := make(chan bool)
	go outRandTen(p)
	for {
		select {
		case val := <-p:
			if val {
				return
			}
		default:
			go privet()
		}
		time.Sleep(time.Second)
	}
}
