//Обчислення площі чотирикутника: Напишіть функцію, яка обчислює площу чотирикутника за заданими довжинами його сторін.

package main

import (
	"fmt"
	"math"
)

func calcHeronFormula(SideA, SideB, SideC float64) float64 {
	P := (SideA + SideB + SideC) / 2
	fmt.Println("Пів периметр:", P)
	return math.Sqrt(P * (P - SideA) * (P - SideB) * (P - SideC))
}

func CalculateAreaOfQuadrilateral(SideA, SideB, SideC, SideD, Diagonal float64) float64 {
	Area1 := calcHeronFormula(SideA, SideB, Diagonal)
	fmt.Println("1 периметр:", Area1)
	Area2 := calcHeronFormula(SideC, SideD, Diagonal)
	fmt.Println("2 периметр:", Area2)
	return Area1 + Area2
}

func main() {
	var SideA, SideB, SideC, SideD, Diagonal float64
	fmt.Println("Введіть Сторони a, b, c, d та diagonal:")
	_, err := fmt.Scan(&SideA, &SideB, &SideC, &SideD, &Diagonal)
	if err != nil {
		fmt.Println("Помилка вводу:", err)
		return
	}
	Area := CalculateAreaOfQuadrilateral(SideA, SideB, SideC, SideD, Diagonal)
	fmt.Println("Обчислення формули Герона:", Area)
}
