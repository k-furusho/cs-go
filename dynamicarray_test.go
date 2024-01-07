package main

import (
	"reflect"
	"testing"
)

// appendToSlice関数のテスト
func TestAppendToSlice(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elements []int
		want     []int
	}{
		{"empty slice", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"non-empty slice", []int{1, 2}, []int{3}, []int{1, 2, 3}},
		{"single element", []int{}, []int{1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appendToSlice(tt.slice, tt.elements...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
