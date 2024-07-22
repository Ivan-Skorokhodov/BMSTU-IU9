package main

func hsort(n int, less func(i, j int) bool, swap func(i, j int)) { //реализация взята с Хабра

	var heapify func(n1, n2 int)

	heapify = func(n, i int) {
		largest := i
		l := 2*i + 1
		r := 2*i + 2

		if l < n && less(largest, l) {
			largest = l
		}

		if r < n && less(largest, r) {
			largest = r
		}

		if largest != i {
			swap(i, largest)
			heapify(n, largest)
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(n, i)
	}

	for i := n - 1; i >= 0; i-- {
		swap(i, 0)
		heapify(i, 0)
	}
}

func main() {
}