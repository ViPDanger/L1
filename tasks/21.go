//	Реализовать паттерн «адаптер» на любом примере.

// 	Коментарий автора
// 	Напишем адаптер для counterUnsync struct,
//	неудачливого брата близнеца структуры counter из 18 задания

package tasks

import "fmt"

// Сама структура counterUnsync

type counterUnsync struct {
	count int
}

// создаёт структуру counterUnsync
func newCounterUnsync() *counterUnsync {
	return &counterUnsync{}
}

// Инкрементирует counterUnsync.count
func (c *counterUnsync) AddUnsync() {
	c.count++
}

// Выводит значение count
func (c *counterUnsync) CountUnsync() int {
	return c.count
}

//	Описываем адаптер для counterUnsync

type adapter struct {
	adaptee *counterUnsync
}

func (c adapter) Add() {
	c.adaptee.AddUnsync()
}
func (c adapter) Count() int {
	return c.adaptee.CountUnsync()
}

//	Так же опишем интерфейс, для которого будет предназначен адаптер
//	Запишем и какую нибудь функцию с этим интерфейсом.
type counterInterface interface {
	Add()
	Count() int
}

func checkCounter(countInterface counterInterface) {
	fmt.Println("Число count - ", countInterface.Count())
	fmt.Println("Add to counter")
	countInterface.Add()
	fmt.Println("Число count - ", countInterface.Count())
}

func Task21_1() {
	// Создадим counter и подадим её в функцию
	countSync := newCounter()
	fmt.Println("Counter")
	checkCounter(countSync)

	// А теперь создадим counterUnsync и подадим её в функцию адаптером
	countUnsync := counterUnsync{}
	fmt.Println("\nCounterUnsync")
	checkCounter(adapter{adaptee: &countUnsync})

}
