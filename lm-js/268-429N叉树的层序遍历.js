/*
429. N 叉树的层序遍历
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

 

示例 1：



输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]
示例 2：



输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
 

提示：

树的高度不会超过 1000
树的节点总数在 [0, 10^4] 之间
*/

/*
    思路：层序遍历还是迭代的方式更好处理一些。注意边界控制.
         递归：我是写着写着突然想到深度记录。深搜的过程，记录用下标记录每一层

    
*/
var levelOrder = function (root) {
    if (!root) return []
    let stack = [root], target = []
    while (stack.length > 0) {
        let len = stack.length
        let curArr = []
        for (let i = 0; i < len; i++) {
            let curNode = stack.shift()
            curArr.push(curNode.val)
            if (curNode.children) {
                for (let j = 0; j < curNode.children.length; j++) {
                    if (curNode.children[j]) {
                        stack.push(curNode.children[j])
                    }
                }
            }
        }
        target.push(curArr)
    }
    return target
};


var levelOrder = function (root) {
    if (!root) return []
    let target = [[root.val]]
    let dfs = (root, depth) => {
        if (!root) return;
        for (let i = 0; i < root.children.length; i++) {
            if (target[depth]) {
                target[depth].push(root.children[i].val)
            }
            else {
                target[depth] = [root.children[i].val]
            }
            dfs(root.children[i], depth + 1)
        }
    }
    dfs(root, 1)
    return target
};

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：前中后序遍历使用递归方便
         层序遍历使用迭代方便
         不过都练习练习，更了解递归和迭代对于二(N)叉树的遍历顺序
*/