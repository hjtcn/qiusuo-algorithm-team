week06小结(20201026～20201030)

### 本周题目：
* [[155]最小栈](https://leetcode-cn.com/problems/min-stack/)
* The next day: [[150]逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)
* Third day: [[844]比较含退格的字符串](https://leetcode-cn.com/problems/backspace-string-compare/)
* Fourth day: [[316]去除重复字母](https://leetcode-cn.com/problems/remove-duplicate-letters/[[503]])
* Fifth day: [[503]下一个更大元素II](https://leetcode-cn.com/problems/next-greater-element-ii/)

我的题解总结：

1.最小栈

主要是想要在常数时间内获取到栈中最小元素，使用辅助栈来记录当下最小元素，然后栈元素变化时，辅助栈同时变化。

思想：空间换时间，栈+栈/队列+队列

2.逆波兰表达式求值

题目本身思路不难，一些小细节可能会引起程序出现bug

如 除法和减法，从栈中弹出元素的时候要注意

3.比较含退格的字符串

栈的方法比较容易相当，双指针法写的时候要想清楚，再去写，不然可能调试会比较久

4.去除重复字母

该题目需要满足三个要求：去重，判断是否使用过，后面元素是否还存在，然后使用单调栈进行解决

单个知识点都不难，但是合起来写出没有bug的代码还是有难度的，工程项目中也是如此，单个部分可能不是很复杂，集成到一起就会思考不过来，这个时候我们可以画图，将题目拆成小模块，更方便去理解。

5.下一个更大元素II

将数组扩展成两倍然后遍历或者遍历两遍数组

都是空间复杂度和时间复杂度的权衡，尽可能地去写出空间复杂度和时间复杂度低的代码


### 本周总结：

1. 利用栈的数据结构，可以减少时间复杂度，可以从O(N*N) 减少到 O(N)

2. 做题目是空间复杂度和时间复杂度之间的 trade-off,做项目也是如此，需要从时间，人力，成本，架构等等元素之间 trade-off，没有完美的人或事情，只有我们想清楚自己要什么，可能会做的更好



