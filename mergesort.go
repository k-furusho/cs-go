package main

// マージソート
// 分割統治法を用いた比較ベースのソートアルゴリズム。
// 配列を半分に分割し、それぞれを再帰的にソートした後、二つのソートされたリストをマージして一つのソートされたリストを生成する。
// マージ操作が等しい要素の相対的な順序を変更しないので、同じ値を持つ要素が入力でどのような順序であっても、出力でも同じ順序で保持される。(つまり、マージソートは安定ソート)

// マージソートを実装する関数
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// 二つの配列をマージする関数
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}