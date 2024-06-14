package minimumtimevisitingallpoints

func minTimeToVisitAllPoints(points [][]int) int {
	t := 0

	for i := 0; i < len(points)-1; i++ {
		t += time(points[i+1], points[i])
	}

	return t
}

func time(p, p2 []int) int {
	dx := p[0] - p2[0]
	if dx < 0 {
		dx = -dx
	}

	dy := p[1] - p2[1]
	if dy < 0 {
		dy = -dy
	}

	if dx > dy {
		return dx
	}

	return dy
}
