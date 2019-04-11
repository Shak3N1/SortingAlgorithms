package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func fileToIntArr(fileName string) []int {
	filedata, _ := ioutil.ReadFile(fileName)
	lines := strings.Split(string(filedata), "\n")
	var v []int
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		for _, number := range numbers {
			n, _ := strconv.Atoi(number)
			v = append(v, n)
		}
	}
	return v
}
func main() {
	fileName := "array"
	v := fileToIntArr(fileName)
	fmt.Println("[-] Before sorted\n")
	fmt.Println(v)
	fmt.Println("\n[-] After Sorted\n")
	//Bubble Sort
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(v)
	fmt.Printf("\n[-] Vector size: %d\n", len(v))
	fmt.Printf("[-]Sorted in %s with BubbleSort\n", elapsedTime)
	//Insertion Sort
	v = nil
	v = fileToIntArr(fileName)
	startTime = time.Now()
	insertionSort(v)
	elapsedTime = time.Since(startTime)
	fmt.Printf("[-]Sorted in %s with InsertionSort\n", elapsedTime)
	//Quick Sort
	v = nil
	v = fileToIntArr(fileName)
	startTime = time.Now()
	quickSort(v)
	elapsedTime = time.Since(startTime)
	fmt.Printf("[-]Sorted in %s with QuickSort\n", elapsedTime)
	//Merge Sort
	v = nil
	v = fileToIntArr(fileName)
	startTime = time.Now()
	mergeSort(v)
	elapsedTime = time.Since(startTime)
	fmt.Printf("[-]Sorted in %s with MergeSort\n", elapsedTime)
}
