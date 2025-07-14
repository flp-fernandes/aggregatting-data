package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	userId := getUserId()

	qtyOfGoroutines := 2
	responseChannel := make(chan any, qtyOfGoroutines)
	wg := &sync.WaitGroup{}
	wg.Add(qtyOfGoroutines)

	go getUserBalance(userId, responseChannel, wg)
	go getUserName(userId, responseChannel, wg)
	wg.Wait()
	close(responseChannel)

	for response := range responseChannel {
		fmt.Println("Resposta recebida:", response)
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("Tempo de execução:", elapsedTime)
}

func getUserId() string {
	time.Sleep(time.Millisecond * 100)

	return "f93581ea-8e76-4d7e-8432-d016e397bb21"
}

func getUserBalance(_ string, responseChannel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	responseChannel <- 400
	wg.Done()
}

func getUserName(_ string, responseChannel chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	responseChannel <- "Felipão"
	wg.Done()
}
