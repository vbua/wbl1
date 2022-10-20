package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout.Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

// N - количество воркеров
const N = 250

func main() {
	secondsChan := make(chan int64)
	quitChan := make(chan os.Signal)
	shutdownChan := make(chan struct{})

	// ловим сигинт (ctrl+c) и сигтерм (kill) и уведомляем в канал
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	for i := 1; i <= N; i++ {
		wg.Add(1)
		go func(i int, secondsChan chan int64, shutdownChan chan struct{}, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Printf("Starting goroutine %v\n", i)
			for {
				select {
				case res := <-secondsChan:
					fmt.Printf("Worker %v; data %v\n", i, res)
				case <-shutdownChan: // тут приходит nil и закрываем горутину
					fmt.Println("Shutdown goroutine: ", i)
					return
				}
			}
		}(i, secondsChan, shutdownChan, &wg)
	}

	// бесконечно шлем текущее юникс время, пока не получим сообщение в quitChan.
	// тогда закрываем канал shutdownChan, что в следствии шлет во все горутины nil.
l:
	for {
		select {
		case secondsChan <- time.Now().Unix():
		case <-quitChan:
			// при закрытии канала летят nil сообщения в канал, поэтому каждая горутина получит свое сообщение
			close(shutdownChan)
			break l
		}
	}

	wg.Wait()
	fmt.Println("Finished all goroutines")
}
