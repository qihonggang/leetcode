package cache

import (
	"container/list"
	"github.com/go-programming-tour-book/cache"
)

// fifo 是一个 FIFO cache， 它不是并发安全的
type fifo struct {
	maxBytes int
	onEvicted func(key string, value interface{})
	usedBytes int
	l1 *list.List
	cache map[string]*list.Element
}

type entry struct {
	key string
	value interface{}
}

func (e *entry) Len() int {
	return cache.CalcLen(e.value)
}

func (f *fifo) Set(key string, value interface{}) {
	if e, ok := f.cache[key]; ok {
		f.l1.MoveToBack(e)
		en := e.Value.(*entry)
		f.usedBytes = f.usedBytes - cache.CalcLen(en.value) + cache.CalcLen(value)
		en.value = value
		return
	}
	en := &entry{key, value}
	e := f.l1.PushBack(en)
	f.cache[key] = e

	f.usedBytes += en.Len()
	if f.maxBytes > 0 && f.usedBytes > f.maxBytes {
		f.DelOldest()
	}
}

func (f *fifo) Get(key string) interface{} {
	if e, ok := f.cache[key]; ok {
		return e.Value.(entry).value
	}
	return nil
}

func (f *fifo) Del(key string) {
	if e, ok := f.cache[key]; ok {
		f.removeElement(e)
	}
}

func (f *fifo) DelOldest() {
	f.removeElement(f.l1.Front())
}

func (f *fifo) Len() int {
	return f.l1.Len()
}

func (f *fifo) removeElement(e *list.Element) {
		if e == nil {
			return
		}
		f.l1.Remove(e)
		en := e.Value.(*entry)
		f.usedBytes -= en.Len()
		delete(f.cache, en.key)
		if f.onEvicted != nil {
			f.onEvicted(en.key, en.value)
		}
}

// 用 New 创建一个新的 Cache，如果maxBytes 是 0，则表示没有容量限制
func New(maxBytes int, onEvicted func(key string, value interface{})) cache.Cache {
	return &fifo{
		maxBytes: maxBytes,
		onEvicted: onEvicted,
		l1: list.New(),
		cache: make(map[string]*list.Element),
	}
}