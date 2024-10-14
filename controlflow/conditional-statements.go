package controlflow

import "fmt"

func ConditionalStatements1() {
	var fruit = "watermelon"
	if fruit == "apple" {
		fmt.Printf("fruit is an %s\n", fruit)
	} else {
		fmt.Printf("fruit \"%s\" is not an apple\n", fruit)
	}
}

func ConditionalStatements2() {
	var fruit = "orange"
	switch fruit {
	case "orange":
		fmt.Printf("%s is Sweet and Sour.\n", fruit)
	case "apple":
		fmt.Printf("%s is Sweet.\n", fruit)
	case "mango":
		fmt.Printf("%s is Sweet and Sour.\n", fruit)
	case "banana":
		fmt.Printf("%s is Sweet and Soft.\n", fruit)
	default:
		fmt.Printf("%s is not in our fruits list\n", fruit)
	}
}

func Loop1() {
	var num = 1
	for num < 11 {
		fmt.Printf("num is %d\n", num)
		num++
	}
}

func Loop2() {
	var fruit = "Watermelon"
	for i, ch := range fruit {
		aux := "th"
		if i%10 == 1 && i%100 != 11 {
			aux = "st"
		}
		if i%10 == 2 && i%100 != 12 {
			aux = "nd"
		}
		if i%10 == 3 && i%100 != 13 {
			aux = "rd"
		}
		fmt.Printf("%c is %d%s character of %s \n", ch, i, aux, fruit)
	}
}

func Loop3() {
	var numbers = [5]int{2, 4, 8, 16, 32}
	for i, num := range numbers {
		aux := "th"
		if i%10 == 1 && i%100 != 11 {
			aux = "st"
		}
		if i%10 == 2 && i%100 != 12 {
			aux = "nd"
		}
		if i%10 == 3 && i%100 != 13 {
			aux = "rd"
		}
		fmt.Printf("%d is %d%s item of %v \n", num, i, aux, numbers)
	}
}
