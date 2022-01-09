package sortingalgorithms

// Insertion sort works similar to the way you sort playing cards in your hands
// The algorithm iterates over an array and checks if all elements is larger than their preceding ones
// If it finds an element that is not larger than it's preceding one it moves the element down until it is
func InsertionSort(_arr []int) ([]int) {
    arr := make([]int, len(_arr))
    copy(arr, _arr)
    // Iterate over each array element
    for i := 1; i < len(arr); i++ {
        // Move the element down until the element before it is smaller than itself and the element after it is bigger than itself
        for j := i; j > 0; j-- {
            if arr[j] > arr[j-1] {
                break
            }
            arr[j], arr[j-1] = arr[j-1], arr[j]
        }
    }
    return arr
}
