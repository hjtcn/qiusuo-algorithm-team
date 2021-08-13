/*
146. LRU 缓存机制
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
*/

/*
    题解：核心点就在于如何找到最久未使用的数据值
    1.双向链表：head,tail虚拟头尾节点。减少prev,next节点为空的处理。
        head.next=tail
        tail.prev=head
        每一次使用的节点都需要将其放到虚拟头节点的后面。
        如果这个节点的key之前出现过，就需要先删除之前链表中的节点，再加到头部
        未出现过，直接加到头部就行
        get,put方法都是需要更新到头部的。
        不仅要更新链表。
        也要更新map,比如，删除最久未使用的节点时，同时需要将节点的key值在map中删去。
    2.map的特殊处理。map删除key，再设置key->value,map.keys()返回迭代器的对象，保持插入遍历，.next()就能拿到最早插入(即最久未使用的数据).

*/

/**
 * @param {number} capacity
 */
function ListNode() {
    this.key = null
    this.val = null
    this.prev = null
    this.next = null
}
var LRUCache = function (capacity) {
    this.capacity = capacity
    this.size = 0
    this.map = new Map()
    this.head = new ListNode()
    this.tail = new ListNode()
    this.head.next = this.tail
    this.tail.prev = this.head
};
LRUCache.prototype.moveToHead = function (node) {
    this.deleteFromList(node)
    this.addToHead(node)
}
LRUCache.prototype.deleteFromList = function (node) {
    let prev = node.prev
    let next = node.next
    prev.next = next
    next.prev = prev
}
LRUCache.prototype.addToHead = function (node) {

    node.prev = this.head;
    node.next = this.head.next;
    this.head.next.prev = node;   //这个地方有点巧妙
    this.head.next = node;
}
LRUCache.prototype.deleteNoUsed = function () {
    let target = this.tail.prev
    this.deleteFromList(target)
    return target
}
/** 
* @param {number} key
* @return {number}
*/
LRUCache.prototype.get = function (key) {
    let node = this.map.get(key)
    if (this.map.has(key)) {
        this.moveToHead(node)
        return node.val
    }
    else {
        return -1
    }
};

/** 
* @param {number} key 
* @param {number} value
* @return {void}
*/
LRUCache.prototype.put = function (key, value) {
    let node = this.map.get(key)
    if (!node) {
        let curNode = new ListNode()
        curNode.key = key
        curNode.val = value
        this.map.set(key, curNode)
        this.addToHead(curNode)
        this.size++
        if (this.size > this.capacity) {
            //删除最久未使用.替换
            let target = this.deleteNoUsed()
            this.size--;
            this.map.delete(target.key)
        }

    }
    else {
        node.val = value
        this.moveToHead(node)
    }
};

/**
* Your LRUCache object will be instantiated and called as such:
* var obj = new LRUCache(capacity)
* var param_1 = obj.get(key)
* obj.put(key,value)
*/

/**
 * @param {number} capacity
 */
 var LRUCache = function(capacity) {
    this.capacity=capacity
    this.size=0
    this.map=new Map()
};

/** 
 * @param {number} key
 * @return {number}
 */
LRUCache.prototype.get = function(key) {
    if(this.map.has(key)){
        let value=this.map.get(key)
        this.map.delete(key)
        this.map.set(key,value)
        return value
    }
    else{
        return -1
    }
};

/** 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
LRUCache.prototype.put = function(key, value) {
    if(this.map.has(key)){
        this.map.delete(key)
    }
    this.map.set(key,value)
    if(this.map.size>this.capacity){
        this.map.delete(this.map.keys().next().value)
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */

/*
    题解结构上基本都满清晰的，主要还是要处理双向链表的时的拼接要全乎点。删除时要链表和map数据都删干净。添加到头部如果出现过，就先删再加，没出现过直接加。

    写不出来的时候多分分块，不要想着一步实现所有功能，有可能有思路缺失的地方就很容易遗漏掉，反而陷到里面出不来。
*/