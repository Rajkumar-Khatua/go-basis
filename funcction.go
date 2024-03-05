package main

import "fmt"

func add(x,y int)int{
	return y+x
}

func swap(x,y int)(int, int){
	return y,x 
}

func main(){
	// Call the functions
	sum := add(10,20);
	fmt.Println("Sum", sum);

	// Call the swap function
	a, b := 5,7
	a, b = swap(a, b)

	fmt.Println("After Swap: ", a, b);
}