package main

import (
	"reflect"
	"testing"
)


func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		arr []int
		want []int
	}{
		{"sorted array", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse array", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"unsorted array", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}, []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9}},
		{"empty array", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
