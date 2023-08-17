package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{
	"test",
}

var wg sync.WaitGroup

// alow one's use at a time
var mut sync.Mutex

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
	}
	for _, web := range websiteList {
		wg.Add(1)
		go getStatusCode(web)
	}
	//waitgroup: stop the method main from being finisher until all the goroutines is done
	wg.Wait()

	fmt.Print(signals)
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)

		fmt.Println(s)
	}
	defer wg.Done()
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
	defer wg.Done()
}
