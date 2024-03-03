package graph

import (
	"fmt"
	"math"
	"sort"

	"github.com/Basu008/DSA-Golang/queue"
	"github.com/Basu008/DSA-Golang/stack"
)

func (g Graph) CountProvinces() (provinces int) {
	visited := make([]int, len(g))
	for i := 1; i < len(g); i++ {
		if visited[i] == 0 {
			provinces++
			g.BFS(i, visited)
		}
	}
	return
}

func CountNoOfIslands(earth [][]int, m, n int) (island int) {
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	q := queue.Queue{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == 0 && earth[i][j] == 1 {
				q.Push([]int{i, j})
				visited[i][j] = 1
				for !q.IsEmpty() {
					pair := q.Pop().([]int)
					for r := pair[0] - 1; r <= pair[0]+1; r++ {
						for c := pair[1] - 1; c <= pair[1]+1; c++ {
							if (r >= 0 && c >= 0) && (r < m && c < n) {
								if earth[r][c] == 1 && visited[r][c] != 1 {
									q.Push([]int{r, c})
									visited[r][c] = 1
								}
							}
						}
					}
				}
				island++
			}
		}
	}
	return
}

func FloodFill(image [][]int, sn, sc, nc int) {
	q := queue.Queue{}
	q.Push([]int{sn, sc})
	initialColor := image[sn][sc]
	image[sn][sc] = nc
	for !q.IsEmpty() {
		pair := q.Pop().([]int)
		for r := pair[0] - 1; r <= pair[0]+1; r++ {
			for c := pair[1] - 1; c <= pair[1]+1; c++ {
				if (r >= 0 && c >= 0) && (r < len(image) && c < len(image[0])) && int(math.Abs(float64(r)-float64(c))) <= 1 {
					if image[r][c] == initialColor {
						q.Push([]int{r, c})
						image[r][c] = nc
					}
				}
			}
		}
	}
	fmt.Println(image)
}

func RottenOranges(basket [][]int, m, n int) (time int) {
	q := queue.Queue{}
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	//Enter rotten one in queue and mark them as visited
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if basket[i][j] == 2 && visited[i][j] != 2 {
				visited[i][j] = 2
				q.Push([]int{i, j, 0})
			}
		}
	}
	//Now we start rotting the array
	for !q.IsEmpty() {
		top := q.Pop().([]int)
		i, j, timeTillNow := top[0], top[1], top[2]
		time = int(math.Max(float64(timeTillNow), float64(time)))
		if visited[i][j] == 2 {
			row := []int{-1, 0, 1, 0}
			col := []int{0, 1, 0, -1}
			for k := 0; k < len(row); k++ {
				r := i + row[k]
				c := j + col[k]
				if (r >= 0 && c >= 0) && (r < m && c < n) && basket[r][c] == 1 && visited[r][c] != 2 {
					visited[r][c] = 2
					q.Push([]int{r, c, timeTillNow + 1})
				}
			}
		}
	}
	//Now to check if all the oranges are rotten
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if basket[i][j] == 1 && visited[i][j] != 2 {
				return -1
			}
		}
	}
	return
}

func (g Graph) CycleBFS() bool {
	//Only for connected graph
	visited := make([]bool, len(g))
	q := queue.Queue{}
	for i := 1; i < len(g); i++ {
		if visited[i] {
			continue
		}
		q.Push([]int{i, -1})
		visited[i] = true
		for !q.IsEmpty() {
			top := q.Pop().([]int)
			node := top[0]
			parent := top[1]
			for _, adjNode := range g[node] {
				if !visited[adjNode.Value] {
					visited[adjNode.Value] = true
					q.Push([]int{adjNode.Value, node})
				} else if parent != adjNode.Value {
					return true
				}
			}
		}
	}
	return false
}

func (g Graph) CycleDFS() bool {
	return g.helperCyclerDFS(1, -1, make([]bool, len(g)))
}

func (g Graph) helperCyclerDFS(src, parent int, visited []bool) bool {
	visited[src] = true
	for _, node := range g[src] {
		if !visited[node.Value] {
			visited[node.Value] = true
			return g.helperCyclerDFS(node.Value, src, visited)
		} else if parent != node.Value {
			return true
		}
	}
	return false
}

