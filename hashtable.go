package main

import (
	"fmt"
	"hash/fnv"
)

// ハッシュテーブルを実装する構造体
type HashTable struct {
	array map[int][]KeyValue
	size  int
}

// キーと値のペアを保持する構造体
type KeyValue struct {
	Key   string
	Value string
}

// 新しいハッシュテーブルを生成する関数
func NewHashTable(size int) *HashTable {
	return &HashTable{
		array: make(map[int][]KeyValue, size),
		size:  size,
	}
}

// キーをハッシュ化する関数
func (h *HashTable) hash(key string) int {
	hf := fnv.New32a()
	hf.Write([]byte(key))
	return int(hf.Sum32()) % h.size
}

// ハッシュテーブルに要素を追加するメソッド
func (h *HashTable) Set(key string, value string) {
	hash := h.hash(key)
	h.array[hash] = append(h.array[hash], KeyValue{Key: key, Value: value})
}

// キーに対応する値を取得するメソッド
func (h *HashTable) Get(key string) (string, bool) {
	hash := h.hash(key)
	for _, kv := range h.array[hash] {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return "", false
}

func main() {
	ht := NewHashTable(10)
	ht.Set("key1", "value1")
	ht.Set("key2", "value2")

	if value, ok := ht.Get("key1"); ok {
		fmt.Println("key1:", value)
	}
}
