package mathSkills

import (
    "fmt"
    "sort"
    "strconv"
)

// Median calculates the median of a set of numbers represented as strings
func Median(slice []string) (float64, error) {
    var numbers []float64 // Stores converted numbers from the input slice
    for _, value := range slice {
        number, err := strconv.ParseFloat(value, 64) // Convert string to float64
        if err != nil {
            return 0, fmt.Errorf("invalid number in slice: %v", value) // Handle invalid data
        }
        numbers = append(numbers, number) // Add converted number to numbers slice
    }
    sort.Float64s(numbers) // Sort the numbers slice in ascending order

    length := len(numbers)
    if length == 0 { // Check for empty slice
        return 0, nil // Handle empty slice gracefully
    }

    var median float64
    if length%2 == 0 { // Even number of elements
        median = (numbers[length/2-1] + numbers[length/2]) / 2 // Median is average of middle two elements
    } else {
        median = numbers[length/2] // Odd number of elements, median is the middle element
    }

    return median, nil
}
