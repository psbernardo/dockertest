package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi patrick")
}

type TestAPI interface {
	Get() (int, error)
}

type TestRest struct {
	BaseURL string
	TestAPI TestAPI
}

func NewTestRest(testAPI TestAPI) *TestRest {
	return &TestRest{
		TestAPI: testAPI,
	}
}

func (r *TestRest) Get() (int, error) {
	return r.TestAPI.Get()
}
