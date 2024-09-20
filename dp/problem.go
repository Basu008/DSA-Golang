package dp

import (
	"math"
)

func CreateDP(size int) []int {
	dp := make([]int, size)
	for i := range dp {
		dp[i] = -1
	}
	return dp
}

func CreateDp2d(m, n int) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return dp
}

func CreateDp2dBool(m, n int) [][]bool {
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	return dp
}

// Basic recursion
func Fibonacci(n int, res []int) int {
	if n <= 1 {
		return n
	}
	if res[n] != -1 {
		return res[n]
	}
	res[n] = Fibonacci(n-1, res) + Fibonacci(n-2, res)
	return res[n]
}

// Tabulation
func FibonacciDP2(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		if i == 0 || i == 1 {
			dp[i] = i
			continue
		}
		dp[i] = -1
	}
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Space optimization
func FibonacciDP3(n int) int {
	prev, prev2, curr := 1, 0, -1
	for i := 2; i <= n; i++ {
		curr = prev + prev2
		prev2 = prev
		prev = curr
	}
	return curr
}

func ClimbingStairs(n int) int {
	if n <= 2 {
		return n
	}
	left := ClimbingStairs(n - 1)
	right := ClimbingStairs(n - 2)
	return left + right
}

func FrogJump(n []int) int {
	prev := 0
	prev1 := 0
	for i := 1; i < len(n); i++ {
		fs := prev + int(math.Abs(float64(n[i]-n[i-1])))
		ss := math.MaxInt
		if i > 1 {
			ss = prev1 + int(math.Abs(float64(n[i]-n[i-2])))
		}
		curr := int(math.Min(float64(fs), float64(ss)))
		prev1 = prev
		prev = curr
	}
	return prev
}

func FrogJumpWithKSteps(n []int, k int) int {
	dp := make([]int, len(n))
	for i := range dp {
		if i == 0 || i == 1 {
			dp[i] = i
			continue
		}
		dp[i] = -1
	}
	for i := 1; i < len(n); i++ {
		minStep := math.MaxInt
		for j := 1; j <= k; j++ {
			if i-j >= 0 {
				jump := dp[i-j] + int(math.Abs(float64(n[i]-n[i-j])))
				minStep = int(math.Min(float64(minStep), float64(jump)))
			}
		}
		dp[i] = minStep
	}
	return dp[len(n)-1]
}

func NonAdjacentSum(dataset []int) int {
	prev := dataset[0]
	prev1 := 0
	for i := 1; i < len(dataset); i++ {
		curr := dataset[i] + prev1
		max := int(math.Max(float64(curr), float64(prev)))
		prev1 = prev
		prev = max
	}
	return prev
}

func HouseRobber(dataset []int) int {
	withFirst := NonAdjacentSum(dataset[:len(dataset)-1])
	withLast := NonAdjacentSum(dataset[1:])
	return int(math.Max(float64(withFirst), float64(withLast)))
}

func NinjaTraining(i, prev int, dataset [][]int) int {
	if i == 0 {
		maxSum := 0
		for j := range dataset[i] {
			if j != prev {
				maxSum = int(math.Max(float64(dataset[i][j]), float64(maxSum)))
			}
		}
		return maxSum
	}
	maxSum := math.MinInt
	for j := range dataset[i] {
		if j == prev {
			continue
		}
		curr := dataset[i][j] + NinjaTraining(i-1, j, dataset)
		maxSum = int(math.Max(float64(curr), float64(maxSum)))
	}
	return maxSum
}

// Memoization
func NinjaTraining2(i, prev int, dataset [][]int, dp [][]int) int {
	if i == 0 {
		maxSum := 0
		for j := range dataset[i] {
			if j != prev {
				maxSum = int(math.Max(float64(dataset[i][j]), float64(maxSum)))
			}
		}
		return maxSum
	}
	if dp[i][prev] != -1 {
		return dp[i][prev]
	}
	maxSum := math.MinInt
	for j := range dataset[i] {
		if j == prev {
			continue
		}
		curr := dataset[i][j] + NinjaTraining2(i-1, j, dataset, dp)
		maxSum = int(math.Max(float64(curr), float64(maxSum)))
	}
	dp[i][prev] = maxSum
	return dp[i][prev]
}

