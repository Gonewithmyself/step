package cmp

import (
	"testing"
)

func TestNewRuneTrie(t *testing.T) {
	tr := NewRuneTrie()

	words := []string{"我靠", "fuck", "cao", "damn", "我去", "我顶你个"}
	for i := range words {
		tr.Put(words[i], i)
	}

	tr.Walk(func(s string, i interface{}) error {
		t.Log(s, i)
		return nil
	})

	tr.walkPrefix("我顶", func(s string, i interface{}) error {
		t.Log(s, i)
		return nil
	})
}
