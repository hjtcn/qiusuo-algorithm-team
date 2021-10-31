给定一个链表的 头节点 head ，请判断其是否为回文链表。

如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。

 示例 1：
输入: head = [1,2,3,3,2,1]
输出: true

示例 2：
输入: head = [1,2]
输出: false

提示：
链表 L 的长度范围为 [1, 105]
0 <= node.val <= 9

进阶：能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

1. Clarification:
1.1：将链表元素放入到数组 中
1.2：对比数组首尾元素是否相同
1.3：如果不相同return false, 如果都相同则return true

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    arr := []int{}
    for head != nil {
        arr = append(arr,head.Val)
        head = head.Next
    }

    return isPalindromeArr(arr)
}

func isPalindromeArr(arr []int) bool {
    if len(arr) == 0 {
        return true
    }
    
    i,j := 0,len(arr) - 1
    for i < j {
        if arr[i] != arr[j] {
            return false
        } 

        i++
        j--
    }

    return true
}

2. 看题解：

递归：

这个是我没有想到的，还有题解里面的函数的写法也是自己项目中会有点懵逼的地方

func isPalindrome(head *ListNode) bool {
    frontPointer := head
    var recursivelyCheck func(*ListNode) bool 
    
    recursivelyCheck = func(curNode *ListNode) bool {
        if curNode != nil {
            if !recursivelyCheck(curNode.Next) {
            	return false
            }
            
            if curNode.Val != frontPointer.Val {
            	return false
            }
            
           	frontPointer = frontPointer.Next
        }
        return true
    }
    
    return recursivelyCheck(head)
}

快慢指针：
找到前半部分链表的尾节点
反转后半部分链表
判断是否回文
恢复链表
返回结果

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n)


4. 总结：
4.1: 越来越觉得写代码其实是一种表达方式，表达自己的想法和对具体问题的解决方式
