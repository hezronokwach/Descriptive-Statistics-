package mathSkills

import (
    "fmt"
    "strconv"
)

// Average calculates the average of a set of numbers represented as strings
func Average(slice []string) (float64, error) {
    var acc float64 // Accumulator to store the sum of all numbers
    if len(slice) == 0 { // Check for empty slice
        return 0, nil // Handle empty slice gracefully
    }
    
    for _, value := range slice {
        number, err := strconv.ParseFloat(value, 64) // Convert string to float64
        if err != nil {
            return 0, fmt.Errorf("invalid number in slice: %v", value) // Handle invalid data
        }
        acc += number // Add the converted number to the accumulator
    }
    
    av := acc / float64(len(slice)) // Calculate average (sum divided by number of elements)
    return av, nil
}
