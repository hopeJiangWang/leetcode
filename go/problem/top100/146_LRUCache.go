package top100

import (
	// "fmt"
	// "container/list"
)

/*
	哈希表+双向链表：因为需要删除操作，删除一个节点不光要得到该节点本身的指针，也需要操作其前驱节点的指针
	而双向链表才能支持直接查找前驱，保证操作的时间复杂度O（1）
*/

/*
	首先，实现链表节点
*/

type Node struct {
	Key, Value int	// 键、值
	Pre, Next *Node // 前、后
}

func initNode(key, val int) *Node {
	return &Node{
		Key: key,
		Value: val,
	}
}

/*
	然后，实现双向链表
*/
type DoubleLinkedList struct {
	Head, Tail *Node	// 虚拟头、尾节点
	Size int			// 链表元素数
}

func initDoubleLinkedList() *DoubleLinkedList {
	res := &DoubleLinkedList{
		Head: initNode(0, 0),
		Tail: initNode(0, 0),
		Size: 0,
	}

	// 双向
	res.Head.Next = res.Tail
	res.Tail.Pre = res.Head

	return res
}

// 在尾部插入新节点 O(1)
func (cache *DoubleLinkedList)addNodeAtLast(n *Node) {
	/*
		n 的next 指向表尾, n 的pre 指向最后一个节点
		最后一个节点的next 指向n ，表尾的pre 指向n
		size加一
	*/
	n.Next = cache.Tail
	n.Pre = cache.Tail.Pre

	cache.Tail.Pre.Next = n
	cache.Tail.Pre = n

	cache.Size++
}

// 删除节点 O(1)
func (cache *DoubleLinkedList)removeNode(n *Node) {
	/*
		n 的前一个节点的next 指向n 的next
		n 的后一个节点的pre 指向n 的pre
		释放n
		size减一
	*/
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre

	n.Next = nil
	n.Pre = nil

	cache.Size--
}

// 删除头部第一个节点 O(1)
func (cache *DoubleLinkedList)removeFirstNode() *Node {
	/*
		找到头部第一个节点，然后删除它
	*/
	// 为空，返回空
	if cache.Head.Next == cache.Tail {
		return nil
	}

	// 否则，直接删除头部第一个节点
	first := cache.Head.Next
	cache.removeNode(first)
	return first
}

/*
	哈希表+双向链表：因为需要删除操作，删除一个节点不光要得到该节点本身的指针，也需要操作其前驱节点的指针
	而双向链表才能支持直接查找前驱，保证操作的时间复杂度O（1）
*/
type LRUCache struct {
	Capacity int			// 容量
	KeyMap map[int]*Node	// 哈希表（key, node）
	Cache *DoubleLinkedList	// 双向链表
}

// 初始化
func Constructor(capacity int) LRUCache {
	return LRUCache {
		Capacity: capacity,
		KeyMap: map[int]*Node{},
		Cache: initDoubleLinkedList(),
	}
}

func (cache *LRUCache) Get(key int) int {
	// 如果key 不存在，则直接返回-1
	if _, ok := cache.KeyMap[key]; !ok {
		return -1
	} else {
		// 如果key 存在，则将对应的值返回（并将node 放到尾部）
		node := cache.KeyMap[key]
		cache.Cache.removeNode(node)
		cache.Cache.addNodeAtLast(node)

		// 这里不需要删除KeyMap
		return cache.KeyMap[key].Value
	}
}

func (cache *LRUCache) Put(key int, value int)  {
	// 如果key 已存在，把旧节点删除，新节点插入到尾部
	if _, ok := cache.KeyMap[key]; ok {
		node := cache.KeyMap[key]
		cache.Cache.removeNode(node)	
		// 这里需要删除KeyMap
		delete(cache.KeyMap, key)

		addNode := initNode(key, value)
		cache.Cache.addNodeAtLast(addNode)
		cache.KeyMap[key] = addNode
	} else {
		// 如果capacity已满，抛出最近未使用的节点，插入新数据到尾部
		addNode := initNode(key, value)
		if cache.Cache.Size >= cache.Capacity {
			lruNode := cache.Cache.removeFirstNode()

			// 这里需要删除KeyMap，同时需要添加新节点的KeyMap
			delete(cache.KeyMap, lruNode.Key)
		}
		// 否则，直接插入新数据到尾部
		cache.Cache.addNodeAtLast(addNode)
		cache.KeyMap[key] = addNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */