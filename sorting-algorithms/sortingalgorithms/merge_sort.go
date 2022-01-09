package sortingalgorithms

func min(a, b int) int{
    if a < b {
        return a
    }
    return b
}

func merge(arr []int, start int, mid int, end int) {
    var sArr1, sArr2 = make([]int, mid-start), make([]int, end+1-mid)
    var arrIdx, sArr1Idx, sArr2Idx = start, 0, 0 // Make index pointers to each of the arrays
    // Make copies of the lower and higher subarrays
    copy(sArr1, arr[start:mid])
    copy(sArr2, arr[mid:end+1])

    // Iterate while the sorted array is not completed
    for {
        if sArr2Idx >= len(sArr2) {
            // If all elements is taken from the second subarray
            // Copy the remaining elements in the first subarray into the sorted array
            copy(arr[arrIdx:end+1], sArr1[sArr1Idx:])
            return
        } else if sArr1Idx >= len(sArr1) {
            // If all elements is taken from the first subarray
            // Copy the remaining elements in the second subarray into the sorted array
            copy(arr[arrIdx:end+1], sArr2[sArr2Idx:])
            return
        }

        if sArr1[sArr1Idx] < sArr2[sArr2Idx] {
            // If the lower subarray's current element is smaller than the higher one's
            // Increment the lower subarray current index
            arr[arrIdx] = sArr1[sArr1Idx]
            sArr1Idx++
        } else {
            // If the higher subarray's current element is smaller than the lower one's
            // Increment the higher subarray current index
            arr[arrIdx] = sArr2[sArr2Idx]
            sArr2Idx++
        }
        arrIdx++ // Increment the solved array index pointer
    }
}

func MergeSort(_arr []int) []int {
    // Make a copy of the input array
    var arr = make([]int, len(_arr))
    copy(arr, _arr)

    var arrLen = len(arr) // Get the length of the array
    // Start merging subarrays of size 1 and continue in powers of 2
    for saSize := 1; saSize < arrLen; saSize *= 2 {
        // Pick starting points for the different subarrays of the current size
        for saIdx := 0; saIdx < arrLen-1; saIdx += saSize*2 {
            var mid = min(saIdx+saSize, arrLen-1) // Get the starting index of the second subarray
            var end = min(saIdx+saSize*2-1, arrLen-1) // Get the ending index of the second subarray
            merge(arr, saIdx, mid, end) // Merge the subarrays
        }
    }
    return arr
}
