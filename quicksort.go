package main

func partition(v []int, lowerPos, highPos int) int {
	pivot := v[highPos]
	i := lowerPos - 1
	for j := lowerPos; j <= highPos-1; j++ {
		if v[j] <= pivot {
			i++
			switchValues(&v[i], &v[j])
		}
	}
	switchValues(&v[i+1], &v[highPos])
	return i + 1
}

func quickSortRecursive(v []int, low, high int) {
	if low < high {
		p := partition(v, low, high)
		quickSortRecursive(v, low, p-1)
		quickSortRecursive(v, p+1, high)
	}
}
func quickSort(v []int) {
	quickSortRecursive(v, 0, len(v)-1)
}
