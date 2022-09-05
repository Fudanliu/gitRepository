package main

import (
	"container/heap"
	"fmt"
	"time"
)

//用map和堆 表实现超时缓存

type Item struct {
	Value    int //key
	Deadline int
}
type Heap []*Item

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i].Deadline < h[j].Deadline
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Heap) Push(x interface{}) {
	item := x.(*Item)
	*h = append(*h, item)
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type TimeoutCache struct {
	cache map[int]interface{}
	h     Heap
}

func NewTimeoutCache(cap int) *TimeoutCache {
	tc := new(TimeoutCache)
	tc.cache = make(map[int]interface{}, cap)
	tc.h = make(Heap, 0, 10)
	heap.Init(&tc.h) //传指针是因为源码中Init函数参数是一个实现了Push和Pop的Interface，而只有指针实现了这两个function
	return tc
}

func (tc *TimeoutCache) Add(key int, value interface{}, life int) {
	tc.cache[key] = value
	now := int(time.Now().Unix())
	heap.Push(&tc.h, &Item{key, now + life})
}

func (tc TimeoutCache) Get(key int) (interface{}, bool) {
	value, exists := tc.cache[key]
	return value, exists
}

func (tc *TimeoutCache) Eliminate() {
	for {
		if tc.h.Len() == 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		now := int(time.Now().Unix())
		top := tc.h[0]
		if now > top.Deadline {
			heap.Pop(&tc.h)
			delete(tc.cache, top.Value)
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}
func Test3() {
	fmt.Printf("=======t3测试========\n")
	tc := NewTimeoutCache(10)
	go tc.Eliminate()
	tc.Add(1, "1", 0)
	tc.Add(2, "2", 3)
	tc.Add(3, "3", 4)
	time.Sleep(3 * time.Second)

	for _, key := range []int{1, 2, 3} {
		_, exists := tc.Get(key)
		fmt.Printf("key %d exists %t\n", key, exists) //1不存在，2 3还存在
	}
}
