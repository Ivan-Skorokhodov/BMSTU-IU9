package main

import (
	"container/heap"
	"fmt"
)

type Elem struct {
	number_vertex, distance, vertex_weight, index int
}

type PriorityQueue []*Elem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	it := x.(*Elem)
	it.index = n
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	it := old[n-1]
	old[n-1] = nil
	it.index = -1
	*pq = old[0 : n-1]
	return it
}

func get_link_v(v, n int, input_data []*Elem) []*Elem {
	var list []*Elem

	if v%n != 0 {
		list = append(list, input_data[v-1])
	}

	if (v+1)%n != 0 {
		list = append(list, input_data[v+1])
	}

	if v-n >= 0 {
		list = append(list, input_data[v-n])
	}

	if v+n <= n*n-1 {
		list = append(list, input_data[v+n])
	}

	return list
}

func main() {
	var n, weight int
	var input_data []*Elem
	list_adjacencies := make(map[*Elem][]*Elem, 20)

	fmt.Scan(&n)

	for i := 0; i < n*n; i++ {
		fmt.Scan(&weight)

		dist := 1000000
		if i == 0 {
			dist = 0
		}
		input_data = append(input_data, &Elem{i, dist, weight, 0})
	}

	for _, elem := range input_data {
		list_adjacencies[elem] = append(list_adjacencies[elem], get_link_v(elem.number_vertex, n, input_data)...)
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, input_data[0])

	for pq.Len() > 0 {
		elem := heap.Pop(&pq).(*Elem)

		for _, new_elem := range list_adjacencies[elem] {
			if new_elem.distance > elem.distance+new_elem.vertex_weight {
				new_elem.distance = elem.distance + new_elem.vertex_weight
				heap.Push(&pq, new_elem)
			}
		}
	}

	fmt.Println(*&input_data[n*n-1].distance + *&input_data[0].vertex_weight)

}