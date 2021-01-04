// 116. 填充每个节点的下一个右侧节点指针
// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

// struct Node {
//   int val;
//   Node *left;
//   Node *right;
//   Node *next;
// }
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

// 初始状态下，所有 next 指针都被设置为 NULL。

 

// 进阶：

// 你只能使用常量级额外空间。
// 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。


/*
 * @lc app=leetcode.cn id=116 lang=javascript
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */

/**
 * @param {Node} root
 * @return {Node}
 */

 //BFS
 var connect = function(root) {
    if(!root){
        return root
    }
    let queue=[root]
    while(queue.length){
        let len=queue.length
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            if(i==len-1){
                curNode.next=null
            }
            else{
                curNode.next=queue[0]
            }
            if(curNode.left){
                queue.push(curNode.left)
            }
            if(curNode.right){
                queue.push(curNode.right)
            }
        }
    }
    return root
};

/** 
 借用队列实现BFS,这个是自己写的，很顺利，思路比较清晰，啦啦啦
 横向进行next连接，仅控制最右部节点的next为null即可，其余当前节点出队列后，则当前节点的next节点为队列头
     时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    借用数组模拟队列，取决于树的节点数
*/

//不借用队列

var connect=function(root){
    if(!root){
        return root
    }
    let leftMost=root
    while(leftMost.left){
        let head=leftMost
        while(head){
            //第一种情况
            head.left.next=head.right
            //第二种情况
            if(head.next){
                head.right.next=head.next.left
            }
            head=head.next
        }
        leftMost=leftMost.left
    }
    return root
}
/** 
 直接借助next指针。则next指针的赋值分为两种情况。
    1.左节点：head.left.next=head.right
    2.右节点需要链接父节点的next的左节点：head.right.next=head.next.left(head的next节点有值)
 key：树的层级扫描则利用最左节点是否存在左节点
    时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(1)
    常量元素记录节点
*/

//递归

var connect=function(root){
    if(!root){
        return root
    }
    let helper=(root)=>{
        if(!root.left&&!root.right){
            return;
        }
        root.left.next=root.right
        if(root.next){
            root.right.next=root.next.left
        }
        helper(root.left)
        helper(root.right)
    }
    helper(root)
    return root
}

/** 
 递归。核心和上面借助next指针基本一直。遍历方式为前序遍历：根->左->右
 递归结束的调节为当前节点没有左右子节点
    时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(1)
    题目中提到：本题中递归程序占用的栈空间不算做额外的空间复杂度。
*/



