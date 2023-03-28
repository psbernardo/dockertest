package model

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
}