func Nearest1(matrix [][]int, m, n int) [][]int {
	q := queue.Queue{}
	visited := make([][]int, m)
	res := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
		res[i] = make([]int, n)
	}
	//Set first set of 1s
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				q.Push([]int{i, j, 0})
				visited[i][j] = 1
			}
		}
	}
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}
	for !q.IsEmpty() {
		top := q.Pop().([]int)
		i, j, distance := top[0], top[1], top[2]
		for k := 0; k < 4; k++ {
			r := i + row[k]
			c := j + col[k]
			if (r >= 0 && c >= 0) && (r < m && c < n) && visited[r][c] != 1 {
				q.Push([]int{r, c, distance + 1})
				res[r][c] = distance + 1
				visited[r][c] = 1
			}
		}
	}
	return res
}

func Convert0s(matrix [][]int, m, n int) {
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 && visited[0][j] == 0 {
			DFSForMatrix(0, j, matrix, visited)
		}
		if matrix[m-1][j] == 0 && visited[m-1][j] == 0 {
			DFSForMatrix(m-1, j, matrix, visited)
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 && visited[i][0] == 0 {
			DFSForMatrix(i, 0, matrix, visited)
		}
		if matrix[i][n-1] == 0 && visited[i][n-1] == 0 {
			DFSForMatrix(i, n-1, matrix, visited)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 && visited[i][j] == 0 {
				matrix[i][j] = 1
			}
		}
	}
}

func NoOfEnclaves(matrix [][]int, m, n int) int {
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	//To check
	for i := 0; i < m; i++ {
		if matrix[i][0] == 1 && visited[i][0] == 0 {
			DFSForMatrix(i, 0, matrix, visited)
		}
		if matrix[i][m-1] == 1 && visited[i][m-1] == 0 {
			DFSForMatrix(i, m-1, matrix, visited)
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 1 && visited[0][j] == 0 {
			DFSForMatrix(0, j, matrix, visited)
		}
		if matrix[n-1][j] == 1 && visited[n-1][j] == 0 {
			DFSForMatrix(n-1, j, matrix, visited)
		}
	}
	var totalCount int
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if matrix[i][j] == 1 && visited[i][j] == 0 {
				totalCount++
			}
		}
	}
	return totalCount
}

func NoOfDistinctIslands(matrix [][]int, m, n int) {
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}
	set := [][][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 && visited[i][j] == 0 {
				points := DFSToCountIslands(i, j, matrix, visited, [][]int{}, &i, &j)
				if len(set) == 0 {
					set = append(set, points)
				} else {
					isValid := false
					for i := range set {
						if !arePointsValid(set[i], points) {
							break
						}
					}
					if isValid {
						set = append(set, points)
					}
				}
			}
		}
	}
	fmt.Println("asDFASDF")
}

func arePointsValid(existing, current [][]int) bool {
	if len(existing) != len(current) {
		return true
	}
	for i := range existing {
		if existing[i][0] == current[i][0] && existing[i][1] == current[i][1] {
			return false
		}
	}
	return true
}

type coloredNode struct {
	node  int
	color int
}

func (g Graph) BipartiteGraph() bool {
	q := queue.Queue{}
	visited := make([]int, len(g))
	for i := range visited {
		visited[i] = -1
	}
	q.Push(coloredNode{1, 0})
	visited[1] = 0
	for !q.IsEmpty() {
		top := q.Pop().(coloredNode)
		color := 0
		if top.color == 0 {
			color = 1
		}
		for _, node := range g[top.node] {
			if visited[node.Value] == top.color {
				return false
			}
			if visited[node.Value] == -1 {
				q.Push(coloredNode{node: node.Value, color: color})
				visited[node.Value] = color
			}
		}
	}
	return true
}

func (g Graph) IsDirectedGraphCyclicBFS() bool {
	indegrees := make([]int, len(g))
	queue := []int{}
	topSort := []int{}
	for i := 1; i < len(g); i++ {
		for _, node := range g[i] {
			indegrees[node.Value]++
		}
	}
	for index, indegree := range indegrees {
		if indegree == 0 && index != 0 {
			queue = append(queue, index)
		}
	}
	for len(queue) != 0 {
		top := queue[0]
		topSort = append(topSort, top)
		queue = queue[1:]
		for _, node := range g[top] {
			indegrees[node.Value]--
			if indegrees[node.Value] == 0 {
				queue = append(queue, node.Value)
			}
		}
	}
	return len(topSort) < len(g)-1
}

