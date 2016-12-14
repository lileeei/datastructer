package ordermap

import (
	"container/list"
	"sync"
)

type Pair struct {
  Key interface{}
  Value interface{}
}

type OrderMap struct {
	sync.RWMutex
	Map  map[interface{}]*Pair
	List list.List
	len  int32
  Compare func(a *Pair, b *Pair) bool
}

// type Compare interface{
//   Compare() bool
// }

func (om *OrderMap) Put(p *Pair) {
	//
}

func (om *OrderMap) Get(key interface{}) interface{} {
	//
}

func (om *OrderMap) GetFirst() *Pair {
	//
}

func (om *OrderMap) GetLast() *Pair {
	//
}

func (om *OrderMap) GetByIndex(index int) *Pair {
	//
}

// func (om *OrderMap) GetAndSet(index int) *Pair {
// 	//
// }

func (om *OrderMap) Delete(key interface{}) {
	//todo
}

func (om *OrderMap) Len() int32 {
	//todo
}

func (om *OrderMap) Traversal() []*Pair {
	//todo
}

func (om *OrderMap) Set() (key interface{}, newvalue interface{}) (oldvalue interface{}) {
	//todo
}




//------------------------------------------------------------------------------
