package main

import (
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	waitGroup = sync.WaitGroup{}
	fundInfo  []byte
	optInfo   []byte
	orderInfo []byte
)

func Get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func getFundInfo(url string) {
	fundInfo = Get(url)
	waitGroup.Done() // signal that OK i'm done, this thread/task is completed
}

func getOTP(url string) {
	optInfo = Get(url)
	waitGroup.Done()
}

func getOrder(url string) {
	orderInfo = Get(url)
	waitGroup.Done()
}

func isAbleToCharge() bool {
	// checking for fundInfo, orderInfo and OTPInfo here
	return false
}

func main() {
	waitGroup.Add(3) // 3 thread that need to be waited, let Add 3

	go getFundInfo("https://.....?signature=xxx")
	go getOTP("https://.....?signature=yyy")
	go getOrder("https://.....?signature=zzz")

	waitGroup.Wait() // tell the main thread to wait
	println("Is able to charge: ", isAbleToCharge())
}
