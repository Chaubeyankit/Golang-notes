package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the order please")
	orderName, _ := reader.ReadString('\n')
	fmt.Println("Your order is", orderName)

	// Conversion

	fmt.Println("Please rate the Order in between 1 to 5")
	Reader := bufio.NewReader(os.Stdin)
	rating, _ := Reader.ReadString('\n')
	fmt.Println("Your rating is", rating)

	// fmt.Println(rating + 1) its throw error because
	// we are trying to add string with int
	// so we need to convert string to int
	// we can do it by using Atoi function
	// Atoi is a function of strconv package
	// strconv is a package of go

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("We added 1 to your rating:", numRating+1)
	}
}
