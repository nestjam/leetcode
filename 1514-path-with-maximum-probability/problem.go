package pathwithmaximumprobability

import "math"

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	dist := make([]float64, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	dist[start_node] = 1

	for n > 0 {
		u := getNext(dist, visited)

		if u == -1 {
			break
		}

		visited[u] = true
		n--

		neighbors := getNeighbors(u, edges, visited)
		for i := 0; i < len(neighbors); i++ {
			v := neighbors[i]
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

func getNeighbors(u int, edges [][]int, visited []bool) []neighbor {
	neighbors := make([]neighbor, 0)

	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		if edge[0] == u && !visited[edge[1]] {
			neighbors = append(neighbors, neighbor{edge: i, node: edge[1]})
		} else if edge[1] == u && !visited[edge[0]] {
			neighbors = append(neighbors, neighbor{edge: i, node: edge[0]})
		}
	}

	return neighbors
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
	edge int
	node int
}
