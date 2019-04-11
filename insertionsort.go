package main

func insertionSort(vector []int) {
	size := len(vector)
	if size <= 1 {
		return
	}
	for i := 0; i < size; i++ {
		j := i - 1
		key := vector[i]
		for ; j >= 0 && vector[j] > key; j-- {
			vector[j+1] = vector[j]
		}
		vector[j+1] = key
	}
}
