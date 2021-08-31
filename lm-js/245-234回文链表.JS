/*
234. 回文链表
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

 

示例 1：


输入：head = [1,2,2,1]
输出：true
示例 2：


输入：head = [1,2]
输出：false
 

提示：

链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9
 

进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
/*
    思路：这条题不难，但是要尽量优化着写还是有很多内容的。
         我自己写了三种方案，但是耗时差距还是非常大的，疑惑。
         1.js的api转化为数组翻转再转化为字符串对比。耗时256ms
         2.链表值存入数组，然后不停头尾出栈，进行对比。耗时7740ms
         3.改变链表为双向链表，加一个prev属性。然后分别从头的next,尾的prev节点进行对比。耗时172ms

         再看之前的提交记录，最小为80ms.（最后又提交了原答案一遍，好吧，和网速有关系，哈哈哈哈）
    题解：快慢指针+翻转链表，不熟悉的还是多练练 
         翻转链表过程还是有点不熟悉。
         今天看翻转又搜索总结了一番：
         两种方案：
         1.就地翻转法
         利用四个节点：dummy->prev->cur->next
         prev.next=cur.next
         cur.next=dummy.next
         dummy.next=cur
         cur=prev.next
         2.作为新节点插入法(这个不太熟悉)
         新建一个头结点，遍历原链表，把每个节点用(头结点)插入到新建链表中。最后，新建的链表就是反转后的链表。
         cur是要插入新链表的节点
         let next=cur.next  //next是临时存放cur.next
         cur.next=dummy.next //接下来这两句相当于把cur插入进要新链表
         dummy.next=cur
         cur=next //插入下一个节点

         

         
       

*/

// 方案一
var isPalindrome = function(head) {
    let arr=[]
    while(head){
        arr.push(head.val)
        head=head.next
    }
    let str=arr.join(',')
    let reverseStr=arr.reverse().toString()
    if(str==reverseStr){
        return true
    }
    return false
};
// @lc code=end

// 方案二
var isPalindrome = function(head) {
    let arr=[]
    while(head){
        arr.push(head.val)
        head=head.next
    }
    while(arr.length>1){
        let top=arr.shift()
        let tail=arr.pop()
        if(top!=tail){
            return false
        }
    }
    return true
};
/*
时间复杂度：O(N)
空间复杂度：O(N)
*/

//方案三
var isPalindrome = function(head) {
    let curNode=head
    let tailNode=null
    while(curNode){
        let prev=curNode
        curNode=curNode.next
        if(curNode){
            curNode.prev=prev
        }
        else{
            tailNode=prev
        }
    }
    while(head&&tailNode){
        if(head.val!=tailNode.val){
            return false
        }
        head=head.next
        tailNode=tailNode.prev
    }
    return true
}

var isPalindrome = function(head) {
    let fast=slow=head,prev=new ListNode()
    while(fast&&fast.next){
        prev=slow
        slow=slow.next
        fast=fast.next.next
    }
    prev.next=null
    let head2=null
    while(slow){
      //一次次借用头结点插入
       let next=slow.next
       slow.next=head2
       head2=slow
       slow=next 
    }
    while(head){
        if(head.val!=head2.val){
            return false
        }
        head=head.next
        head2=head2.next
    }
    return true  
}
/*
时间复杂度：O(N)
空间复杂度：O(1)
*/

/*
    思考：翻转链表一不小心就把自己绕进去了。
*/