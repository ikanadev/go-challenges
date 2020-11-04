package main

import "fmt"

func FilterUnique(developers []Developer) []string {
	devNames := make([]string, 0)
	for _, dev := range developers {
		finded := false
		for _, devName := range devNames {
			if dev.Name == devName {
				finded = true
				break
			}
		}
		if !finded {
			devNames = append(devNames, dev.Name)
		}
	}
	return devNames
}

func goChall4() {
	devs := []Developer{
		{Name: "Elliot"},
		{Name: "Alan"},
		{Name: "Jennifer"},
		{Name: "Graham"},
		{Name: "Paul"},
		{Name: "Alan"},
	}
	fmt.Println(FilterUnique(devs))
}
