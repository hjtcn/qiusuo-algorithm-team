/*
给出一个完全二叉树，求出该树的节点个数。
说明：
完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。
示例:
输入: 
    1
   / \
  2   3
 / \  /
4  5 6
输出: 6
*/

Clarfication:
完全二叉树，求出该树的节点个数
感觉这道题目是要根据提示用到 2 ^ h 个，可是我想不到怎么用它
我想到的是遍历整个二叉树节点然后统计个数：时间复杂度是O(n)

Time/Space Complexity:
Time Complexity: 遍历二叉树每个节点 O(n)
      Space Complexity: O(n) 递归栈用到的空间

Coding:
前序遍历统计节点个数：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    ret := 0
    helper(root,&ret)
    return ret
}
func helper(root *TreeNode,ret *int) {
    if root == nil {
        return
    }
    *ret++
    helper(root.Left,ret)
    helper(root.Right,ret)
}
广度优先遍历统计节点个数：
func countNodes(root *TreeNode) int {
    ret := 0
    if root == nil {
        return ret
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:len(queue)]
        ret++
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue,node.Right)
        }
    }
    return ret
}

这两块代码都可以写出来了，还是没有想到这道题目想要的解决方法，看题解了。。。
利用完全二叉树的特性递归求解：比较左右子树的高度，判断那边是完全二叉树，然后使用公司 2 ^h 次方进行求解，没有-1是将父节点计算进去了
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := countLevel(root.Left)
    right := countLevel(root.Right)
    if left == right {
        return countNodes(root.Right) + (1 << left)
    }else {
        return countNodes(root.Left) + (1 <<  right)
    }
}
func countLevel(root *TreeNode)int {
    level := 0
    for root != nil {
        level++
        root = root.Left
    }
    return level
}

题解中的利用两次二分查找还是值得品味的。人脑计算的话，我们计算出树的高度为 h，节点n= 2 ^h - 1 + 最后一层节点的个数，问题就是需要判断最后一层有多少个节点，节点是否存在。
主函数的二分查找是缩小边界，再使用另外一个二分查找判断是否存在。
第一个二分查找的退出边界还是比较有意思的，left > right 的时候退出，这个点自己也思考了一下
判断是否存在的二分查找也是蛮巧妙的，让我自己写，可能写不出来，也想不清楚的。
题解中Java的代码: 计算高度的代码也是少了一层，方便下面代码进行计算

class Solution {
  // Return tree depth in O(d) time.
  public int computeDepth(TreeNode node) {
    int d = 0;
    while (node.left != null) {
      node = node.left;
      ++d;
    }
    return d;
  }
  // Last level nodes are enumerated from 0 to 2**d - 1 (left -> right).
  // Return True if last level node idx exists. 
  // Binary search with O(d) complexity.
  public boolean exists(int idx, int d, TreeNode node) {
    int left = 0, right = (int)Math.pow(2, d) - 1;
    int pivot;
    for(int i = 0; i < d; ++i) {
      pivot = left + (right - left) / 2;
      if (idx <= pivot) {
        node = node.left;
        right = pivot;
      }
      else {
        node = node.right;
        left = pivot + 1;
      }
    }
    return node != null;
  }
  public int countNodes(TreeNode root) {
    // if the tree is empty
    if (root == null) return 0;
    int d = computeDepth(root);
    // if the tree contains 1 node
    if (d == 0) return 1;
    // Last level nodes are enumerated from 0 to 2**d - 1 (left -> right).
    // Perform binary search to check how many nodes exist.
    int left = 1, right = (int)Math.pow(2, d) - 1;
    int pivot;
    while (left <= right) {
      pivot = left + (right - left) / 2;
      if (exists(pivot, d, root)) left = pivot + 1;
      else right = pivot - 1;
    }
    // The tree contains 2**d - 1 nodes on the first (d - 1) levels
    // and left nodes on the last level.
    return (int)Math.pow(2, d) - 1 + left;
  }
}

总结：
1. 完全二叉树的概念容易理解，计算节点个数也容易推导和计算，转化为代码还是有难度的

2. 二分查找并不简单，细节是魔鬼，一个大于号或者小于号写错了，代码就凉凉了

3. 题解文字如果看不懂的话，就去看图，理解了大概，就去看代码，不要去纠结思路，动手做，如果实在看不懂，就动手，不要停留在原地，一直想，是没有结果的
4. 空杯心态，不要题目做出来了，就不去看好的解法和好的代码，也不要害怕，如果一直这样，是不可能进步和提高的
