package main

import (
	task "Arrays/Task"
	// "fmt"
)

func main() {
	// Using ... to let the compiler determine the array size:
	
	// arr := [...]int{1, 2, 3, 4, 5}
	// fmt.Println("First element:", arr[0])
	// fmt.Println("Second element:", arr[1])

	// for index, value := range arr {
	// 	fmt.Println("Index:", index, "Value:", value)
	// }

	//Slices
	/*
		Slices are more flexible and powerful than arrays. A slice is a dynamically-sized, flexible view into the elements of an array.
		In practice, slices are much more common than arrays. Slices are like smart pointers into
		an array that allow you to read and write the elements. Slices are like references to
		arrays; they do not store any data, they describe a section of an underlying array.
	*/

	//Declaring Slices
	// slice := arr[1:4]
	// slice[0] = 10
	// fmt.Println(slice)

	//1. Appending to a Slice:
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 4, 5)
	// fmt.Println(slice)

	//2. Copying Slices:
	src := []int{1, 2, 3}

	dest := make([]int, len(src))
	copy(dest, src)
	// fmt.Println(dest)

	//Task

	task.Task1()
	task.Task2()
	task.Task3()

	//Output Explanation

	/*
		Task 1: Appending to Slices

		1. The append function takes a slice as its first argument and one or more values as
		its subsequent arguments. It returns a new slice that is the result of appending the
		values to the original slice. The original slice is not modified.

		2. Initially, the slice has a length and capacity of 3. After appending elements, the capacity doubles when it exceeds the original capacity, showing how Go handles dynamic resizing of slices.

		Task 2: Using append and copy

		1. Appending elements to a slice creates a new slice if the original capacity is exceeded.
		2. The copy function copies elements from one slice to another. Modifying the copied slice does not affect the original slice.

		Task 3: Slices from Arrays

		1. Slices are references to arrays. Modifying a slice modifies the underlying array because they share the same memory.
	*/

}
