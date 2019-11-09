package main

import "fmt"

//同一パッケージ内で使用できる変数
var (
	example = []string{
		"golang",
		"hands-on",
		"in",
		"kagawa",
		"presentation",
		"fast",
		"compile",
	}
)

func main() {
	var data []string
	data = example
	//for文で繰り返し
	for _, v := range data {
		// ○×だけならこの下の文は必要ない
		fmt.Println(v)

		// 値の入力
		fmt.Print("input: ")
		var ans string
		fmt.Scanln(&ans) //&が必要 ... 値のコピー
		// 入力の値の判定
		if v == ans {
			fmt.Println("○")
		} else {
			fmt.Println("×")
		}
	}
	fmt.Println("終了")
}
