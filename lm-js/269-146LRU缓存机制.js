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
    思路：这次没有用数组模拟了，细抠使用记录还是可以写出来的，鉴于上次做这道题没过去多久，有记忆处理map数组。
    map核心原理：
    (1)每次get，put都先进行一次删除，再进行设置，来保持插入遍历。即之前的被删除，她的key,value位置都会新插入一遍放在最后面。
    (2)获取最早未使用的key
       map.keys().next().value
    这样一旦超过容量，就删它。

    加强记忆。。。主要还是map的使用太浅

    题解:链表。链表的删除新增还是用的不熟练，尽管它的时间复杂度很优。删除新增多练习用链表
    伪代码在这个题上就可以多练习，别害怕。
    看题目，想要什么方法。
    这道题需求还是比较明确的。
    难点在put:1.能记录最久未使用的key
             2.能判断是否超出容量。
             3.是新增，还是更新
    2，3其实还好，主要问题就在如何记录使用的新旧记录。
    因此我们可以想到使用数组，链表等去记录。
    考验我们能力也就是从数组到链表的优化了。
    养成使用链表遍历记录的过程

*/

/**
 * @param {number} capacity
 */
 var LRUCache = function(capacity) {
    this.capacity=capacity
    this.userList=new Map()
};

/** 
 * @param {number} key
 * @return {number}
 */
LRUCache.prototype.get = function(key) {
    if(this.userList.has(key)){
        let value=this.userList.get(key)
        this.userList.delete(key)
        this.userList.set(key,value)
        return value
    }
    return -1
};

/** 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
LRUCache.prototype.put = function(key, value) {
    if(this.userList.has(key)){
        this.userList.delete(key)
    }
    this.userList.set(key,value)
    if(this.userList.size>this.capacity){
        this.userList.delete(this.userList.keys().next().value)
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
 var NodeList=function(){
    this.prev=null
    this.next=null
    this.val=null
    this.key=null
}
var LRUCache = function(capacity) {
   this.size=0
   this.capacity=capacity
   this.map=new Map()
   //双向链表
   this.head=new NodeList()
   this.tail=new NodeList()
   this.head.next=this.tail
   this.tail.prev=this.head
};
LRUCache.prototype.moveToHead= function(node) {
   this.deleteFromList(node)
   this.addToHead(node)
};
LRUCache.prototype.addToHead = function(node) {
   node.prev=this.head
   node.next=this.head.next
   this.head.next.prev=node
   this.head.next=node
};
LRUCache.prototype.deleteFromList= function(node) {
   let prev=node.prev
   let next=node.next
   prev.next=next
   next.prev=prev
};
LRUCache.prototype.deleteNoUsed= function(node) {
   let target=this.tail.prev
   this.deleteFromList(target)
   return target
};
/** 
* @param {number} key
* @return {number}
*/
LRUCache.prototype.get = function(key) {
   if(this.map.has(key)){
       let node=this.map.get(key)
       this.moveToHead(node)
       return node.val
   }
   return -1
};

/** 
* @param {number} key 
* @param {number} value
* @return {void}
*/
LRUCache.prototype.put = function(key, value) {
   //先构建put，再get
   if(!this.map.has(key))
   {let node=new NodeList()
   node.key=key
   node.val=value
   this.map.set(key,node)
   this.addToHead(node)
   this.size++;
   if(this.size>this.capacity){
       let target=this.deleteNoUsed()
       this.size--;
       this.map.delete(target.key)
   }
   }
   else{
       let node=this.map.get(key)
       node.val=value
       this.moveToHead(node)
   }
};

/*
    时间复杂度：O(1)
    空间复杂度：O(N)
*/

/*
    思考：map不止set,get,has.
    还有keys().next().value。
    迭代器的使用，我知道的太少了。。。

    看完题解才想到双向链表，知识储备不行，而且链表掌握的水平让我更习惯用数组去处理

*/