# Trimmed Mean Calculation of a Population in Go

This Go program generates a population of random integers and floats, and calculates the trimmed mean of the population. It utilizes a trimed mean package I created previously.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Code Explanation](#code-explanation)
- [Comparison](#comparison)
- [Summary](#summary)

## Introduction

The main purpose of this program is to showcase the usage of packages in Go. I have created a package that calculates the trimmed mean of a given slice, which can be found [here](https://github.com/hamodikk/trimmedmean). This program requires the package to function, as I import and use the package in calculating the trimmed mean on the generated population.

Trimmed mean is a statistical measure that removes the outliers by a specific percentage before calculating the mean.

## Features

- Generates a population of integers and floats.
- Imports the trimmedmean package.
- Calculates the trimmed mean of the population using trimmedmean package.
- Logs the population, population size, and trimmed mean of the population.
- Allows user input for population size.
- Defaults population size if executable is run.

## Requirements

Ensure that Go is installed. [Install go here](https://go.dev/doc/install).

This program requires the installation of trimmedmean package. The repository for this package can be found [here](https://github.com/hamodikk/trimmedmean).

To install the package, open terminal and run:
```bash
go get github.com/hamodikk/trimmedmean
```

This will install the package, allowing you to import it into the program without issues.

It is also important to note than the package needs to be added in `go.mod` file as follows:
```go
require (
    github.com/hamodikk/trimmedmean v1.0.0
)
```

However, `go tidy` will automatically add this line in your `go.mod` file if you included the package in the import section of `main.go`

## Installation

To use the package, first clone the repository in your directory:
```bash
git clone https://github.com/hamodikk/poptrimmean.git
```
Change directory to the repository:
```bash
cd <path/to/repository/directory>
```

## Usage

You can run the code either from the [executable](poptrimmean.exe) or by running the following in your terminal.
Running the executable will default the population size to 200. Running the program using the following command will give you the option to choose the population size.
```bash
go run .\main.go [population_size]
```

Example command to calculate the trimmed mean of a randomly generated population of 500 numbers:
```bash
go run .\main.go 500
```

## Code Explanation

There are two parts to this program. There is a helper function that generates the population of random integers and floats, and the main function that calculates the trimmed mean for the generated population. The main function includes error handling, conversion of user input into an integer, calling the population generation function, logging (which includes converting the dataset to a format that is usable for cross-checking the results) and finally, the trimmed mean calculation using the package trimmedmean.

- Helper function that generates the population
```go
func generatePopulation(populationSize int) []interface{} {
	// seed for random number generator
	rand.Seed(int64(42))

	// generate population
	population := make([]interface{}, populationSize)
	for i := 0; i < populationSize; i++ {
		if rand.Intn(2) == 0 {
			population[i] = rand.Intn(100)
		} else {
			population[i] = rand.Float64() * 100
		}
	}

	// return population
	return population
}
```

- Handling the user input and potential errors.
```go
    // accept the population size from the user
	if len(os.Args) >= 3 {
		log.Fatal("Usage: poptrimmean <population size>")
	}

	// convert the input to an integer
	var populationSize int
	var err error
	if len(os.Args) == 2 {
		populationSize, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		populationSize = 200
	}
```

- Log setup.
```go
	logFile, err := os.OpenFile("trimmedmean.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
```

- Conversion of []interface{} dataset to []string for log reporting. Logging directly from []interface{} results in a space separated list of numbers, which we cannot feed into a web based trimmed mean calculator to check our results. This conversion allows us to log the generated dataset as a comma separated list.
```go
	var populationStr []string
	for _, v := range population {
		populationStr = append(populationStr, fmt.Sprintf("%v", v))
	}

	// Log the population with a comma separator
	log.Printf("Population: %s\n", strings.Join(populationStr, ", "))
```

- Calculate and log the trimmed mean. First example is the symmetrical trimming and the second one is the asymmetrical trimming.
```go
	// calculate trimmed mean (symmetrically, 5%)
	populationTrimmedMeanSym, _ := trimmedmean.TrimmedMean(population, 0.05)
	log.Printf("Trimmed mean (5%%): %.4f\n", populationTrimmedMeanSym)

	// calculate trimmed mean (asymmetrically, 10% lower, 5% upper)
	populationTrimmedMeanAsym, _ := trimmedmean.TrimmedMean(population, 0.10, 0.05)
	log.Printf("Trimmed mean (10%% lower, 5%% upper): %.4f\n", populationTrimmedMeanAsym)
```

## Observations

- Random population generation uses a fixed seed for reproducability.
- Default population size for `poptrimmean.exe` is 200, which can be changed from within the code and updated using `go build`.
- The trim amounts can be changed within the code, but by default will give two results, one with 5% trimming symmetrically and the other with 10% lower and 5% upper trim asymmetrically for any given population size.

## Comparison

I compared the results of this program with that of a code ran in R. I copied the generated population from running the executable, and created a dataset in R before running the trimmed mean code.

### Code Length

Following is the code I used in R to calculate the trimmed mean of the population with 5% symmetrical trim.
```R
data <- c([insert_generated_population_here])

# Calculate the trimmed mean with a 5% symmetrical trim
trimmed_mean <- mean(data, trim = 0.05)

# Print the result
print(trimmed_mean)
```

Comparing this code length to the Go program, we can see that we can accomplish the same purpose with a lot less lines. It is important to note that the R code does not have a population generator. It also utilizes built in function `mean`, but even then, the code looks a lot more concise in comparison to the Go functions.

### Accuracy

Both Go and R program returned 49.9952 as the trimmed mean for the same population, so the accuracy of both programs are the same, although the calculations performed are not complex.

### Efficiency

No performance metrics were measured for either program, but can be added to compare the execution times.

## Summary

This program successfully generates a population of random integer and float numbers. It then trims this population symmetrically and asymmetrically and logs the results for comparison and cross-check. The results of this program are comparable to a similar function ran in R.