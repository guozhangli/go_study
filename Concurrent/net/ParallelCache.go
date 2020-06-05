package net

import (
	"sync"
	"time"
)

const MAX_LIVING_TIME_MILLIS = 60000 //60s

var interrupt = false

type CacheItem struct {
	command    string
	data       string
	createTime int64
	accessTime int64 //最后访问时间
}

type ParallelCache struct {
	concurrentMap sync.Map
}

func NewParallelCache() *ParallelCache {
	pc := new(ParallelCache)
	go func() {
		for !interrupt {
			time.Sleep(10 * time.Second)
			pc.cleanCache()
		}
	}()
	return pc
}

func (pc *ParallelCache) put(key string, value string) {
	cacheItem := &CacheItem{
		command:    key,
		data:       value,
		createTime: time.Now().Unix(),
		accessTime: 0,
	}
	pc.concurrentMap.Store(key, cacheItem)
}

func (pc *ParallelCache) get(key string) (string, bool) {
	if v, ok := pc.concurrentMap.Load(key); ok {
		v.(*CacheItem).accessTime = time.Now().Unix()
		return v.(*CacheItem).data, ok
	}
	return "", false
}

func (pc *ParallelCache) cleanCache() {
	pc.concurrentMap.Range(func(key, value interface{}) bool {
		v := value.(*CacheItem)
		if v.accessTime-v.createTime > MAX_LIVING_TIME_MILLIS {
			pc.concurrentMap.Delete(key)
		}
		return true
	})
}

func (pc *ParallelCache) shutDown() {
	interrupt = true
}

func (pc *ParallelCache) getItemCount() int {
	var count int
	pc.concurrentMap.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}
