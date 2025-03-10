  func main() {
    // Initialize an array with some numbers
    var nums = []int{1, 2, 3, 4, 5}

    // Declare a variable to store the sum
    var total int

    // Iterate over the array and calculate the sum
    for _, num := range nums {
      total += num
    }

    // Print the sum
    fmt.Println(total)
  }