func (g Graph) IsDirectedGraphCyclic() bool {
	isCyclic := false
	visited := make([]int, len(g))
	for i := 1; i < len(g); i++ {
		isCyclic = g.dfsForCyclic(i, visited)
		if isCyclic {
			return true
		}
	}
	return isCyclic
}

func (g Graph) dfsForCyclic(src int, visited []int) bool {
	visited[src] = 2
	isCyclic := false
	for _, node := range g[src] {
		if visited[node.Value] == 2 {
			return true
		}
		if visited[node.Value] == 0 {
			isCyclic = g.dfsForCyclic(node.Value, visited)
		}
	}
	visited[src]--
	return isCyclic
}

func (g Graph) AllSafeNodes() []int {
	visited := make([]int, len(g))
	safeNodes := make([]int, len(g))
	index := len(g) - 1
	for i := 0; i < len(g); i++ {
		if g[i] == nil || visited[i] != 0 {
			continue
		}
		g.safeNodeDFS(i, visited, safeNodes, &index)
	}
	return safeNodes[index+1:]
}

func (g Graph) safeNodeDFS(src int, visited, safeNodes []int, index *int) bool {
	visited[src] = 2
	isSafeNode := true
	for _, node := range g[src] {
		if visited[node.Value] == 2 {
			return false
		}
		if visited[node.Value] == 0 {
			isSafeNode = g.safeNodeDFS(node.Value, visited, safeNodes, index)
		}
		if !isSafeNode {
			return false
		}
	}
	visited[src]--
	safeNodes[*index] = src
	*index--
	return true
}

func (g Graph) SafeNodesBFS() {
	adjRev := make([][]int, len(g)-1)
	indegrees := make([]int, len(g)-1)
	safeEdges := []int{}
	for a := range adjRev {
		adjRev[a] = []int{}
	}
	for i := 0; i < len(g)-1; i++ {
		for _, node := range g[i] {
			adjRev[node.Value] = append(adjRev[node.Value], i)
			indegrees[i]++
		}
	}
	q := queue.Queue{}
	for i, indegree := range indegrees {
		if indegree == 0 {
			q.Push(i)
		}
	}
	for !q.IsEmpty() {
		top := q.Pop().(int)
		safeEdges = append(safeEdges, top)
		for _, node := range adjRev[top] {
			indegrees[node]--
			if indegrees[node] == 0 {
				q.Push(node)
			}
		}
	}
	sort.IntSlice.Sort(safeEdges)
	for _, edge := range safeEdges {
		fmt.Printf("%d ", edge)
	}
}

func (g Graph) TopologicalSort() stack.Stack {
	visited := make([]int, len(g))
	s := stack.Stack{}
	for i := 0; i < len(g)-1; i++ {
		if visited[i] == 0 {
			g.dfsForSort(i, visited, &s)
		}
	}
	return s
}

func (g Graph) dfsForSort(src int, visited []int, stack *stack.Stack) {
	visited[src] = 1
	for _, node := range g[src] {
		if visited[node.Value] == 0 {
			g.dfsForSort(node.Value, visited, stack)
		}
	}
	stack.Push(src)
}

func (g Graph) TopologicalSortBFS() []int {
	indegrees := make([]int, len(g))
	ans := []int{}
	q := queue.Queue{}
	for i := 1; i < len(g); i++ {
		for _, node := range g[i] {
			indegrees[node.Value]++
		}
	}
	for node, indegree := range indegrees {
		if indegree == 0 {
			q.Push(node)
		}
	}
	for !q.IsEmpty() {
		top := q.Pop().(int)
		ans = append(ans, top)
		for _, node := range g[top] {
			indegrees[node.Value]--
			if indegrees[node.Value] == 0 {
				q.Push(node.Value)
			}
		}
	}
	return ans[1:]
}

func AlienDictionary(dict []string, k int) []string {
	adj := make([][]Pair, k+1)
	for i := 1; i <= k; i++ {
		adj[i] = []Pair{}
	}
	intToLetterMap := make(map[int]string)
	letterToIntMap := getMapOfAlphabets(k)
	for i := 0; i < k-1; i++ {
		str1 := dict[i]
		str2 := dict[i+1]
		size := int(math.Min(float64(len(str1)), float64(len(str2))))
		for j := 0; j < size; j++ {
			if str1[j] != str2[j] {
				u := letterToIntMap[string(str1[j])]
				intToLetterMap[u] = string(str1[j])
				v := letterToIntMap[string(str2[j])]
				intToLetterMap[v] = string(str2[j])
				pair := Pair{v, 0}
				adj[u] = append(adj[letterToIntMap[string(str1[j])]], pair)
				break
			}
		}
	}
	pattern := []string{}
	var g Graph = adj
	asd := g.TopologicalSortBFS()
	for _, node := range asd {
		pattern = append(pattern, intToLetterMap[node])
	}
	return pattern
}

