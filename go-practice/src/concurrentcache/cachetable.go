package concurrentcache

import "sync"

type CacheTable struct {
	sync.RWMutex
	// key is CacheItem.Key, value is CacheItem
	table map[interface{}]CacheItem

	// 添加item的回调函数
	addedCallBack func(item *CacheItem)

	// 删除item的回调函数
	deletedCallBack func(item *CacheItem)
}
