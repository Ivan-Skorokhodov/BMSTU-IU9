package main

import "fmt"

func my_min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func my_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func add(a, b []int32, p int) []int32 {
	var result []int32
	var res int32
	var counter int32 = 0
	last_index := 0
	len_a := len(a)
	len_b := len(b)

	for i := 0; i < my_min(len_a, len_b); i++ {
		res = b[i] + a[i] + counter

		if res >= int32(p) {
			res %= int32(p)
			counter = 1
		} else {
			counter = 0
		}

		result = append(result, res)
		last_index++
	}

	for last_index <= my_max(len_a, len_b) {
		if last_index <= len_a-1 {

			res = a[last_index] + counter
			if res >= int32(p) {
				res %= int32(p)
				counter = 1
			} else {
				counter = 0
			}
			result = append(result, res)

		} else if last_index <= len_b-1 {

			res = b[last_index] + counter
			if res >= int32(p) {
				res %= int32(p)
				counter = 1
			} else {
				counter = 0
			}
			result = append(result, res)
		}

		last_index++
	}

	if counter == 1 {
		result = append(result, 1)
	}

	return result
}

func main() {
	var n1, n2 int
	fmt.Scanf("%d %d", &n1, &n2)
	var a []int32
	var b []int32

	for i := 0; i < n1; i++ {
		var elem int32
		fmt.Scan(&elem)
		a = append(a, elem)
	}

	for i := 0; i < n2; i++ {
		var elem int32
		fmt.Scan(&elem)
		b = append(b, elem)
	}

	fmt.Println(add(a, b, 7))
}