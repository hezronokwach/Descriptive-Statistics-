package mathSkills

import "testing"

func TestStandardDeviation(t *testing.T) {
  // This function defines unit tests for the StandardDeviation function

  // Define test cases with various input data and expected results
  tests := []struct {
    name string
    data []string
    want float64
  }{
    {name: "empty slice", data: []string{}, want: 0},
    {name: "single value", data: []string{"5"}, want: 0},                     // Standard deviation is 0 for a single value
    {name: "multiple values", data: []string{"3", "5", "7"}, want: 1.632993161855452},
    {name: "with decimals", data: []string{"3.14", "2.72", "1.59"}, want: 0.6545397025560957},
  }

  // Loop through each test case
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      // Run the StandardDeviation function with test data
      got, err := StandardDeviation(tt.data)
      if err != nil {
        t.Errorf("StandardDeviation(%v) error = %v, want err = nil", tt.data, err)
        return
      }

      // Check for errors and compare the obtained standard deviation with the expected value
      if got != tt.want {
        t.Errorf("StandardDeviation(%v) = %v, want %v", tt.data, got, tt.want)
      }
    })
  }
}
