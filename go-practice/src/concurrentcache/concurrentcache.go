package concurrentcache

import "sync"

var (
	concurrentcache = make(map[string]*CacheTable)
	mutex           sync.RWMutex
)

func Cache(tablename string) *CacheTable {
	mutex.RLock()
	table, exist := concurrentcache[tablename]
	mutex.RUnlock()

	if !exist {
		mutex.Lock()
		// double check
		table, exist = concurrentcache[tablename]
		if !exist {
			table = &CacheTable{
				name:  tablename,
				items: make(map[interface{}]*CacheItem),
			}
			concurrentcache[tablename] = table
		}
		mutex.Unlock()
	}

	return table
}
