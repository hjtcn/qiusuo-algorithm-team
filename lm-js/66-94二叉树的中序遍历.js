// 94. 二叉树的中序遍历
// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。

 

// 示例 1：


// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
// 示例 2：

// 输入：root = []
// 输出：[]
// 示例 3：

// 输入：root = [1]
// 输出：[1]
// 示例 4：


// 输入：root = [1,2]
// 输出：[2,1]
// 示例 5：


// 输入：root = [1,null,2]
// 输出：[1,2]
 

// 提示：

// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
 

// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？



/*
 * @lc app=leetcode.cn id=94 lang=javascript
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var inorderTraversal = function(root) {
    if(!root){
        return []
    }
    let res=[]
    const inorder=(node)=>{
        if(!node) return;
        //先深层调用左子树
        inorder(node.left)
        // 存储当前节点val值
        res.push(node.val)
        //再深层调用右子树
        inorder(node.right)
    }
    inorder(root)
    return res
};
// @lc code=end

/**
 * 递归。开始先了解了一下什么是中序遍历.左-根-右，
 * 1.遇到存在左子树，继续深层调用左子树.直到节点为空为止，则将此时的节点push入res数组
 * 2.此时节点push到res数组后，如果它还存在右子树，继续深层调用右子树，直到此时节点为空为止

       时间复杂度：O(N)
    取决于树的节点个数，每个节点只会被入栈和出栈一次。
    空间复杂度：O(N)
    取决于树的深度
 
 */


var inorderTraversal = function(root) {
    if(!root){
        return []
    }
    let res=[]
    let stack=[]
   while(root||stack.length){
    // root节点有值，就被push,且指针指向左子树
       while(root){
           stack.push(root)
           root=root.left
       }
       //pop方式将root先左后根的出栈
       root=stack.pop()
       res.push(root.val)
    //  一层左根操作完，指针指向根的右子树。
       root=root.right
   }
   return res
};

/** 
  迭代，这是接触二叉树以来，第一次除递归和bfs,接触到迭代方法。
  看官方题解，了解到的，递归是用了隐形的栈空间，而迭代则是将栈显式的模拟出来
  中序遍历逻辑还是左-根-右，
  1.循环控制根节点root，如果当前root不为空，则将先根后左的形式入栈。
  2.跳出root循环后，栈开始pop,root变为先左后根，当前节点值push进入res数组，root指针开始指向右子树。

  写着题解把自己写迷糊了。
  res第一次push的是最左边的节点值left1，此时root指针指向它的右子树，则root为空，
  栈接着弹出，root变为left1的根节点，res被push此根节点的值，root指针指向右子树。
  会进入while(root)的循环会一直先根后左的入栈所有节点，直到，左子树木有了，栈开始接着pop,方向变为先左后根
  res被push的也是先左后根再右

   时间复杂度：O(N)
    取决于树的节点个数，每个节点只会被入栈和出栈一次。
    空间复杂度：O(N)
    取决于树的深度
*/