package mathSkills

import (
    "fmt"
    "math"
    "strconv"
)

// StandardDeviation calculates the standard deviation of a set of numbers represented as strings
func StandardDeviation(slice []string) (float64, error) {
    if len(slice) == 0 {
        return 0, nil // Handle empty slice gracefully
    }
    variance, err := calculateVariance(slice)
    if err != nil {
        return 0, err
    }
    return math.Sqrt(variance), nil // Standard deviation is square root of variance
}

// calculateVariance calculates the variance of a set of numbers represented as strings
func calculateVariance(slice2 []string) (float64, error) {
    var result float64 // Stores the sum of squared deviations from the mean
    mean, err := Average(slice2)
    if err != nil {
        return 0, err 
    }
    for _, value := range slice2 {
        number, err := strconv.ParseFloat(value, 64)
        if err != nil {
            return 0, fmt.Errorf("invalid number in slice: %v", value) 
        }
        deviation := number - mean // Calculate deviation from the mean
        result += deviation * deviation // Square the deviation and add to the sum
    }
    deviation := result / float64(len(slice2)) // Calculate average squared deviation (variance)
    return deviation, nil
}
