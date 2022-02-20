package main

// thread could step on each other, this is the example for it,
// run it without lock Mutex then you will see the money at the result may will be diffirent from 100

import (
	"sync"
	"time"
)

var (
	stocks = 100
	lock   = sync.Mutex{}
)

func takeProduct() {
	for i := 0; i <= 10000; i++ {
		// lock.Lock() // lock, then the other thread can't touch to the `stocks` variable
		stocks += 1
		time.Sleep(1 * time.Millisecond)

		// lock.Unlock() // Now we unlock it so the other can't worl with `stocks`
	}
	println("Sell 1")
}

func addProduct() {
	for i := 0; i <= 10000; i++ {
		// lock.Lock()
		stocks -= 1
		time.Sleep(1 * time.Millisecond)

		// lock.Unlock()
	}
	println("Adding 1")
}

func main() {
	go takeProduct()
	go addProduct()
	time.Sleep(20000 * time.Millisecond)
	println(stocks)
}
