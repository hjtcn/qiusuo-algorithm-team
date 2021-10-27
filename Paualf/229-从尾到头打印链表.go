输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：
输入：head = [1,3,2]
输出：[2,3,1]
 
限制：
0 <= 链表长度 <= 10000

1. Clarification:
1.1: 遍历链表，将元素放到数组中
1.2: 遍历数组，交换元素，i,j 进行交换
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    arr := []int{}

    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }

    arr = reverseArr(arr)

    return arr
}

func reverseArr(arr []int) []int {
    i,j := 0,len(arr) - 1

    for i < j {
        arr[i],arr[j] = arr[j],arr[i]
        i++
        j--
    }

    return arr
}

每次都新建一个切片，然后将原有元素放到这个切片的后面
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    arr := []int{}

    for head != nil {
        arr = append([]int{head.Val}, arr...)
        head = head.Next
    }

    return arr
}

2. 看题解：
递归进行求解
递归本身就是一个栈么
func reversePrint(head *ListNode) ([]int) {
	if head == nil {return nil}
	return append(reversePrint(head.Next), head.Val)
}

反转链表，然后再进行遍历，将元素放入数组中

3. 复杂度分析：
时间复杂度：O(n)
空间复杂度：O(n) 使用数组进行存储

4. 总结：
4.1: 简单题的解法还是蛮多的，开阔自己的思维
