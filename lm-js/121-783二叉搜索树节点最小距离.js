// 783. 二叉搜索树节点最小距离
// 给定一个二叉搜索树的根节点 root，返回树中任意两节点的差的最小值。

 

// 示例：

// 输入: root = [4,2,6,1,3,null,null]
// 输出: 1
// 解释:
// 注意，root是树节点对象(TreeNode object)，而不是数组。

// 给定的树 [4,2,6,1,3,null,null] 可表示为下图:

//           4
//         /   \
//       2      6
//      / \    
//     1   3  

// 最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
 

// 注意：

// 二叉树的大小范围在 2 到 100。
// 二叉树总是有效的，每个节点的值都是整数，且不重复。
// 本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/ 相同


/** 
 思路：刚开始题意有点迷，不知道怎么更新最小值，想在递归中直接更新，尝试了一下，发现有点乱。
 后来就决定从最老的方法开始做，先中序遍历，存至数组，再遍历数组更新最小值
 后来尝试第一个方法，就发现结构就很清晰了，然后在中序遍历过程中，更新差最小值。
 发现我第一次失败的原因是我把上次的值lastVal也作为dfs的参数了
*/


var minDiffInBST = function(root) {
    let min=Infinity,arr=[]
    let dfs=(root)=>{
        if(!root) return ;
        dfs(root.left)
        arr.push(root.val)
        dfs(root.right)
    }
    dfs(root)
    for(let i=1;i<arr.length;i++){
        if(arr[i]-arr[i-1]<min){
            min=arr[i]-arr[i-1]
        }
    }
    return min
};
/** 
 时间复杂度：O(N)
 递归遍历子节点及一层循环遍历
 空间复杂度：O(N)
 将节点存至数组
*/


var minDiffInBST = function(root) {
    let min=Infinity,lastVal=Infinity
    let dfs=(root)=>{
        if(!root) return ;
        dfs(root.left)
        //第一次将lastVal也当作dfs的参数了，然后把自己写懵了，要一步一步来
        if(Math.abs(root.val-lastVal)<min){
            min=Math.abs(root.val-lastVal)
        }
        lastVal=root.val
        dfs(root.right)
    }
    dfs(root)
    return min
};
// 时间复杂度：O(N)
// 递归遍历子节点
// 空间复杂度：O(N)
// 递归栈的节点占据空间