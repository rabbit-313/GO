package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	defer panic("runtime error")
	fmt.Println("Start")
}
