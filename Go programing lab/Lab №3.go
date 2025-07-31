//Виведіть всі парні числа від 2 до 100.

package main

import "fmt"

func main() {
	/*for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			if 100 <= i {
				break
			}
		}
	}*/
	i := 1
	for i <= 100 {
		if i%2 == 0 {
			fmt.Println(i)
		}
		i++
	}

	/*for i := 2; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}*/
}
