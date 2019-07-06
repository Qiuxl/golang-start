package Mgoruntine


func QuickSortQzh(array []int) {

	if len(array) == 1 || len(array) == 0 || array == nil {
		return
	}
	if len(array) <= 3 {
		insertSort(array)
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

	QuickSortQzh(array[:left])
	QuickSortQzh(array[left:])
	return
}

/**
 当数量长度小于3的时候采用插入排序
 */
func insertSort( arr []int){
	if(len(arr) <= 1){
		return;
	}
	for end := 1; end < len(arr); end ++ {
		temp := arr[end]
		var index int
		for index = end -1 ; index >= 0 && arr[index] > temp; index-- {
			arr[index+1] = arr[index]
		}
		arr[index + 1] = temp
	}
}


func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}