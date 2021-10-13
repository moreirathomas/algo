package sort

// TODO make it generic or name it as specialized for ints.

func Binary(arr []int, target int) int {
	low, high := 0, len(arr)

	for low < high {
		mid := (low + high) / 2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
