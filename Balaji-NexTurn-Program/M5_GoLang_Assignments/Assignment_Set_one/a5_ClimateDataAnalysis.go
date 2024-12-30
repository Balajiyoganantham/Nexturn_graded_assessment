// Case Study:5
// You are tasked with analyzing climate data from multiple cities. The data includes the
// city name, average temperature (°C), and rainfall (mm).
// 1. Data Input: Create a slice of structs to store data for each city. Input data can be
// hardcoded or taken from the user.
// 2. Highest and Lowest Temperature: Write functions to find the city with the
// highest and lowest average temperatures. Use conditions for comparison.
// 3. Average Rainfall: Calculate the average rainfall across all cities using loops. Use
// type casting if necessary.
// 4. Filter Cities by Rainfall: Use loops to display cities with rainfall above a certain
// threshold. Prompt the user to enter the threshold value.
// 5. Search by City Name: Allow users to search for a city by name and display its
// data.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type City struct {
	Name        string
	AverageTemp float64
	Rainfall    float64
}

var climateData = []City{
	{"New York", 13.0, 1200.5},
	{"Los Angeles", 18.5, 384.6},
	{"Chicago", 10.0, 990.7},
	{"Houston", 20.0, 1319.8},
	{"Phoenix", 23.0, 210.3},
}

func cityWithHighestTemp(data []City) City {
	highest := data[0]
	for _, city := range data {
		if city.AverageTemp > highest.AverageTemp {
			highest = city
		}
	}
	return highest
}

func cityWithLowestTemp(data []City) City {
	lowest := data[0]
	for _, city := range data {
		if city.AverageTemp < lowest.AverageTemp {
			lowest = city
		}
	}
	return lowest
}

func averageRainfall(data []City) float64 {
	totalRainfall := 0.0
	for _, city := range data {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(data))
}

func filterCitiesByRainfall(data []City, threshold float64) []City {
	filteredCities := []City{}
	for _, city := range data {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

func searchCityByName(data []City, name string) (City, bool) {
	for _, city := range data {
		if strings.EqualFold(city.Name, name) {
			return city, true
		}
	}
	return City{}, false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Climate Data Analysis")

	// Highest and Lowest Temperature
	highest := cityWithHighestTemp(climateData)
	lowest := cityWithLowestTemp(climateData)
	fmt.Printf("\nCity with the highest temperature: %s (%.2f°C)\n", highest.Name, highest.AverageTemp)
	fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowest.Name, lowest.AverageTemp)

	// Average Rainfall
	avgRainfall := averageRainfall(climateData)
	fmt.Printf("\nAverage Rainfall: %.2f mm\n", avgRainfall)

	// Filter Cities by Rainfall
	fmt.Print("\nEnter rainfall threshold (mm) to filter cities: ")
	scanner.Scan()
	thresholdInput := scanner.Text()
	var threshold float64
	fmt.Sscanf(thresholdInput, "%f", &threshold)
	filteredCities := filterCitiesByRainfall(climateData, threshold)
	if len(filteredCities) > 0 {
		fmt.Println("\nCities with rainfall above threshold:")
		for _, city := range filteredCities {
			fmt.Printf("%s: %.2f mm\n", city.Name, city.Rainfall)
		}
	} else {
		fmt.Println("\nNo cities found with rainfall above the threshold.")
	}

	// Search by City Name
	fmt.Print("\nEnter city name to search: ")
	scanner.Scan()
	cityName := scanner.Text()
	city, found := searchCityByName(climateData, cityName)
	if found {
		fmt.Printf("\nCity found: %s\nTemperature: %.2f°C\nRainfall: %.2f mm\n", city.Name, city.AverageTemp, city.Rainfall)
	} else {
		fmt.Println("\nCity not found in the database.")
	}
}
