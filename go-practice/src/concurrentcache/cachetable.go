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
func (table *CacheTable) Add(key interface{}, value interface{}, lifeSpan time.Duration) *CacheItem {
	item := NewCacheItem(key, value, lifeSpan)

	table.Lock()
	table.addItem(item)

	return item
}

// 获取值
func (table *CacheTable) Value(key interface{}) (*CacheItem, error) {
	table.RLock()
	value, ok := table.items[key]
	table.RUnlock()

	if ok {
		value.KeepAlive()
		return value, nil
	}
	// 返回未找到Key
	return nil, ErrKeyNotFound
}

// 调用addItem方法，必须先将table锁住，运行回调函数时解锁
func (table *CacheTable) addItem(item *CacheItem) {
	table.items[item.key] = item
	expireDuration := table.cleanupInterval
	addedCallBack := table.addedCallBack
	table.Unlock()

	// 回调函数调用
	if addedCallBack != nil {
		addedCallBack(item)
	}

	if item.lifeSpan > 0 && (expireDuration == 0 || item.lifeSpan < expireDuration) {
		// TODO 补充过期检查
	}
}
