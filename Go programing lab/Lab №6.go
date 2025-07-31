// Напишіть програму, яка сортує кожен рядок у двовимірному масиві розміром  3x3 за спаданням.

package main

import "fmt"

func main() {
	const rows = 3
	const cols = 3
	var matrix [rows][cols]int
	matrix[0][0] = 9
	matrix[0][1] = 6
	matrix[0][2] = 3
	matrix[1][0] = 2
	matrix[1][1] = 1
	matrix[1][2] = 8
	matrix[2][0] = 5
	matrix[2][1] = 4
	matrix[2][2] = 7
	fmt.Print("Initial array:\n")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
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
