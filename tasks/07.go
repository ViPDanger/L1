// Реализовать конкурентную запись данных в map.

// Коментарий автора
// т.к. в задании чётко написано о конкурентной записи в Map а не sync.Map, будем писать свою sync.Map
package tasks

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// горутины записи в мапу
func gorootine(channel chan string, mutex *sync.Mutex, ctx context.Context, str string) {
	for i := 0; ctx.Err() == nil; i++ {
		channel <- str
	}
}
func Task7_1() {

	var Mutex sync.Mutex
	// Graceful Shutdown context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Map для конкурентной записи
	channel := make(chan string)
	Map := make(map[int]string, 100)
	// Запуск трёх горутин, что будут конкуретно записывать данные в канал
	go gorootine(channel, &Mutex, ctx, "First")
	go gorootine(channel, &Mutex, ctx, "Second")
	go gorootine(channel, &Mutex, ctx, "Third")
	// Чтение канала.
	for i := 0; ctx.Err() == nil; i++ {
		Map[i] = <-channel
		time.Sleep(200 * time.Millisecond)
	}
	<-ctx.Done()
	fmt.Println(Map)
}
