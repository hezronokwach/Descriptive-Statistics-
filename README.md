
  

# Statistics Calculator
![enter image description here](https://cdn.pixabay.com/photo/2018/05/18/14/36/statistics-3411311_1280.jpg)

  

## Description:

  

This program calculates basic statistics (average, median, variance, and standard deviation) for a set of numerical data provided in a text file.

  

## Features:

  

- Calculates average, median, variance, and standard deviation of numerical data.

- Handles decimal values.

- Checks for empty files and prevents errors during calculations.

  
  

## Data Structure:

  

The program consists of the following files and directories:

```

Project Directory/

|

|-- functions/ # Directory containing function files

| |-- <functionfile.go> # Funtion files format

| +-- <functionfile_test.go> # Test files format

|

|-- main.go # Main program file

|-- data.txt # Example data file

|-- README.md # Readme file

+-- go.mod # Go module file

```

  
  

## How to Use:

  

1. Save the program and the data file in the same directory.

2. Modify the `main.go` function to specify the filename of your data file. (e.g., `fileName := "data.txt"`)

3. Open a terminal or command prompt.

4. Navigate to the directory where the program is saved.

5. Run the program by executing the following command:

  

~~~shell

go  run  .  main.go <file.text>

~~~

  

## Example Data File:

  

11

3

5

13

## Output:

  

The program will print the calculated average, median, variance, and standard deviation of the data in the file.

  

~~~ shell

Average:  8

Median:  8

Variance:  17

Standard  Deviation:  4

~~~

  

## Dependencies:

  

- Go programming language and its standard library.

- Go module

  

### Creating Go Module:

  

1. Open a terminal or command prompt.

2. Navigate to the directory where the program is saved.

3. Run the following command to initialize the Go module:

~~~shell

go  mod  init <module_name>

~~~

  

## Testing:

  

The code includes unit tests for the functions to ensure the functionality of the statistics calculations.

  

`average_test.go`

`median_test.go`

`variance_test.go`

`standarddeviation_test.go`

  

## Contributing:

  

Feel free to fork the repository and contribute improvements to the code.
