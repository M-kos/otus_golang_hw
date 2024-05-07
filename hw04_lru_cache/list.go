package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}, key Key) *ListItem
	PushBack(v interface{}, key Key) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Key   Key
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	zero ListItem
	len  int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}

	return l.zero.Next
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}

	return l.zero.Prev
}

func (l *list) PushFront(v interface{}, key Key) *ListItem {
	item := newListItem(v, key)

	if l.len == 0 {
		l.zero.Prev = item
	}

	if l.zero.Next != nil {
		item.Next = l.zero.Next
		l.zero.Next.Prev = item
	}

	l.zero.Next = item
	l.len++

	return item
}

func (l *list) PushBack(v interface{}, key Key) *ListItem {
	item := newListItem(v, key)

	if l.len == 0 {
		l.zero.Next = item
	}

	if l.zero.Prev != nil {
		item.Prev = l.zero.Prev
		l.zero.Prev.Next = item
	}

	l.zero.Prev = item
	l.len++

	return item
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.zero.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Next = nil
	i.Prev = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	i.Prev.Next = i.Next

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.zero.Prev = i.Prev
	}

	i.Next = l.zero.Next

	l.zero.Next.Prev = i
	l.zero.Next = i
	i.Prev = nil
}

func NewList() List {
	return new(list)
}

func newListItem(v interface{}, key Key) *ListItem {
	return &ListItem{Value: v, Key: key}
}
