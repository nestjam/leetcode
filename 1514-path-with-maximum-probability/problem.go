package pathwithmaximumprobability

import "math"

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	dist := make([]float64, n)
	visited := make([]bool, n)

	graph := initGraph(n, edges)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	dist[start_node] = 1

	for n > 0 {
		u := getNext(dist, visited)

		if u == -1 || u == end_node {
			break
		}

		visited[u] = true
		n--

		neighbors := graph[u]
		for i := 0; i < len(neighbors); i++ {
			v := neighbors[i]
			if visited[v.node] {
				continue
			}
			alt := dist[u] * succProb[v.edge]
			if alt > dist[v.node] {
				dist[v.node] = alt
			}
		}
	}

	if dist[end_node] == -1 {
		return 0
	}

	return round(dist[end_node])
}

func initGraph(n int, edges [][]int) map[int]neighbors {
	g := make(map[int]neighbors, n)

	for i := 0; i < n; i++ {
		g[i] = make(neighbors, 0)
	}

	for i := 0; i < len(edges); i++ {
		start := edges[i][0]
		end := edges[i][1]
		g[start] = append(g[start], neighbor{node: int16(end), edge: int16(i)})
		g[end] = append(g[end], neighbor{node: int16(start), edge: int16(i)})
	}

	return g
}

func getNext(dist []float64, visited []bool) int {
	max := float64(0)
	n := -1

	for i := 0; i < len(visited); i++ {
		if visited[i] {
			continue
		}

		if dist[i] >= max {
			max = dist[i]
			n = i
		}
	}

	return n
}

func round(v float64) float64 {
	d := float64(100_000)
	return math.Round(v*d) / d
}

type neighbor struct {
	edge int16
	node int16
}

type neighbors []neighbor