package main

func main() {
	// switch statement is used to compare a variable with multiple values and execute the corresponding case block
	// it is more efficient than if else statement when we have multiple conditions to check

	// this is imple switch
	// i := 2
	// switch i {
	// case 1:
	// 	println("i is 1")

	// case 2:
	// 	println("i is 2")

	// case 3:
	// 	println("i is 3")

	// default: // this is optional, it will be executed if none of the above cases are matched
	// 	println("i is not 1,2,3")
	// }

	// this is switch with multiple cases

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's workday")
	// }

	// type switch

	// whoAmi := func(i interface{}) { // any type can be passed as argument
	// 	switch t := i.(type) {
	// 	case int:
	// 		println("i am an integer")
	// 	case string:
	// 		println("i am a string")
	// 	case bool:
	// 		println("i am a boolean")
	// 	default:
	// 		fmt.Println("other", t)
	// 	}
	// }

	// whoAmi(1)
	// whoAmi("hello")
	// whoAmi(true)
	// whoAmi(3.14)
}