func getMapOfAlphabets(k int) map[string]int {
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "x", "y", "z"}
	data := make(map[string]int)
	for i := 1; i <= k; i++ {
		data[alphabets[i-1]] = i
	}
	return data
}

func (g Graph) ShortestPathInDAG() []int {
	shortestPaths := make([]int, len(g)-1)
	for i := 1; i < len(shortestPaths); i++ {
		shortestPaths[i] = math.MaxInt
	}
	topoSort := g.TopologicalSort()
	var dist int
	for !topoSort.IsEmpty() {
		top := topoSort.Pop().(int)
		dist = shortestPaths[top]
		for _, node := range g[top] {
			weight := node.Weight + dist
			shortestPaths[node.Value] = int(math.Min(float64(shortestPaths[node.Value]), float64(weight)))
		}

	}
	return shortestPaths
}

func (g Graph) ShortesPathInUAG() []int {
	shortestPaths := make([]int, len(g)-1)
	for i := 1; i < len(shortestPaths); i++ {
		shortestPaths[i] = math.MaxInt
	}
	q := queue.Queue{}
	q.Push([]int{0, 0})
	for !q.IsEmpty() {
		top := q.Pop().([]int)
		for _, node := range g[top[0]] {
			dist := 1 + top[1]
			distSaved := shortestPaths[node.Value]
			if dist < distSaved {
				shortestPaths[node.Value] = dist
				q.Push([]int{node.Value, dist})
			}
		}
	}
	return shortestPaths
}

type wordWithLevel struct {
	word  string
	level int
}

func WordLadder1(startWord, endWord string, words []string) int {
	wordMap, selectedWordMap := make(map[string]bool), make(map[string]bool)
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "x", "y", "z"}
	for _, word := range words {
		wordMap[word] = true
	}
	q := queue.Queue{}
	q.Push(wordWithLevel{startWord, 1})
	for !q.IsEmpty() {
		top := q.Pop().(wordWithLevel)
		word := top.word
		level := top.level
		if word == endWord {
			return level
		}
		for i := range word {
			for _, alphabet := range alphabets {
				newWord := word[0:i] + alphabet + word[i+1:]
				if wordMap[newWord] && newWord != word && !selectedWordMap[newWord] {
					delete(wordMap, newWord)
					selectedWordMap[newWord] = true
					q.Push(wordWithLevel{newWord, level + 1})
				}
			}
		}
	}
	return 0
}

type comboWithLevel struct {
	combo []string
	level int
}

func WordLadder2(startWord, endWord string, words []string) [][]string {
	wordMap, _ := make(map[string]bool), make(map[string]bool)
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "x", "y", "z"}
	for _, word := range words {
		wordMap[word] = true
	}
	q := queue.Queue{}
	q.Push(comboWithLevel{[]string{startWord}, 1})
	wordsToBeDeleted := []string{}
	acceptedAns := [][]string{}
	level := 1
	shortestLength := math.MaxInt
	for !q.IsEmpty() {
		top := q.Pop().(comboWithLevel)
		if len(top.combo) > shortestLength {
			break
		}
		if level != top.level {
			for _, wordToBeDeleted := range wordsToBeDeleted {
				delete(wordMap, wordToBeDeleted)
			}
			wordsToBeDeleted = []string{}
			level = top.level
		}
		word := top.combo[len(top.combo)-1]
		if word == endWord {
			shortestLength = len(top.combo)
			acceptedAns = append(acceptedAns, top.combo)
			continue
		}
		acceptedWords := []string{}
		for i := range word {
			for _, alphabet := range alphabets {
				newWord := word[0:i] + alphabet + word[i+1:]
				if wordMap[newWord] && newWord != startWord {
					acceptedWords = append(acceptedWords, newWord)
					wordsToBeDeleted = append(wordsToBeDeleted, newWord)
				}
			}
		}
		for _, acceptedWord := range acceptedWords {
			combo := append(top.combo, acceptedWord)
			q.Push(comboWithLevel{combo: combo, level: top.level + 1})
		}
	}
	return acceptedAns
}

