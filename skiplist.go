package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	SKIPLIST_MAXLEVEL = 16
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

	head := createSkipListNode(nil, nil, SKIPLIST_MAXLEVEL)
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

		for i, _ := range insertnode.level {
			skiplist.head.level[i].levelForward = insertnode
			insertnode.level[i].levelBack = skiplist.head
			insertnode.level[i].span = 0
		}

		return
	}

	node := skiplist.SearchInsertNodeBack(insertnode)
	for i, e := range node.level {
		insertnode.level[i].levelBack = node
		if e.levelForward != nil {
			insertnode.level[i].levelForward = e.levelForward
			insertnode.level[i].span = 1
			e.levelForward.level[i].span -= 1
			e.levelForward = insertnode
		}
	}

	leninsertnodelevel := len(insertnode.level)
	currentlevel := len(insertnode.level[0].levelBack.level)

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
func (skiplist *SkipList) UpdateNode(node *SkipListNode) bool {
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
func (skiplist *SkipList) SearchInsertNodeBack(standernode *SkipListNode) *SkipListNode {
	if standernode == nil {
		panic("SearchFirstGreater standernode is nil!!!!")
	}

	lenlevel := len(standernode.level)
	fmt.Printf("SearchInsertNodeBack\nstandernode.level: %v\nstandernode.key: %v\n", len(standernode.level), standernode.Key)

	if skiplist.tail == nil {
		fmt.Println("true")
		return skiplist.head
	}

	forwardnode := skiplist.head.level[lenlevel].levelForward

	for forwardnode != nil && skiplist.Compare(standernode, forwardnode) {
		forwardnode = forwardnode.level[lenlevel].levelForward
	}

	return forwardnode.level[lenlevel].levelBack
}

func (skiplist *SkipList) compare(standernode, node *SkipListNode, currentlevel uint) *SkipListNode {
	currentlevel := len(skiplist.head)
	currentnode := skiplist.head
	for currentlevel >= 0 {
		forwardnode := currentnode.level[currentlevel].levelForward
		for forwardnode != nil && skiplist.Compare(standernode, forwardnode) {
			currentnode = forwardnode
			forwardnode = forwardnode.level[currentlevel].levelForward
			if forwardnode != nil && !skiplist.Compare(standernode, forwardnode) {
				currentlevel--
				break
			}
		}
		currentlevel--
	}
	
	return currentnode
}

//Traversal ..
func (skiplist *SkipList) Traversal() {
	if skiplist.head.level[0].levelForward == nil {
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
		e.span = 0
	}

	return node
}

//-----------------------------common function
//randLevel ...
func randLevel() uint {
	var level uint = 1
	rand.Seed(time.Now().Unix())
	for (rand.Int() & SKIPLIST_MAXLEVEL) < int(SKIPLIST_MAXLEVEL*SKIPLIST_P) {
		level++
	}

	if level < SKIPLIST_MAXLEVEL {

		return level
	}

	return (level & SKIPLIST_MAXLEVEL)
}
