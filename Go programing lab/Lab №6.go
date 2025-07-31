// Напишіть програму, яка сортує кожен рядок у двовимірному масиві розміром  3x3 за спаданням.

package main

import "fmt"

func main() {
	const rows = 3
	const cols = 3
	matrix := [rows][cols]int{
		{6, 3, 9},
		{2, 5, 8},
		{1, 7, 4},
	}
	fmt.Print("Початковий масив:\n")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols-1; j++ {
			for k := 0; k < cols-1; k++ {
				if matrix[i][k] < matrix[i][k+1] {
					temp := matrix[i][k]
					matrix[i][k] = matrix[i][k+1]
					matrix[i][k+1] = temp
				}
			}
		}
	}
	fmt.Println("Sorted array:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
