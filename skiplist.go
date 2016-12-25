package skiplist

import (
	"fmt"
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

		panic("InsertNode insertnode is nil!!")
	}

	defer func() {
		skiplist.length++
	}()

	//no node here
	if skiplist.tail == nil {
		skiplist.tail = insertnode

		for i, _ := range insertnode.level {
			skiplist.head.level[i].levelForward = insertnode
			insertnode.level[i].levelBack = skiplist.head
			insertnode.level[i].span = 1
		}

		return
	}

	leninsertnodelevel := len(insertnode.level)
	node := skiplist.SearchNodeBack(insertnode)
	//fmt.Printf("backnode:\n\tkey: %s\n\tlevel: %d\n\n", node.Key, len(node.level))
	for i, _ := range node.level {
		// fmt.Printf("backnode.key: %s, backnode.level: %d\n", node.Key, i)
		if i >= leninsertnodelevel {
			break
		}

		insertnode.level[i].levelBack = node
		if node.level[i].levelForward != nil {
			// fmt.Println("enter if")
			insertnode.level[i].levelForward = node.level[i].levelForward
			insertnode.level[i].span = 1
			node.level[i].levelForward.level[i].span -= 1
			node.level[i].levelForward = insertnode
		} else {
			// fmt.Println("enter else")
			insertnode.level[i].levelForward = nil
			node.level[i].levelForward = insertnode
			insertnode.level[i].span = 1
		}
	}

	if leninsertnodelevel <= len(insertnode.level[0].levelBack.level) {

		return
	}

	skiplist.fixInsert(insertnode)

}

// func (skiplist *SkipList)insertnode(standernode *SkipListNode) {
// 	currentlevel := len(skiplist.head.level) - 1
// 	currentnode := skiplist.head
// 	for currentlevel >= 0 {
// 		//fmt.Printf("currentnode: %#v\n", currentnode)
// 		//fmt.Printf("currentlevel: %d\n", currentlevel)
// 		//fmt.Printf("currentnode.level[%d]: %#v\n", currentlevel, currentnode.level[currentlevel])
// 		forwardnode := currentnode.level[currentlevel].levelForward
// 		for forwardnode != nil && skiplist.Compare(standernode, forwardnode) {
// 			currentnode = forwardnode
// 			forwardnode = forwardnode.level[currentlevel].levelForward
// 			if forwardnode != nil && !skiplist.Compare(standernode, forwardnode) {
// 				currentlevel--
// 				break
// 			}
// 		}

// 		if forwardnode != nil {
			
// 		}
// 		currentlevel--
// 	}

// 	return currentnode
// }

func (skiplist *SkipList) fixInsert(insertnode *SkipListNode) {
	var span uint = 0
	currentlevel := len(insertnode.level[0].levelBack.level)
	leninsertnodelevel := len(insertnode.level)

	backnode := insertnode.level[0].levelBack
	for backnode != nil {
		span++
		for len(backnode.level) > currentlevel {
			if currentlevel < leninsertnodelevel {
				insertnode.level[currentlevel].levelForward = backnode.level[currentlevel].levelForward

				if insertnode.level[currentlevel].levelForward != nil {
					backnode.level[currentlevel].levelForward.level[currentlevel].levelBack = insertnode
					// insertnode.level[currentlevel].levelForward.level[currentlevel].levelBack = insertnode
					insertnode.level[currentlevel].levelForward.level[currentlevel].span -= span
				}

				backnode.level[currentlevel].levelForward = insertnode
				insertnode.level[currentlevel].levelBack = backnode
				insertnode.level[currentlevel].span = span
				currentlevel++
			} else {
				break
			}
		}

		if currentlevel >= leninsertnodelevel {
			break
		} else {
			backnode = backnode.level[len(backnode.level)-1].levelBack
		}

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
func (skiplist *SkipList) SearchNodeBack(standernode *SkipListNode) *SkipListNode {
	if standernode == nil {
		panic("SearchFirstGreater standernode is nil!!!!")
	}

	//	lenlevel := len(standernode.level)
	//fmt.Printf("SearchInsertNodeBack\nstandernode.level: %v\tstandernode.key: %v\n", len(standernode.level), standernode.Key)

	if skiplist.tail == nil {
		fmt.Println("true")
		return skiplist.head
	}

	// forwardnode := skiplist.head.level[lenlevel].levelForward

	// for forwardnode != nil && skiplist.Compare(standernode, forwardnode) {
	// 	forwardnode = forwardnode.level[lenlevel].levelForward
	// }

	// return forwardnode.level[lenlevel].levelBack
	return skiplist.compare(standernode)
}

func (skiplist *SkipList) compare(standernode *SkipListNode) *SkipListNode {
	currentlevel := len(skiplist.head.level) - 1
	currentnode := skiplist.head
	for currentlevel >= 0 {
		//fmt.Printf("currentnode: %#v\n", currentnode)
		//fmt.Printf("currentlevel: %d\n", currentlevel)
		//fmt.Printf("currentnode.level[%d]: %#v\n", currentlevel, currentnode.level[currentlevel])
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
	fmt.Printf("\nTraversal:\n")
	for node != nil {
		fmt.Printf("key: %v\tData: %#v\tlevellen: %d\n", node.Key, node.Data, len(node.level))
		node = node.level[0].levelForward
	}
	fmt.Printf("skiplist.len: %d\n\n", skiplist.length)
}

//GetHead ...
func (skiplist *SkipList)GetHead() *SkipListNode {
	if skiplist.head == nil {
		panic("the skiplist head is nil!!!")
	}

	return skiplist.head
}

//GetTail ...
func (skiplist *SkipList)GetTail() *SkipListNode {
	if skiplist.tail == nil {
		panic("the skiplist tail is nil!!")
	}

	return skiplist.tail
}

//-----------------------------operation for skipnode
//CreateSkipListNode ...
func CreateSkipListNode(key interface{}, data interface{}) *SkipListNode {
	level := randLevel()
	//fmt.Printf("level: %d\n", level)

	return createSkipListNode(key, data, level+1)
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

//GetNLevelForward ...
func  (slnode *SkipListNode)GetNLevel(n int) *SkipListLevel{
	if n > len(slnode) {
		fmt.Printf("GetNLevelForward n > len(slnode)")
		return nil
	}

	return slnode.level[uint(n)]
}


//-----------------------------SkipListLevel function
//GetForward ...
func (sllevel *SkipListLevel)GetForward() *SkipListNode{
	return sllevel.levelForward
}

//GetBack ...
func (sllevel *SkipListLevel)GetBack() *SkipListNode{
	return sllevel.levelBack
}

//-----------------------------common function
//randLevel ...
func randLevel() uint {
	var level uint = 1
	rand.Seed(time.Now().UnixNano())
	jugeline := SKIPLIST_MAXLEVEL * SKIPLIST_P
	for (rand.Int() & SKIPLIST_MAXLEVEL) < int(jugeline) {
		level++
	}
	// for (level < SKIPLIST_MAXLEVEL) && (rand.Int()%4 == 0) {
	// 	level++
	// }
	if level < SKIPLIST_MAXLEVEL {

		return level
	}

	return (level & SKIPLIST_MAXLEVEL)

	// return level
}
