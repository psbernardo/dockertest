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

func NewMockRequestRouter() MockRequestRouter {
	return MockRequestRouter{
		Router: make(map[string]*mockrequest.MockRequest),
	}
}

func main() {
	router := NewMockRequestRouter()
	for _, mockdata := range mockrequest.MockRequestList {
		if _, ok := router.Router[mockdata.GetKey()]; !ok {
			router.Router[mockdata.GetKey()] = mockdata
		} else {
			log.Fatalf("duplicate request key: %s", mockdata.GetKey())
		}
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := fmt.Sprintf("%s%s", r.Method, r.RequestURI)
		if mockdata, ok := router.Router[key]; ok {
			mockdata.SetResponseWriter(w)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{
				"message": "500 - request key not found: %s"
			}`, key)))
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
