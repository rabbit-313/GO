package main

import "fmt"

func main() {
	switch n := 1; n {

	case 1, 3:
		fmt.Println("1 or 3")

	default:
		fmt.Println("default")
	}
}
