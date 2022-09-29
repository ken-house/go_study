package main

func AddUpper(num int) int {
	res := 0
	for i := 1; i <= num; i++ {
		res += i
	}
	return res
}

func SliceAppend(num int) []int {
	s := make([]int, 0)
	for i := 0; i < num; i++ {
		s = append(s, i)
	}
	return s
}

func MapAppend(num int) map[int]int {
	m := make(map[int]int, 0)
	for i := 0; i < num; i++ {
		m[i] = i
	}
	return m
}
