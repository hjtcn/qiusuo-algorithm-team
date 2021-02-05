// 98. 验证二叉搜索树
// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。

// 假设一个二叉搜索树具有如下特征：

// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
// 示例 1:

// 输入:
//     2
//    / \
//   1   3
// 输出: true
// 示例 2:

// 输入:
//     5
//    / \
//   1   4
//      / \
//     3   6
// 输出: false
// 解释: 输入为: [5,1,4,null,null,3,6]。
//      根节点的值为 5 ，但是其右子节点值为 4 。


/* 
   果然练得多了，看见二叉搜索树一点都不慌了
   思路：中序遍历，节点会递增排列，递归过程中查看是否有前一个数比当前数大的，有就打一个标记，有标记则返回false,没有就返回true
   迭代的中序遍历还是不是那么清晰，它和bfs的框架不一样，多做几遍，形成框架
*/


var isValidBST=function(root){
    let min=-Infinity,flag=0
  let dfs=(root)=>{
      if(!root) return ;
      dfs(root.left)
      if(min>=root.val){
          flag=1
      }
      min=root.val
      dfs(root.right)
  }
  dfs(root)
  if(flag){
      return false
  }
  return true
}


var isValidBST=function(root){
    let min=-Infinity
    let queue=[]
    while(queue.length||root){
       while(root){
           queue.push(root)
           root=root.left
       }
       root=queue.pop()
       if(min>=root.val){
           return false
       }
       min=root.val
       root=root.right
    }
    return true
}
/*
 时间复杂度：O(N)
 递归和迭代都是O(N),需要遍历节点个数
 空间复杂度：O(N)
 递归的隐式空间，迭代的队列模拟。

*/