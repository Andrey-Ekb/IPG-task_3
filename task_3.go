/*Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
-Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
-Произведение: следующая горутина умножает квадрат числа на 2.
-При вводе «стоп» выполнение программы останавливается.*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

func square(c chan int) {
	num := <-c
	c <- num * num
}

func composition(c chan int) {
	num := <-c
	c <- num * num * 2
}

func main() {
	var numTest int
	var strTest string
	var ex string
	ex = "0"

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	fmt.Println()

	for {
		if ex != "1" {
			fmt.Println("Введите натуральное число и нажмите Enter")
			fmt.Scan(&numTest)
			fmt.Println()
			composChan := make(chan int)
			squareChan := make(chan int)

			go square(squareChan)
			go composition(composChan)

			squareChan <- numTest
			composChan <- numTest

			fmt.Printf("число: %d\n", numTest)

			squareVal := <-squareChan
			composVal := <-composChan
			fmt.Printf("квадрат числа = %d\n", squareVal)
			fmt.Printf("произведение = %d\n", composVal)
			fmt.Println()

			go func() {
				<-sigchan
				ex = "1"
			}()
		}
		if ex != "1" {
			fmt.Println("Введите stop для выхода или p + Enter для продолжения")
			fmt.Scan(&strTest)
		}

		fmt.Println()
		if strTest == "stop" || ex == "1" {
			fmt.Println()
			log.Println("Выход из программы")
			fmt.Println()
			break
		}
	}
}
