package main

import "fmt"

func main() {
	num := 10
	ptr := &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)
	fmt.Println("Value using pointer:", *ptr)

	*ptr = 20
	fmt.Println("New value of num:", num)
}
