package main

import "fmt"

// インターフェースはすべての型になりうる
func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	//fmt.Println(x+1) エラー

	x = "A"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

}
