//Напишіть програму, яка приймає від користувача рядок і замінює всі пробіли на дефіс.

package main

import (
	"bufio" // - бібліотека для зручного читання тексту з консолі
	"fmt"
	"os" // - бібліотека для роботи зі стандартним введенням (stdin)
)

func main() {
	var userInput string
	fmt.Print("Введіть текст: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput = scanner.Text()
	userInputRunes := []rune(userInput) // - Перетворення рядка у масив символів
	for index := 0; index < len(userInputRunes); index++ {
		if userInputRunes[index] == ' ' {
			userInputRunes[index] = '-'
		}
	}
	userInput = string(userInputRunes)

	fmt.Println("Рядок після заміни:", userInput)
}
