// Напишіть програму, яка замінює всі додатні значення в масиві з 10 елементів на нулі.

package main

import "fmt"

func main() {
	array := [10]int{-7, 2, -9, 6, -4, 5, -1, 8, -3, 10}
	fmt.Println("Array:", array)
	size := len(array)
	for i := 0; i < size; i++ {
		if array[i] > 0 {
			array[i] = 0
		}
	}
	fmt.Println("Array after replacing positive numbers with 0:")
	for i := 0; i < size; i++ {
		fmt.Print(array[i], " ")
	}
}
