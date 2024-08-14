package main

import (
	"fmt"
	"net/http"
	"sync"
)

func get(url string, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	post, err := http.Get(url)
	if err != nil {
		result <- "Error when call URL"
	}
	fmt.Println(post.Body)
	result <- "Success"

}

func main() {
	listURl := []string{"https://jsonplaceholder.typicode.com/todos/1", "https://jsonplaceholder.typicode.com/todos/2"}
	var wg sync.WaitGroup
	result := make(chan string, len(listURl))
	for _, url := range listURl {
		wg.Add(1)
		go get(url, result, &wg)
	}
	wg.Wait()

	for item := range result {
		fmt.Println("Item", item)
	}
}
