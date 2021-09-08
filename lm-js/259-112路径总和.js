/*
112. 路径总和
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。

叶子节点 是指没有子节点的节点。

 

示例 1：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
示例 2：


输入：root = [1,2,3], targetSum = 5
输出：false
示例 3：

输入：root = [1,2], targetSum = 0
输出：false
 

提示：

树中节点的数目在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000

*/

/*
    思路：递归参数记录剩余差值和当前节点，题目不是很难，但是坑比较多。
        坑(1)：一开始对于差值的判断放在if(!root)了。后面才get到它并不能代表是叶子节点。
            那么边界应该是什么呢？
            if(!root) return false
        坑(2):也是由于坑1，如果节点一开始空节点呢？应该返回false,因此我借用了帮助函数dfs。
        坑(3):节点值可能为负数，因此，差值为负的时候不能提前返回

*/


var hasPathSum = function(root, sum) {
    if(!root) return false;
    if(root&&!root.left&&!root.right&&sum==root.val){
        return true
    }
    return hasPathSum(root.left,sum-root.val)||hasPathSum(root.right,sum-root.val)
}


var hasPathSum = function(root, sum) {
    if(!root) return false;
    let queue=[[root,sum]]
    while(queue.length){
        let len=queue.length
        for(let i=0;i<len;i++){
            let [curNode,diff]=queue.shift()
            if(!curNode.left&&!curNode.right&&diff==curNode.val){
                return true
            }
            if(curNode.left){
                queue.push([curNode.left,diff-curNode.val])
            }
            if(curNode.right){
                queue.push([curNode.right,diff-curNode.val])
            }

        }
    }
    return false
}


/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：注意细节，边界
*/