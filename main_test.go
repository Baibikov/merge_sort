package main

import "testing"

func Test_recursiveMerge(t *testing.T) {

	var tests = [][]int{
		{4,2,34,423,42,31,123},
		{42,312,123,354,42,122,432},
		{5,65,2,234,2,4,43,234,6534436,86},
		{5,2,34,6,63,134224},
	}

	for _, v := range tests {
		recursiveMerge(v, 0, len(v))
	}

	for _, v := range tests {
		if !check(v) {
			t.Errorf("массив или слайз не является отсортированным")
		}
	}
	
}

func check(n []int) bool {
	for i, v := range n {
		if i+1 < len(n) {
			return true
		}

		if v > n[i+1] {
			return false
		} 
	}

	return true
}