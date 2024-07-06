package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Clear() {
	l.items = make(map[Key]*ListItem, l.capacity)
	l.queue = NewList()
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if i, ok := l.items[key]; ok {
		l.queue.MoveToFront(i)

		if v, ok := i.Value.(ListItemValue); ok {
			return v.Value, ok
		}

		return i.Value, ok
	}

	return nil, false
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if i, ok := l.items[key]; ok {
		i.Value = ListItemValue{Key: key, Value: value}
		l.queue.MoveToFront(i)
		return true
	}

	item := l.queue.PushFront(ListItemValue{Key: key, Value: value})

	if l.queue.Len() > l.capacity {
		lastItem := l.queue.Back()
		l.queue.Remove(lastItem)

		if item, ok := lastItem.Value.(ListItemValue); ok {
			delete(l.items, item.Key)
		}
	}

	l.items[key] = item

	return false
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
