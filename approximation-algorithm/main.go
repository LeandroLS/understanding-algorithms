package main

import (
	"fmt"
)

func intersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {

		if hash[e] {
			inter = append(inter, e)
		}
	}

	return
}

func main() {
	var finalStations []string

	cities := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	kOneCities := []string{"id", "nv", "ut"}
	kTwoCities := []string{"wa", "id", "mt"}
	kThreeCities := []string{"or", "nv", "ca"}
	kFourCities := []string{"nv", "ut"}
	kFiveCities := []string{"ca", "az"}

	radioStations := map[string][]string{
		"kone":   kOneCities,
		"ktwo":   kTwoCities,
		"kthree": kThreeCities,
		"kfour":  kFourCities,
		"kfive":  kFiveCities,
	}

	for len(cities) > 1 {
		bestStation := make(map[string][]string)
		var coveredCities []string
		for radioStation, city := range radioStations {
			covered := intersection(cities, city)

			if len(covered) > len(coveredCities) {
				bestStation[radioStation] = city
				coveredCities = append(coveredCities, covered...)
			}
		}
		for _, valCovered := range coveredCities {
			for i, valCity := range cities {
				if valCovered == valCity {
					cities = append(cities[:i], cities[i+1:]...)
				}
			}
		}
		for station := range bestStation {
			finalStations = append(finalStations, station)
		}
	}

	for v, i := range finalStations {
		fmt.Println(v, i)
	}
}
