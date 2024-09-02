package main

import (
	"fmt"

	"github.com/ViP.Danger/L1/tasks"
)

func main() {
	var task_id int
	for {
		fmt.Println("Task manager. To quit enter value 0")
		fmt.Print("Task id(1 to 26)= ")
		fmt.Scan(&task_id)
		switch task_id {
		case 0:
			fmt.Println("Exit")
			return
		case 1:
			tasks.Task1_1()
		case 2:
			tasks.Task2_1()
		case 3:
			tasks.Task3_1()
		case 4:
			tasks.Task4_1()
		case 5:
			tasks.Task5_1()
		case 6:
			tasks.Task6_1()
		case 7:
			tasks.Task7_1()
		case 8:
			tasks.Task8_1()
		case 9:
			fmt.Println("Task 9_1")
			tasks.Task9_1()
			fmt.Println("Task 9_2")
			tasks.Task9_2()
		case 10:
			fmt.Println("Task 10_1")
			tasks.Task10_1()
			fmt.Println("Task 10_2")
			tasks.Task10_2()
		case 11:
			fmt.Println("Task 11_1")
			tasks.Task11_1()
			fmt.Println("Task 11_2")
			tasks.Task11_2()
			fmt.Println("Task 11_3")
			tasks.Task11_3()
		case 12:
			fmt.Println("Task 12_1")
			tasks.Task12_1()
			fmt.Println("Task 12_2")
			tasks.Task12_2()
		case 13:
			fmt.Println("Task 13_1")
			tasks.Task13_1()
			fmt.Println("Task 13_2")
			tasks.Task13_2()
		case 14:
			tasks.Task14_1()
		case 15:
			tasks.Task15_1()
		case 16:
			tasks.Task16_1()
		case 17:
			tasks.Task17_1()
		case 18:
			tasks.Task18_1()
		case 19:
			tasks.Task19_1()
		case 20:
			tasks.Task20_1()
		case 21:
			tasks.Task21_1()
		case 22:
			tasks.Task22_1()
		case 23:
			fmt.Println("Task 23_1")
			tasks.Task23_1()
			fmt.Println("Task 23_2")
			tasks.Task23_2()
		case 24:
			tasks.Task24_1()
		case 25:
			tasks.Task25_1()
		case 26:
			tasks.Task26_1()
		default:
			fmt.Println("Unknown task id. Try again")
		}

	}
}
