package main

import "fmt"

func main() {
	// 通常の宣言 型指定
	var x, y int
	// 代入 便利！
	x, y = 1, 2
	fmt.Printf("%d, %d\n", x, y)

	// デバッグに便利
	fmt.Printf("データ型埋め込み %v, リテラル表現 %#v, 型情報 %T\n", [...]int{1, 2, 3}, [...]int{1, 2, 3}, [...]int{1, 2, 3})

	// 宣言と代入一緒にする、返り値が変数の型になる
	n := one()
	fmt.Println(n)
}

func one() int {
	return 1
}
