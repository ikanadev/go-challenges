package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Items []Flight
}

func (q *Queue) Pop() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue empty!")
	}
	last := q.Items[0]
	q.Items = q.Items[1:]
	return last, nil
}

func (q *Queue) Push(f Flight) {
	q.Items = append(q.Items, f)
}

func (q *Queue) Peek() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue empty!")
	}
	return q.Items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func goChall6() {
	fmt.Println("Go Queue Implementation")
}
