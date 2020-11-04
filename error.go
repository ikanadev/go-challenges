package main

import (
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		fmt.Printf("Fatal Error: %s\n", e.Error())
		os.Exit(1)
	}
}
