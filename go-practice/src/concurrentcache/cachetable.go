package concurrentcache

import (
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex

	// table 的名字
	name string
	// key is CacheItem.Key, value is CacheItem
	items map[interface{}]*CacheItem

	// 定时
	cleanupTimer    *time.Timer
	cleanupInterval time.Duration

	// 添加item的回调函数
	addedCallBack func(item *CacheItem)

	// 删除item的回调函数
	deletedCallBack func(item *CacheItem)
}

// 设置添加item回调函数
func (table *CacheTable) SetAddedCallBack(f func(item *CacheItem)) {
	table.Lock()
	defer table.Unlock()
	table.addedCallBack = f
}

// 设置删除item回调函数
func (table *CacheTable) SetDeletedCallBack(f func(item *CacheItem)) {
	table.Lock()
	defer table.Unlock()
	table.deletedCallBack = f
}

// 添加缓存
func (table *CacheTable) Add(key interface{}, value interface{}, lifeSpan time.Duration) {
	item := NewCacheItem(key, value, lifeSpan)
	table.Lock()
	defer table.Unlock()

	table.addItem(item)
}

// 获取值
func (table *CacheTable) Value(key interface{}) {

}

func (table *CacheTable) addItem(item *CacheItem) {

}
