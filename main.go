package main

import (
	"fmt"
)

var version string = "0.1"
var operation int
var num1 int
var num2 int

func main() {
	fmt.Println("Welcome to gcalc")
	fmt.Println("[1] Addition")
	fmt.Println("[2] Subtraction")
	fmt.Println("[3] Multiplication")
	fmt.Println("[4] Division")
	fmt.Print("Select your operation:- ")
	fmt.Scanf("%d", &operation)
	switch operation {
	case 1:
		fmt.Print("Enter 1st number:- ")
		fmt.Scanf("%d", &num1)
		fmt.Print("Enter 2nd number:- ")
		fmt.Scanf("%d", &num2)
		fmt.Println("Answer = ", num1+num2)
	case 2:
		fmt.Print("Enter 1st number:- ")
		fmt.Scanf("%d", &num1)
		fmt.Print("Enter 2nd number:- ")
		fmt.Scanf("%d", &num2)
		fmt.Println("Answer = ", num1-num2)
	case 3:
		fmt.Print("Enter 1st number:- ")
		fmt.Scanf("%d", &num1)
		fmt.Print("Enter 2nd number:- ")
		fmt.Scanf("%d", &num2)
		fmt.Println("Answer = ", num1*num2)
	case 4:
		fmt.Print("Enter 1st number:- ")
		fmt.Scanf("%d", &num1)
		fmt.Print("Enter 2nd number:- ")
		fmt.Scanf("%d", &num2)
		fmt.Println("Answer = ", num1/num2)
	default:
		fmt.Println("Error: Invalid Operation")
	}
}
