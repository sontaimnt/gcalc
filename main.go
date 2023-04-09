package main

import (
	"fmt"
)

var operation string
var num1 int
var num2 int 

func main() {
	fmt.Println("Welcome to gcalc")
	fmt.Println("[a] Addition")
	fmt.Println("[s] Subtraction")
	fmt.Println("[m] Multiplication")
	fmt.Println("[d] Division")
	fmt.Print("Select your operation:- ")
	fmt.Scanf("%d" , &operation)
	switch operation{
		case "a":
			fmt.Print("Enter 1st number")
	}
}