func (g Graph) Dijsktra(src int) {
	//Can't be done without set/PQ and I'm not building basic DS again and again. Maa ki chut graph ki
	//So instead I will use Queue and sort it at every stage
	q := queue.Queue{}
	dis := make([]int, len(g)-1)
	for i := range dis {
		if i == src {
			dis[src] = 0
			continue
		}
		dis[i] = math.MaxInt
	}
	for !q.IsEmpty() {
		top := q.Pop().([]int)
		dist, node := top[0], top[1]
		for _, adjNode := range g[node] {
			if dist+adjNode.Weight < dis[node] {
				dis[node] = dist + adjNode.Weight
				q.Push([]int{dist + adjNode.Weight, adjNode.Value})
			}
		}
	}
}

func (g Graph) ShortesPath(src, destination int) []int {
	q := queue.Queue{}
	dis := make([]int, len(g)-1)
	parent := make([]int, len(g)-1)
	for i := range dis {
		parent[i] = i
		if i == src {
			dis[src] = 0
			continue
		}
		dis[i] = math.MaxInt
	}
	for !q.IsEmpty() {
		top := q.Pop().([]int)
		dist, node := top[0], top[1]
		for _, adjNode := range g[node] {
			if dist+adjNode.Weight < dis[node] {
				parent[adjNode.Value] = node
				dis[node] = dist + adjNode.Weight
				q.Push([]int{dist + adjNode.Weight, adjNode.Value})
			}
		}
	}
	ans := []int{}
	node := destination
	for parent[node] != node {
		ans = append(ans, node)
		node = parent[node]
	}
	ans = append(ans, src)
	return ans
}

func ShortesDistanceBinaryPath(matrix [][]int, src, dest []int) int {
	m, n := len(matrix), len(matrix[0])
	distance := make([][]int, m)
	for i := 0; i < m; i++ {
		distance[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == src[0] && j == src[1] {
				distance[i][j] = 0
				continue
			}
			distance[i][j] = math.MaxInt
		}
	}
	q := queue.Queue{}
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}
	q.Push([]int{0, src[0], src[1]})
	for !q.IsEmpty() {
		top := q.Pop().([]int)
		dist := top[0]
		i, j := top[1], top[2]
		for k := 0; k < 4; k++ {
			r := i + row[k]
			c := j + col[k]
			if (r >= 0 && c >= 0) && (r < m && c < n) && matrix[r][c] == 1 && dist+1 < distance[r][c] {
				distance[r][c] = dist + 1
				q.Push([]int{dist + 1, r, c})
			}
		}
	}
	if distance[dest[0]][dest[1]] == math.MaxInt {
		return -1
	}
	return distance[dest[0]][dest[1]]
}

func PathWithMinEffort(matrix [][]int, m, n int) int {
	distance := make([][]int, m)
	for i := 0; i < m; i++ {
		distance[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				distance[i][j] = 0
				continue
			}
			distance[i][j] = math.MaxInt
		}
	}
	q := queue.Queue{}
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}
	q.Push([]int{0, 0, 0})
	for !q.IsEmpty() {
		top := q.Top().([]int)
		diff, i, j := top[0], top[1], top[2]
		for k := 0; k < 4; k++ {
			r := i + row[k]
			c := j + col[k]
			currDiff := int(math.Abs(float64(matrix[i][j] - matrix[r][c])))
			if (r >= 0 && c >= 0) && (r < m && c < n) && matrix[r][c] == 1 && currDiff > diff && currDiff < distance[r][c] {
				q.Push([]int{currDiff, r, c})
			}
		}
	}
	ans := matrix[m-1][n-1]
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func (g Graph) CheapestFlight(src, dest, k int) int {
	q := queue.Queue{}
	distance := make([]int, len(g)-1)
	for i := range distance {
		if i == 0 {
			continue
		}
		distance[i] = math.MaxInt
	}
	q.Push([]int{0, src, 0})
	for !q.IsEmpty() {
		top := q.Pop().([]int)
		stop, node, dis := top[0], top[1], top[2]
		if node == dest {
			continue
		}
		for _, adjNode := range g[node] {
			if stop <= k && dis+adjNode.Weight < distance[adjNode.Value] {
				distance[adjNode.Value] = dis + adjNode.Weight
				q.Push([]int{stop + 1, adjNode.Value, dis + adjNode.Weight})
			}
		}
	}
	return distance[dest]
}

