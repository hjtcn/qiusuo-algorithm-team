/*
1325. 删除给定值的叶子节点
给你一棵以 root 为根的二叉树和一个整数 target ，请你删除所有值为 target 的 叶子节点 。

注意，一旦删除值为 target 的叶子节点，它的父节点就可能变成叶子节点；如果新叶子节点的值恰好也是 target ，那么这个节点也应该被删除。

也就是说，你需要重复此过程直到不能继续删除。

 

示例 1：



输入：root = [1,2,3,2,null,2,4], target = 2
输出：[1,null,3,null,4]
解释：
上面左边的图中，绿色节点为叶子节点，且它们的值与 target 相同（同为 2 ），它们会被删除，得到中间的图。
有一个新的节点变成了叶子节点且它的值与 target 相同，所以将再次进行删除，从而得到最右边的图。
示例 2：



输入：root = [1,3,3,3,2], target = 3
输出：[1,3,null,null,2]
示例 3：



输入：root = [1,2,null,2,null,2], target = 2
输出：[1]
解释：每一步都删除一个绿色的叶子节点（值为 2）。
示例 4：

输入：root = [1,1,1], target = 1
输出：[]
示例 5：

输入：root = [1,2,3], target = 1
输出：[1,2,3]
 

提示：

1 <= target <= 1000
每一棵树最多有 3000 个节点。
每一个节点值的范围是 [1, 1000] 。
*/

/*
    思路：递归：写的时候就想先把递归左右节点的返回值。
        什么时候删除？叶子节点等于target：也就是没有子节点了
        比如left=dfs(root.left)
        当left&&!left.left&&!left.right&&left.val==target时
        可以删除了。root.left=null
        right同理。
        最终返回root
        踩坑：case:[1,1,1]  target=1
        此时要将所有节点删除完。但是我的思路是删除左右节点呀？
        最后想了一下，决定将root拿到递归函数外面处理。
        如果最后就剩一个根节点了，且val还等于target,则返回null
        其它返回root就可以了。

    题解：看过题解，感觉将root.left=dfs(root.left);root.right=dfs(root.right)更巧妙一些。目标就盯准root去处理，去返回


*/

var removeLeafNodes = function(root, target) {
    let dfs=(root)=>{
        if(!root) return null;
        let left=dfs(root.left)
        let right=dfs(root.right)
        if(left&&left.val===target&&!left.left&&!left.right){
            root.left=null
        }
        if(right&&right.val===target&&!right.left&&!right.right){
            root.right=null
        }
        return root
    }
    dfs(root)
    if(root.val==target&&!root.left&&!root.right){
            return null
    }
    else{
        return root
    }
};

var removeLeafNodes = function(root, target) {
    let dfs=(root)=>{
        if(!root) return null;
        root.left&&(root.left=dfs(root.left))
        root.right&&(root.right=dfs(root.right))
        if(root.val===target&&!root.left&&!root.right){
            return null
        }
        return root
    }
    return dfs(root)
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
/*
    思考：如果前中后序遍历不够熟悉，就先分左右子节点，结构分散开去理解，去模拟。
         好吧，我承认，这三种遍历是真的记不住，目前就记住根左右，左根右，左右根。
*/