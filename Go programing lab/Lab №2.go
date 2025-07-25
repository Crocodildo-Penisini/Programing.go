//Вибір гравця в грі "Камінь, ножиці, папір": Реалізувати гру "Камінь, ножиці, папір" і засновувати переможця на виборі користувача та комп'ютера.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
  
	var playerChoice int
	computerChoice := rng.Intn(3) + 1
	fmt.Println("Випадкове число:", computerChoice)
	fmt.Println("Гра: камінь, ножиці, папір")
	fmt.Println("Зробіть свій вибір:")
	fmt.Println(" 1 - Stone")
	fmt.Println(" 2 - Scissors")
	fmt.Println(" 3 - Paper")

	_, err := fmt.Scan(&playerChoice)
	if err != nil {
		fmt.Println("Помилка вводу:", err)
		return
	}
	switch playerChoice {
	case 1:
		fmt.Println("You have chosen a stone")
		break
	case 2:
		fmt.Println("You have chosen scissors")
		break
	case 3:
		fmt.Println("You have selected the paper")
		break
	}
  
	switch computerChoice {
	case 1:
		fmt.Println("The computer chose the stone")
		break
	case 2:
		fmt.Println("The computer has selected the scissors")
		break
	case 3:
		fmt.Println("The computer has selected paper")
		break
	}
	if playerChoice == computerChoice {
		fmt.Println("Draw")
	} else {
		if playerChoice == 1 && computerChoice == 2 || playerChoice == 2 && computerChoice == 3 || playerChoice == 3 && computerChoice == 1 {
			fmt.Println("You win")
		} else {
			fmt.Println("The computer win")
		}
	}
}
