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
    思路：第二次做这道题了。还是踩了无数的坑。
         题目有5个需求
         1.从头部插入
         2.从尾部插入(如果原链表为空)
         3.从index插入(牵扯到下标，一定要明确是从0开始，还是从1开始,以及各种边界值，边打印，边完善逻辑，不太可行，不能确认每一个代码的边界是否是准确的)
         4.从index删除
         5.活动当前下标对应val

*/

var MyLinkedList = function() {
    this.head=null
    this.node=function(val){
        this.val=val
        this.next=null
    }
    this.size=0
};

/** 
 * @param {number} index
 * @return {number}
 */
MyLinkedList.prototype.get = function(index) {
    let head=this.head
     if(index<0||index>=this.size){
        return -1;
    }
    for(let i=0;i<index;i++){
        head=head.next
    }
    return head.val
};

/** 
 * @param {number} val
 * @return {void}
 */
MyLinkedList.prototype.addAtHead = function(val) {
    let curNode=new this.node(val)
    let head=this.head
    this.head=curNode
    curNode.next=head
    this.size++;
};

/** 
 * @param {number} val
 * @return {void}
 */
MyLinkedList.prototype.addAtTail = function(val) {
    let curNode=new this.node(val)
    if(!this.head){
        this.head=curNode
        this.size++;
        return ;
    }
    let head=this.head
    for(let i=0;i<this.size-1;i++){
        head=head.next
    }
    head.next=curNode
     this.size++;
};

/** 
 * @param {number} index 
 * @param {number} val
 * @return {void}
 */
MyLinkedList.prototype.addAtIndex = function(index, val) {
    if(index>this.size) return;
let curNode=new this.node(val)
    if(!this.head){
        this.head=curNode
        this.size++;
        return ;
    }
    if(index<=0){
        this.addAtHead(val)
        return;
    }
    let head=this.head
    for(let i=0;i<index-1&&i<=this.size;i++){
        head=head.next
    }
    let next=head.next
    head.next=curNode
    curNode.next=next
    this.size++;
};

/** 
 * @param {number} index
 * @return {void}
 */
MyLinkedList.prototype.deleteAtIndex = function(index) {
    if(index<0||index>=this.size){
        return;
    }
    let head=this.head
    if(index==0){
        this.head=head.next
        this.size--;
        return;
    }
     for(let i=0;i<index-1;i++){
        head=head.next
      }
      let next=head.next
      if(next){
           head.next=next.next
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
    时间复杂度:O(N)
    空间复杂度：O(N)
*/

/*
    思考：这个才是真的考细节，核心不在于把题目做出来，而是能明确每一个边界和兜底是否做到位。上班摸鱼真的不适合做这道题。
*/