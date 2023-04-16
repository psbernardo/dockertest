package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	internal "github.com/psbernardo/dockertest/internal"
)

func main() {
	fmt.Println("patrick")

}

type handler struct {
	useCase *internal.UseCase
}

func NewHanlder(useCase *internal.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) CreatePerson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	person, err := h.useCase.FetchAndCreate(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, person)
}
