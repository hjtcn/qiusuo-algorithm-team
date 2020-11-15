week08 (20201109~20201113)

## 
* First day: [[100]相同的树](https://leetcode-cn.com/problems/same-tree/)
* The next day: [[104]二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)
* Third day: [[107]二叉树的层次遍历 II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)
* Fourth day: [[226]翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)
* Fifth day: [[101]对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/)

## 总结：

1.相同的树：
考查树的结构，结构相同 => 同时遍历两个树，判断节点是否相同

2.二叉树的最大深度

dfs: 分析清楚问题的重复性在那里，用计算机帮我们解决这些重复的问题。首先还是我们找到那个是重复的子问题。
bfs: 使用队列实现遍历，水波纹一样。卡壳的地方是在于记录当前层的节点个数  size = len(queue)

3.二叉树的层次遍历II

bfs :使用队列进行广度优先遍历，注意结果数组形式

4.翻转二叉树

找问题的重复性：将节点的右节点和左节点交换

5.对称二叉树

用两个指针进行遍历：思维还是没有打开

使用队列进行存储，本质上还是广度优先遍历，换了一些进队列的元素，就不会了，说明还是没有掌握

总结：
1. 改善自己的思维：找问题的重复性，解决小问题，小问题的结果 => 大问题
2. 做一遍是不会的，所以目标不是做出这一道题，而是通过这一道题我是否学习到东西，暴露出来自己的问题



