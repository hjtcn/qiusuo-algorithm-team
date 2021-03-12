/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
 
示例 1：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true
示例 2：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false
 
提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/

1. Clearfication:
先判断是否在一行里面，如果在一行里面的话，然后在这一行里面去进行二分查找
伪代码：
m := len(matrix)
n := len(matrix[0])
for i := 0;i < m;i++ {
if target >= matrix[0][0] && target <=  matrix[0][n-1] {
return binary_search(matrix[i],target)
}
}
func binary_search(nums []int, target) bool {
left,right := 0,len(nums) - 1
for left <= right {
mid :=  left + (right - left) / 2
if nums[mid] == target {
return true
}else if nums[mid] > target {   
right = mid
}else if nums[mid] < target {
left = mid
}
}
return false
}

2. Coding:
func searchMatrix(matrix [][]int, target int) bool {
    m,n := len(matrix),len(matrix[0])
    if m == 0 || n == 0 {
        return false
    }
    
    for i := 0;i < m;i++ {
        if target >= matrix[i][0] && target <= matrix[i][n-1] {
            return binary_search(matrix[i], target)
        }
    }
    
    return false
}

func binary_search(nums []int, target int) bool {
    left,right := 0,len(nums) - 1
    
    for left <= right {
        mid :=  left + (right - left) / 2
        if nums[mid] == target {
             return true
        }else if nums[mid] > target {   
             right = mid - 1
        }else if nums[mid] < target {
             left = mid + 1
        }
    }
    
    return false
}

3. 看题解：
将二维数组根据升序数据结构抽象成一维的数据结构, 主要是抽象后的下标的获取
func searchMatrix(matrix [][]int, target int) bool {
    m,n := len(matrix),len(matrix[0])
    
    if m == 0 || n == 0 {
        return false
    }
    
    left,right := 0,m * n - 1 
    
    for left <= right {
        pivotIdx := left + (right - left) / 2
        pivotElement := matrix[pivotIdx / n][pivotIdx % n]
        
        if target == pivotElement {
            return true
        }else if target > pivotElement {
            right = mid - 1
        }else if target < pivotElement {
            left = mid + 1
        }
    }
    
    return false
}

4. 复杂度分析
时间复杂度：O(logN)
空间复杂度：O(1)

5. 总结：
5.1: 将二维数组抽象成一维数组，然后计算的时候再转换为一维数组，还是蛮巧妙的
