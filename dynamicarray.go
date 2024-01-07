package main

// 動的配列
// 配列のサイズが動的に変更できるデータ構造で、連続したメモリ領域に要素を格納する。
// 動的配列はランダムアクセスが可能だが、サイズ変更にはコストがかかる場合がある。

// スライスに要素を追加する関数
// この関数は既存のスライスと追加する要素を受け取り、新しい要素が追加されたスライスを返す。
func appendToSlice(slice []int, elements ...int) []int {
	return append(slice, elements...)
}
