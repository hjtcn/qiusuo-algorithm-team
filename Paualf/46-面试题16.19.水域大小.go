/*
你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
示例：
输入：
[
  [0,2,1,0],
  [0,1,0,1],
  [1,1,0,1],
  [0,1,0,1]
]
输出： [1,2,4]
提示：
• 0 < len(land) <= 1000
• 0 < len(land[i]) <= 1000
*/

1. Clearfiction:
思路：初始化池子数量为0，遍历二维数组，如果元素为0，则判断该元素上下左右斜对角线的元素是否为0，如果为0，则继续该操作，递归遍历该操作，直到元素周围没有为0的元素
伪代码：
# 八联通定义
(x,y)
index := [
   -1,0 // 左
   1,0  // 右
   0,-1 // 上
   0,1  // 下
   1,-1 // 左下
   1,1  // 右下
   -1,-1 // 左上
    -1,+1 // 右上
];
direction := [][]int{
    {-1,0}, // 左
    {1,0},  // 右
    {0,-1}, // 上
    {0,1},  // 下
    {1,-1}, // 左下
    {1,1},  // 右上
    {-1,-1},  // 左上
    {-1,1},   // 右上
}
func pondSizes(land [][]int)[]int {
    ret := 0
     
    for i := 0;i < len(land);i++ {
        for j := 0;j < len(land[i]);j++ {
            if land[i][j] == 0 {
                ret ++
                
            }   
        }
    }    
}
func dfs(land [][]int,i,j int) {
}

还是写了一点点
不会的地方在于，二维数组定义不熟练， 是 {}号
还有不知道如何递归继续去使用八联通图去遍历，代码写的及其不熟练
看了题解：

var direction = [][]int{
    {-1,0}, // 左
    {1,0},  // 右
    {0,-1}, // 上
    {0,1},  // 下
    {1,-1}, // 左下
    {1,1},  // 右上
    {-1,-1},  // 左上
    {-1,1},   // 右上
}
func pondSizes(land [][]int)[]int {
    ans := make([]int, 0)
    rows := len(land)
    if rows == 0 {
        return ans
    }
    cols := len(land[0])
    
    for x := 0;x < rows;x++ {
        for y := 0;y < cols;y++ {
            if land[x][y] == 0 {
                if cnt := dfs(land,x,y);cnt > 0 {
                    ans = append(ans,cnt)
                }
            }
        }
    }
    
    sort.Slice(ans, func(i,j int)bool {
        return ans[i] < ans[j]
    })
    
    return ans
}

func dfs(land [][]int,x,y int) int {
    cnt := 0
    if land[x][y] == 0 {
        cnt++
        land[x][y] = -1
        
        for i := 0;i < 8;i++ {
            xx,yy := x + direction[i][0],y + direction[i][1]
            
            if xx >= 0 && xx < len(land) && yy >= 0 && yy < len(land[0]) {
                cnt += dfs(land,xx,yy)
            }
        }
    }else {
        land[x][y] = -1
    }
    return cnt
}

看了题解中广度优先遍历的代码

type point struct {
    m,n int
}

func pondSizes(land [][]int)[]int {
        M,N := len(land),len(land[0])
    dy := [8][2]int{
        {-1,-1},
      {-1,0},
      {-1,1},
      {0,-1},
      {0,1},
      {1,-1},
      {1,0},
      {1,1},
    }
    szs := []int{}
    
    for m := 0;m < M;m++ {
          for n := 0;n < N;n++ {
                if land[m][n] == 0 {
                    land[m][n] = -1
                q := []point{}
                
                fptr := 0
                q = append(q, point{m,n})
                for fptr != len(q) {
                        i0,j0 := q[fptr].m,q[fptr].n
                    fptr++
                    
                    for x := 0;x < 8;x++ {
                          i,j := i0 + dy[x][0],j0 + dy[x][1]
                        
                        if i > -1 && i < M && j > -1 && j < N && land[i][j] == 0 {
                                land[i][j] = -1
                            q = append(q, point{i,j})
                        }
                    }
                }
                szs = append(szs,len(q))
            }
        }
    }
    sort.Ints(szs)
    return szs
}

BFS的时候是每次找到元素为0的元素就开始进行BFS,使用队列进行模拟，中间使用fptr变量标记是否终止

复杂度分析：
时间复杂度：O(n*n) ,矩阵中每个元素都要被访问
空间复杂度：O(n) 递归栈空间，方向数组，返回结果数组

总结:
1. 数组中的DFS和BFS感觉距离真的会用还差好多

2. 了解并查集的实现，不要急
