# 第三周

## 20200928-20201004

> 本周主要刷的是链表，由我找的五个题，个人觉得偏基础，但是质量也不错！难易程度为简单！

### 本周题目

-  [面试题 02.02] [返回倒数第 k 个节点](https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/)
-  [2] [两数相加](https://leetcode-cn.com/problems/add-two-numbers/description/)
-  [21] [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/description/)
-  [445] [两数相加 II](https://leetcode-cn.com/problems/add-two-numbers-ii)
-  [160] [相交链表](https://leetcode-cn.com/problems/intersection-of-two-linked-lists)


### 我的题解

1. [返回倒数第 k 个节点](../15-剑指offer22.链表中倒数第k个节点.php)
    - 快慢指针
      - 定义两个指针，fast快，slow慢，一开始都是同一个位置。
      - 现在fast从头移动k个长度的节点位置，那么也就是后面还剩余 len-k
      - 然后slow再和fast一起走，直到fast走到结尾部分(len-k长度)，那么slow也就走len-k，因此slow最后的位置就是距离尾部k的位置。
      - 返回slow指向的节点即可。

2. [两数相加](../16-2.两数相加.php)
    - 链表遍历
    - 其实就是两个数字相加，像我们平时自己计算一样，不过这里的链表顺序跟计算顺序刚好相同，个位相加，然后十位相加
      - 进位的话，两个个位数相加之和一定小于20
      - 那么只需要判断当前对应位置之和大于等于10，carray就可以进位+1
    - 头指针
      - 新定义一个cur链表，用来指向连个链表之和。
      - 当我们两个链表都遍历结束的时候，当前cur指针已经指向尾端。
      - 额外定义一个头指针out，采用赋值的方式等于cur指针，对象的赋值默认是引用类型
      - 最后可以直接用out，这个就是整个链表
    
3. [合并两个有序链表](../17-2.合并两个有序链表.php)
    - 题意重点
      - 合并两个有序链表，其实跟合并两个有序数组差不多
      - 都是升序链表，相对简单一些。
      - 直接开干
    - 遍历法
      - 遍历两个链表，比较当前值，谁小谁先加入合并链表。
      - 这里两种思路，一种是&&判断循环条件，一种是||判断循环条件。（老黑的思路很棒，采用||效率会更高一些）
      - 链表是否为空，采用!=null或者!==null效率也会不一样。（如果采用全等判断，性能也会慢下来）
    - 递归法
      - 代码量就简洁
       
4. [两数相加 II](../18-445.两数相加%20II.php)
    - 与第2天的题类似
    - 当然与上次不同的是，上次两数相加本来就是低位在链表的头节点位置，这次是高位在链表头节点
    - 所以要么反转链表，要么借助栈来存储逆序的链表。（栈 先进后出）
    - PHP这里采用数组来实现栈的功能。 
       - 我们可以将所有数字压入栈中，然后再依次取出来进行相加求和，同时考虑是否进位。
    - 头插法
      - 插入新的节点，先将当前节点插入链表头部，然后再用链表头部指向
      - 这样顺序也是正的
       
5. [相交链表](../19-160.相交链表.php)
    - 题意
      - 如果两个链表相交，第一个相交的节点就是起始节点，后面相交的长度是一样的，也就是值都相等。意思是后面会重合到一起。最后一个节点一定相等。
    - PHP
      - 采用全等===来判断，当前结点地址是否完全一样，比较的是对象地址
    - 暴力法
      - 这个就是遍历A链表，里面遍历B链表，然后看每一个A链表结点，是否存在B链表结点是一样的。 时间复杂度高，放弃！
    - 哈希表法
      - 遍历A，存储链表A每个节点的地址，追加到map中
      - 遍历B，判断看看有没有节点在map里面，有就相交重合。
    - 双指针法(最佳！！！)
      - 创建两个指针pA pB ，然后都指向两个链表的头节点，然后逐步向后遍历。
      - 当pA链表到达A链表的尾部的时候，重定向到B的头节点，
      - 当pB链表到达B链表的尾部的时候，重定向到A的头节点，
      - 如果在某一时刻相遇 pA == pB ，那么当前就是相交节点。
       

### 总结一下

#### 链表

- General
  - 链表的遍历。
  - 链表的递归。
  - 两条链表的合并=>两个数组的合并=>双指针
  - 两条链表相加=>两数相加
  - 链表的反转=>栈的使用=>PHP用数组模拟栈
- Special
  - 链表的头插法
  - 链表的头结点，利用变量指向头结点，最终用于返回整个链表
  - 链表的倒数第k个节点=>快慢指针
  - while循环中尽可能用||而不用&&, 因为||判断更快一些。

#### 感受

- 时间。本周是中间经历国庆假期，所以时间上也是比较难抽空做题的，所以可能建议之后假期就刷题。然后下周把题补一补，
- 坚持。原本以为一天刷一道题，比较轻松，或者有精力还能刷两道题。现在看来，好好坚持刷好一道题就挺不容易了。小涛，牛逼！



> 继续加油！奥力给！相信未来！