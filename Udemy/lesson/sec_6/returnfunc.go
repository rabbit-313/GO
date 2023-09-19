//関数を返す関数
package main

import "fmt"

func ReturnFunc() func() {
	return func() {
		fmt.Println("I`m funciton")
	}
}

func main() {
	f := ReturnFunc()
	f()
}
