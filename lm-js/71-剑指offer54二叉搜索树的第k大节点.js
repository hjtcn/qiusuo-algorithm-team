// 剑指 Offer 54. 二叉搜索树的第k大节点
// 给定一棵二叉搜索树，请找出其中第k大的节点。



// 示例 1:

// 输入: root = [3,1,4,null,2], k = 1
//    3
//   / \
//  1   4
//   \
//    2
// 输出: 4
// 示例 2:

// 输入: root = [5,3,6,2,4,null,null,1], k = 3
//        5
//       / \
//      3   6
//     / \
//    2   4
//   /
//  1
// 输出: 4


// 限制：

// 1 ≤ k ≤ 二叉搜索树元素个数


var kthLargest = function (root, k) {
    /*根据提示，root不会为空 */
    // if(!root){
    //     return null
    // }
    let res = []
    const dfs = (root) => {
        if (!root) {
            return;
        }
        dfs(root.right)
        res.push(root.val);
        dfs(root.left)
    }
    dfs(root)
    return res[k - 1];
};

/**
 * 题解:自己没有做出来，看到题解中说明中序遍历(左根右)是递增序列
 * 题目要求：第k大的节点，因此可以中序遍历倒序(右根左)往数组中push节点,然后返回res[k-1]即可
 *  时间复杂度：O(N)
    取决于节点个数
    
    空间复杂度：O(N)
    取决于队列节点数，不会超过n
 */

var kthLargest = function (root, k) {
    let res = null, num = 0
    const dfs = (root) => {
        if (!root) {
            return;
        }
        dfs(root.right)
        num++;
        if (num == k) {
            res = root.val
            return;
        }
        dfs(root.left)
    }
    dfs(root)
    return res
};
//计数法，提前退出递归，且优化了空间复杂度，不用数组存储节点


// 迭代
var kthLargest = function (root, k) {
    let stack = [root], res = null, num = 0
    while (root || stack.length) {
        while (root) {
            stack.push(root)
            root = root.right
        }
        root = stack.pop()
        num++;
        if (num == k) {
            return root.val
        }
        root = root.left
    }
};

/**
    这个是上周四刷的中序遍历拓展出来的方法。
    其实这道题主要还是中序遍历：
    1.中序遍历是递增序列
    2.了解中序遍历的解法，多加练习
    时间复杂度：O(N)
    取决于节点个数

    空间复杂度：O(N)
    取决于队列节点数，不会超过n
 */