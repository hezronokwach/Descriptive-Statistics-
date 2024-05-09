package mathSkills

import "testing"

func TestMedian(t *testing.T) {
    // This function defines unit tests for the Median function

    // Define test cases with various input data and expected results
    tests := []struct {
        name string
        data []string
        want float64
    }{
        {name: "empty slice", data: []string{}, want: 0},
        {name: "single value", data: []string{"5"}, want: 5},
        {name: "even elements", data: []string{"3", "5", "7"}, want: 5},
        {name: "odd elements", data: []string{"3", "5", "7", "10"}, want: 6},
        {name: "with decimals", data: []string{"3.14", "2.72", "1.59"}, want: 2.72},
    }

    // Loop through each test case
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Run the Median function with test data
            got, err := Median(tt.data)
            if err != nil {
                t.Errorf("Median(%v) error = %v, want err = nil", tt.data, err)
                return
            }

            // Check for errors and compare the obtained median with the expected value
            if got != tt.want {
                t.Errorf("Median(%v) = %v, want %v", tt.data, got, tt.want)
            }
        })
    }
}
