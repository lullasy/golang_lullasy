package main

import "fmt"

// パッケージ変数実質グローバル
var pkg = 100

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

	pkg = pkg + 1
	fmt.Println(pkg)

	a := 1
	// b := 2 error
	// c := 3 error
	fmt.Println(a)
}

func one() int {
	return 1
}
