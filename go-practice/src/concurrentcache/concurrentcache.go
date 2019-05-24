package concurrentcache

import "sync"

var (
	concurrentCache = make(map[string]*CacheTable)
	mutex           sync.RWMutex
)

func Cache(tableName string) *CacheTable {
	mutex.RLock()
	table, exist := concurrentCache[tableName]
	mutex.RUnlock()

	if !exist {
		mutex.Lock()
		// double check
		table, exist = concurrentCache[tableName]
		if !exist {
			table = &CacheTable{
				name:           tableName,
				items:          make(map[interface{}]*CacheItem),
				expireDuration: 0,
			}
			concurrentCache[tableName] = table
		}
		mutex.Unlock()
	}

	return table
}
