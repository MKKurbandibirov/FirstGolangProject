package main

import (
	"fmt"
)

type CoffeeMachine struct {
	water	int
	milk	int
	beans	int
	cups	int
	money	int
}

func remaining(machine *CoffeeMachine) {
	fmt.Println("\nThe coffee machine has:")
	fmt.Printf("%d ml of water\n", machine.water)
	fmt.Printf("%d ml of milk\n", machine.milk)
	fmt.Printf("%d g of coffee beans\n", machine.beans)
	fmt.Printf("%d disposable cups\n", machine.cups)
	fmt.Printf("$%d of money\n\n", machine.money)
}

func syrup(machine *CoffeeMachine) {
	var number int

	fmt.Println("\nWould you like a syrup for coffee: 1 - vanilla, 2 - caramel, 3 - strawberry or none:")
	fmt.Scan(&number)
	if number == 1 {
		machine.money -= 2
	} else if number == 2 {
		machine.money -= 3
	} else if number == 3 {
		machine.money -= 1
	}
}

func buy(machine *CoffeeMachine) {
	var number int

	fmt.Println("\nWhat do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&number)
	if number == 1 {
		if machine.water - 250 < 0 {
			fmt.Println("Sorry, not enough water!")
		} else if machine.beans - 16 < 0 {
			fmt.Println("Sorry, not enough beans!")
		} else if machine.cups == 0 {
			fmt.Println("Sorry, not enough cups!")
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			machine.water -= 250
			machine.beans -= 16
			machine.money += 4
			machine.cups--
			syrup(machine)
		}
	} else if number == 2 {
		if machine.water - 350 < 0 {
			fmt.Println("Sorry, not enough water!")
		} else if machine.milk - 75 < 0{
			fmt.Println("Sorry, not enough milk!")
		} else if machine.beans - 20 < 0 {
			fmt.Println("Sorry, not enough beans!")
		} else if machine.cups == 0 {
			fmt.Println("Sorry, not enough cups!")
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			machine.water -= 350
			machine.milk -= 75
			machine.beans -= 20
			machine.money += 7
			machine.cups--
			syrup(machine)
		}
	} else if number == 3 {
		if machine.water - 200 < 0 {
			fmt.Println("Sorry, not enough water!")
		} else if machine.milk - 100 < 0{
			fmt.Println("Sorry, not enough milk!")
		} else if machine.beans - 12 < 0 {
			fmt.Println("Sorry, not enough beans!")
		} else if machine.cups == 0 {
			fmt.Println("Sorry, not enough cups!")
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			machine.water -= 200
			machine.milk -= 100
			machine.beans -= 12
			machine.money += 6
			machine.cups--
			syrup(machine)
		}
	}
	fmt.Println()
}

func fill(machine *CoffeeMachine) {
	var add int

	fmt.Println("\nWrite how many ml of water you want to add:")
	fmt.Scan(&add)
	machine.water += add
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&add)
	machine.milk += add
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&add)
	machine.beans += add
	fmt.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scan(&add)
	machine.cups += add
	fmt.Println()
}

func take(machine *CoffeeMachine) {
	fmt.Printf("\nI gave you $%d\n", machine.money)
	machine.money = 0
	fmt.Println()
}

func chooseAction(machine *CoffeeMachine) {
	var action string

	for true {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)
		if action == "buy" {
			buy(machine)
		} else if action ==  "fill" {
			fill(machine)
		} else if action == "take" {
			take(machine)
		} else if action == "remaining" {
			remaining(machine)
		} else if action == "exit"{
			break ;
		}
	}
}

func main() {
	var machine *CoffeeMachine
	machine = new(CoffeeMachine)
	machine.water = 400
	machine.milk = 540
	machine.beans = 120
	machine.cups = 9
	machine.money = 550
	chooseAction(machine)
}