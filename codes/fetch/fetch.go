package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	rsp, _ := http.Get(url)
	b, _ := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()
	ch <- string(b)
}
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
