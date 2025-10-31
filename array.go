package main

import "fmt"

func main() {
	var numbers [5]int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Array elememts:", numbers)

	fmt.Println("Array elememt at index 2:", numbers[2])

	fmt.Println("iterating through the array:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index %d\n", i, numbers[i])
	}

	colors := [3]string{"red", "green", "blue"}
	fmt.Println("Colors array:", colors)

}
