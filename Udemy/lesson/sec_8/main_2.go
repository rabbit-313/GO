package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "1000"

	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("i = %T\n", i)
}