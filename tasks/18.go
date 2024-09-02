// Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.
package tasks

import (
	"fmt"
	"sync"
	"time"
)

// Структура counter с mutex и встроенными методами new,add и count
type counter struct {
	count int
	mutex *sync.RWMutex
}

// создаёт структуру counter
func newCounter() *counter {
	return &counter{
		mutex: &sync.RWMutex{},
	}
}

// Инкрементирует counter.count
func (c *counter) Add() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

// Выводит значение count
func (c *counter) Count() int {
	return c.count
}

func Task18_1() {
	// Инициализация counter
	c := newCounter()
	// Добавим потоки, которые будут инкрементировать counter
	go func() {
		for {
			c.Add()
		}
	}()
	go func() {
		for {
			c.Add()
		}
	}()
	go func() {
		for {
			c.Add()
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Counter.Count = ", c.Count())
}