// Tabulation
func NinjaTraining3(dataset [][]int) int {
	dp := CreateDp2d(len(dataset), 4)
	dp[0][0] = int(math.Max(float64(dataset[0][1]), float64(dataset[0][2])))
	dp[0][1] = int(math.Max(float64(dataset[0][0]), float64(dataset[0][2])))
	dp[0][2] = int(math.Max(float64(dataset[0][1]), float64(dataset[0][0])))
	dp[0][3] = int(math.Max(float64(dataset[0][1]), math.Max(float64(dataset[0][2]), float64(dataset[0][0]))))
	for day := 1; day < len(dataset); day++ {
		for prev := 0; prev < 4; prev++ {
			dp[day][prev] = 0
			max := 0
			for task := 0; task < 3; task++ {
				if task == prev {
					continue
				}
				curr := dataset[day][task] + dp[day-1][task]
				max = int(math.Max(float64(curr), float64(max)))
			}
			dp[day][prev] = max
		}
	}
	return dp[len(dataset)-1][3]
}

func NinjaTraining4(dataset [][]int) int {
	dp := CreateDP(4)
	dp[0] = int(math.Max(float64(dataset[0][1]), float64(dataset[0][2])))
	dp[1] = int(math.Max(float64(dataset[0][0]), float64(dataset[0][2])))
	dp[2] = int(math.Max(float64(dataset[0][1]), float64(dataset[0][0])))
	dp[3] = int(math.Max(float64(dataset[0][1]), math.Max(float64(dataset[0][2]), float64(dataset[0][0]))))
	for day := 1; day < len(dataset); day++ {
		temp := make([]int, 4)
		for prev := 0; prev < 4; prev++ {
			max := 0
			for task := 0; task < 3; task++ {
				if task == prev {
					continue
				}
				curr := dataset[day][task] + dp[day-1]
				max = int(math.Max(float64(curr), float64(max)))
			}
			temp[day] = max
		}
		dp = temp
	}
	return dp[len(dataset)-1]
}

// Recursion
func GridUniquePaths(i, j int) int {
	if i == 0 && j == 0 {
		return 1
	}
	var left, top int
	if j > 0 {
		left = GridUniquePaths(i, j-1)
	}
	if i > 0 {
		top = GridUniquePaths(i-1, j)
	}
	return top + left
}

// Memoization
func GridUniquePaths2(i, j int, dp [][]int) int {
	if i == 0 && j == 0 {
		return 1
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	var left, top int
	if j > 0 {
		left = GridUniquePaths2(i, j-1, dp)
	}
	if i > 0 {
		top = GridUniquePaths2(i-1, j, dp)
	}
	dp[i][j] = top + left
	return dp[i][j]
}

// Tabulation
func GridUniquePaths3(m, n int) int {
	dp := CreateDp2d(m, n)
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			var left, top int
			if j > 0 {
				left = dp[i][j-1]
			}
			if i > 0 {
				top = dp[i-1][j]
			}
			dp[i][j] = left + top
		}
	}
	return dp[m-1][n-1]
}

// Tabulation
func GridUniquePathsWithObstacles(course [][]int) int {
	m, n := len(course), len(course[0])
	dp := CreateDp2d(m, n)
	if course[0][0] == -1 {
		return -1
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if course[i][j] == -1 {
				continue
			}
			var left, top int
			if j > 0 {
				left = dp[i][j-1]
				if left == -1 {
					left = 0
				}
			}
			if i > 0 {
				top = dp[i-1][j]
				if top == -1 {
					top = 0
				}
			}
			dp[i][j] = left + top
		}
	}
	return dp[m-1][n-1]
}

// Recursion
func GridMinPathSum(i, j int, grid [][]int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	leftPathWeight, topPathWeight := math.MaxInt, math.MaxInt
	if i > 0 {
		topPathWeight = GridMinPathSum(i-1, j, grid)
	}
	if j > 0 {
		leftPathWeight = GridMinPathSum(i, j-1, grid)
	}
	return grid[i][j] + int(math.Min(float64(topPathWeight), float64(leftPathWeight)))
}

// Memoisation
func GridMinPathSum2(i, j int, grid, dp [][]int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	leftPathWeight, topPathWeight := math.MaxInt, math.MaxInt
	if i > 0 {
		topPathWeight = GridMinPathSum2(i-1, j, grid, dp)
	}
	if j > 0 {
		leftPathWeight = GridMinPathSum2(i, j-1, grid, dp)
	}
	dp[i][j] = grid[i][j] + int(math.Min(float64(topPathWeight), float64(leftPathWeight)))
	return dp[i][j]
}

