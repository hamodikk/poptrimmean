package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/hamodikk/trimmedmean"
)

// Helper function to generate a population of random integers and floats
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

func main() {
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

	// set up logging (mainly to recover the population for cross-checking)
	logFile, err := os.OpenFile("trimmedmean.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// generate population
	population := generatePopulation(populationSize)

	// Convert the population []interface{} to []string for log reporting with a comma separator
	// This is so we can feed it onto a trimmed mean calculator online to compare results
	var populationStr []string
	for _, v := range population {
		populationStr = append(populationStr, fmt.Sprintf("%v", v))
	}

	// Log the population with a comma separator
	log.Printf("Population: %s\n", strings.Join(populationStr, ", "))
	log.Printf("Population size: %d\n", populationSize)

	// calculate trimmed mean (symmetrically, 5%)
	populationTrimmedMeanSym, _ := trimmedmean.TrimmedMean(population, 0.05)
	log.Printf("Trimmed mean (5%%): %.4f\n", populationTrimmedMeanSym)

	// calculate trimmed mean (asymmetrically, 10% lower, 5% upper)
	populationTrimmedMeanAsym, _ := trimmedmean.TrimmedMean(population, 0.10, 0.05)
	log.Printf("Trimmed mean (10%% lower, 5%% upper): %.4f\n", populationTrimmedMeanAsym)
}
