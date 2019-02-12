package main

import (
	f "fmt"
	t "time"
)

func show_term(m map[int]byte, arr []int, size int) {
	for i := 0; i < size; i++ {
		f.Printf("%c", m[arr[i+1]])
	}
}

func last_term(arr []int, size int) bool {
	j := size
	for i := 1; i < size+1; i++ {
		if arr[i] != j {
			return true
		}
		j--
	}
	return false
}

func swap(arr []int, i int, j int) (int, int) {
	return arr[j], arr[i]
}

func sort(arr []int, j int, end int) []int {
	for start := j; start < end+1; start++ {
		for i := start + 1; i < end+1; i++ {
			if arr[start] > arr[i] {
				arr[start], arr[i] = swap(arr, start, i)
			}
		}
	}
	return arr
}

func check(arr []int, begin int, end int) int {
	item := arr[begin]
	for i := end; i > begin+1; i-- {
		if arr[i] > item {
			return i
		}
	}
	return 0
}

func next_term(arr []int, size int) []int {
	var i int
	for i = size - 1; i > 0; i-- {
		if arr[i] < arr[i+1] {
			index := check(arr, i, size)
			if index != 0 {
				arr[i], arr[index] = swap(arr, i, index)
			} else {
				arr[i], arr[i+1] = swap(arr, i, i+1)
			}
			break
		}
	}
	return sort(arr, i+1, size)
}

func main() {

	var name string
	assign_map := make(map[int]byte)
	f.Scan(&name)
	arr := make([]int, len(name)+1)
	var start t.Time
	for i, _ := range name {
		start = t.Now()
		assign_map[i+1] = name[i]
		arr[i+1] = i + 1
	}
	counter := 0
	for last_term(arr, len(name)) {
		f.Println("\ncounter: ", counter)
		counter++
		show_term(assign_map, arr, len(name))
		arr = next_term(arr, len(name))
	}
	f.Println("\ncounter: ", counter)
	show_term(assign_map, arr, len(name))

	elapsed := t.Since(start)
	f.Println("\n", elapsed)
}
