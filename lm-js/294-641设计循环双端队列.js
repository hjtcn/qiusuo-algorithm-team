/*
641. 设计循环双端队列
设计实现双端队列。
你的实现需要支持以下操作：

MyCircularDeque(k)：构造函数,双端队列的大小为k。
insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
isEmpty()：检查双端队列是否为空。
isFull()：检查双端队列是否满了。
示例：

MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
circularDeque.insertLast(1);			        // 返回 true
circularDeque.insertLast(2);			        // 返回 true
circularDeque.insertFront(3);			        // 返回 true
circularDeque.insertFront(4);			        // 已经满了，返回 false
circularDeque.getRear();  				// 返回 2
circularDeque.isFull();				        // 返回 true
circularDeque.deleteLast();			        // 返回 true
circularDeque.insertFront(4);			        // 返回 true
circularDeque.getFront();				// 返回 4
 
 

提示：

所有值的范围为 [1, 1000]
操作次数的范围为 [1, 1000]
请不要使用内置的双端队列库。

*/


/*
    思路：记录当前长度和可存储空间
         同时更新头部节点和尾部节点。用双向链表存储

         注意细节：1.当前长度随着增删的加减
                 2.头尾节点增加时，将原本没有节点的情况单拎出来
                 头尾节点删除时，将原本只剩一个节点的情况单拎出来

*/


/**
 * @param {number} k
 */
 function NodeList(val){
    this.val=val
    this.next=null
}
var MyCircularDeque = function(k) {
   this.size=k
   this.head=null
   this.tail=null
   this.len=0
};

/** 
* @param {number} value
* @return {boolean}
*/
MyCircularDeque.prototype.insertFront = function(value) {
    let newHead=new NodeList(value)
   if(this.len==0){
       this.tail=this.head=newHead 
         this.len=this.len+1
       return true
   }
   else if(this.len<this.size){
       let curHead=this.head
       newHead.next=curHead
       curHead.prev=newHead
         this.len=this.len+1
       this.head=newHead
       return true
   }
   else{
       return false
   }
};

/** 
* @param {number} value
* @return {boolean}
*/
MyCircularDeque.prototype.insertLast = function(value) {
   let newHead=new NodeList(value)
   if(this.len==0){
       this.tail=this.head=newHead 
       this.len=this.len+1
       return true
   }
   else if(this.len<this.size){
       let tail=this.tail
       tail.next=newHead
       newHead.prev=tail
       this.tail=newHead
        this.len=this.len+1
       return true
   }
   else{
       return false
   }
};

/**
* @return {boolean}
*/
MyCircularDeque.prototype.deleteFront = function() {
   if(this.len==1){
       this.head=this.tail=null
       this.len=this.len-1
       return true
   }
   else if(this.len>0){
       let next=this.head.next
       this.head=next
        this.len=this.len-1
       return true
   }
   return false
};

/**
* @return {boolean}
*/
MyCircularDeque.prototype.deleteLast = function() {
   if(this.len==1){
       this.head=this.tail=null
       this.len=this.len-1
       return true
   }
   else if(this.len>0){
       let tail=this.tail.prev
       this.tail=tail
        this.len=this.len-1
       return true
   }
   return false
};

/**
* @return {number}
*/
MyCircularDeque.prototype.getFront = function() {
   return this.head?this.head.val:-1
};

/**
* @return {number}
*/
MyCircularDeque.prototype.getRear = function() {
   return this.tail?this.tail.val:-1
};  

/**
* @return {boolean}
*/
MyCircularDeque.prototype.isEmpty = function() {
   return this.len<=0
};

/**
* @return {boolean}
*/
MyCircularDeque.prototype.isFull = function() {
   return this.len==this.size
};

/**
* Your MyCircularDeque object will be instantiated and called as such:
* var obj = new MyCircularDeque(k)
* var param_1 = obj.insertFront(value)
* var param_2 = obj.insertLast(value)
* var param_3 = obj.deleteFront()
* var param_4 = obj.deleteLast()
* var param_5 = obj.getFront()
* var param_6 = obj.getRear()
* var param_7 = obj.isEmpty()
* var param_8 = obj.isFull()
*/

/*
    时间复杂度：O(1)
    空间复杂度：O(N)
*/

/*
    思考：链表的增删改还是比较有优势的，多练习，培养这种思考方向。
*/