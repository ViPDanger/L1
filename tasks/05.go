// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.
package tasks

import (
	"fmt"
	"log"
	"time"
)

func Task5_1() {
	var N int
	// N, кол-во секунд работы программы
	fmt.Print("N: ")
	_, err := fmt.Scan(&N)

	// Канал для последовательной записи/прочтения. Для удобства воспользуемся буфферизированным каналом
	сhannel := make(chan int, 1)
	if err != nil || N < 1 {
		log.Fatalln(err)
	}
	// Горутина чтения канала
	go func(c chan int) {
		for {
			fmt.Println("Writer: ", <-c)
		}
	}(сhannel)

	// Горутина записи в канал
	go func(c chan int) {
		for i := 0; true; i++ {
			c <- i
			time.Sleep(250 * time.Millisecond)
		}
	}(сhannel)
	time.Sleep(time.Duration(N) * time.Second)

}