func MinimumMultiplication(start, end int, arr []int) int {
	steps := make([]int, 10000)
	for i := 0; i < len(steps); i++ {
		if i == start {
			steps[i] = 0
		}
		steps[i] = math.MaxInt
	}
	q := queue.Queue{}
	q.Push([]int{0, start})
	for !q.IsEmpty() {
		top := q.Pop().([]int)
		step, node := top[0], top[1]
		for _, n := range arr {
			d := (n * node) % 100000
			if d == end {
				return step + 1
			}
			if step+1 < steps[d] {
				steps[d] = step + 1
				q.Push([]int{step + 1, d})
			}
		}
	}
	return -1
}

func (g Graph) NoOfWaysToReachDest(n int) int {
	distance, paths := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			distance[i] = 0
		}
		distance[i] = math.MaxInt
	}
	q := queue.Queue{}
	q.Push([]int{0, 0})
	for !q.IsEmpty() {
		top := q.Pop().([]int)
		currentDistance, node := top[0], top[1]
		for _, adjNode := range g[node] {
			dis := currentDistance + adjNode.Weight
			if dis < distance[adjNode.Value] {
				distance[adjNode.Value] = dis
				paths[adjNode.Value] = paths[node]
				q.Push([]int{dis, adjNode.Value})
			} else if dis == distance[adjNode.Value] {
				paths[adjNode.Value] += paths[node] % 10000000
			}
		}
	}
	return paths[n-1] % 10000000
}

func BellmanFord(edges [][]int, nodes int) []int {
	dis := []int{0}
	for i := 1; i < nodes; i++ {
		dis = append(dis, math.MaxInt)
	}
	for i := 0; i < nodes-1; i++ {
		for _, edge := range edges {
			u, v, wt := edge[0], edge[1], edge[2]
			if dis[u] != math.MaxInt {
				if dis[u]+wt < dis[v] {
					dis[v] = dis[u] + wt
				}
			}
		}
	}
	for _, edge := range edges {
		u, v, wt := edge[0], edge[1], edge[2]
		if dis[u] != math.MaxInt {
			if dis[u]+wt < dis[v] {
				return []int{-1}
			}
		}
	}
	return dis
}

func FloyWarshall(nodes int, edges [][]int) [][]int {
	adjMatrix := make([][]int, nodes)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, nodes)
		for j := 0; j < nodes; j++ {
			if i == j {
				adjMatrix[i][j] = 0
			} else {
				adjMatrix[i][j] = math.MaxInt
			}
		}
	}
	for _, edge := range edges {
		adjMatrix[edge[0]][edge[1]] = edge[2]
	}
	for n := 0; n < nodes; n++ {
		for i := 0; i < nodes; i++ {
			for j := 0; j < nodes; j++ {
				if adjMatrix[i][n] != math.MaxInt && adjMatrix[n][j] != math.MaxInt {
					adjMatrix[i][j] = int(math.Min(float64(adjMatrix[i][j]), float64(adjMatrix[i][n]+adjMatrix[n][j])))
				}
			}
		}
	}
	for i := 0; i < nodes; i++ {
		if adjMatrix[i][i] < 0 {
			return nil
		}
	}
	return adjMatrix
}

func CitiesAtDistance(edges [][]int, nodes, threshold int) int {
	adjMatrix := FloyWarshall(nodes, edges)
	minCityReached := nodes
	cityNo := -1
	for city := 0; city < nodes; city++ {
		var cityCount int
		for adjCity := 0; adjCity < nodes; adjCity++ {
			if city == adjCity {
				continue
			}
			if adjMatrix[city][adjCity] <= threshold {
				cityCount++
			}
		}
		if cityCount < minCityReached {
			minCityReached = cityCount
			cityNo = city
		}
	}
	return cityNo
}

func (g Graph) PrimsAlgo() int {
	visited := make([]bool, len(g)-1)
	pq := queue.Queue{}
	pq.Push([]int{0, 0})
	//in pq, the data is already sorted, but in golang PQ doesn't exist so fuck it
	var sum int
	for !pq.IsEmpty() {
		top := pq.Pop().([]int)
		node := top[1]
		wt := top[0]
		if visited[node] {
			continue
		}
		visited[node] = true
		sum += wt
		for _, adj := range g[node] {
			if !visited[adj.Value] {
				pq.Push([]int{adj.Weight, adj.Value})
			}
		}
	}
	return sum
}

