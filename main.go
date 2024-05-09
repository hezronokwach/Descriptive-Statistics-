package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "mathSkills/functions"
)

func main() {
    // Check for missing filename
    if len(os.Args) < 2 {
        fmt.Println("Missing data file")
        return
    }

    // Open and check file
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Error accessing data file: %v\n", err)
        return
    }
    defer file.Close()

    // Check for empty file
    stat, err := file.Stat()
    if err != nil || stat.Size() == 0 {
        fmt.Println("Invalid data file")
        return
    }

    // Read data from file
    scanner := bufio.NewScanner(file)
    data := []string{}
    for scanner.Scan() {
        data = append(data, scanner.Text())
    }

    // Calculate and print statistics
    average, err := mathSkills.Average(data)
    if err != nil {
        fmt.Println("Error calculating statistics")
        return
    }
    fmt.Println("Average:", int(math.Round(average))) 

    median, err := mathSkills.Median(data)
    if err != nil {
        fmt.Println("Error calculating statistics")
        return
    }
    fmt.Println("Median:", int(math.Round(median))) 

    variance, err := mathSkills.Variance(data)
    if err != nil {
        fmt.Println("Error calculating statistics")
        return
    }
    fmt.Println("Variance:", int(math.Round(variance))) 
    std, err := mathSkills.StandardDeviation(data)
    if err != nil {
        fmt.Println("Error calculating statistics")
        return
    }
    fmt.Println("Standard Deviation:", int(math.Round(std))) 
}
