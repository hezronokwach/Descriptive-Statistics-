package mathSkills

import "testing"

func TestVariance(t *testing.T) {
    // This function defines unit tests for the Variance function

    // Define test cases with various input data and expected results
    tests := []struct {
        name string
        data []string
        want float64
    }{
        {name: "empty slice", data: []string{}, want: 0},
        {name: "single value", data: []string{"5"}, want: 0},                     // Variance is 0 for a single value
        {name: "multiple values", data: []string{"3", "5", "7"}, want: 3},        // Expected variance
        {name: "with decimals", data: []string{"3.14", "2.72", "1.59"}, want: 0}, // Expected variance (rounded to integer)
    }

    // Loop through each test case
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Run the Variance function with test data
            got, err := Variance(tt.data)
            if err != nil {
                t.Errorf("Variance(%v) error = %v, want err = nil", tt.data, err)
                return
            }

            // Check for errors and compare the obtained variance with the expected value
            if got != (tt.want) {
                t.Errorf("Variance(%v) = %v, want %v", tt.data, got, tt.want)
            }
        })
    }
}
