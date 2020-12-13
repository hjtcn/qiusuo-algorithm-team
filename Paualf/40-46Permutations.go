/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。
示例:
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

Clearfication:
3个盒子，3个数，有多少种不一样的排列
1开头 or 2开头 or 3开头
然后判断它是否已经使用过，使用一个map[int]int 判断它是否已经使用过
最后返回生成的全排列序列

预估复杂度：
时间复杂度：O(2^n)
空间复杂度: O(n) 递归栈空间的复杂度

整体思路：
1. 递归进行循环
2. 判断每一层是否使用过该元素，没有使用过将该元素放入到数组中，然后下探到下一层
细节点：
1. 开辟二维数组进行存储结果数组
2. 开辟一维数组进行存储当前层结果
3. 开辟map记录元素是否使用过
4. 使用level,num 记录终止条件
5. 使用for循环遍历每一层

发现细节点自己都知道，但是代码还是不一定可以写出来。。。
估计这也就是为什么要白板coding的原因，思路大家都知道，但是真正做出来的时候要对这些细枝末节了解的非常清楚，才能写出来可以使用的代码

半半成品，知道思路，也写不出来。。。。。。
func permute(nums []int) [][]int {
    ret := [][]int{}
    res := []int{}
    hash := map[int]int{}
    helper(nums,0,3,&ret,res)
    return ret
}
func helper(nums []int,level,num int,ret **[],res []int,hash map[int]int) {
    // terminator
    if level >= num {
        *ret = append(*ret, res)
        return
    }
    // process current logic
    for i := level;i < num;i++ {
        // 使用hash判断该元素是否使用过了，没有使用过的话放入当前数组中
        if exist,err := hash[i];err != nil {
        }
        res = append(res,i)
    }
     
    // drill down
    // restore current status
}


    
func permute(nums []int)[][]int {
    result := [][]int{}
    list := []int{}
    visited := make([]bool, len(nums))
    pos := 0
    helper(nums,pos,list,visited,&result)
    return result
}
        
func helper(nums []int,pos int,list []int,visited []bool,result *[][]int) {
    // terminator
    if pos == len(nums) {
        ans := make([]int, len(list))
        copy(ans, list)
        *result = append(*result, ans)
        return
    }
    
    for i := 0;i < len(nums);i++ {
        if visited[i] {
            continue
        }
        
        // process current logic
        list = append(list,nums[i])
        visited[i] = true
        
        // dirll down
        helper(nums, pos + 1,list,visited,result)
        
        // reverse the current status
        list = list[0:len(list)-1]
        visited[i] = false
    }
}

ans := make([]int, len(list))
copy(ans, list)
*result = append(*result, ans)
这里看到题解里面有人说，切片底层公用数据，所以要copy，也是有点懵逼的

总结：
1. 做过的题目还是不会，里面很多细节，包括变量的使用还是有点不清晰的感觉

2. 知道思路，还是不一定写的出来，说明还是不会，不熟练

3. 回溯，将当前层的效果去掉还是挺巧妙的
