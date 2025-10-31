package main

import "fmt"

func main() {
	defer fmt.Println("This is Printed last.")
	fmt.Println("This is Printed first.")
	fmt.Println("This is Printed second.")
}
