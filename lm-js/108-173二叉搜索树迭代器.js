// 173. 二叉搜索树迭代器
// 实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。

// 调用 next() 将返回二叉搜索树中的下一个最小的数。

 

// 示例：



// BSTIterator iterator = new BSTIterator(root);
// iterator.next();    // 返回 3
// iterator.next();    // 返回 7
// iterator.hasNext(); // 返回 true
// iterator.next();    // 返回 9
// iterator.hasNext(); // 返回 true
// iterator.next();    // 返回 15
// iterator.hasNext(); // 返回 true
// iterator.next();    // 返回 20
// iterator.hasNext(); // 返回 false
 

// 提示：

// next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
// 你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。

// // @lc code=start
// /**
//  * Definition for a binary tree node.
//  * function TreeNode(val, left, right) {
//  *     this.val = (val===undefined ? 0 : val)
//  *     this.left = (left===undefined ? null : left)
//  *     this.right = (right===undefined ? null : right)
//  * }
//  */
// /**
//  * @param {TreeNode} root
//  */
let stack=[],res=null
var BSTIterator = function(root) {
    while(root){
        stack.push(root)
        root=root.left
    }
};

/**
 * @return {number}
 */
BSTIterator.prototype.next = function() {
    let p=stack.pop()
    let res=p.val
    p=p.right
    while(p){
        stack.push(p)
        p=p.left
    }
    return res
};

// /**
//  * @return {boolean}
//  */
BSTIterator.prototype.hasNext = function() {
    return stack.length
};

/** 
 利用栈进行迭代。鉴于二叉搜索树左<根<右，目标是要找下一个最小的数。
 首先把左节点存入栈中，然后寻找next节点时，栈pop()，当前节点为最小节点target。
 此时找最小节点的右节点p，接着循环找p的最左的左节点，且入栈，还是在找最小值

 感觉这个利用从小到大入栈有点巧妙，尤其部分放到构造函数，部分放到next函数中，自己想想不出来。
 看到这个题的时候还是感觉不好入手，对这种类型题真的有种莫名恐惧
 时间复杂度:O(N)
 迭代入栈，取决于节点数
 空间复杂度:O(N)
 stack模拟栈，取决于节点数
*/

// /**
//  * Your BSTIterator object will be instantiated and called as such:
//  * var obj = new BSTIterator(root)
//  * var param_1 = obj.next()
//  * var param_2 = obj.hasNext()
//  */
// // @lc code=end

let stack=[]
var BSTIterator = function(root) {
  let dfs=(root)=>{
     if(!root) return ;
     dfs(root.left)
     stack.push(root.val)
     dfs(root.right)
  }
  dfs(root)
};

/**
 * @return {number}
 */
BSTIterator.prototype.next = function() {
    return stack.shift()
};

/**
 * @return {boolean}
 */
BSTIterator.prototype.hasNext = function() {
    return stack.length
};

/** 
 递归的思路还是比较简单的.
 首先进行左根右的中序遍历，数组中从小到达存储节点val,next则求栈底元素，并shift(),hasNext看栈的长度
    时间复杂度:O(N)
    取决于节点数
    空间复杂度:O(N)
    stack存储val，取决于节点数
*/

