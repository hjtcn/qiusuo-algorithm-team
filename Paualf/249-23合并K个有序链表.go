给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]

1. Clarification:
合并两个有序链表的时候可以用 if l1.Val <= l2.Val 合并多个有序链表的时候怎么判断呢？好像不能这样判断了

2. 看题解：
2.1 顺序合并：
使用合并两个链表的方式，一个一个的合并

上层是 
ans := &ListNode{}
for i := 0;i < len(list);i++ {
ans = mergeTwoLists(ans,list[i])
}
底层是
 mergeTwoLists(list1 *ListNode,list2 *ListNode) {

}

方法二：分治合并
优化顺序合并，用分治的方法进行合并
将k个链表配对并将同一对中的链表合并
第一轮合并以后，k个链表被合并成了k/2个链表，平均长度为2n/k
重复这个过程，直到我们得到了最终的有序链表

方法三：使用优先队列合并：
维护当前每个链表没有被合并的元素的最前面一个,k个元素最多有k个满足这样条件的元素，每次在这些元素里面选取val属性最小的元素合并到答案中。在选取最小元素的时候，我们可以用优先队列来优化这个过程


分治的代码看着还是有一丢丢看不懂，写的话不一定可以写出来
func mergeKLists( lists []*ListNode ) *ListNode {
    length := len(lists)
    if length == 0 {return nil}
    if length == 1 {return lists[0]}
    mid := length>>1                                              //分治
    return meger(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))  //归并
}

func meger(l1 *ListNode, l2 *ListNode) *ListNode{
    tmp := new(ListNode)
    ans := tmp
    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
            tmp.Next = l2
            l2 = l2.Next
            tmp = tmp.Next
        }else{
            tmp.Next = l1
            l1 = l1.Next
            tmp = tmp.Next
        }
    }
    if l2 != nil {
        tmp.Next = l2
    }
    if l1 != nil {
        tmp.Next = l1
    }
    return ans.Next
}
go 最小堆
type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := new(minHeap)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
		pre.Next = tmp
		pre = pre.Next
	}

	return dummyHead.Next
}

3. 复杂度分析：
时间复杂度：O(nk x logK），分治和优先队列的复杂度现在还不会分析，分治还不太熟悉，可以多注意下
空间复杂度：O(logK)

4. 总结：
4.1: 当问题的复杂度变多了，程序也就不好维护了
4.2: 复杂的问题可以将它转化成比较简单的问题，然后一点点的解决
