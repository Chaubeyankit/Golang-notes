package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Let's study the time in golang ⌛")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}
