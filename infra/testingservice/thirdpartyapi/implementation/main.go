package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/psbernardo/dockertest/infra/testingservice/thirdpartyapi/implementation/mockrequest"
)

// Request Response  - will response if match specific request else return request url key not found

type MockRequestRouter struct {
	Router map[string]*mockrequest.MockRequest
}

func NewMockRequestRouter() (*MockRequestRouter, error) {
	MockRequestRouter := &MockRequestRouter{
		Router: make(map[string]*mockrequest.MockRequest),
	}

	if err := LoadMockRequestList(MockRequestRouter,
		// Add here all mock request
		//mockrequest.MockRequestPersonList,
		mockrequest.MockRequestLocusList,
	); err != nil {
		return nil, err
	}

	return MockRequestRouter, nil
}

func LoadMockRequestList(router *MockRequestRouter, mockRequest ...mockrequest.MockRequestList) error {
	for _, mockdataList := range mockRequest {
		for _, mockdata := range mockdataList {
			if _, ok := router.Router[mockdata.GetKey()]; !ok {
				router.Router[mockdata.GetKey()] = mockdata
			} else {
				return fmt.Errorf("duplicate request key: %s", mockdata.GetKey())
			}
		}
	}

	return nil

}

func main() {
	router, err := NewMockRequestRouter()
	if err != nil {
		log.Fatal(err.Error())
	}

	//all request will be handle here
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := fmt.Sprintf("%s%s", r.Method, r.RequestURI)
		if mockdata, ok := router.Router[key]; ok {
			mockdata.SetResponseWriter(w)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{
				"message": "500 - request key not found: %s"
			}`, key)))
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
