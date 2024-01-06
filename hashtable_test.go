package main

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	ht := NewHashTable(10)

	// Set メソッドのテスト
	ht.Set("key1", "value1")
	ht.Set("key2", "value2")

	// Get メソッドのテスト
	value, exists := ht.Get("key1")
	if !exists || value != "value1" {
		t.Errorf("Get(\"key1\") = %v, %v; want \"value1\", true", value, exists)
	}

	value, exists = ht.Get("key2")
	if !exists || value != "value2" {
		t.Errorf("Get(\"key2\") = %v, %v; want \"value2\", true", value, exists)
	}

	value, exists = ht.Get("key3")
	if exists || value != "" {
		t.Errorf("Get(\"key3\") = %v, %v; want \"\", false", value, exists)
	}
}