package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}
type ByPrice []Flight

func (bp ByPrice) Len() int           { return len(bp) }
func (bp ByPrice) Swap(i, j int)      { bp[i], bp[j] = bp[j], bp[i] }
func (bp ByPrice) Less(i, j int) bool { return bp[i].Price < bp[j].Price }

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	sort.Sort(ByPrice(flights))
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("%d ", flight.Price)
	}
	fmt.Printf("\n")
}

func goChall3() {
	// an empty slice of flights
	var flights []Flight = []Flight{
		{Price: 32}, {Price: 34}, {Price: 20}, {Price: 44}, {Price: 24}, {Price: 0}, {Price: 34}, {Price: 12}, {Price: 65}, {Price: 120}, {Price: -5}, {Price: 1}, {Price: 35},
	}
	fmt.Println("Original")
	printFlights(flights)
	sortedList := SortByPrice(flights)
	fmt.Println("Sorted")
	printFlights(sortedList)
}
