package task

import "fmt"

// 1. Experiment with appending to slices and observe how the underlying array changes

func Task1() {
	fmt.Printf(">>>>>>>------------->>>>>>> Task 1 <<<<<<<--------------<<<<<<<\n")
	// Initial slice with a capacity of 3
	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3

	fmt.Println("Initial slice:", slice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

	// Append elements to the slice
	slice = append(slice, 4)
	fmt.Println("Slice after appending 4:", slice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

	slice = append(slice, 5)
	fmt.Println("Slice after appending 5:", slice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

	slice = append(slice, 6)
	fmt.Println("Slice after appending 6:", slice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))
}

// Task 2: Use the append and copy functions to manipulate slices and observe their behavior
func Task2() {
	fmt.Printf(">>>>>>>------------->>>>>>> Task 2 <<<<<<<--------------<<<<<<< \n")
	// Initial slice
	original := []int{1, 2, 3}
	fmt.Println("Original slice:", original)

	// Use append to add elements
	appendedSlice := append(original, 4, 5)
	fmt.Println("Appended slice:", appendedSlice)

	// Use copy to copy elements from one slice to another
	copiedSlice := make([]int, len(original))
	copy(copiedSlice, original)
	fmt.Println("Copied slice:", copiedSlice)

	// Modify the copied slice
	copiedSlice[0] = 10

	fmt.Println("Modified copied slice:", copiedSlice)
	fmt.Println("Original slice after modifying copied slice:", original)
}

// Task 3: Create a slice from an array and modify its elements. Print both the array and the slice to see the effect
func Task3() {
	fmt.Printf(">>>>>>>------------->>>>>>> Task 3 <<<<<<<--------------<<<<<<<\n")
	// Create an array
	arr := [5]int{1, 2, 3, 4, 5}

	// Create a slice from the array
	slice := arr[1:4]
	fmt.Println("Initial array:", arr)
	fmt.Println("Initial slice:", slice)

	// Modify the slice
	slice[0] = 10
	slice[1] = 20
	fmt.Println("Modified slice:", slice)
	fmt.Println("Array after modifying slice:", arr)
}
