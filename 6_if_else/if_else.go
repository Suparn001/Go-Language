package main

import "fmt"

func main() {
	// age := 16

	// if age >= 21 {
	// 	fmt.Println("person is above 21")
	// } else if age >= 13 && age <= 19 {
	// 	fmt.Println("person is teen")
	// } else {
	// 	fmt.Println("person is child")
	// }

	// we can declare a variable in if statement and it will be only accessible in that if statement
	if age := 16; age >= 21 {
		fmt.Println("person is above 21")
	} else if age >= 13 && age <= 19 {
		fmt.Println("person is teen")
	} else {
		fmt.Println("person is child")
	}

	// go does not have ternary operator but we can use if else statement to achieve the same result
	age := 16
	result := "person is child"
	if age >= 21 {
		result = "person is above 21"
	} else if age >= 13 && age <= 19 {
		result = "person is teen"
	}
	fmt.Println(result)
}
