/*
222. 完全二叉树的节点个数
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

 

示例 1：


输入：root = [1,2,3,4,5,6]
输出：6
示例 2：

输入：root = []
输出：0
示例 3：

输入：root = [1]
输出：1
 

提示：

树中节点的数目范围是[0, 5 * 104]
0 <= Node.val <= 5 * 104
题目数据保证输入的树是 完全二叉树
 

进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？
*/

/*  
    1.思路：两种遍历方式
    （1)递归遍历
     (2)迭代遍历
    迭代的框架写起来有点不是那么顺手了，应该是隔的时间太长了。
    2.题解。看了二进制的看不太懂，找了代码随想录的题解，感觉更好理解一点。
    利用完全二叉树的特性
    1.如果是满二叉树，个数就是2^N-1
    2.如果不是满二叉树，那么它的左右子节点肯定有是满二叉树的，直到遇到满二叉树为止，一个节点也是满二叉树
    为什么它的时间复杂度是O(logN*logN)？？？
    判断满二叉树通过判断
    left的left...就是一直在最左边
    right的right....一直在最右边
    且深度还相同，就是满二叉树。O(logN)
    不是满二叉树，接着递归左子树或又子树。O(logN)
*/

var countNodes = function (root) {
    if (!root) return 0
    let dfs = (root) => {
        if (!root) return 0
        let left = 0, right = 0
        if (root.left) {
            left = dfs(root.left) + 1
        }
        if (root.right) {
            right = dfs(root.right) + 1
        }
        return left + right
    }
    return dfs(root) + 1
}
/*
    时间复杂度：O(N)
    空间复杂度：O(logN)
*/

var countNodes = function (root) {
    if (!root) return 0
    let stack = [root], count = 1
    while (stack.length) {
        let len = stack.length
        for (let i = 0; i < len; i++) {
            let curNode = stack.shift()
            count++;
            if (curNode.left) {
                stack.push(curNode.left)
            }
            if (curNode.right) {
                stack.push(curNode.right)
            }
        }

    }
    return count

}
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

var countNodes = function (root) {
    if(!root) return 0
    let left=root.left,leftHeight=1
    let right=root.right,rightHeight=1
    while(left){
        left=left.left
        leftHeight++
    }
    while(right){
        right=right.right
        rightHeight++
    }
    if(leftHeight==rightHeight){
        return Math.pow(2,leftHeight)-1
    }
    return countNodes(root.left)+countNodes(root.right)+1
}
/*
    时间复杂度：O(logN * logN)
    空间复杂度：O(logN)

*/

/*
    思考：先分成大块，再分成小块，直到为一个节点。判断满二叉树也有点厉害
*/