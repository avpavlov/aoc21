package utils

type PriorityQueue struct {
	Items      []interface{}
	Comparator func(i, j interface{}) bool
}

func (pq PriorityQueue) Len() int { return len(pq.Items) }

func (pq PriorityQueue) Less(i, j int) bool { return pq.Comparator(pq.Items[i], pq.Items[j]) }

func (pq PriorityQueue) Swap(i, j int) { pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i] }

func (pq *PriorityQueue) Push(x interface{}) { pq.Items = append(pq.Items, x) }

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.Items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	pq.Items = old[0 : n-1]
	return item
}
