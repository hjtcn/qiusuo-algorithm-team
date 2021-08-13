/*
707. 设计链表
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
 

示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3
 

提示：

所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 LinkedList 库。

*/

/*
    思路：差点崩溃系列。
        1.定义this.node.这个还属于比较常见的。
        this.node=function(val){
            this.val=val
            this.next=null
        }
        2.顾头顾尾系列
        1）定义了this.size.想着for循环遍历节点，但是总是添加方法中忘记size+1,删除方法中忘记size-1.
        2)addAtIndex和deleteAtIndex要考虑头部的添加删除。同时提前reture也要控制this.size
        3)列举边界case,去测试自己的代码
        我自己这次提交了10次，就是刚开始已经混乱了，感觉逻辑是通的，但是细节处理太难了。
        4）还是this.size,在for循环遍历中，基本都是在this.size-1就已经开始处理。但是在addAtIndex中，如果index大于链表长度，不会插入节点,需要提前判断的是this.size。
        也就是需要了解，节点停留处理的位置和提前处理超出边界的位置。

*/

/**
 * Initialize your data structure here.
 */
 var MyLinkedList = function(val) {
    this.Node=function(val){
        this.val=val
        this.next=null
    }
    this.size=0
};

/**
 * Get the value of the index-th node in the linked list. If the index is invalid, return -1. 
 * @param {number} index
 * @return {number}
 */
MyLinkedList.prototype.get = function(index) {
    if(index<0||index>=this.size){
        return -1
    }
    let cur=this.head
    for(let i=0;i<index;i++){
        cur=cur.next
    }
    return cur.val
};

/**
 * Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. 
 * @param {number} val
 * @return {void}
 */
MyLinkedList.prototype.addAtHead = function(val) {
    let cur=this.head
    this.head=new this.Node(val)
    let head=this.head
    head.next=cur
    this.size++;
};

/**
 * Append a node of value val to the last element of the linked list. 
 * @param {number} val
 * @return {void}
 */
MyLinkedList.prototype.addAtTail = function(val) {
    let head=new this.Node(val)
    if(!this.head){
        this.head=head
        this.size++;
        return;
    }
    let cur=this.head
    for(let i=0;i<this.size-1;i++){
        cur=cur.next
    }
    this.size++;
    cur.next=head
};

/**
 * Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. 
 * @param {number} index 
 * @param {number} val
 * @return {void}
 */
MyLinkedList.prototype.addAtIndex = function(index, val) {
    if(index>this.size){
        return;
    }
    if(index<=0){
        this.addAtHead(val)
        return;
    }
    let cur=this.head
    for(let i=0;i<index-1&&i<=this.size;i++){
        cur=cur.next
    }
    this.size++;
    let next=cur.next
    cur.next=new this.Node(val)
    if(next)
    {cur.next.next=next
    }
};

/**
 * Delete the index-th node in the linked list, if the index is valid. 
 * @param {number} index
 * @return {void}
 */
MyLinkedList.prototype.deleteAtIndex = function(index) {
    if(index<0||index>=this.size){
        return ;
    }
    if(index==0){
        this.head=this.head.next
        this.size--;
        return;
    }
    let cur=this.head
    for(let i=0;i<index-1;i++){
        cur=cur.next
    }
    let target=cur.next
    if(target){
        cur.next=target.next
    }
    this.size--
};

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * var obj = new MyLinkedList()
 * var param_1 = obj.get(index)
 * obj.addAtHead(val)
 * obj.addAtTail(val)
 * obj.addAtIndex(index,val)
 * obj.deleteAtIndex(index)
 */

/*
    思考：细节在自己不太熟悉的位置，真的会顾及不到，虽然逻辑很简单，但是想法不会那么完善。
*/