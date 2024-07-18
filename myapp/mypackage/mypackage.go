package mypackage

import "fmt"

func Greet(name string) {
	fmt.Println("Hello,", name)
}

func IsStringPalindrome(str string) bool {
	if len(str) == 0 {
		return true
	}
	for i := 0; i < (len(str))/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}
