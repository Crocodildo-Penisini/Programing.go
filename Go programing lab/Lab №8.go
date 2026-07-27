// Напишіть програму, яка приймає від користувача рядок і замінює всі пробіли на дефіс, використовуючи вказівники.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func replaceSpace(str *string) { // - str *адреса змінної
	runes := []rune(*str)
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			runes[i] = '-'
		}
	}
	*str = string(runes)
}

func main() {
	var str string
	fmt.Print("Enter string: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()

	replaceSpace(&str)
	fmt.Println("Рядок після заміни пробілів на дефіси:", str)
}
