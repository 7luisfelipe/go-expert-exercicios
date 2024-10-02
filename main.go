package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func get(url string, quantity int, wg *sync.WaitGroup, ch chan map[string]int) {
	defer wg.Done()
	result := make(map[string]int)
	for i := 0; i < quantity; i++ {
		req, _ := http.Get(url)

		//Se já teve algum retorno com esse status
		if _, exists := result[req.Status]; exists {
			result[req.Status] = result[req.Status] + 1
		} else {
			result[req.Status] = 1
		}

	}
	ch <- result
}

func main() {
	url := os.Args[1]
	if url == "" {
		panic(errors.New("url inválida"))
	}

	requests, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	concurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	requestPerConcurrency := requests / concurrency

	wg := sync.WaitGroup{}
	wg.Add(concurrency)

	ch := make(chan map[string]int)

	//Resultados
	result := make(map[string]int)

	start := time.Now()
	end := time.Now()
	for i := 0; i < concurrency; i++ {
		go get(url, requestPerConcurrency, &wg, ch)
	}
	go func() {
		wg.Wait()
		end = time.Now()
		close(ch)
	}()

	//Resultado
	for list := range ch {
		for k, v := range list {
			if _, exists := result[k]; exists {
				result[k] = result[k] + v
			} else {
				result[k] = v
			}
		}
	}

	fmt.Println("Tempo total: ", (end.Sub(start).Seconds()), " segundos")
	for k, v := range result {
		fmt.Println(k + " - " + strconv.Itoa(v))
	}

}
