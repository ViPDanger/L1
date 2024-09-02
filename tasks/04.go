// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.
package tasks

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// Воркер,читающий данные из канала и выводящий их в stdout. (пакет fmt в выводит данные в stdout в функциях без указания io.Writer)
// Закрываем воркер по контексту сигнала системы
func worker(ctx context.Context, workerId int, in chan int) {
	for ctx.Err() == nil {
		fmt.Println("Worker ", workerId, ": ", <-in)
	}
	fmt.Println("Worker ", workerId, ": Done")
}

func Task4_1() {
	var N int
	//	signal.NotifyContext позволяет назначать в качестве context.CancelFunc системные сигналы о прерывании и принудительном завершении
	//	Используем контекст для отлова данного сигнала и адекватной работы программы при его получении (без критических выходов, ошибок и прочего.)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Канал внутри воркера. Буфферезированный, для удобства
	jobsChannel := make(chan int, 1)
	// Выбор кол-ва воркеров при старте
	fmt.Print("Кол-во воркеров: ")
	_, err := fmt.Scan(&N)
	if err != nil || N < 1 {
		log.Fatalln(err)
	}
	// Запуск воркеров
	for N > 0 {
		go worker(ctx, N, jobsChannel)
		N--
	}
	// Постоянная запись данных в канал из главного потока
	for i := 0; ctx.Err() == nil; i++ {
		jobsChannel <- i
		time.Sleep(200 * time.Millisecond)
	}
	<-ctx.Done()
	close(jobsChannel)
	time.Sleep(1 * time.Second)

}
