package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	person1 := Person{
		Name: "Sri",
		Age:  25,
		City: "Parimi",
	}

	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("City:", person1.City)

	fmt.Println("Full struct:", person1)
}
