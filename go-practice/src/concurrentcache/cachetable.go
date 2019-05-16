package concurrentcache

import (
	"fmt"
	"sync"
	"text/tabwriter"
	"time"
)

type CacheTable struct {
	sync.RWMutex

	// table 的名字
	name string
	// key is CacheItem.Key, value is CacheItem
	items map[interface{}]*CacheItem

	// 过期检查定时器
	expirationCheckTimer *time.Timer
	expireDuration       time.Duration

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
	table.items[key] = item
	addedCallBack := table.addedCallBack
	expireDuration := table.expireDuration
	table.Unlock()

	if addedCallBack != nil {
		addedCallBack(item)
	}

	if lifeSpan > 0 && (0 == expireDuration || lifeSpan < expireDuration) {
		table.expirationCheck()
	}

	return item
}

// 未查找到添加
func (table *CacheTable) AddIfNotFound(key interface{}, value interface{}, liftSpan time.Duration) *CacheItem {
	//TODO
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

func (table *CacheTable) Delete(key interface{}) {

}

// 设置定时机制，过期检查
// 遍历所有的item集合，获取到最迫切smallExpireDuration; 设置expireDuration后检查的Timer
func (table *CacheTable) expirationCheck() {
	table.Lock()
	defer table.Unlock()

	// 1. 先关闭原来的
	if table.expirationCheckTimer != nil {
		table.expirationCheckTimer.Stop()
	}

	// 2. 记录最小的过期检查周期
	smallExpireDuration := 0 * time.Second
	now := time.Now()
	for key, item := range table.items {
		item.RLock()
		liftSpan := item.lifeSpan
		lastVisitedTime := item.lastVisitedTime
		item.RUnlock()

		if liftSpan == 0 {
			continue
		}
		if now.Sub(lastVisitedTime) > liftSpan {
			// 过期了
			table.Delete(key)
		} else {
			if 0 == smallExpireDuration || liftSpan-now.Sub(lastVisitedTime) < smallExpireDuration {
				smallExpireDuration = liftSpan - now.Sub(lastVisitedTime)
			}
		}
	}

	// 3. 设置定时器
	table.expireDuration = smallExpireDuration
	if smallExpireDuration > 0 {
		table.expirationCheckTimer = time.AfterFunc(smallExpireDuration, func() {
			go table.expirationCheck()
		})
		fmt.Printf("expiration check after %s, table name %s", smallExpireDuration, table.name)
	}
}
