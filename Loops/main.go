package main

import "fmt"

func main() {
	/*
		Go only has the for loop (no while or do-while loops). The for loop has three forms: the basic form, the condition-only form, and the range form.
		The basic form of the for loop is the most common. It consists of three parts:

		1. The initialization part is executed before the loop starts. This part is usually used to
		initialize a loop counter variable.

		2. The condition part is evaluated before each iteration of the loop. If the condition is
		true, the loop continues. If the condition is false, the loop ends.

		3. The increment part is executed after each iteration of the loop. This part is usually
		used to increment a loop counter variable.

		The condition-only form of the for loop is similar to the basic form, but it does
		not have an initialization part. This form is useful when you want to loop over a
		range of values that are already initialized.

		The range form of the for loop is used to iterate over the elements of a slice or
		map. It consists of two parts: the iteration variable and the value of the element
		being iterated over. The iteration variable is used to access the elements of the
		slice or map, and the value of the element is used in the loop body.
	*/

	// 1. Basic Form:
	fmt.Println("Basic Form ---->")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 2. Condition-Only Form:
	fmt.Println("Condition-Only Form ---->")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// 3. Infinite Loop:
	fmt.Println("Infinite Loop ---->")
	n := 1
	for {
		if n > 5 {
			break
		}
		fmt.Println(n)
		n++
	}

	// 4. Range Form:
	fmt.Println("Range Form ---->")
	s := []string{"a", "b", "c", "d", "e"}
	for i, v := range s {
		fmt.Println("Index:", i, "Value:", v)
	}

}
