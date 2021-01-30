// 24. 两两交换链表中的节点
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

 

// 示例 1：


// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
// 示例 2：

// 输入：head = []
// 输出：[]
// 示例 3：

// 输入：head = [1]
// 输出：[1]
 

// 提示：

// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
 

// 进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）

/** 
 思路：两两分组进行交换。
    递归：要提前记录next节点，然后再交换，最后返回next节点，此时next节点已经是头节点了。
         而串起来节点形成链表的过程，可利用递归root.next=dfs(next.next)


    迭代：迭代需要只存储prev节点来串联链表，比如：[1，2，3，4]
    第一步：prev为null,先1和2交换，然后prev.next节点指向2了.null->2->1串起来了
    第二步：prev指向1，此时轮到3和4交换，交换后将prev.next指向4，此时1->4->3串起来了。
    依次循环。。。。（循环的条件即head=head.next）



    其实交换都知道，但是记录节点，串接链表，返回链表就不好弄了,来来回回把链表又忘了，害。
*/

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function(head) {
    let dfs=(root)=>{
        if(!root||!root.next) return root;
        let next=root.next
        root.next=dfs(next.next)
        next.next=root
        return next
    }
   return  dfs(head)

};
/** 
 时间复杂度：O(N)
 递归和链表节点个数有关
 空间复杂度：O(N)
 递归的隐用栈。和链表节点个数有关
*/

var swapPairs = function(head) {
    let res=new ListNode()
    res.next=head
    let prev=res
    while(head&&head.next){
        let next=head.next
        head.next=next.next
        next.next=head
        prev.next=next //这个是串联的过程

        prev=head   //更新prev节点
        head=head.next  //保持链表遍历
    }
    return res.next
};

/** 
 时间复杂度：O(N)
 递归和链表节点个数有关
 空间复杂度：O(1)
 只声明了节点变量
 */