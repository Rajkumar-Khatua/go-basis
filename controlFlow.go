package main

import "fmt"

func main(){

	// If statement

	age := 25
    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }

	// For loop

	for i:=0; i< 5; i++{
		fmt.Println("Iteration ", i)
	}

	// Switch Statement

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday": 
	fmt.Println("It's Monday")
default:
	fmt.Println("It's another day.")
	
}
}
