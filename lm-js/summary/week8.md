week08 (20201109~20201115)

## 
* First day: [[100]相同的树](https://leetcode-cn.com/problems/same-tree/)
* The next day: [[104]二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)
* Third day: [[107]二叉树的层次遍历 II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)
* Fourth day: [[226]翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)
* Fifth day: [[101]对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/)

## 总结：

1. [相同的树](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/57-100%E7%9B%B8%E5%90%8C%E7%9A%84%E6%A0%91.js)

    - 递归。这道题递归的方式还是非常简单的，就是比较节点是否都存在，val是否相等
    - 广度优先搜索(bfs)。利用数组模拟队列。循环判断队首元素，然后判断队首元素的左右节点是否存在，一方不存在则返回false,都存在则入队。

2. [二叉树的最大深度](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/58-104%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%9C%80%E5%A4%A7%E6%B7%B1%E5%BA%A6.js)

    - 递归。重复调用左右子树，更新max+1作为返回值
    - 广度优先搜索(bfs)。横向记录节点个数，将当前层每个节点出队，然后其左右节点入队。

3. [二叉树的层次遍历II](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/59-107%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%B1%82%E6%AC%A1%E9%81%8D%E5%8E%86II.js)

    - 广度优先搜索(bfs).利用数组模拟队列入队出队，依次队首元素出队，其左右子树入队。

4. [翻转二叉树](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/60-226%E7%BF%BB%E8%BD%AC%E4%BA%8C%E5%8F%89%E6%A0%91.js)

    - 广度优先搜索(bfs)。在二叉树层次遍历的基础上，提前将左右子树借用变量进行交换
    - 递归。先重复调用左右子树，获取左右节点，再进行交换，最终返回root
    

5. [对称二叉树](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/lm-js/61-101%E5%AF%B9%E7%A7%B0%E4%BA%8C%E5%8F%89%E6%A0%91.js)

    - 递归。重复子问题为：检查左右节点是否存在且相等&&左节点的左孩子==右节点的右孩子&&左节点的右孩子==右节点的左孩子
    - 广度优先搜索(bfs)。构建队列格式，让需要对比的节点紧挨着放入队列。而需要对比的节点规律同上重复子问题。


## 心得体会：

1. 其实看我题解蛮清晰的，核心解决方案两种，递归和bfs,为什么只有这两种呢，因为我当时认为递归就是dfs，但又觉得两者还是不同的，后来查询了一下递归，回溯，dfs，动态规划的区别。

    - 递归是一种实现方式。自我调用
    - 回溯是一种通用算法，可用递归解决
    - dfs。则为回溯用于树时也就称为深度优先搜素了。
    - 动态规划。回溯可以用于所有用穷举法可以解决的问题，而DP只用于具有最优子结构的问题（需要记忆，存储子问题的解）。
2. 递归的核心在于终止方式和递进方式，还要多加练习，抓住发散思维的方向
3. bfs写着写着还是有一些规律可循，核心还是队列入队出队，去对比，去判断