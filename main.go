package main

func recursiveMerge(n []int, l, r int) {
	if l+1 >= r {
		return
	}

	m := (l+r) / 2
	recursiveMerge(n, l, m)
	recursiveMerge(n, m, r)
	merge(n, m, l, r)
}

func merge(n []int, m, l, r int) {
	var (
		it1, it2 int
		result = make([]int, r-l)
	)

	for l+it1 < m && m + it2 < r {
		if n[l+it1] < n[m+it2] {
			result[it1+it2] = n[l+it1]
			it1++
		}else{
			result[it1+it2] = n[m+it2]
			it2++
		}
	}

	for l+it1 < m {
		result[it1+it2] = n[l+it1]
		it1++
	}

	for m+it2 < r {
		result[it1+it2] = n[m+it2]
		it2++
	}

	for i := 0; i < it1+it2; i++ {
		n[l+i] = result[i]
	}
}


