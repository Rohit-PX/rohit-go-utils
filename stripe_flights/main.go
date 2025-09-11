package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	data := "UK:US:FedEx:4,UK:FR:Jet1:2,UK:US:FedEx:3,US:UK:RyanAir:8,CA:UK:CanadaAir:8"
	src := "UK"
	dest := "US"
	minCost, err := GetFlightCost(data, src, dest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Min cost between %s and %s is %d\n", src, dest, minCost)

	lookup, err := CreateFlightLookup(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Min cost in O(1) between %s and %s is %d\n", src, dest, lookup[src][dest])

}

func GetFlightCost(data, src, dest string) (int, error) {
	parsedData := strings.Split(data, ",")
	if len(parsedData) < 1 {
		return -1, fmt.Errorf("empty input")
	}
	minCost := -1
	for _, flight := range parsedData {
		fl := strings.Split(flight, ":")
		if len(fl) != 4 || fl[0] == "" || fl[1] == "" || fl[2] == "" || fl[3] == "" {
			return -1, fmt.Errorf("invalid flight data")
		}
		currSrc := fl[0]
		currDest := fl[1]
		currCost, err := strconv.Atoi(fl[3])
		if err != nil {
			return -1, fmt.Errorf("invalid flight data")
		}

		if currSrc == src && currDest == dest {
			if minCost == -1 || currCost < minCost {
				minCost = currCost
			}
		}
	}

	if minCost == -1 {
		return minCost, fmt.Errorf("no direct flights found between: %s and %s", src, dest)
	}
	return minCost, nil
}

func CreateFlightLookup(data string) (map[string]map[string]int, error) {
	minLookup := make(map[string]map[string]int)
	parsedData := strings.Split(data, ",")
	if len(parsedData) < 1 {
		return nil, fmt.Errorf("empty initial flight data input")
	}

	for _, fl := range parsedData {
		flight := strings.Split(fl, ":")
		if len(flight) != 4 || flight[0] == "" || flight[1] == "" || flight[2] == "" || flight[3] == "" {
			return nil, fmt.Errorf("invalid flight data")
		}
		currSrc := flight[0]
		currDest := flight[1]
		currCost, err := strconv.Atoi(flight[3])
		if err != nil {
			return nil, fmt.Errorf("invalid flight cost data")
		}

		if _, ok := minLookup[currSrc][currDest]; !ok {
			minLookup[currSrc] = map[string]int{currDest: currCost}
		} else {
			if currCost < minLookup[currSrc][currDest] {
				minLookup[currSrc][currDest] = currCost
			}
		}
	}
	return minLookup, nil

}
