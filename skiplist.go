package skiplist

import (
	"math/rand"
	"time"
)

const (
	SKIPLIST_MAXLEVEL = 32
	SKIPLIST_P        = 0.25
)

type CompareFunc func(a *SkipListNode, b *SkipListNode) bool

//SkipListLevel ...
type SkipListLevel struct {
	levelForward *SkipListNode
	levelBack    *SkipListNode
	span         uint
}

//SkipListNode ...
type SkipListNode struct {
	Key     interface{}
	Data    interface{}
	Forward *SkipListNode
	level   []SkipListLevel
}

//SkipList ...
type SkipList struct {
	head    *SkipListNode
	tail    *SkipListNode
	Level   uint
	length  uint
	Compare func(a *SkipListNode, b *SkipListNode) bool
}

//-----------------------------operation for skiplist
func CreateSkipList() *SkipList {
	skiplist := new(SkipList)

	return skiplist
}

func (skiplist *SkipList) SetCompareFunc(compare CompareFunc) {
	skiplist.Compare = compare
}

func (skiplist *SkipList) InsertNode(node *SkipListNode) {
	//
}

//-----------------------------operation for skipnode
/*
 *todo
 *
**/

func createSkipNode(key interface{}, data interface{}, level uint) *SkipListNode {
	node := new(SkipListNode)
	node.Key = key
	node.Data = data
	node.level = make([]SkipListLevel, level)

	return node
}

//-----------------------------common function
func randLevel() uint {
	var level = 1
	rand.Seed(time.Now().Unix())
	for (rand.Int() & 0xffff) < (SKIPLIST_P * 0xffff) {
		level++
	}

	if level < SKIPLIST_MAXLEVEL {

		return level
	}

	return (level & SKIPLIST_MAXLEVEL)
}
