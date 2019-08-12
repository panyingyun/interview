package main

//###########题目 LRU Cache##################
//## LRU 最近最少使用算法，LRU算法主要用于缓存淘汰。
//## 主要目的就是把最近最少使用的数据移除内存，以加载其他数据

//===========原理=============
//添加元素时，放到链表头
//缓存命中，将元素移动到链表头
//缓存满了之后，将链表尾的元素删除

//==========LRUCache对外暴露三个接口==============
//func NewLRUCache(capacity int) *LRUCache
//func (l *LRUCache) Put(k interface{}, v interface{})
//func (l *LRUCache) Get(k interface{}) interface{}

//
//举例
// 大小为3的LRU Cache,P1为一个二元结构(k1,v1)
// PUT P1 ----> P1
// PUT P2 ----> P2  P1
// PUT P3 ----> P3  P2  P1
// PUT P4 ----> P4  P3  P2  //加入P4,淘汰最近最久未使用的P1
// GET P2 ----> P2  P4  P3  //P2获取一次，将P2位置放到首位
// PUT P5 ----> P5  P2  P4  //加入P5，淘汰P3

import (
	"container/list"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	lru := NewLRUCache(3)
	lru.Set("P1", "V1")
	lru.PrintAll()
	lru.Set("P2", "V2")
	lru.PrintAll()
	lru.Set("P3", "V3")
	lru.PrintAll()
	lru.Set("P4", "V4")
	lru.PrintAll()
	lru.Get("P2")
	lru.PrintAll()
	lru.Set("P5", "V5")
	lru.PrintAll()
}

// An AtomicInt is an int64 to be accessed atomically.
type AtomicInt int64

// LRUCache is an LRU cache. It is safe for concurrent access.
type LRUCache struct {
	mutex       sync.RWMutex
	maxItemSize int
	cacheList   *list.List
	cache       map[interface{}]*list.Element
	hits, gets  AtomicInt
}

//return status of chache
type CacheStatus struct {
	Gets        int64
	Hits        int64
	MaxItemSize int
	CurrentSize int
}

type entry struct {
	key   interface{}
	value interface{}
}

//NewMemCache If maxItemSize is zero, the cache has no limit.
//if maxItemSize is not zero, when cache's size beyond maxItemSize,start to swap
func NewLRUCache(maxItemSize int) *LRUCache {
	return &LRUCache{
		maxItemSize: maxItemSize,
		cacheList:   list.New(),
		cache:       make(map[interface{}]*list.Element),
	}
}

//Status return the status of cache
func (c *LRUCache) Status() *CacheStatus {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return &CacheStatus{
		MaxItemSize: c.maxItemSize,
		CurrentSize: c.cacheList.Len(),
		Gets:        c.gets.Get(),
		Hits:        c.hits.Get(),
	}
}

//Get value with key
func (c *LRUCache) PrintAll() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	for e := c.cacheList.Front(); e != nil; e = e.Next() {
		fmt.Printf("[%v, %v]  ", e.Value.(*entry).key, e.Value.(*entry).value)
	}
	fmt.Println()
	return
}

//Get value with key
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	c.gets.Add(1)
	if ele, hit := c.cache[key]; hit {
		c.hits.Add(1)
		c.cacheList.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return nil, false
}

//Set a value with key
func (c *LRUCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.cacheList = list.New()
	}

	if ele, ok := c.cache[key]; ok {
		c.cacheList.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}

	ele := c.cacheList.PushFront(&entry{key: key, value: value})
	c.cache[key] = ele
	if c.maxItemSize != 0 && c.cacheList.Len() > c.maxItemSize {
		c.RemoveOldest()
	}
}

//Delete delete the key
func (c *LRUCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cache == nil {
		return
	}
	if ele, ok := c.cache[key]; ok {
		c.cacheList.Remove(ele)
		key := ele.Value.(*entry).key
		delete(c.cache, key)
		return
	}
}

//RemoveOldest remove the oldest key
func (c *LRUCache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.cacheList.Back()
	if ele != nil {
		c.cacheList.Remove(ele)
		key := ele.Value.(*entry).key
		delete(c.cache, key)
	}
}

// Add atomically adds n to i.
func (i *AtomicInt) Add(n int64) {
	atomic.AddInt64((*int64)(i), n)
}

// Get atomically gets the value of i.
func (i *AtomicInt) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}
