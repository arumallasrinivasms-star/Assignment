package main

import (
	"fmt"
)

//Defer

func defercondition() {
	fmt.Println("Start of the function")

	defer fmt.Println("This is deferred") // will run at the end of main()

	fmt.Println("End of the function")
}

//For loop

func forloop() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Number:", i)
	}
}

//Ifelse

func ifloop() {
	number := 7

	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}
}

//While loop

func whileloop() {
	i := 1

	// while i <= 5
	for i <= 5 {
		fmt.Println("Number:", i)
		i++
	}
}

//Switch

func switchabc() {
	num := -5

	switch {
	case num > 0:
		fmt.Println("Positive number")
	case num < 0:
		fmt.Println("Negative number")
	default:
		fmt.Println("Zero")
	}
}

//String

type Person struct {
	Name string
	Age  int
	City string
}

func sristring() {
	p1 := Person{Name: "Amrutha", Age: 25, City: "Hyderabad"}

	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	fmt.Println("City:", p1.City)

	p1.City = "Bangalore"
	fmt.Println("Updated City:", p1.City)
}

//Pointers

func point() {
	x := 5
	p := &x // p is a pointer to x

	fmt.Println("x =", x)   // prints value of x
	fmt.Println("p =", p)   // prints address of x
	fmt.Println("*p =", *p) // prints value stored at address p (same as x)

	// change value using the pointer
	*p = 10
	fmt.Println("x after change =", x)
}

//Arrays

func Arr() {
	// Declare an array of 5 integers
	var numbers [5]int

	// Assign values
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// Print all elements
	fmt.Println("Array elements:", numbers)

	// Access specific element
	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[4])

	// Using a loop to print all elements
	fmt.Println("Using for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index", i, "=", numbers[i])
	}

}

//fibonacci series

func fib() {

	n := 8
	a, b := 0, 1

	fmt.Println("Fibonacci Series: ")
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

// Max array number

func maxarr() {
	arr := []int{10, 25, 5, 78, 34}
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	fmt.Println("Maximum number in array:", max)
}

// palindrome

func palin() {
	var num, temp, rev int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	temp = num
	for num > 0 {
		rem := num % 10
		rev = rev*10 + rem
		num /= 10
	}

	if temp == rev {
		fmt.Println("Palindrome number")
	} else {
		fmt.Println("Not a palindrome")
	}
}

func main() {

	defercondition()
	forloop()
	ifloop()
	whileloop()
	switchabc()
	sristring()
	point()
	maxarr()
	fib()
	maxarr()
	palin()
}
