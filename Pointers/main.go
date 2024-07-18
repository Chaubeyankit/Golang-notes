package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var a int = 10
	var p *int = &a
	fmt.Println("Address of a:", p)
	fmt.Println("Value of a:", *p)

	//Modifying Values through Pointers:
	*p = 20
	fmt.Println("Value of a:", a)

	//Pointers to Structs:
	p1 := Person{"Alice", 25}
	p2 := &p1

	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)

	p2.age = 30
	fmt.Println("Updated Person 1:", p1)

}
