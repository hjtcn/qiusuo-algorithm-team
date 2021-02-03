// 897. 递增顺序查找树
// 给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。

 

// 示例 ：

// 输入：[5,3,6,2,4,null,8,1,null,null,null,7,9]

//        5
//       / \
//     3    6
//    / \    \
//   2   4    8
//  /        / \ 
// 1        7   9

// 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

//  1
//   \
//    2
//     \
//      3
//       \
//        4
//         \
//          5
//           \
//            6
//             \
//              7
//               \
//                8
//                 \
//                  9  
 

// 提示：

// 给定树中的结点数介于 1 和 100 之间。
// 每个结点都有一个从 0 到 1000 范围内的唯一整数值。
// 通过

/*
    思路：
    这个题要想ac还是比较简单的，我上去就直接遍历，用数组记录数据，然后又sort排序，之后才查找右节点构建树。
    后来看到题解感觉到我对于中序遍历了解的层次还是比较浅。
    1.首先，直接在遍历的过程就可以直接中序遍历，得到升序的数组，再构建树也行
    2.其次，对于这道题来说，直查找右节点，就可以遍递归遍构建

*/


var increasingBST = function(root) {
    let arr=[]
    let dfs=root=>{
        if(!root) return ;
        dfs(root.left)
        arr.push(root.val)//放在这个位置递归可直接进行中序遍历
        dfs(root.right)
    }
    dfs(root)
    let newRoot=new TreeNode(arr[0]),headRoot=newRoot
    for(let i=1;i<arr.length;i++){
        newRoot.right=new TreeNode(arr[i])
        newRoot=newRoot.right
    }
    return headRoot
};


//不借用数组的构建方式
var increasingBST = function(root) {
    let rootRight=new TreeNode(),headRoot=rootRight
    let dfs=root=>{
        if(!root) return ;
        dfs(root.left)
        rootRight.right=new TreeNode(root.val)
        rootRight=rootRight.right
        dfs(root.right)
    }
    dfs(root)
    return headRoot.right
};

//BFS中序遍历
var increasingBST = function(root) {
    if(!root){
        return []
    }
    let rootRight=new TreeNode(),headRoot=rootRight
    let stack=[]
   while(root||stack.length){
       while(root){
           stack.push(root)
           root=root.left
       }
       root=stack.pop()
       rootRight.right=new TreeNode(root.val)
       rootRight=rootRight.right
       root=root.right
   }
   return headRoot.right
};

/*
    时间复杂度：O(N)
    取决于树的节点个数
    空间复杂度:O(N)
    借用数组从新构建
*/