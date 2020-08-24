package dllist

type List struct { //List shows us a structure of a doubly linked list
	root   Item
	length int
}

func (l *List) Len() int { // Len return list length
	return l.length
}

func (l *List) First() *Item { //First returns first element or nil
	if l.length == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Last() *Item { //Last returns last element or nil
	if l.length == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) inserter(i, it *Item) *Item { //inserter inserts i after it, increments list length and returns added element
	i.prev = it
	i.next = it.next
	i.prev.next = i
	i.next.prev = i
	i.list = l
	l.length++
	return i
}

func (l *List) ValueInserter(v interface{}, it *Item) *Item { //ValueInserter is a wrapper for inserter(&Item{Value: v}, it)
	return l.inserter(&Item{Value: v}, it)
}

func (l *List) PushFront(v interface{}) *Item { // PushFront inserts a new item i value v at the front of list l and returns i
	return l.ValueInserter(v, &l.root)
}

func (l *List) PushBack(v interface{}) *Item { // PushBack inserts a new item i value v at the back of list l and returns i
	return l.ValueInserter(v, l.root.prev)
}

func (l *List) remover(i *Item) *Item { // remover removes i from its list, decrements list length and returns i
	i.prev.next = i.next
	i.next.prev = i.prev
	i.next = nil
	i.prev = nil
	i.list = nil
	l.length--
	return i
}

func (l *List) Remove(i *Item) interface{} { //Remove removes i from l if i is an element of list l and returns the item value i.Value
	if i.list == l {
		l.remover(i)
	}
	return i.Value
}

type Item struct { //item is an element of a linked list
	Value interface{}
	next  *Item
	prev  *Item
	list  *List
}

func (i *Item) Next() *Item { // Next returns the next list element or nil
	if p := i.next; i.list != nil && p != &i.list.root {
		return p
	}
	return nil
}

func (i *Item) Prev() *Item { // Next returns the previous list element or nil
	if p := i.prev; i.list != nil && p != &i.list.root {
		return p
	}
	return nil
}

func (l *List) InitList() *List { // InitList initializes or clears list l
	l.root.next = &l.root
	l.root.prev = &l.root
	l.length = 0
	return l
}

func NewList() *List { // NewList returns an initialized list
	return new(List).InitList()
}
