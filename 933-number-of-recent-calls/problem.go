package numberofrecentcalls

type RecentCounter struct {
	q queue
}

func Constructor() RecentCounter {
	return RecentCounter{
		q: newQueue(),
	}
}

func (c *RecentCounter) Ping(t int) int {
	c.q.push(t)

	for c.q.peek() < t-3000 {
		c.q.pop()
	}

	return c.q.len()
}

type queue []int

func (q queue) len() int {
	return len(q)
}

func (q *queue) push(x int) {
	*q = append(*q, x)
}

func (q *queue) peek() int {
	return (*q)[0]
}

func (q *queue) pop() int {
	h := *q
	var i int
	l := len(h)
	i, *q = h[0], h[1:l]
	return i
}

func newQueue() queue {
	return queue{}
}
