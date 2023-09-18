package main

import "fmt"

func main() {
Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
			fmt.Println("処理しない")
		}
		fmt.Println("処理しない")
	}
	fmt.Println("END")
}
