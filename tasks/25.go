// Реализовать собственную функцию sleep.
package tasks

import (
	"fmt"
	"time"
)

// Функция MySleep без использования пакета time
// Минус - неявное время, проведённое внутри ожидания
func MySleep(time uint64) {
	time = time * 1000000000
	c := make(chan struct{})
	go func() {
		for time > 1 {
			time--
		}
		c <- struct{}{}
	}()
	<-c
}

// Функция TimeSleep с использованием пакета time
func TimeSleep(t time.Duration) {
	timer := time.NewTimer(t) //Инициализируем таймер
	<-timer.C                 //Ждём, когда в канал передастся нужное значение
}

func Task25_1() {
	fmt.Println(time.Now())
	MySleep(1)
	fmt.Println(time.Now())
	TimeSleep(1 * time.Second)
	fmt.Println(time.Now())
}
