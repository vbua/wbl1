package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

/*
С помощью контекста. Можно использовать WithCancel, WithDeadline, WithTimeout.
В данном случае используется WithTimeout
*/
func stopGoRoutineWithContext() {
	ctx, close := context.WithTimeout(context.Background(), 5*time.Second)
	defer close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Shutting down go routine with context")
				return
			}
		}
	}(ctx, &wg)

	wg.Wait()
}

/*
Создаем канал для сигнала о завершении. Как только получаем сообщение, сразу завершаем горутину.
*/
func stopGoRoutineWithChannel() {
	quit := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup, quit <-chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("Shutting down go routine with channel")
				return
			}
		}
	}(&wg, quit)

	time.Sleep(5 * time.Second)
	quit <- struct{}{}
	wg.Wait()
}

/*
Закрываем канал и фор рендж в горутине завершается.
*/
func stopGoRoutineWithCloseChannel() {
	reader := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup, reader <-chan struct{}) {
		defer wg.Done()
		for v := range reader {
			fmt.Println(v)
		}

		fmt.Println("Shutting down when channel closed")
	}(&wg, reader)

	time.Sleep(5 * time.Second)
	close(reader)
	wg.Wait()
}

func main() {
	stopGoRoutineWithContext()
	stopGoRoutineWithChannel()
	stopGoRoutineWithCloseChannel()
}
