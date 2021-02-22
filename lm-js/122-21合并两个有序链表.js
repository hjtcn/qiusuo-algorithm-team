// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

 

// 示例 1：


// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 示例 2：

// 输入：l1 = [], l2 = []
// 输出：[]
// 示例 3：

// 输入：l1 = [], l2 = [0]
// 输出：[0]
 

// 提示：

// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列


/*
思路：两个链表遍历，谁小，谁的指针指向next，同时赋值目标链表，跳出循环后，没遍历完的链接到目标链表
    这道题写的很顺利，第二个方法，不是自己想出来的，但是感觉理解起来也比之前顺畅多了。还是有进步的
    第二个方法，使用递归思路：哪条链表节点当前值小，就将那条链表作为目标链表，不断递归，直到结束
*/


/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {
    let root=new ListNode(),res=root
    while(l1&&l2){
        if(l1.val<l2.val){
            root.next=new ListNode(l1.val)
            l1=l1.next
        }
        else{
            root.next=new ListNode(l2.val)
            l2=l2.next
        }
        root=root.next    
    }
    if(l1){
        root.next=l1
    }
    if(l2){
        root.next=l2
    }
    return res.next
};

/** 
    复杂度分析：
    时间复杂度是:O(N)
    一次循环将该最短链表遍历完毕。
    空间复杂度：O(N+M）
    定义了res和newRes变量，存放新的和链表.
*/


var mergeTwoLists = function(l1, l2) {
    if(!l1) return l2
    if(!l2) return l1
    if(l1.val<l2.val){
        l1.next=mergeTwoLists(l1.next,l2)
        return l1
    }
    else{
        l2.next=mergeTwoLists(l1,l2.next)
        return l2
    }
 };

 /** 
    复杂度分析：
    时间复杂度是:O(N+M)
    遍历两条链表结束
    空间复杂度：O(N+M）
    第一次分析空间复杂度是O(1),原因是没有声明额外变量，但是现在感觉是O(N+M)了，由于递归的隐式栈
*/