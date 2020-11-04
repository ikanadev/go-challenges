package main

import (
	"errors"
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	var dev Developer
	strName, okName := name.(string)
	intAge, okAge := age.(int)
	if !okName || !okAge {
		checkErr(errors.New("cant convert"))
	}
	dev.Age = intAge
	dev.Name = strName
	return dev
}

func goChall1() {
	var name interface{} = "Elliot"
	var age interface{} = 26

	dynamicDev := GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
}
