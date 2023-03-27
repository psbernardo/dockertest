package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hi patrick")
}

type Rest struct {
	BaseURL string
}

func (r *Rest) Get() (int, error) {
	response, err := http.Get(fmt.Sprintf("%s/health", r.BaseURL))

	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}
