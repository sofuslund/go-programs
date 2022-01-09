package sortingalgorithms

// Selection sort sorts an array by repeatedly finding the minimum element from the unsorted part and placing it at the beginning. It maintains two subarrays in a given array.
// 1) The subarray which is already sorted.
// 2) Remaining subarray which is unsorted.
// In every iteration of selection sort, the smallest value from the unsorted subarray is picked and moved to the sorted subarray.
func SelectionSort(_arr []int) ([]int) {
    arr := make([]int, len(_arr))
    copy(arr, _arr)
    // Rapidly increase the starting index of the unsorted part while solving
    for i, _ := range arr {
        // Find the smallest value in the unsorted part
        minIdx := i
        for j := i; j < len(arr); j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        // Move the minimum element to the sorted part
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
    return arr
}
