// Реализовать все возможные способы остановки выполнения горутины.
package tasks

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Task6_1() {
	var (
		mutex     sync.Mutex     // остановка мьютексом
		rwMutex   sync.RWMutex   // остановка .RLock() (Блокируются только .Lock(), .RLock() могут проходить дальше). Используется для паралельного чтения горутинами
		waitGroup sync.WaitGroup // остановка WaitGroup
		atomicInt int32          // "остановка" атомиком
		channel   chan int       // остановка каналом
	)

	mutex.Lock() //
	rwMutex.RLock()
	waitGroup.Add(1)
	channel = make(chan int)
	go func() {
		//
		mutex.Lock()
		fmt.Println("Mutex was unlocked!")
		rwMutex.Lock()
		fmt.Println("RWMutex was unlocked!")
		waitGroup.Wait()
		fmt.Println("waitGroup was unlocked!")
		// Чисто технически, это не является "Остановкой" в полном понимании. Однако это один из вариантов применения атомиков,т.к они безопасны в горутинах
		for atomic.LoadInt32(&atomicInt) < 1 {
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("atomic was unlocked!")
		<-channel
		fmt.Println("channel was unlocked!")
	}()
	time.Sleep(500 * time.Millisecond)
	// Разблокировка мьютекса
	mutex.Unlock()
	time.Sleep(500 * time.Millisecond)
	// Разблокировка RW мьютекса
	rwMutex.RUnlock()
	time.Sleep(500 * time.Millisecond)
	// Разблокировка waitgroup
	waitGroup.Done()
	time.Sleep(500 * time.Millisecond)
	// "разблокировка" atomic
	atomic.StoreInt32(&atomicInt, 500)
	time.Sleep(500 * time.Millisecond)
	// "разблокировка" channel
	channel <- 1
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done")
}
