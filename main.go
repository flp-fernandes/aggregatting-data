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

	balanceChannel := make(chan int, 1)
	nameChannel := make(chan string, 1)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go getUserBalance(userId, balanceChannel, wg)
	go getUserName(userId, nameChannel, wg)
	wg.Wait()

	balance := <-balanceChannel
	name := <-nameChannel

	fmt.Printf("Saldo: %d, Nome: %s\n", balance, name)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tempo de execução:", elapsedTime)
}

func getUserId() string {
	time.Sleep(time.Millisecond * 100)

	return "f93581ea-8e76-4d7e-8432-d016e397bb21"
}

func getUserBalance(_ string, balanceChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(balanceChannel)

	time.Sleep(time.Millisecond * 100)

	balanceChannel <- 400
}

func getUserName(_ string, nameChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(nameChannel)

	time.Sleep(time.Millisecond * 150)

	nameChannel <- "Felipão"
}
