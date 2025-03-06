func main() {
	// Declare and initialize an array of integers
	arr := []int{1, 2, 3, 4, 5}

	// Print the sum of all elements in the array
	fmt.Println(sum(arr))
}

func sum(arr []int) int {
	var total int

	for _, num := range arr {
		total += num
	}

	return total
}
