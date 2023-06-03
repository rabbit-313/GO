package main

import "fmt"

func main() {
	// var 変数名 型 = 値
	var i int = 200
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "sog"
	fmt.Println(i3, s3)

	//暗黙的

	i4 := 400
	fmt.Println(i4)
	// i4:= 500 はできない（再定義）
	// 暗黙的な定義は関数の外で定義できない

}
