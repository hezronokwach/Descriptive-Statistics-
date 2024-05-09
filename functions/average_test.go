package mathSkills

import (
    "testing"
)

func TestAverage(t *testing.T) {
    // This function defines unit tests for the Average function

    // Define test cases with various input data and expected results
    tests := []struct {
        name string
        data []string
        want float64
    }{
        {name: "empty slice", data: []string{}, want: 0},
        {name: "single value", data: []string{"5"}, want: 5},
        {name: "multiple values", data: []string{"3", "5", "7"}, want: 5},
        {name: "with decimals", data: []string{"3.14", "2.72"}, want: 2.93},
    }

    // Loop through each test case
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Run the Average function with test data
            got, err := Average(tt.data)
            if err != nil {
                t.Errorf("Average(%v) error = %v, want err = nil", tt.data, err)
                return
            }

            // Check for errors and compare the obtained average with the expected value
            if got != tt.want {
                t.Errorf("Average(%v) = %v, want %v", tt.data, got, tt.want)
            }
        })
    }
}
