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

//Put ...
func (om *orderMap) Put(key interface{}, value inteface{}) {
		om.Lock()
  defer om.Unlock()

  if om.List == nil {
    om.List = list.New()
  }

  for e := om.List.Front(); e != nil; e = e.Next() {
    if !om.Compare(p, e) {
      om.List.InsertBefore(p, e)
    }
  }

}

//PutPair ...
func (om *OrderMap) PutPair(p *Pair) {
	om.Lock()
  defer om.Unlock()

  if om.List == nil {
    om.List = list.New()
  }

  for e := om.List.Front(); e != nil; e = e.Next() {
    if !om.Compare(p, e) {
      om.List.InsertBefore(p, e)
    }
  }
}

//Get ...
func (om *OrderMap) Get(key interface{}) interface{} {
	om.RLock()
  defer om.RUnlock()

  if v, ok := om.Map[key]; ok {

    return v.Value
  }

  return nil
}

//GetFirst ...
func (om *OrderMap) GetFirst() *Pair {
  om.RLock()
  defer om.RUnlock()

  return om.List.Front()
}

//GetLast ...
func (om *OrderMap) GetLast() *Pair {
  om.RLock()
  defer om.RUnlock()

  return om.List.Back()
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

//返回ordermap当前元素的数量
func (om *OrderMap) Len() int32 {
	om.RLock()
  	defer om.RUnlock()

	return om.len
}

func (om *OrderMap) Traversal() []*Pair {
	//todo
}

func (om *OrderMap) Set() (key interface{}, newvalue interface{}) (oldvalue interface{}) {
	//todo
}




//------------------------------------------------------------------------------
