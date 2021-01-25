// 面试题 04.03. 特定深度节点链表
// 给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。

 

// 示例：

// 输入：[1,2,3,4,5,null,7,8]

//         1
//        /  \ 
//       2    3
//      / \    \ 
//     4   5    7
//    /
//   8

// 输出：[[1],[2,3],[4,5,7],[8]]

/** 
 读完这道题，想着害，这道题也太简单了吧，bfs，卡卡一顿写，然后执行代码。
    TypeError: [[1],[2,3],[4,5,7],[8]] is not valid value for the expected return type ListNode[]
然后静下心来重新读了读题，原来它是要链表啊，不是数组存储就行。
然后往答题区看了下，果然有一个ListNode的构造函数。

怎么记录链表呢？回忆一下子，首先声明变量headRoot记录头指针root，然后头指针root进行next遍历，然后统计时利用headRoot进行处理

题目思路：
首先声明变量headRoot记录头指针root，然后在BFS的for循环中开始移动头指针root进行next遍历，循环结束直接push记录下来的头指针headRoot变量

*/


/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {TreeNode} tree
 * @return {ListNode[]}
 */
var listOfDepth = function(tree) {
    if(!tree) return []
    let queue=[tree],res=[]
    while(queue.length){
        let len=queue.length,root=new ListNode(null),headRoot=root
        for(let i=0;i<len;i++){
              let curNode=queue.shift()
              root.next=new ListNode(curNode.val)
              root=root.next
              if(curNode.left){
                  queue.push(curNode.left)
              }
              if(curNode.right){
                  queue.push(curNode.right)
              }
        }
        res.push(headRoot.next)
    }
    return res
  };