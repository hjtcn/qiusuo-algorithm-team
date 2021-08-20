/*
706. 设计哈希映射
不使用任何内建的哈希表库设计一个哈希映射（HashMap）。

实现 MyHashMap 类：

MyHashMap() 用空映射初始化对象
void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key 已经存在于映射中，则更新其对应的值 value 。
int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。
 

示例：

输入：
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
输出：
[null, null, null, 1, -1, null, 1, null, -1]

解释：
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // myHashMap 现在为 [[1,1]]
myHashMap.put(2, 2); // myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(1);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
myHashMap.get(3);    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
myHashMap.put(2, 1); // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
myHashMap.get(2);    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
myHashMap.remove(2); // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
myHashMap.get(2);    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]
 

提示：

0 <= key, value <= 106
最多调用 104 次 put、get 和 remove 方法
 

进阶：你能否不使用内置的 HashMap 库解决此问题？

*/

/*
    思路：一开始就利用数组遍历，但是一直溢出
    题解：拉链法，利用某种hash函数将key映射到下标空间，下标相同的用链表存储
*/

function linkNode(key,val){
    this.key=key;
    this.val=val;
    this.next=null;
}

var MyHashMap = function() {
   const N=2048; 
   const hash=new Array(N);
   for(let i=0;i<hash.length;i++){
       hash[i]=new linkNode();
   }
   this.hash=hash;
};

/**
 * value will always be non-negative. 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
MyHashMap.prototype.put = function(key, value) {
    const N=2048;
    const index=key%N;
    let head=this.hash[index];
    let cur=head;
    while(cur && cur.next){
        if(cur.next.key===key){
            cur.next.val=value;
            return ;
        }
        cur=cur.next;
    }
    //没找到就新增
    const node=new linkNode(key,value);
    cur.next=node;
};

/**
 * Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key 
 * @param {number} key
 * @return {number}
 */
MyHashMap.prototype.get = function(key) {
   const N=2048;
   const index=key%N;
   let head=this.hash[index];
   let cur=head;
   while(cur && cur.next){
       if(cur.next.key===key){
           return cur.next.val;
       }
       cur=cur.next;
   }
   return -1;


};

/**
 * Removes the mapping of the specified value key if this map contains a mapping for the key 
 * @param {number} key
 * @return {void}
 */
MyHashMap.prototype.remove = function(key) {
   const N=2048;
   const index =key%N;
   let head=this.hash[index];
   let cur=head;
   while(cur && cur.next){
       if(cur.next.key===key){
           cur.next=cur.next.next;
       }
       cur=cur.next;
   }
};

/**
 * Your MyHashMap object will be instantiated and called as such:
 * var obj = new MyHashMap()
 * obj.put(key,value)
 * var param_2 = obj.get(key)
 * obj.remove(key)
 */

/*
    思考:链地址法真的有点酷，对hash更有感觉了，且链表这样简单存储关联有点厉害
*/