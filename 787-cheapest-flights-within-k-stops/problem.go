package cheapestflightswithinkstops

var f cheapestFlight

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	f = newCheapestFlight(n, flights, dst, k)

	forward(newPath(src))

	if f.path == nil {
		return -1
	}
	return f.path.score
}

func forward(path path) {
	if path.len-2 > f.k || (f.path != nil && path.score >= f.path.score) {
		return
	}

	if path.n == f.dst {
		f.path = &path
		return
	}

	for i := 0; i < len(f.graph[path.n]); i++ {
		newPath := path.append(f.graph[path.n][i])
		if !path.contains(newPath.n) {
			forward(newPath)
		}
	}
}

type cheapestFlight struct {
	path  *path
	graph graph
	dst   int
	k     int
}

func newCheapestFlight(n int, flights [][]int, dst int, k int) cheapestFlight {
	return cheapestFlight{
		graph: newGraph(n, flights),
		dst:   dst,
		k:     k,
	}
}

type path struct {
	prev  *path
	score int
	n     int
	len   int
}

func newPath(n int) path {
	return path{
		n:   n,
		len: 1,
	}
}

func (p path) append(e edge) path {
	return path{
		n:     e.n,
		score: p.score + e.score,
		len:   p.len + 1,
		prev:  &p,
	}
}

func (p *path) contains(n int) bool {
	for p != nil {
		if p.n == n {
			return true
		}
		p = p.prev
	}
	return false
}

type edge struct {
	n     int
	score int
}

type graph [][]edge

func newGraph(n int, edges [][]int) graph {
	g := make([][]edge, n)
	for i := 0; i < n; i++ {
		g[i] = make([]edge, 0)
	}

	for i := 0; i < len(edges); i++ {
		e := edges[i]
		g[e[0]] = append(g[e[0]], edge{n: e[1], score: e[2]})
	}
	return g
}