func KruskalAlgo(v int, edges [][]int) int {
	ds := NewDisjointSet(v)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	var sum int
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		wt := edge[2]
		if ds.UltimateParent(u) != ds.UltimateParent(v) {
			sum += wt
			ds.Union([]int{u, v})
		}
	}
	return sum
}

func NoOfProvinces(v int, adjMatrix [][]int) (count int) {
	ds := NewDisjointSet(v)
	for i := 1; i < len(adjMatrix); i++ {
		for j := 1; j < len(adjMatrix[i]); j++ {
			if adjMatrix[i][j] == 1 {
				ds.Union([]int{i, j})
			}
		}
	}
	for i := 1; i <= v; i++ {
		if ds.UltimateParent(i) == i {
			count++
		}
	}
	return
}

func NoOfOperations(nodes int, edges [][]int) int {
	ds := NewDisjointSet(nodes - 1)
	var extraEdges, ans, comp int
	// ans := nodes - 1
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if ds.UltimateParent(u) == ds.UltimateParent(v) {
			extraEdges++
		} else {
			ds.Union(edge)
		}
	}
	for i := 0; i < nodes; i++ {
		if ds.UltimateParent(i) == i {
			comp++
		}
	}
	ans = comp - 1
	if extraEdges >= ans {
		return extraEdges
	}
	return -1
}

func AccountsMerge(data [][]string) [][]string {
	emailToIndex := make(map[string]int)
	ds := NewDisjointSet(len(data) - 1)
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data[i]); j++ {
			if pos, ok := emailToIndex[data[i][j]]; ok {
				ds.Union([]int{i, pos})
			} else {
				emailToIndex[data[i][j]] = i
			}
		}
	}
	ans := make([][]string, len(data))
	for email, index := range emailToIndex {
		pos := ds.UltimateParent(index)
		if ans[pos] == nil {
			ans[pos] = append(ans[pos], data[pos][0])
		}
		ans[pos] = append(ans[pos], email)
	}
	return ans
}

func NoOfIslandsII(m, n int, operations [][]int) []int {
	ds := NewDisjointSet(m*n - 1)
	var islandCount int
	ans := make([]int, len(operations))
	islandMatrix := make([][]int, m)
	for i := range islandMatrix {
		islandMatrix[i] = make([]int, n)
	}
	row := []int{-1, 0, 1, 0}
	column := []int{0, 1, 0, -1}
	for i, operation := range operations {
		u, v := operation[0], operation[1]
		if islandMatrix[u][v] != 1 {
			islandMatrix[u][v] = 1
			islandCount++
			for j := 0; j < 4; j++ {
				r := u + row[j]
				c := v + column[j]
				if (r >= 0 && c >= 0) && (r < m && c < n) && islandMatrix[r][c] == 1 {
					node1 := u*n + v
					node2 := r*n + c
					if ds.UltimateParent(node1) != ds.UltimateParent(node2) {
						islandCount--
						ds.Union([]int{node1, node2})
					}
				}
			}
			ans[i] = islandCount
		} else {
			ans[i] = islandCount
		}
	}
	return ans
}

func LargestIsland(m, n int, matrix [][]int) int {
	ds := NewDisjointSet(m*n - 1)
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				for k := range row {
					r := i + row[k]
					c := j + col[k]
					if (r >= 0 && c >= 0) && (r < m && c < n) && matrix[r][c] == 1 {
						u, v := (i*n)+j, (r*n)+c
						ds.UnionBySize([]int{u, v})
					}
				}
			}
		}
	}
	maxIsland := math.MinInt
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				continue
			}
			islandMap := make(map[int]bool)
			for k := range row {
				r := i + row[k]
				c := j + col[k]
				if (r >= 0 && c >= 0) && (r < m && c < n) && matrix[r][c] == 1 {
					islandMap[ds.UltimateParent(r*n+c)] = true
				}
			}
			var currentMax int
			for parentNode := range islandMap {
				currentMax += ds.size[parentNode]
			}
			currentMax++
			if currentMax > maxIsland {
				maxIsland = currentMax
			}
		}
	}
	return maxIsland
}

