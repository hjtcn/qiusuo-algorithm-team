// 100. 相同的树
// 给定两个二叉树，编写一个函数来检验它们是否相同。

// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

// 示例 1:

// 输入:       1         1
//           / \       / \
//          2   3     2   3

//         [1,2,3],   [1,2,3]

// 输出: true
// 示例 2:

// 输入:      1          1
//           /           \
//          2             2

//         [1,2],     [1,null,2]

// 输出: false
// 示例 3:

// 输入:       1         1
//           / \       / \
//          2   1     1   2

//         [1,2,1],   [1,1,2]

// 输出: false


/*
 * @lc app=leetcode.cn id=100 lang=javascript
 *
 * [100] 相同的树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {boolean}
 */
var isSameTree = function (p, q) {
    //为false的三种情况：1.p和q节点都有值却不想等；2.p有值q无；3:q有值p无
    if ((p && q && p.val !== q.val) || (p && !q) || q && !p) {
        return false
    }
    //当p没值，q也没值的时候，不会再深层递归下去，返回true
    if (!p && !q) return true
    //深层递归，分别对比左子树的值和右子树的值
    return p && q && isSameTree(p.left, q.left) && isSameTree(p.right, q.right)
};
// @lc code=end

/*
 比较搞笑的是，我上去以为输入信息就是数组，就以为题目老简单了，对比两个数组是否相同。提交之后看到红色的解答错误，哈哈哈哈

 回归正题哈，鉴于以前坐链表题的经验，看到上面声明treeNode结构的时候，已经不慌张了。

 递归。缕清结束递归的几种情况
 1.返回false:
            1)两个节点都有值，但值不想等
            2）一个节点有值，一个节点没值。
 2.返回true:
            p，q都没有值

  时间复杂度：O(min(m,n))
  递归，对比左子树和右子树的值
  空间复杂度：O(min(m,n))
  这个真的是迷惑了，不曾借助额外空间，但是树节点会占据空间，空间大小取决于节点数的多少
*/


/*
    上面是我自己写的，只想到了递归，没有想到分为深度优先或是广度优先。
    看了看题解，可以了解到，如果将思路规划为深搜或者广搜的方式，代码会更为整洁。
*/
//深度优先搜索dfs
var isSameTree = function (p, q) {
    if (p == null && q == null) //两个节点都为空，跳出递归，返回true,对比到最深处
        return true;
    if (p == null || q == null)//走到这一步，说明p或着q其中一个为空，返回false
        return false;
    if (p.val != q.val)//走到这一步说明，p,q都不为空，此时来对比val是否相等&&left是否相等&&right是否相等
        return false;
    //都不用进行p&&q的防止为空操作了。
    return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
};


//广度优先搜索bfs

var isSameTree = function (p, q) {
    if (p == null && q == null)
        return true
    if (p == null || q == null)
        return false
    let queueP = [p], queueQ = [q]
    //将节点置于队列，通过头部shift，进行对比每一个节点本身，左子树，右子树
    while (queueP.length && queueQ.length) {
        let nodeP = queueP.shift(), nodeQ = queueQ.shift()
        if (nodeP.val != nodeQ.val)
            return false
        if (nodeP.left && nodeQ.left) {
            //如果p,q节点都有左子树，则push进入各自队列。
            queueP.push(nodeP.left)
            queueQ.push(nodeQ.left)
        }
        else if (nodeP.left || nodeQ.left) {
            //其中一个为空，返回false
            return false
        }

        //同上
        if (nodeP.right && nodeQ.right) {
            queueP.push(nodeP.right)
            queueQ.push(nodeQ.right)
        }
        else if (nodeP.right || nodeQ.right) {
            return false
        }
    }
    //对比到最后两个队列都为空。
    return !queueP.length && !queueQ.length
};
