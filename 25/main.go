package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func sleep(sec int) {
	// возвращает канал, на котором блокируется на указанное время пока не придет сообщение
	<-time.After(time.Second * time.Duration(sec))
}

func main() {
	fmt.Println("Start sleeping!")
	sleep(10)
	fmt.Println("Woke up!")
}
