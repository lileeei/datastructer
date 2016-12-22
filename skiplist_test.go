package skiplist

import (
	"fmt"
	"testing"
)

func Compare(a *SkipListNode, b *SkipListNode) bool {
	value1 := a.Key.(string)
	value2 := b.Key.(string)

	if value1 > value2 {
		return true
	}

	return false
}

func TestSkipList(t *testing.T) {
	skiplist := CreateSkipList(SKIPLIST_MAXLEVEL, nil)
	skiplist.SetCompareFunc(Compare)

	for i := 0; i < 100; i++ {
		node := CreateSkipListNode(fmt.Sprintf("%d", i), i)
		skiplist.InsertNode(node)
	}

	skiplist.Traversal()
}
