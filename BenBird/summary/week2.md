## 第二周总结

#### 题目列表

* First day: [[1535]找出数组游戏的赢家](https://leetcode-cn.com/problems/find-the-winner-of-an-array-game)

* The next day: [[1438]绝对差不超过限制的最长连续子数组](https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit)

* Third day: [[1267]统计参与通信的服务器](https://leetcode-cn.com/problems/count-servers-that-communicate)

* Fourth day: [[840]矩阵中的幻方](https://leetcode-cn.com/problems/magic-squares-in-grid)

* Fifth day: [[969]煎饼排序](https://leetcode-cn.com/problems/pancake-sorting/)

* Sixth day: [[1424]对角线遍历](https://leetcode-cn.com/problems/diagonal-traverse-ii)

#### 我的题解

1. [找出数组游戏的赢家](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/8-1535%E6%89%BE%E5%87%BA%E6%95%B0%E7%BB%84%E6%B8%B8%E6%88%8F%E7%9A%84%E8%B5%A2%E5%AE%B6.go)
    
    - 解题思路：
    
          这道题把题分析清楚，就很容易做了，首先把最大数存入变量max_num，后面在遍历arr时，只需更新最大值的变量即可，无需移动元素位置
          遍历过程中，如果max_num大于当前值
        	则获胜次数times 自增1，并判断获胜次数是否等于k
        		若等于k，则直接返回max_num
          若max_num小于当前值
        	max_num 更改为当前值，获胜次置为1
        
          若循环结束(即times < k)，还未有返回值，则获胜的元素为max_num 
    
2. [绝对差不超过限制的最长连续子数组](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/9-1438%E7%BB%9D%E5%AF%B9%E5%B7%AE%E4%B8%8D%E8%B6%85%E8%BF%87%E9%99%90%E5%88%B6%E7%9A%84%E6%9C%80%E9%95%BF%E8%BF%9E%E7%BB%AD%E5%AD%90%E6%95%B0%E7%BB%84.go)
    
    -  滑动窗口：
            
            维护滑动窗口的最大值和最小值
            只要保证最大值和最小值之差不大于limit即可
            
            遍历nums时不断更新最大值和最小值
            并判断最大值和最小值之差
            若之差 <= limit
            	则更新当前当前滑动窗口的长度
            若之差 > limit
            	则更新最大值、最小值、最长连续子数组的长度
           
3. [统计参与通信的服务器](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/10-1267%E7%BB%9F%E8%AE%A1%E5%8F%82%E4%B8%8E%E9%80%9A%E4%BF%A1%E7%9A%84%E6%9C%8D%E5%8A%A1%E5%99%A8.go)

    - 暴力求解：
    
           要保证某个节点上的服务器能与其他服务器通信的条件：为这台服务器的所在节点的 行服务器数 > 1 或 列服务器数 > 1即可
           
           若节点(r, w)存在服务器，则在row_arr[r]++, col_arr[c]++, 然后把grid每行每列遍历一下
           
           然后在重新把grid遍历一遍，当查到一台服务器(r, w)，判断该服务器的行或列是否存在其他的服务器(即：row_arr[r] > 1 或 col_arr[c] > 1)，
           	若成立则，能与其他服务器通信的服务器数量 +1
           	若不成立则，继续向下遍历

4. [矩阵中的幻方](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/11-840%E7%9F%A9%E9%98%B5%E4%B8%AD%E7%9A%84%E5%B9%BB%E6%96%B9.go)

    - 枚举求解:
    
           要想满足每一行、每一列、两个对角线的和相等，则和都等于15，且中心数字必是5
           再则就是这些满足条件的幻方，其实都可以有一种幻方进行旋转可得到，
           所以只记录满足条件的顺时针数据为 81672943
           
           所以首先循环遍历grid，查找中心元素为5的坐标，
           	若发现存在中心元素为5的3x3矩阵，则获取改矩阵外围数时针数字，并以此存入切片around
           
           将获取3x3矩阵和标准幻方进行对比
           	若和标准幻方相等，则count++
           	若无幻方与其相等，则继续遍历下一个中心元素

5. [煎饼排序](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/12-969%E7%85%8E%E9%A5%BC%E6%8E%92%E5%BA%8F.go)

    - 递归求解：
    
          首先找到需要排序元素的最大值，并获取其下标
          然后将首位元素到最大值中间的所有元素进行反转一次，将最大值置于第一位
          需要排序元素个数为n，将前n个元素进行反转一次，此时最大值就在它应该在位置上了
          
          按照此思路依次循环往复，直到需要排序元素个数为1，即排序完成

6. [对角线遍历](https://github.com/hjtcn/qiusuo-algorithm-team/blob/master/BenBird/13-1412%E5%AF%B9%E8%A7%92%E7%BA%BF%E9%81%8D%E5%8E%86.go)

    - 解题思路：
    
          通过规律可发现每个对角线上纵横坐标和是相等的
          所以在遍历nums时，将nums[i][j] 存入map m[i+j][] 中
          
          因为Go没有自带的在数组/切片/map 头部插入元素的内置函数
          所以为解决对m 展开遍历获取返回值中对角线元素排序的问题，在遍历nums时，是从最后一个元素开始向第一个元素进行遍历
          
          所以，最后直接对m 进行遍历即可获取正确的返回值

#### 总结

本周六道题难度皆为中等，刷题稍微有点压力，读完题，大都没有思路，大部分题是看题解的，不过有些题解也很难看懂，不过，还是坚持了下来
Go作为我刷题的语言，稍微有点吃力，好多方法都需要自己实现，加油

坚持坚持再坚持

