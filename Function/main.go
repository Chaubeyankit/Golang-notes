package main

import (
	"fmt"
)

// 1. Basic Function:
func add(a int, b int) int {
	return a + b
}

// 2. Multiple Return Values:
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")

	}
	return a / b, nil
}

// 3. Named Return Values:
func rectangleDimensions(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func main() {
	sum := add(1, 2)
	fmt.Println("sum :", sum)

	result, err := divide(4.5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("result :", result)
	}

	area, perimeter := rectangleDimensions(5.0, 3.0)
	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)
}
