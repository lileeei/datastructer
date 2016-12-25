package ordermap

import (
  "ordermap/skiplist"
	"sync"
)

type Pair struct {
  Key interface{}
  Value interface{}
}

type OrderMap struct {
	sync.RWMutex
	Map  map[interface{}]interface{}
	List *skiplist.SkipList
	len  int32
  Compare skiplist.CompareFunc
}

//CreateOrdeMap ...
func CreateOrdeMap(compare Compare) *OrderMap {
  om := new(OrderMap)
  om.Compare = compare
  om.Map := make(map[interface{}]interface{})
  om.List := skiplist.CreateSkipList(skiplist.SKIPLIST_MAXLEVEL, compare)
  om.len = 0
}

//Put ...
func (om *orderMap) Put(key interface{}, value inteface{}) {
  if _, ok := om.Map[key]; ok {
    fmt.Printf("the key %v has been existed!!\n", key)

    return
  }

  putPair := &Pair{Key: key, Value: value}
  om.putPair(putPair)
}

//PutPair ...
func (om *OrderMap) putPair(p *Pair) {
	om.Lock()
  defer om.Unlock()
  if om.List == nil {
    panic("List is nil!!!!")
  }
  node := skiplist.CreateSkipListNode(p.Key, p.Value)
  List.InsertNode(node)
  List.Map[Key] = node
  om.len++
}

//Get ...
func (om *OrderMap) GetValueByKey(key interface{}) interface{} {
	om.RLock()
  defer om.RUnlock()

  value, isexist := om.Map[key]
  if isexist != nil {
    
    return value
  }

  return nil
}

//GetFirst ...
func (om *OrderMap) GetFirst() *Pair {
  om.RLock()
  defer om.RUnlock()

  head := om.List.GetHead()
  if head != nil {
    level0 := head.GetNLevel(0)
    if level0 != nil {
      node := level0.GetForward()

      return &Pair{Key: node.Key, Value: node.Data}
    }
  }

  return nil
}

//GetLast ...
func (om *OrderMap) GetLast() *Pair {
  om.RLock()
  defer om.RUnlock()

  tail := om.List.GetTail()
  if tail != nil {
  
    return &Pair{Key: tail.Key, Value: tail.Data}
  }

  return nil
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
