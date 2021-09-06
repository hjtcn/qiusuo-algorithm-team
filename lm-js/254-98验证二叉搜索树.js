/*
    98. 验证二叉搜索树
    给定一个二叉树，判断其是否是一个有效的二叉搜索树。

    假设一个二叉搜索树具有如下特征：

    节点的左子树只包含小于当前节点的数。
    节点的右子树只包含大于当前节点的数。
    所有左子树和右子树自身必须也是二叉搜索树。
    示例 1:

    输入:
        2
    / \
    1   3
    输出: true
    示例 2:

    输入:
        5
    / \
    1   4
        / \
        3   6
    输出: false
    解释: 输入为: [5,1,4,null,null,3,6]。
        根节点的值为 5 ，但是其右子节点值为 4 。
*/

/*
    思路：(1)中序遍历单调递增。先增序的将节点放入arr中，然后遍历确认是否单调递增
        优化：不借用数组，中序可直接更新最大值max就行，如果出现小于等于最大值max的，就返回false
    题解：
    1.递归：
    先序遍历，根左右，记录最小值min和最大值max。根如果出现小于min,大于max的直接返回false.
        递归返回值及参数的意义：
        dfs(root,min,max)
        边界：if(!root) return true
        判断 if(root.val<=min||root.val>=max) return false
        返回:return dfs(root.left,min,root.val)&&dfs(root.right,root.val,max)
    对于返回有两个疑问：
    1.dfs(root.left,min,root.val)代表着更新最大值。因为往左子树跑了。最大值就变为root.val
    2.左右子树的dfs为什么是&&而不是||
    如果出现返回值为false,就该结束了，直接为false了。

    2.迭代



*/

var isValidBST = function(root) {
    let arr=[]
    let dfs=(root)=>{
        if(!root) return;
        dfs(root.left)
        arr.push(root.val)
        dfs(root.right)
    }
    dfs(root)
    for(let i=0;i<arr.length-1;i++){
        if(arr[i]>=arr[i+1]){
            return false
        }
    }
    return true
}


// 优化空间复杂度
var isValidBST = function(root) {
    let max=-Infinity
    let dfs=(root)=>{
        if(!root) return true;
        let left=dfs(root.left)
        if(max<root.val){
            max=root.val
        }
        else{
            return false
        }
        let right=dfs(root.right)
        return left&&right
    }
   return  dfs(root)
}


var isValidBST = function(root) {
    let arr=[]
    let dfs=(root,min,max)=>{
        if(!root) return true;
        if(root.val<=min||root.val>=max) return false
        return dfs(root.left,min,root.val)&&dfs(root.right,root.val,max)
    }
    return dfs(root,-Infinity,Infinity)
}

var isValidBST = function(root) {
    let queue=[],pre=null
    while(root||queue.length>0){
        if(root){
            queue.push(root)
            root=root.left
        }
        else{
            root=queue.pop()
            if(pre&&pre.val>=root.val){
                return false
            }
            pre=root
            root=root.right
        }
    }
    return true
}
/*
    时间复杂度：O(N)
    空间复杂度：O(N)

*/
/*
    最近一直专注在递归的前中序遍历，迭代也不能放弃，都多练练，最重要还是要深入了解各种遍历方式
*/