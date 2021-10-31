/*
    450. 删除二叉搜索树中的节点
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
 

示例 1:



输入：root = [5,3,6,2,4,null,7], key = 3
输出：[5,4,6,2,null,null,7]
解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
另一个正确答案是 [5,2,6,null,4,null,7]。


示例 2:

输入: root = [5,3,6,2,4,null,7], key = 0
输出: [5,3,6,2,4,null,7]
解释: 二叉树不包含值为 0 的节点
示例 3:

输入: root = [], key = 0
输出: []
 

提示:

节点数的范围 [0, 104].
-105 <= Node.val <= 105
节点值唯一
root 是合法的二叉搜索树
-105 <= key <= 105
 

进阶： 要求算法时间复杂度为 O(h)，h 为树的高度。
*/

/*
    思路：这个没有思考出来，只想到了，右子树上位，左子树移到右子树的叶子结点。
    步子跨的有点大了。只记得这个，忘了区分各种情况。
    1.无子树
    2.有左子树，无右子树
    3.有右子树，无左子树
    4.左右子树都有。然后再根据上面的思路去走。

    左子树迁移的时候也有问题，首先要迁到右子树的叶子结点上。所以。遍历到右子树时，要记得深层遍历，找到叶子结点再拼接。


    而进阶要求时间复杂度为O(h),也就是借用二叉搜索树的特性。
    如果root.val<key,向右子树找
    如果root.val>key,向左子树找
    如果相等，说明找到了要删除的节点。再利用上面的思路，进行上位，拼接，删除。




*/


var deleteNode = function (root, key) {
    if (!root) return null
    if (root.val < key) {
        root.right = deleteNode(root.right, key)
    }
    else if (root.val > key) {
        root.left = deleteNode(root.left, key)
    }
    else {
        if (!root.left) {
            root = root.right
        }
        else if (!root.right) {
            root = root.left
        }
        else {
            let cur = root.right
            while (cur.left) {
                cur = cur.left
            }
            cur.left = root.left
            root = root.right
            delete root
            return root
        }
    }
    return root
}


var deleteNode = function (root, key) {
    if (!root) return root
    let queue = [[root, 0]], prev = null
    while (queue.length) {
        let [curNode, flag] = queue.shift()
        if (curNode.val < key) {
            if (curNode.right) {
                prev = curNode
                queue.push([curNode.right, 2])
            }
        }
        else if (curNode.val > key) {
            if (curNode.left) {
                prev = curNode
                queue.push([curNode.left, 1])
            }
        }
        else {
            if (!curNode.left && !curNode.right) {
                curNode = null
            }
            else if (!curNode.left) {
                curNode = curNode.right
            }
            else if (!curNode.right) {
                curNode = curNode.left
            }
            else {
                let right = curNode.right
                while (right.left) {
                    right = right.left
                }
                right.left = curNode.left
                curNode = curNode.right
            }
            if (flag == 0) {
                return curNode
            }
            if (flag == 1) {
                prev.left = curNode
            }
            if (flag == 2) {
                prev.right = curNode
            }
        }
    }
    return root
}

/*
    时间复杂度：O(H)
    空间复杂度：O(H)
*/

/*
    思考：一步一步去拓展思路。就算记住了难点或想到了难点，同时也要往回把所有边界情况给想到，否则因为记住难点，而把基础思路给弄乱，反而有点得不偿失。
*/