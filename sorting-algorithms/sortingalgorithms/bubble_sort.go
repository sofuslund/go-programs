package sortingalgorithms

// Bubble sort sorts an array by continuously swapping adjacent elements if they are in the wrong order
// It's done if it can iterate over the whole array without the need for swapping i.e. arr[n] < arr[n+1]
func BubbleSort(_arr []int) ([]int) {
    arr := make([]int, len(_arr))
    copy(arr, _arr)
    // Do swapping iterations through the array while it is not sorted (done)
    var isDone = false
    for !isDone {
        isDone = true
        // One by one move the index at which two elements possibly will be swapped
        for i := 0; i < len(arr)-1; i++ {
            // Swap the elements if they are in wrong order
            if arr[i] > arr[i+1] {
                arr[i], arr[i+1] = arr[i+1], arr[i]
                // If there is need for swapping the algorithm is not done
                isDone = false
            }
        }
    }
    return arr
}
