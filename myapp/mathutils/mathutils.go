package mathutils

import "fmt"

func FibonacciSequence(n int) {
	a := 0
	b := 1
	var (
		c, i int
	)
	if n == 0 || n == 1 {
		fmt.Println(n)
		return
	}
	for i = 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	fmt.Println(b)
}
