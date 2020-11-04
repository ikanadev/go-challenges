package main

import (
	"errors"
	"fmt"
)

func GetMinMax(flights []Flight) (int, int, error) {
	if len(flights) == 0 {
		return 0, 0, errors.New("No flights provided")
	}
	min := flights[0].Price
	max := flights[0].Price
	for _, flight := range flights {
		if min > flight.Price {
			min = flight.Price
		}
		if max < flight.Price {
			max = flight.Price
		}
	}
	return min, max, nil
}

func goChall7() {
	fmt.Println("Getting the Minimum and Maximum Flight Prices")
}
