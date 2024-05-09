package mathSkills

import (
    "math"
)

// Variance calculates the variance of a set of numbers represented as strings
// (assuming StandardDeviation function handles conversion)
func Variance(slice []string) (float64, error) {
    stdDev, err := StandardDeviation(slice)
    if err != nil {
        return 0, err
    }
    variance := stdDev * stdDev
    roundedVariance := math.Round(variance) 
    return roundedVariance, nil
}
