package main

import (
	"fmt"
	"net/http"
	"sync"
)

//同时开三个协程去请求网页， 等三个请求都完成后才继续 Wait 之后的工作
func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"https://studygolang.com/",
		"http://www.baidu.com/",
		"http://www.google.com/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
			fmt.Println(url)
			fmt.Println(http.Get(url))
		}(url)
	}

	wg.Wait()
}
