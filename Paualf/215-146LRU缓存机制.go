运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 
进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例：
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：
1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put

1. Clarification:
运用已掌握的数据结构，设计和实现一个 LRU (最近最少使用）缓存机制。

get:
如果关键字 key 存在与缓存中，则返回关键字的值，否则返回-1

put(int key,int value):
如果关键字已经存在，则变更其数据值
如果关键字不存在，则插入该组 关键字 - 值
当缓存达到上限时，它应该在写入新数据之前删除最久未使用的数据值，为新数据值留出空间。

目前想到的会不会是 map 和链表的组合，然后组合到一起去

2. 看题解：
对这种设计题目还是有点怕怕的，同时也反应出来基础知识不是很扎实。

哈希表 + 双向链表

双向链表按照被使用的顺序存储了这些键值对，靠近头部的键值对是最近使用的，而靠近尾部的键值对是最久未使用的。

哈希表即为普通的哈希映射（HashMap),通过缓存数据的键映射其在双向链表中的位置。

对于 get 操作，首先判断key是否存在：
如果key不存在，则返回 -1
如果key存在，则key对应的节点是最近被使用的节点。通过哈希表定位到该节点在表中的位置，并将其移动到双向链表的头部，最后返回该节点的值

func (this *LRUCache) Get(key int) int {
    if _,ok := this.cache[key];!ok {
    	return -1
    }
    node := this.cache[key]
    this.moveToHead(node)
    return node.value
}

func (this *LRUCache)moveToHead(node *DLinkedNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache)removeNode(node *DLinkedNode) {
	node.prev.next = node.next
    node.next.prev = node.prev
}
移动到头节点的时候，先把节点删除，然后把节点加入到头部
每一种操作都是O(1)

这里就有一个疑问了：
这里为什么要用 双向链表呢？

单链表可以不？
单链表也是可以定位到要删除的节点，但是定位到以后，删除的时间复杂度是O(n), 你使用map 定位到了，没有办法立即删除，你要找到单链表中要删除的节点的前一个节点，然后才能删除这个节点

对于 put操作，首先判断 key是否存在：
 	如果key不存在，使用key和value 创建一个新的节点，在双向链表的头部添加该节点，并将key和该节点添加到哈希表中。
  然后判断双向链表的节点数是否超出容量，如果超出容量，则删除双向链表的尾部节点，并删除哈希表中的项

 如果key存在，则与get操作类似，先通过哈希表定位，再将对应的节点的值更新为value，并将该节点移到双向链表的头部。

func (this *LRUCache) Put(key int,value int) {
    if _,ok := this.cache[key];!ok {
        node := initDLinkedNode(key,value)
        this.cache[key] = node
        this.addToHead(node)
        this.size++
        if this.size > this.capacity {
            removed := this.removeTail()
            delete(this.cache, removed.key)
            this.size--
        }
    }else {
        node := this.cache[key]
        node.value = value
        this.moveToHead(node)
    }
}
初始化的数据结构：
type LRUCache struct {
	size int
    capacity int
    cache map[int]*DLinkedNode
    head,tail *DLinkedNode
}

type DLinkedNode struct {
	key,value int
    prev,next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        cache: map[int]*DLinkedNode{},
        head: initDlinkedNode(0,0),
        tail: initDlinkedNode(0,0),
        capacity:capacity,
    }
    l.head.next = l.tail
    l.tal.prev = l.head
    return l
}

基础的 lru 是这样，但是MySQL 和 Redis 中的lru 也是这样么？实际的是这样用的么，印象中不是的


3. 复杂度分析：
时间复杂度：O(1)
空间复杂度：O(1)


4. 总结：
4.1: 通过这道题目，认识到自己的基础知识还是不扎实，不断练习，看书，然后不断总结
