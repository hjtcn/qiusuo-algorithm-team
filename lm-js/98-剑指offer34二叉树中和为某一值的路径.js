// 剑指 Offer 34. 二叉树中和为某一值的路径
// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

 

// 示例:
// 给定如下二叉树，以及目标和 sum = 22，

//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \    / \
//         7    2  5   1
// 返回:

// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]


var pathSum = function(root, sum) {
   if(!root){
       return [];
   }
   let res=[]
  let dfs=(root,sum,path)=>{
      if(!root){
          return ;
      }
       path.push(root.val)
       if(sum-root.val==0&&!root.left&&!root.right){
          //这里踩了一个坑，必须到叶节点，差为0，上次也这里错了
          res.push([...path])
      }
      if(root.left){
          dfs(root.left,sum-root.val,path)
      }
      if(root.right){
          dfs(root.right,sum-root.val,path)
      }
      path.pop()
  }
  dfs(root,sum,[])
  return res
};

/**
 * 递归，记录路径和当前差。
 * 时间复杂度：O(N)
    取决于树的节点数
    
    空间复杂度：O(N)
    取决于树节点数
 */


var pathSum = function(root, sum) {
   if(!root){
       return [];
   }
   let res=[],queue=[[root,[],sum]]
   while(queue.length){
         let [curNode,path,count]=queue.shift()
         path.push(curNode.val)
        if(count-curNode.val==0&&!curNode.left&&!curNode.right){
            res.push([...path])
         }
         if(curNode.left){
            queue.push([curNode.left,[...path],count-curNode.val])
         }
         if(curNode.right){
            queue.push([curNode.right,[...path],count-curNode.val])
         }
   }
   return res
}

/**
 * BFS.数组模拟队列，这个还是比较顺利的，但是记录path，也要利用结构赋值进行深拷贝。
 *  时间复杂度：O(N)
    取决于树的节点数
    
    空间复杂度：O(N*M)
    取决于树节点数,队列记录节点和路径以及差值
 */