// Tabulation
func GridMinPathSum3(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := CreateDp2d(len(grid), len(grid[0]))
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			leftPathWeight, topPathWeight := math.MaxInt, math.MaxInt
			if i > 0 {
				topPathWeight = dp[i-1][j]
			}
			if j > 0 {
				leftPathWeight = dp[i][j-1]
			}
			dp[i][j] = grid[i][j] + int(math.Min(float64(topPathWeight), float64(leftPathWeight)))
		}
	}
	return dp[m-1][n-1]
}

// Recursion
func Triangle(i, j int, dataset [][]int) int {
	if i == len(dataset)-1 {
		return dataset[i][j]
	}
	bottom := Triangle(i+1, j, dataset)
	right := Triangle(i+1, j+1, dataset)
	return dataset[i][j] + int(math.Min(float64(bottom), float64(right)))
}

// Memoization
func Triangle2(i, j int, dataset, dp [][]int) int {
	if i == len(dataset)-1 {
		return dataset[i][j]
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	bottom := Triangle(i+1, j, dataset)
	right := Triangle(i+1, j+1, dataset)
	dp[i][j] = dataset[i][j] + int(math.Min(float64(bottom), float64(right)))
	return dp[i][j]
}

// Tabulation
func Triangle3(dataset [][]int, m, n int) int {
	dp := CreateDp2d(m, n)
	dp[0][0] = dataset[0][0]
	min := math.MaxInt
	for i := 1; i < len(dataset); i++ {
		for j := 0; j <= i; j++ {
			top, left := math.MaxInt, math.MaxInt
			if i != j {
				top = dp[i-1][j]
			}
			if j > 0 {
				left = dp[i-1][j-1]
			}
			dp[i][j] = dataset[i][j] + int(math.Min(float64(top), float64(left)))
			if i == len(dataset)-1 {
				if dp[i][j] < min {
					min = dp[i][j]
				}
			}
		}
	}
	return min
}

// Recursion
func MaxFallingPathSum(dataset [][]int) int {
	max := math.MinInt
	for j := 0; j < len(dataset[0]); j++ {
		path := maxFallingPathSumHelper(len(dataset)-1, j, dataset)
		if path > max {
			max = path
		}
	}
	return max
}

func maxFallingPathSumHelper(i, j int, dataset [][]int) int {
	if i == 0 {
		return dataset[i][j]
	}
	var left, right = math.MaxInt, math.MaxInt
	top := maxFallingPathSumHelper(i-1, j, dataset)
	if j > 0 {
		left = maxFallingPathSumHelper(i-1, j-1, dataset)
	}
	if j < len(dataset[0])-1 {
		right = maxFallingPathSumHelper(i-1, j+1, dataset)
	}
	return dataset[i][j] + int(math.Max(math.Max(float64(left), float64(right)), float64(top)))
}

// Memoisation
func MaxFallingPathSum2(dataset [][]int) int {
	m, n := len(dataset), len(dataset[0])
	dp := CreateDp2d(m, n)
	max := math.MinInt
	for j := 0; j < n; j++ {
		path := maxFallingPathSum2Helper(m-1, j, dataset, dp)
		if path > max {
			max = path
		}
	}
	return max
}

func maxFallingPathSum2Helper(i, j int, dataset, dp [][]int) int {
	if i == 0 {
		return dataset[i][j]
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	var left, right = math.MaxInt, math.MaxInt
	top := maxFallingPathSum2Helper(i-1, j, dataset, dp)
	if j > 0 {
		left = maxFallingPathSum2Helper(i-1, j-1, dataset, dp)
	}
	if j < len(dataset[0])-1 {
		right = maxFallingPathSum2Helper(i-1, j+1, dataset, dp)
	}
	dp[i][j] = dataset[i][j] + int(math.Max(math.Max(float64(left), float64(right)), float64(top)))
	return dp[i][j]
}

// Tabulation
func MaxFallingPathSum3(dataset [][]int) int {
	m, n := len(dataset), len(dataset[0])
	dp := CreateDp2d(m, n)
	for j := 0; j < n; j++ {
		dp[0][j] = dataset[0][j]
	}
	max := math.MinInt
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			var left, right = math.MaxInt, math.MaxInt
			top := dp[i-1][j]
			if j > 0 {
				left = dp[i-1][j-1]
			}
			if j < n-1 {
				right = dp[i-1][j+1]
			}
			dp[i][j] = dataset[i][j] + int(math.Max(math.Max(float64(left), float64(right)), float64(top)))
			if i == m-1 {
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max
}

// Recursion
func CherryPicking(i, j1, j2 int, dataset [][]int) int {
	if i < 0 || j1 < 0 || j2 < 0 {
		return 0
	}
	if i == len(dataset)-1 {
		if j1 == j2 {
			return dataset[i][j1]
		}
		return dataset[i][j1] + dataset[i][j2]
	}
	possibleMoves := []int{-1, 0, 1}
	var max int
	for _, move1 := range possibleMoves {
		for _, move2 := range possibleMoves {
			if j1 == j2 {
				max = int(math.Max(float64(max), float64(dataset[i][j1]+CherryPicking(i+1, j1+move1, j2+move2, dataset))))
			} else {
				max = int(math.Max(float64(max), float64(dataset[i][j1]+dataset[i][j2]+CherryPicking(i+1, j1+move1, j2+move2, dataset))))
			}
		}
	}
	return max
}

// Memoisation
func CherryPicking2(i, j1, j2 int, dataset [][]int, dp [][][]int) int {
	if i < 0 || j1 < 0 || j2 < 0 {
		return 0
	}
	if dp[i][j1][j2] != -1 {
		return dp[i][j1][j2]
	}
	if i == len(dataset)-1 {
		if j1 == j2 {
			return dataset[i][j1]
		}
		return dataset[i][j1] + dataset[i][j2]
	}
	possibleMoves := []int{-1, 0, 1}
	var max int
	for _, move1 := range possibleMoves {
		for _, move2 := range possibleMoves {
			if j1 == j2 {
				max = int(math.Max(float64(max), float64(dataset[i][j1]+CherryPicking2(i+1, j1+move1, j2+move2, dataset, dp))))
			} else {
				max = int(math.Max(float64(max), float64(dataset[i][j1]+dataset[i][j2]+CherryPicking2(i+1, j1+move1, j2+move2, dataset, dp))))
			}
		}
	}
	dp[i][j1][j2] = max
	return dp[i][j1][j2]
}

// Recursion
func SubsetToK(i, target int, dataset []int) bool {
	if target == 0 {
		return true
	}
	if i == 0 {
		return target == dataset[i]
	}
	notTaken := SubsetToK(i-1, target, dataset)
	if notTaken {
		return true
	}
	var taken bool
	if target >= dataset[i] {
		taken = SubsetToK(i-1, target-dataset[i], dataset)
	}
	return taken || notTaken
}

// Memoisation
func SubsetToK2(i, target int, dataset []int, dp [][]int) int {
	if target == 0 {
		return 1
	}
	if i == 0 {
		if target == dataset[i] {
			return 1
		}
		return -1
	}
	if dp[i][target] != -1 {
		return 1
	}
	notTaken := SubsetToK(i-1, target, dataset)
	if notTaken {
		return 1
	}
	var taken bool
	if target >= dataset[i] {
		taken = SubsetToK(i-1, target-dataset[i], dataset)
	}
	if taken || notTaken {
		dp[i][target] = 1
	}
	return dp[i][target]
}

func SubsetToK3(dataset []int, target int) bool {
	dp := CreateDp2dBool(len(dataset), target)
	for i := range dp {
		dp[i][0] = true
	}
	dp[0][dataset[0]] = true
	for i := 1; i < len(dataset); i++ {
		for t := 1; t <= target; t++ {
			notTaken := dp[i-1][t]
			var taken bool
			if dataset[i] >= t {
				taken = dp[i-1][t-dataset[i]]
			}
			dp[i][t] = taken || notTaken
		}
	}
	return dp[len(dataset)-1][target]
}

func Partition(dataset []int) bool {
	var sum int
	for _, data := range dataset {
		sum += data
	}
	if sum%2 == 1 {
		return false
	}
	return SubsetToK3(dataset, sum/2)
}

func CountSubset(dataset []int) int {
	//same as subest wala problem
	return -1
}