func MostStonesRemoved(m, n int, matrix [][]int) int {
	ds := NewDisjointSet(m + n - 1)
	var componentCount, stones int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				stones++
				u := i
				v := j + m
				ds.UnionBySize([]int{u, v})
			}
		}
	}
	for i := 0; i < m+n; i++ {
		if ds.UltimateParent(i) == i && ds.size[i] > 1 {
			componentCount++
		}
	}
	return stones - componentCount
}

func (g Graph) StronglyConnectedComponents() int { //Kosaraju's Algo
	queue := queue.Queue{}
	visited := make([]int, len(g)-1)
	for i := 0; i < len(visited); i++ {
		visited[i] = 0
	}
	//Sort nodes by end time
	g.dfsForEndTime(0, visited, &queue)
	//Reverse the node
	newGraph := make([][]int, len(g)-1)
	for i := 0; i < len(g)-1; i++ {
		visited[i] = 0
		for _, node := range g[i] {
			adjListForNode := newGraph[node.Value]
			if adjListForNode != nil {
				newGraph[node.Value] = append(adjListForNode, i)
			} else {
				newGraph[node.Value] = []int{i}
			}
		}
	}
	var strong int
	for !queue.IsEmpty() {
		top := queue.IntTop()
		queue.Pop()
		if visited[top] == 0 {
			strong++
			helperDFS(g, visited, top)
		}
	}
	return strong
}

func (g Graph) dfsForEndTime(src int, visited []int, queue *queue.Queue) {
	visited[src] = 1
	for _, node := range g[src] {
		if visited[node.Value] == 0 {
			g.dfsForEndTime(node.Value, visited, queue)
		}
	}
	queue.Push(src)
}

var timer int

func (g Graph) dfsForInsertTime(src int, parent int, visited, insertionTime, lowestIT []int, bridges *[][]int) {
	visited[src] = 1
	insertionTime[src] = timer
	lowestIT[src] = timer
	timer++
	for _, node := range g[src] {
		if node.Value == parent {
			continue
		}
		if visited[node.Value] == 0 {
			g.dfsForInsertTime(node.Value, src, visited, insertionTime, lowestIT, bridges)
			lowestIT[src] = int(math.Min(float64(lowestIT[node.Value]), float64(lowestIT[src])))
			if lowestIT[node.Value] > insertionTime[src] {
				*bridges = append(*bridges, []int{src, node.Value})
			}
		} else {
			lowestIT[src] = int(math.Min(float64(lowestIT[node.Value]), float64(lowestIT[src])))
		}
	}
}

func (g Graph) Bridges() [][]int {
	visited, insertionTime, lowestIT := make([]int, len(g)-1), make([]int, len(g)-1), make([]int, len(g)-1)
	bridges := [][]int{}
	timer = 1
	for i := 0; i < len(visited); i++ {
		visited[i] = 0
	}
	g.dfsForInsertTime(0, -1, visited, insertionTime, lowestIT, &bridges)
	return bridges
}

func (g Graph) dfsForArticulation(node, parent int, visited, insertionTime, lowestIT []int, articulationNodes map[int]bool) {
	visited[node] = 1
	lowestIT[node], insertionTime[node] = timer, timer
	timer++
	if parent == -1 && len(g[node]) > 1 {
		articulationNodes[node] = true
	}
	for _, it := range g[node] {
		if it.Value == parent {
			continue
		}
		if visited[it.Value] == 1 {
			lowestIT[node] = int(math.Min(float64(insertionTime[it.Value]), float64(lowestIT[node])))
		} else {
			g.dfsForArticulation(it.Value, node, visited, insertionTime, lowestIT, articulationNodes)
			lowestIT[node] = int(math.Min(float64(lowestIT[it.Value]), float64(lowestIT[node])))
			if lowestIT[it.Value] >= insertionTime[node] && parent != -1 {
				articulationNodes[node] = true
			}
		}
	}
}

func (g Graph) ArticulationPoint() map[int]bool {
	visited, insertionTime, lowestIT := make([]int, len(g)-1), make([]int, len(g)-1), make([]int, len(g)-1)
	timer = 1
	ans := make(map[int]bool)
	for i := 0; i < len(g)-1; i++ {
		g.dfsForArticulation(i, -1, visited, insertionTime, lowestIT, ans)
	}
	return ans
}
