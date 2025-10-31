package main

import "fmt"

func main() {
	num := 10

	if num > 0 {
		fmt.Println("It is a positive number.")
	} else if num == 0 {
		fmt.Println("It is a Zero.")
	} else {
		fmt.Println("It is a negative number.")
	}
}
