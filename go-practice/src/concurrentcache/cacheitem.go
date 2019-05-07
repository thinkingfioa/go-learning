package concurrentcache

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex
	// item 的键
	key interface{}
	// item 的值
	value interface{}

	// 存活时间
	lifeSpan time.Duration
	// 创建时间
	createdTime time.Time
	// 最近访问时间
	lastVisitedTime time.Time
	// 访问次数
	visitedCount int64
}

// 创建一个CacheItem变量，返回指针
func NewCacheItem(key interface{}, value interface{}, leftSpan time.Duration) *CacheItem {
	t := time.Now()
	return &CacheItem{
		key:             key,
		value:           value,
		lifeSpan:        leftSpan,
		createdTime:     t,
		lastVisitedTime: t,
		visitedCount:    0,
	}
}

// 更新最新访问时间按
func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()
	item.lastVisitedTime = time.Now()
	item.visitedCount++
}

func (item *CacheItem) Key() interface{} {
	// 不可变
	return item.key
}

func (item *CacheItem) Value() interface{} {
	// 不可变
	return item.value
}

func (item *CacheItem) LifeSpan() time.Duration {
	// 不可变
	return item.lifeSpan
}

func (item *CacheItem) CreatedTime() time.Time {
	return item.createdTime
}

func (item *CacheItem) VisitedTime() time.Time {
	return item.VisitedTime()
}

func (item *CacheItem) VisitedCount() int64 {
	return item.visitedCount
}
