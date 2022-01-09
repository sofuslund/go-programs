package sortingalgorithms

// Partition() chooses a pivot and partitions an array into two parts one with
// elements smaller than the pivot and the other one with elements bigger than
// the pivot it then returns the new position of the pivot. The function uses
// the Lomuto partition scheme
func partition(arr []int, start int, end int) int {
    var pivot = arr[end] // Pick the last element as pivot
    // Maintain where the part of smaller elements than the pivot ends
    // This will increase as the function moves elements smaller than the pivot down to the lower part of the array
    var i = start
    // Loop through all the array elements besides the last one which is the pivot
    for j := start; j < end; j++ {

        // If the current element is smaller than the pivot move it to the lower part and increment where the lower part ends
        if arr[j] < pivot {

            arr[i], arr[j] = arr[j], arr[i]
            i++
        }
    }
    // Move the pivot to the correct place
    arr[i], arr[end] = arr[end], arr[i]
    // Return at what index the pivot is now
    return i
}

// Quick sort sorts an array by partitioning the array into two smaller arrays
// where one holds values smaller than a specified value called the pivot and the
// other one holds values greater than the pivot. It then partitions each of the
// two smaller arrays and repeats the process on the resulting arrays until the
// arrays to be partitioned has a size of one
func quickSort2(arr []int, start int, end int) {
    // Do nothing if the part to be sorted has length one since it's then already sorted
    if start >= end {
        return
    }
    // Let the partition function choose a pivot and divide the array into a
    // part with elements less than the pivot and a part with elements greater
    // than the pivot. Finally get the index of the pivot
    var index = partition(arr, start, end)
    quickSort2(arr, start, index-1) // Sort the part below the pivot index
    quickSort2(arr, index+1, end) // Sort the part above the pivot index
}

// Convenience function to fill in the arguments to the real quicksort function
func QuickSort(_arr []int) []int {
    // Copy the input array
    var arr = make([]int, len(_arr))
    copy(arr, _arr)
    // fmt.Println(arr)
    // Sort the copied array
    quickSort2(arr, 0, len(arr)-1)
    // fmt.Println(partition(arr, 0, len(arr)-1))
    return arr // Return the copied array
}
