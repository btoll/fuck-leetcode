package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Key  int
	Val  string
	Next *Node
	Prev *Node
}

type LinkedList struct {
	root  Node
	count int
}

func (l *LinkedList) Add(k int, v string) *Node {
	node := &Node{
		Key:  k,
		Val:  v,
		Next: l.root.Next,
		Prev: &l.root,
	}
	node.Prev.Next = node
	node.Next.Prev = node
	l.count += 1
	return node
}

func (l *LinkedList) Back() *Node {
	return l.root.Prev
}

func (l *LinkedList) Front() *Node {
	return l.root.Next
}

func (l *LinkedList) MoveToFront(node *Node) {
	// Unhook from previous position.
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	// Insert after the root.
	node.Next = l.root.Next
	node.Prev = &l.root

	// Point its new neighbors to itself.
	node.Prev.Next = node
	node.Next.Prev = node
}

func (l *LinkedList) Remove(node *Node) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	node.Next = nil
	node.Prev = nil
	l.count -= 1
}

func (l *LinkedList) Traverse() {
	node := &l.root
	for node.Next != &l.root {
		node = node.Next
		fmt.Println(node.Key, node.Val)
	}
}

func NewLinkedList() *LinkedList {
	l := new(LinkedList)
	l.root.Next = &l.root
	l.root.Prev = &l.root
	return l
}

type LRUCache struct {
	data map[int]*Node
	list *LinkedList
	cap  int
	mu   sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		data: make(map[int]*Node),
		list: NewLinkedList(),
		cap:  capacity,
		mu:   sync.Mutex{},
	}
}

func (l *LRUCache) Get(k int) string {
	defer l.mu.Unlock()
	l.mu.Lock()
	if node, found := l.data[k]; found {
		l.list.MoveToFront(node)
		return node.Val
	}
	return ""
}

func (l *LRUCache) Put(k int, v string) {
	defer l.mu.Unlock()
	l.mu.Lock()
	if node, found := l.data[k]; found {
		node.Val = v
		l.list.MoveToFront(node)
		return
	}
	if l.cap == l.list.count {
		// Pop the last node.
		toRemove := l.list.Back()
		l.list.Remove(toRemove)
		delete(l.data, toRemove.Key)
	}
	l.data[k] = l.list.Add(k, v)
}

func main() {
	lru := NewLRUCache(5)
	lru.Put(5, "five")
	lru.Put(11, "eleven")
	lru.Put(7, "seven")
	lru.Put(2, "two")
	lru.Put(3, "three")
	lru.Put(14, "fourteen")
	lru.Put(17, "seventeen")
	lru.Put(3, "3")
	lru.list.Traverse()
}
