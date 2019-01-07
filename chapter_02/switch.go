package main

import (
	"fmt"
)

func main() {

	a := 10

	switch a == 10 {
	case true:
		fmt.Println("Yes")
	case false:
		fmt.Println("No!")
	}

	a = 12
	switch a {
	case 10:
		fmt.Println("10")
	case 9:
		fmt.Println("9")
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("None of these")
	}
}
