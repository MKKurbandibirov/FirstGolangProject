package main

import (
	"fmt"
)

func remaining(water, milk, beans, cups, money int) {
	fmt.Println("\nThe coffee machine has:")
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", beans)
	fmt.Printf("%d disposable cups\n", cups)
	fmt.Printf("$%d of money\n\n", money)
}

func syrup(money *int) {
	var number int

	fmt.Println("\nWould you like a syrup for coffee: 1 - vanilla, 2 - caramel, 3 - strawberry or none:")
	fmt.Scan(&number)
	if number == 1 {
		*money -= 2
	} else if number == 2 {
		*money -= 3
	} else if number == 3 {
		*money -= 1
	}
}

func buy(water, milk, beans, cups, money *int) {
	var number int

	fmt.Println("\nWhat do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&number)
	if number == 1 {
		if *water - 250 < 0 {
			fmt.Println("Sorry, not enough water!")
		} else if *beans - 16 < 0 {
			fmt.Println("Sorry, not enough beans!")
		} else if *cups == 0 {
			fmt.Println("Sorry, not enough cups!")
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			*water -= 250
			*beans -= 16
			*money += 4
			*cups--
		}
		syrup(money)
	} else if number == 2 {
		if *water - 350 < 0 {
			fmt.Println("Sorry, not enough water!")
		} else if *milk - 75 < 0{
			fmt.Println("Sorry, not enough milk!")
		} else if *beans - 20 < 0 {
			fmt.Println("Sorry, not enough beans!")
		} else if *cups == 0 {
			fmt.Println("Sorry, not enough cups!")
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			*water -= 350
			*milk -= 75
			*beans -= 20
			*money += 7
			*cups--
		}
		syrup(money)
	} else if number == 3 {
		if *water - 200 < 0 {
			fmt.Println("Sorry, not enough water!")
		} else if *milk - 100 < 0{
			fmt.Println("Sorry, not enough milk!")
		} else if *beans - 12 < 0 {
			fmt.Println("Sorry, not enough beans!")
		} else if *cups == 0 {
			fmt.Println("Sorry, not enough cups!")
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			*water -= 200
			*milk -= 100
			*beans -= 12
			*money += 6
			*cups--
		}
		syrup(money)
	}
	fmt.Println()
}

func fill(water, milk, beans, cups *int) {
	var add int

	fmt.Println("\nWrite how many ml of water you want to add:")
	fmt.Scan(&add)
	*water += add
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&add)
	*milk += add
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&add)
	*beans += add
	fmt.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scan(&add)
	*cups += add
	fmt.Println()
}

func take(money *int) {
	fmt.Printf("\nI gave you $%d\n", *money)
	*money = 0
	fmt.Println()
}

func chooseAction(water, milk, beans, cups, money *int) {
	var action string

	for true {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)
		if action == "buy" {
			buy(water, milk, beans, cups, money)
		} else if action ==  "fill" {
			fill(water, milk, beans, cups)
		} else if action == "take" {
			take(money)
		} else if action == "remaining" {
			remaining(*water, *milk, *beans, *cups, *money)
		} else if action == "exit"{
			break ;
		}
	}
}

func main() {
	var water, milk, beans, cups, money int

	water = 400
	milk = 540
	beans = 120
	cups = 9
	money = 550
	chooseAction(&water, &milk, &beans, &cups, &money)
}