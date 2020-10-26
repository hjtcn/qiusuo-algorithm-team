暴力法，一开始想到的是这种方法，从nums1中元素，找到在nums2 数组的下标，然后开始进行查找，查找到比它大的返回。一开始没看出来和栈有啥关系。。。
/*
    思路：
        1. 从 nums1 中的每个数字通过 hashmap 找到 nums2 的下标 start
        2. 从nums2 中下标 start 开始寻找start 后面是否有比该元素大的数字，如果有的话返回数字
        3. 返回结果
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    hash := map[int]int{}
    ret := []int{}
   
    for i,v := range nums2 {
        hash[v] = i
    }
    for i := 0;i < len(nums1);i++ {
        tmp := -1
        if v,ok := hash[nums1[i]]; ok {
            for j := v + 1;j < len(nums2);j++ {
                if nums2[j] > nums1[i] {
                    tmp = nums2[j]
                    break
                } 
            }
        }
        ret = append(ret, tmp)
    }
    return ret
}

单调栈是看题解后知道是使用单调栈，为什么没有想到呢？
对栈还是不是很熟悉，这种思维还是没有完全掌握，特别是栈中存储数组下标还是存储数组的值，很是不清晰，说明对问题还是没有理解清楚

/*
    单调栈：
        1. 将nums2 数据元素放入到栈中，然后如果元素比栈顶元素小则push到栈中，如果大于栈顶元素，将栈顶元素和将要push到栈中元素构建一个hash map 放入到map中
        2. 遍历num1 数组，寻找map中对应元素
        3. 返回结果数组
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    ret := []int{}
    dict := map[int]int{}
    stack := []int{}
    for i := 0;i < len(nums2);i++ {
        for len(stack) > 0 && nums2[i] > stack[len(stack) - 1] {
            // 构建hashmap
            dict[stack[len(stack) - 1]] = nums2[i]
            // pop
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, nums2[i])
    }
    // stack 还有元素
    for len(stack) > 0 {
        dict[stack[len(stack) - 1]] = -1
        stack = stack[:len(stack) - 1]
    }
    for i := 0;i < len(nums1);i++ {
        value := dict[nums1[i]]
        ret = append(ret, value)
    }
    return ret
}

总结：
1. 对问题没有理解清楚，所以没有想到使用单调栈

2. map 使用不熟悉，定义以及查找map元素是否存在

3. 多看题解，看别人的代码，形成闭环和反馈，才知道自己的代码和思路是真的好，还是仅仅在自嗨，自嗨有毒，培养开放平和的心态，对技术和生活都不要有偏见。平行世界，多元生活
