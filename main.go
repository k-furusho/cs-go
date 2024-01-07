package main

import (
	"fmt"
)

func main() {
	// ハッシュテーブルのテスト
	ht := NewHashTable(10)
	ht.Set("key1", "value1")
	value, exists := ht.Get("key1")
	if exists {
		fmt.Printf("Key: key1, Value: %s\n", value)
	}

	// マージソートのテスト
	arr := []int{3, 2, 5, 1, 6, 4}
	sortedArr := MergeSort(arr)
	fmt.Printf("Original array: %v\n", arr)
	fmt.Printf("Sorted array: %v\n", sortedArr)

	// 動的配列のテスト
	dynamicArray := make([]int, 0)
	dynamicArray = appendToSlice([]int{1, 2}, 3, 5)
	fmt.Printf("Dynamic array: %v\n", dynamicArray)
}
