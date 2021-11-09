/*
1367. 二叉树中的列表
给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。

如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。

一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。

 

示例 1：



输入：head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：true
解释：树中蓝色的节点构成了与链表对应的子路径。
示例 2：



输入：head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：true
示例 3：

输入：head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：false
解释：二叉树中不存在一一对应链表的路径。
 

提示：

二叉树和链表中的每个节点的值都满足 1 <= node.val <= 100 。
链表包含的节点数目在 1 到 100 之间。
二叉树包含的节点数目在 1 到 2500 之间。   

*/

/*
    思路:想到了递归，但是自己犯了一个错，就是一开始没有想好递归函数返回值就开始写，写着写着，发现不能包裹所有的情况。

        其实这道题还是不简单的，之前写的递归，都是递归一个函数，而这个不仅需要dfs递归，同时外层函数也需要递归左右子节点。

        这道题的大致思路就是。
        1.该题返回dfs(head,root)||isSubPath(head,root.left)||isSubPath(head,root.right)
        2.dfs内部函数：
            (1)边界为如果head链表指向末尾，则出现了完整的链表，返回true
            (2)如果还没返回true,此时root跳过叶子结点，指向空了，就返回false
            (3)如果root.val!=head.val 返回false
            (4)如果相等，就开始往下一个节点比较。return dfs(root.left,head.next)||dfs(root.right,head.next)

*/
var isSubPath = function(head, root) {
    if(!root) return false
    let cur=head
    let dfs=(root,cur)=>{
        if(!cur) return true
        if(!root) return false;
        if(root.val!==cur.val) return false
        return dfs(root.left,cur.next)||dfs(root.right,cur.next)
    }
    //这里不太好想到。
    return dfs(root,cur)||isSubPath(cur,root.left)||isSubPath(cur,root.right)
};

/*
    时间复杂度：O(n*min(2^len+1,n))
    空间复杂度：O(H)
*/

/*
    思考：又get新技能了，同时忘了构建递归的考虑过程，细节细节。
*/