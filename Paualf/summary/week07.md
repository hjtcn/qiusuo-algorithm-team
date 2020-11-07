第7周总结：

week07 小结(20201102~200201106)

## 

* First day: [[剑指Offer59]滑动窗口的最大值](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)

* The next day: [[622]设计循环队列](https://leetcode-cn.com/problems/design-circular-queue/)

* Third day: [[933]最近的请求次数](https://leetcode-cn.com/problems/number-of-recent-calls/)

* Fourth day: [[641]设计双端循环队列](https://leetcode-cn.com/problems/design-circular-deque/)

* Fifth day: [[面试题17.09]第k个数](https://leetcode-cn.com/problems/get-kth-magic-number-lcci/)

## 我的题解

1.滑动窗口的最大值

单调递减队列，需要对队列比较熟悉，然后使用递减队列实现，最后使用数组将数据结构和算法模拟出来

2.设计循环队列

想明白 队空，队满，队中元素个数的关系

队空：rear == front

队满：(rear + 1) % size == front

队中元素个数：

rear 有可能大于 front

当 rear 在下面，也有可能小于front，所以队列元素个数为：

(size - front + rear - 0)  % size => (rear - front + size) % size

3.最近的请求个数

将题目意思理解清楚然后使用程序实现，也是需要下功夫的，和工作和工程中做需求是一样的

4.设计双端循环队列

将第二题中的条件搞清楚以后，多了向首部添加元素和从尾部取出元素的操作，数组实现的话注意下标的变化

链表实现还是不是很熟悉，需要多多注意和练习，将数据结构抽象出来

5.第k个数

感觉更像是数学题，数学基础有点弱，后面需要多注意

总结：

1. 好的数据结构+好的算法 = 程序
2. 程序的灵魂在算法，算法的灵活在数学
3. 多思考的同时要多写代码，不要害怕不会和写错
