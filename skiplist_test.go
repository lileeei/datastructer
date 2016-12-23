package skiplist

import (
	"testing"
)

func Compare(a *SkipListNode, b *SkipListNode) bool {
	value1 := a.Key.(int)
	value2 := b.Key.(int)

	if value1 > value2 {
		return true
	}

	return false
}

func TestSkipList(t *testing.T) {
	skiplist := CreateSkipList(SKIPLIST_MAXLEVEL, nil)
	skiplist.SetCompareFunc(Compare)

	for i := 0; i < 1000000; i++ {
		node := CreateSkipListNode(i, i)
		skiplist.InsertNode(node)
	}

	skiplist.Traversal()
}
