// Напишіть програму, яка замінює всі додатні значення в масиві і слайсі з 10 елементів на нулі.

package main

import "fmt"

func main() {
	array := []int{-7, 2, -9, 6, -4, 5, -1, 8, -3, 10} // - масив

	fmt.Println("Array:", array, "\n")

	size1 := len(array) // - Визначення кількості елементів масиву

	fmt.Println("Size array:", size1, "\n")

	for i := 0; i < size1; i++ {
		if array[i] > 0 {
			array[i] = 0
		}
	}

	fmt.Println("Array after replacing positive numbers with 0:")

	for i := 0; i < size1; i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Println("\n")

	slice := []int{-7, 2, -9, 6, -4, 5, -1, 8, -3, 10}

	fmt.Println("Slice:", slice, "\n")

	size2 := len(slice) // - Визначення кількості елементів слайсу

	fmt.Println("Size slice:", size2, "\n")

	for i := 0; i < size2; i++ {
		if slice[i] > 0 {
			slice[i] = 0
		}
	}

	fmt.Println("Slice after replacing positive numbers with 0:")

	for i := 0; i < size2; i++ {
		fmt.Print(slice[i], " ")
	}
	fmt.Println("")
}
