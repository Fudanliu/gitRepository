package main

import (
	"container/list"
	"fmt"
	"strconv"
)

//用map和链表实现LRU缓存
type LRU struct {
	cache map[int]string //相当于key和内容
	lst   list.List
	cap   int
}

func NewLRU(cap int) *LRU {
	lru := new(LRU)
	lru.lst = list.List{}
	lru.cache = make(map[int]string, cap)
	lru.cap = cap
	return lru
}

//向缓存中添加
func (lru *LRU) Add(key int, value string) {
	//如果缓存没满，将键值对添加进map，键添加进list
	if len(lru.cache) < lru.cap {
		lru.cache[key] = value
		lru.lst.PushFront(key)
	} else { //满了则删除最后一个
		diposibleKey := lru.lst.Back()
		delete(lru.cache, diposibleKey.Value.(int))
		lru.lst.Remove(diposibleKey)
		lru.lst.PushFront(key)
		lru.cache[key] = value
	}
}

//查看是否命中缓存
func (lru *LRU) Find(key int) *list.Element {
	if lru.lst.Len() == 0 {
		return nil
	}
	head := lru.lst.Front()
	for {
		if head == nil {
			return nil
		} else {
			if head.Value == key {
				return head
			} else {
				head = head.Next()
			}
		}
	}
	return nil
}

//是否命中缓存后的决断
func (lru *LRU) Get(key int) (string, bool) {
	value, exist := lru.cache[key]
	ele := lru.Find(key)
	if ele != nil { //命中 将命中的元素放到队首
		lru.lst.MoveToFront(ele)
	} else { //未命中 添加
		lru.Add(key, strconv.Itoa(key))
	}
	if exist {
		return value, exist
	}
	return strconv.Itoa(key), false
}

func (lru LRU) traverse() {
	head := lru.lst.Front()
	for head != nil {
		fmt.Print(head.Value, " ")
		head = head.Next()
	}
	fmt.Println()
}

//一般情况下key和内容不一样，简便起见将内容设为key的string类型
//Get函数即为LRU算法的核心， 当请求页面命中缓存时直接从缓存中获取， 而未命中则应将页面（此处方便起见并非真添加页面）添加入缓存
func Test2() {
	fmt.Printf("=======t2测试========\n")
	lru := NewLRU(10)
	for i := 0; i < 10; i++ {
		lru.Add(i, strconv.Itoa(i)) //9 8 7 6 5 4 3 2 1 0
	}
	lru.traverse()

	for i := 0; i < 10; i += 2 {
		lru.Get(i) //8 6 4 2 0 9 7 5 3 1
	}
	lru.traverse()

	for i := 10; i < 15; i++ {
		lru.Add(i, strconv.Itoa(i)) //14 13 12 11 10 8 6 4 2 0
	}
	lru.traverse()

	for i := 0; i < 10; i++ {
		_, exists := lru.Get(i)
		fmt.Printf("key %d exists %t\n", i, exists) //9 7 5 3 1不存在，8 6 4 2 0存在
		lru.traverse()
	}
}
