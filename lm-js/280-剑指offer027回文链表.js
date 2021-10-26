/*
    剑指 Offer II 027. 回文链表
    给定一个链表的 头节点 head ，请判断其是否为回文链表。

    如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。

    

    示例 1：



    输入: head = [1,2,3,3,2,1]
    输出: true
    示例 2：



    输入: head = [1,2]
    输出: false
    

    提示：

    链表 L 的长度范围为 [1, 105]
    0 <= node.val <= 9
    
*/

/*
    思路：1.遍历，存储，反转比较。
    题解：链表分为前后两段，后段反转。与前段比较是否一致。
    过程中可以通过快慢指针分段。
*/

/*
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
var isPalindrome = function(head) {
    let str1='',str2=''
    let dfs=(head)=>{
        if(!head) return ;
        str1+=head.val+''
        str2=head.val+str2+''
        dfs(head.next)
    }
    dfs(head)
    if(str1==str2){
        return true
    }
    else{
        return false
    }
};
/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

//快慢指针
let getMidNode=(root)=>{
    let slow=root,fast=root
    while(fast.next&&fast.next.next){
        slow=slow.next
        fast=fast.next.next
    }
    return slow
}
let  reverseNode=(root)=>{
    let prev=null
    let cur=root
    while(cur){
      let next=cur.next
      cur.next=prev
      prev=cur
      cur=next
    }
    return prev
}
var isPalindrome = function(head) {
    if(!head)return true
    let midNode=getMidNode(head)
    let rightReverseNode=reverseNode(midNode.next)
    //判断是否是回文
    let p1=head,p2=rightReverseNode
    let result=true
    while(p1&&p2){
        if(p1.val!=p2.val) return false
        p1=p1.next
        p2=p2.next
    }
    //这里还可以还原链表..当然可以在return false的时候，仅将result赋值，不进行提前返回，这样还原就可以生效了。
    midNode.next=reverseNode(rightReverseNode)
    return result
}


/*
    时间复杂度：O(N)
    空间复杂度：O(1)
*/

/*
    思考：别看用数组存储，反转看起来挺简单的，但是如果空间复杂度O(1)去构建的话，就满复杂了。
    反转链表和快慢指针都很厉害，这个思路都不敢想。
*/