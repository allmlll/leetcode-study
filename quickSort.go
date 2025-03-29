package main

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	left := 0
	right := len(arr) - 1
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[left], arr[right] = arr[right], arr[left]
	}

	arr[left] = pivot
	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

//

func qSort(arr []int) {
	if len(arr) < 1 {
		return

	}

	pivot := arr[0]
	l := 0
	r := len(arr) - 1
	for l < r {
		if l < r && arr[r] >= pivot {
			r--
		}
		arr[l] = arr[r]
		if l < r && arr[l] <= pivot {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = arr[pivot]

	quickSort(arr[:l])
	quickSort(arr[l+1:])
}

func quSort(arr []int) {

	n := len(arr)
	if n < 1 {
		return
	}
	pivot := arr[0]
	l := 0
	r := n - 1
	for l < r {
		for l < r && arr[r] >= pivot {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= pivot {
			l++
		}
		arr[r] = arr[l]

	}
	arr[l] = pivot
	quSort(arr[:l])
	quSort(arr[l+1:])
}
