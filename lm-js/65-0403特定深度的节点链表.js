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
        //构建队列，入队出队，层层操作
        let len=queue.length,floorArr=new ListNode(null),headArr=floorArr
        //floorArr初始化链表结构，headarr记录头节点
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            //构建链表的过程，赋值next,且指针后移
            floorArr.next=new ListNode(curNode.val)
            floorArr=floorArr.next
            //入队使左右子树变为下一轮的当前节点
            curNode.left&&queue.push(curNode.left)
            curNode.right&&queue.push(curNode.right)
        }
        res.push(headArr.next)
    }
    return res
};

/*
    bfs.这道题最主要的就是读清题。返回的是数组，放入的元素是链表。
    广度优先搜索，横向队列入队出队，来获取一层的节点，并构建为链表，然后push到结果数组
    时间复杂度：O(N)
    每个节点只会执行入队出队一次
    空间复杂度：O(N)
    包含树的节点数和链表空间。

*/