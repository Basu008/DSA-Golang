package graph

import (
	"fmt"

	"github.com/Basu008/DSA-Golang/queue"
)

type Pair struct {
	Value  int
	Weight int
}

type GraphInput struct {
	Nodes    int
	Edges    [][]int
	Directed bool
}

type Graph [][]Pair

func CreateGraph(input GraphInput) (graph Graph) {
	graph = make(Graph, input.Nodes+1)
	for _, edge := range input.Edges {
		weight := 1
		if len(edge) > 2 {
			weight = edge[2]
		}
		graph[edge[0]] = append(graph[edge[0]], Pair{edge[1], weight})
		if !input.Directed {
			graph[edge[1]] = append(graph[edge[1]], Pair{edge[0], weight})
		}
	}
	return
}

func (g Graph) String() string {
	for i := 1; i < len(g); i++ {
		fmt.Printf("%d -> ", i)
		for _, v := range g[i] {
			fmt.Printf("%d ", v.Value)
		}
		fmt.Println("")
	}
	return ""
}

func (g Graph) BFS(node int, visited []int) {
	q := queue.Queue{}
	if len(visited) == 0 {
		visited = make([]int, len(g))
	}
	q.Push(node)
	visited[node] = 1
	for !q.IsEmpty() {
		top := q.Pop().(int)
		fmt.Printf("%d, ", top)
		for _, v := range g[top] {
			if visited[v.Value] != 1 {
				q.Push(v.Value)
				visited[v.Value] = 1
			}
		}
	}
}

func (g Graph) DFS(node int) {
	helperDFS(g, make([]int, len(g)), node)
}

func helperDFS(g Graph, visited []int, node int) {
	visited[node] = 1
	for _, v := range g[node] {
		if visited[v.Value] != 1 {
			helperDFS(g, visited, v.Value)
		}
	}
}

func DFSForMatrix(i, j int, matrix, visited [][]int) {
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}
	visited[i][j] = 1
	m := len(matrix)
	n := len(matrix[0])
	for k := 0; k < 4; k++ {
		r := i + row[k]
		c := j + col[k]
		if (r >= 0 && c >= 0) && (r < m && c < n) && visited[r][c] == 0 && matrix[r][c] == matrix[i][j] {
			DFSForMatrix(r, c, matrix, visited)
		}
	}
}

func DFSToCountIslands(i, j int, matrix, visited, set [][]int, baseI, baseJ *int) [][]int {
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}
	visited[i][j] = 1
	if len(set) == 0 {
		set = append(set, []int{0, 0})
	}
	m := len(matrix)
	n := len(matrix[0])
	for k := 0; k < 4; k++ {
		r := i + row[k]
		c := j + col[k]
		if (r >= 0 && c >= 0) && (r < m && c < n) && visited[r][c] == 0 && matrix[r][c] == matrix[i][j] {
			set = append(set, []int{r - *baseI, c - *baseJ})
			DFSForMatrix(r, c, matrix, visited)
		}
	}
	return set
}

type DisjointSet struct {
	rank   []int
	size   []int
	parent []int
}

func NewDisjointSet(n int) *DisjointSet {
	rank, size := make([]int, n+1), make([]int, n+1)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	for i := range size {
		size[i] = 1
	}
	return &DisjointSet{
		rank:   rank,
		parent: parent,
		size:   size,
	}
}

func (ds *DisjointSet) UltimateParent(node int) int {
	if ds.parent[node] == node {
		return node
	}
	ds.parent[node] = ds.UltimateParent(ds.parent[node])
	return ds.parent[node]
}

func (ds *DisjointSet) Union(edge []int) {
	u, v := edge[0], edge[1]
	pu, pv := ds.UltimateParent(u), ds.UltimateParent(v)
	if pu == pv {
		return
	}
	ru, rv := ds.rank[pu], ds.rank[pv]
	if ru < rv {
		ds.parent[pu] = pv
	} else if ru > rv {
		ds.parent[pv] = pu
	} else {
		ds.rank[pu]++
		ds.parent[pv] = pu
	}
}

func (ds *DisjointSet) UnionBySize(edge []int) {
	u, v := edge[0], edge[1]
	pu, pv := ds.UltimateParent(u), ds.UltimateParent(v)
	if pu == pv {
		return
	}
	su := ds.size[pu]
	sv := ds.size[pv]
	if su < sv {
		ds.parent[pu] = pv
		ds.size[pv] += ds.size[pu]
	} else {
		ds.parent[pv] = pu
		ds.size[pu] += ds.size[pv]
	}
}
