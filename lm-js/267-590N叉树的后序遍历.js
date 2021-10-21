/*
590. N 叉树的后序遍历
给定一个 N 叉树，返回其节点值的 后序遍历 。

N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

 

进阶：

递归法很简单，你可以使用迭代法完成此题吗?

 

示例 1：



输入：root = [1,null,3,2,4,null,5,6]
输出：[5,6,3,2,4,1]
示例 2：



输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[2,6,14,11,7,3,12,8,4,13,9,10,5,1]
 

提示：

N 叉树的高度小于或等于 1000
节点总数在范围 [0, 10^4] 内

*/


/*
    思路：左右根
        递归：对于N叉树也就是先从左到右遍历递归子节点，再遍历根
        迭代：先捋清遍历思路。然后再控制出栈入栈。
        最好看着N叉树的图可以更好理解
        stack  [root]
        出栈pop    但是根要放在最后，target.unshift(root.val)从头插入。
        children入栈  从左到右push




*/

/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node|null} root
 * @return {number[]}
 */
var postorder = function (root) {
    let target = []
    let dfs = (root) => {
        if (!root) return;
        let len = root.children.length
        for (let i = 0; i < len; i++) {
            dfs(root.children[i])
        }
        target.push(root.val)
    }
    dfs(root)
    return target
};


var postorder = function (root) {
    if (!root) return []
    let target = []
    let stack = [root]
    while (stack.length) {
        let curNode = stack.pop()
        target.unshift(curNode.val)
        let len = curNode.children.length
        for (let i = 0; i < len; i++) {
            stack.push(curNode.children[i])
        }
    }
    return target
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
/*
    思考：还是先确定了解节点遍历顺序，再伪代码模拟
*/