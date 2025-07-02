package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// PingPong должен запускать две горутины "ping" и "pong",
// которые поочередно выводят строки пять раз каждая.
// Реализуйте синхронизацию через каналы и ожидание завершения.
func PingPong(w io.Writer) {
	// TODO: реализовать обмен сообщениями между горутинами
	const amount = 5
	pingCh := make(chan bool)
	pongCh := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < amount; i++ {
			<-pingCh
			fmt.Fprintln(w, "ping")
			pongCh <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < amount; i++ {
			<-pongCh
			fmt.Fprintln(w, "pong")
			if i != (amount - 1) {
				pingCh <- true
			}
		}
	}()
	pingCh <- true
	wg.Wait()
}

func main() {
	PingPong(os.Stdout)
}
