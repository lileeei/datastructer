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
	Key  interface{}
	Data interface{}
	//Forward *SkipListNode
	level []SkipListLevel
}

//SkipList ...
type SkipList struct {
	head    *SkipListNode
	tail    *SkipListNode
	level   uint
	length  uint
	Compare func(a *SkipListNode, b *SkipListNode) bool
}

//-----------------------------operation for skiplist
//CreateSkipList ...
func CreateSkipList(level uint, comparefunc CompareFunc) *SkipList {
	skiplist := new(SkipList)
	if level > SKIPLIST_MAXLEVEL {
		level = SKIPLIST_MAXLEVEL
	}

	head := createSkipListNode(nil, nil, level)
	head.Forward = nil
	for _, e := range head.level {
		e.levelBack = nil
		e.levelForward = nil
		e.span = 0
	}

	skiplist.head = head
	skiplist.tail = nil
	skiplist.level = level
	skiplist.length = 0
	skiplist.Compare = comparefunc

	return skiplist
}

//SetCompareFunc ...
func (skiplist *SkipList) SetCompareFunc(compare CompareFunc) {
	if compare == nil {
		panic("compare function is nil!!")
	}

	skiplist.Compare = compare
}

//InsertNode ...
func (skiplist *SkipList) InsertNode(insertnode *SkipListNode) {
	if insertnode == nil {

		return
	}

	//no node here
	if skiplist.tail == nil {
		skiplist.tail = insertnode

		for i, e := range insertnode.level {
			skiplist.head.level[i].levelForward = insertnode
			insertnode.level[i].levelBack = skiplist.head
			insertnode.level[i].span = 0
		}

		return
	}

	node := skiplist.SearchFirstGreater(insertnode)
	if node == nil { //insertnode should be apended to tail
		for i, e := range skiplist.tail.level {
			e.levelForward = insertnode
			insertnode.level[i].levelBack = skiplist.tail
		}

		skiplist.tail = node
	} else {
		for i, e := range node.level {
			e.levelBack.level[i].levelForward = insertnode
			insertnode.level[i].levelBack = node.level[i].levelForward

			insertnode.level[i].levelForward = node
			e.levelBack = insertnode
		}
	}

	leninsertnodelevel := len(insertnode.level)
	currentlevel := len(insertnode.level[0].levelBack)

	if leninsertnodelevel <= currentlevel {

		return
	}

	var span uint = 0
	backnode := insertnode.level[0].levelBack
	for backnode != nil {
		span++
		if len(backnode.level) >= currentlevel {
			for currentlevel <= leninsertnodelevel {
				insertnode.level[currentlevel].levelForward = backnode.level[currentlevel].levelForward
				backnode.level[currentlevel].levelForward.level[currentlevel].levelBack = insertnode
				backnode.level[currentlevel].levelForward.level[currentlevel].span -= span
				backnode.level[currentlevel].levelForward = insertnode
				insertnode.level[currentlevel].levelBack = backnode
				insertnode.level[currentlevel].span = span
				currentlevel++
			}
		}

		backnode = backnode.level[currentlevel].levelBack
	}

}

//DelNode ...
func (skiplist *SkipList) DelNode(node *SkipListNode) bool {
	if node == nil {
		fmt.Println("DelNode node is nil!")

		return false
	}

	for i, e := range node.level {
		backnode := node.level[i].levelBack
		forward := node.level[i].levelForward
		if forward == nil {
			backnode.level[i].levelForward = nil
		} else {
			backnode.level[i].levelForward = forward
			forward.level[i].levelBack = backnode
			forward.level[i].span += e.span
		}
	}

	return true
}

//UpdateNode ...
func (skiplist *SkipListNode) UpdateNode(node *SkipListNode) bool {
	if node == nil {
		fmt.Println("UpdateNode node is nill!!!")

		return false
	}

	if !skiplist.DelNode(node) {
		fmt.Println("UpdateNode DelNode failed!!!")

		return false
	}

	skiplist.InsertNode(node)

	return true
}

//SearchFirstGreater ...
func (skiplist *SkipList) SearchFirstGreater(standernode *SkipListNode) *SkipListNode {
	if skiplist.tail == nil || standernode == nil {

		return nil
	}

	var node *SkipListNode = nil

	lenlevel := len(standernode.level)
	if skiplist.head.level[lenlevel] == nil {

		return skiplist.head
	}

	backnode := skiplist.head.level[lenlevel].levelForward

	for backnode != nil && skiplist.Compare(standernode, backnode) {
		backnode = backnode.levelForward
	}

	return backnode.level[lenlevel].levelBack
}

//Traversal ..
func (skiplist *SkipList) Traversal() {
	if skiplist.head.level[0] == nil {
		fmt.Println("the skiplist is nil!")

		return
	}

	node := skiplist.head.level[0].levelForward
	for node != nil {
		fmt.Printf("key: %v\tData: %#v\n", node.Key, node.Data)
	}
}

//-----------------------------operation for skipnode
//CreateSkipListNode ...
func CreateSkipListNode(key interface{}, data interface{}) *SkipListNode {
	level := randLevel()

	return createSkipListNode(key, data, level)
}

//createSkipListNode ...
func createSkipListNode(key interface{}, data interface{}, level uint) *SkipListNode {
	node := new(SkipListNode)
	node.Key = key
	node.Data = data
	node.level = make([]SkipListLevel, level)

	for _, e := range node.level {
		e.levelBack = nil
		e.levelForward = nil
		e.span = nil
	}

	return node
}

//-----------------------------common function
//randLevel ...
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
