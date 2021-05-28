输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
示例 1：
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：
输入：arr = [0,1,2,1], k = 1
输出：[0]
 
限制：
0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000

1. Clearfication:
1.1. Clearfication: 第一想法是排序后，然后获取前k个
1.2. 如果数量特别多，内存中排不了序，需要使用堆排序去比较

2. Coding:
func getLeastNumbers(arr []int, k int) []int {
    sort.Ints(arr)
    ret := []int{}
    for i := 0;i < k;i++ {
        ret = append(ret, arr[i])
    }
    return ret
}

3. 看题解：
quickSearch
func getLeastNumbers(arr []int, k int) []int {
    if len(arr) == 0 || k == 0 {
        return nil
    }
    
    return quicksearch(arr,0,len(arr) - 1,k)
}
func quicksearch(nums []int,i,j,k int) []int {
    t := partition(nums,i,j)
    
    if t == k - 1 {
        return nums[:k]
    }
    
    if t < k - 1 {
        return quicksearch(nums,t + 1,j,k)
    }
    
    return quicksearch(nums, i ,t - 1,k)
}
func partition(nums []int,i,j int) int {
    l,m,r:=i,i,j
    for l<r {
        for l<r && nums[r]>=nums[m] {
                r--
        }
        for l<r && nums[l]<=nums[m] {
                l++
        }
        if l<r {
            nums[l],nums[r]=nums[r],nums[l]
        }
    }
    nums[m],nums[l]=nums[l],nums[m]
    
    return l
}
quicksort
func getLeastNumbers(arr []int, k int) []int {
    if len(arr) == 0 || k == 0 {
        return nil
    }
    
     quicksort(arr, 0, len(arr)-1)
     return arr[:k]
}
func quicksort(nums []int,i,j int) {
    if i>=j {
        return
    }
    m:=partition(nums, i,j)
    quicksort(nums,i,m-1)
    quicksort(nums,m+1,j)
}
func partition(nums []int,i,j int) int {
    l,m,r:=i,i,j
    for l<r {
        for l<r && nums[r]>=nums[m] {
                r--
        }
        for l<r && nums[l]<=nums[m] {
                l++
        }
        if l<r {
            nums[l],nums[r]=nums[r],nums[l]
        }
    }
    nums[m],nums[l]=nums[l],nums[m]
    
    return l
}

堆
func getLeastNumbers(arr []int, k int) []int {
    if len(arr)==0 || k==0 {
        return nil
    }
    //方法三  建堆，大根堆
    d:=&heapInt{}
    for _,v:=range arr {
        if d.Len()<k {
            heap.Push(d,v)
        }else {
            if d.Peek()>v {
                heap.Pop(d)
                heap.Push(d,v)
            }
        }
    }
    return *d
}

type heapInt []int
//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt)Less(i,j int)bool {return (*h)[i]>(*h)[j]}
func (h *heapInt)Swap(i,j int) {(*h)[i],(*h)[j]=(*h)[j],(*h)[i]}
func (h *heapInt)Len() int {return len(*h)}
func (h *heapInt)Push(x interface{}) {
    *h=append(*h,x.(int))
}
func (h *heapInt)Pop() interface{} {
    t:=(*h)[len(*h)-1]
    *h=(*h)[:len(*h)-1]
    return t
}
func (h *heapInt)Peek() int {
    return (*h)[0]
}

4. 复杂度分析：
时间复杂度：O(nlogn)
空间复杂度：O(n)

5. 总结：
5.1: 堆不是很熟悉，快排核心的 partion 自己写的话估计也写不出来，要理解还要多动手多画图
