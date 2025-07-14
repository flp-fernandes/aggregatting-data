package main

import (
	"fmt"
	"sync"
	"time"
)

type Response struct {
	Type  string
	Value any
}

func main() {
	startTime := time.Now()

	userId := getUserId()

	qtyOfGoroutines := 2
	responseChannel := make(chan Response, qtyOfGoroutines)
	wg := &sync.WaitGroup{}
	wg.Add(qtyOfGoroutines)

	go getUserBalance(userId, responseChannel, wg)
	go getUserName(userId, responseChannel, wg)
	wg.Wait()
	close(responseChannel)

	var balance int
	var name string

	for response := range responseChannel {
		switch response.Type {
		case "balance":
			balance = response.Value.(int)
		case "name":
			name = response.Value.(string)
		}
	}

	fmt.Printf("Saldo: %d, Nome: %s\n", balance, name)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tempo de execução:", elapsedTime)
}

func getUserId() string {
	time.Sleep(time.Millisecond * 100)

	return "f93581ea-8e76-4d7e-8432-d016e397bb21"
}

func getUserBalance(_ string, responseChannel chan Response, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 100)

	responseChannel <- Response{Type: "balance", Value: 400}
}

func getUserName(_ string, responseChannel chan Response, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 150)

	responseChannel <- Response{Type: "name", Value: "Felipão"}
}
