package Mgoruntine

import "sync"

var WaitGroup = sync.WaitGroup{}

func QuickSortRoutineVer(array []int) {

	if len(array) == 1 || len(array) == 0 || array == nil {
		// 这里是啥意思？
		WaitGroup.Done()
		return
	}
	if len(array) <= 3 {
		insertSort(array)
		WaitGroup.Done()
		return
	}

	// 三数取中法则
	left := 0
	right := len(array)-1
	mid := len(array) / 2
	if array[left] > array[mid] {
		swap(&array[left],&array[mid])
	}
	if array[left] > array[right] {
		swap(&array[left],&array[right])
	}
	if array[mid] > array[right] {
		swap(&array[mid],&array[right])
	}
	// 中间的数藏起到倒数第二个
	swap(&array[mid],&array[right-1])

	pivot := array[right-1]

	left = 0
	right = len(array) - 2

	for left < right {
		for array[left] <= pivot && left < right{
			left ++
		}
		for array[right] >= pivot && right > left {
			right --
		}
		if left < right && right >= 0{
			swap(&array[left],&array[right])
		}
	}
	swap(&array[left],&array[right])
	WaitGroup.Add(2)
	go QuickSortQzh(array[:left])
	go QuickSortQzh(array[left:])
	WaitGroup.Done()
	return
}
