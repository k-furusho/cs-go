package main

import (
	"fmt"
)

// 動的配列（スライス）の例
func main() {
	// 空のスライスを作成
	dynamicArray := make([]int, 0)

	// 要素を追加
	dynamicArray = append(dynamicArray, 1)
	dynamicArray = append(dynamicArray, 2)
	dynamicArray = append(dynamicArray, 3)

	// スライスの内容を表示
	fmt.Println(dynamicArray)
}
