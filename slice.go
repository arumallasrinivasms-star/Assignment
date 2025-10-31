package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Initial slice:", numbers)
	fmt.Println("Length:", len(numbers), "Capacity:", cap(numbers))

	fmt.Println(numbers[0])

	fmt.Println(numbers[2])

}
