//Відомо площу рівностороннього трикутника. Знайти його периметр. Значення площі ввести з клавіатури.

package main

import (
	"fmt"
	"math"
)

func main() {
	var Area float64
	fmt.Println("Enter area: ")
	_, err := fmt.Scan(&Area)
	if err != nil {
		fmt.Println("Помилка вводу:", err)
		return
	}

	SideA := math.Sqrt((Area * 4) / math.Sqrt(3))
	fmt.Println("Side a: ", SideA)

	Perimeter := 3 * SideA
	fmt.Println("Perimeter: ", Perimeter)
}
