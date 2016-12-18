package skiplist

//SkipListLevel ...
type SkipListLevel struct {
	levelForward *SkipListNode
	span uint
}

//SkipListNode ...
type SkipListNode struct {
	Data inteface{}
	Forward *SkipListNode
	level []SkipListLevel
}

//SkipList ...
type SkipList struct {
	head *SkipListNode
	tail *SkipListNode
	Level uint
	length uint
	Compare func(a *SkipListNode, b *SkipListNode) bool
}

//-----------------------------operation for skiplist



//-----------------------------operation for skipnode

