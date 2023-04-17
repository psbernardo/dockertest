package mockapi

import (
	"fmt"
	"log"
	"net/http"
)

type testMockAPIServer struct {
	router  map[string]*MockRequest
	Error   error
	running bool
}

var (
	mockAPI *testMockAPIServer
)

func NewMockAPIServer() *testMockAPIServer {

	if mockAPI == nil {
		mockAPI = &testMockAPIServer{
			router: make(map[string]*MockRequest),
		}
	}

	return mockAPI
}

func (m *testMockAPIServer) LoadDefaultMockDataTest() *testMockAPIServer {
	if err := loadMockRequestList(m,
		// Add here all mock request
		MockRequestPersonList,
	); err != nil {
		m.Error = err
	}
	return m
}

func (m *testMockAPIServer) LoadMockData(mockRequest ...MockRequestList) *testMockAPIServer {
	if err := loadMockRequestList(m,
		// Add here all mock request
		MockRequestPersonList,
	); err != nil {
		m.Error = err
	}
	return m
}

func (m *testMockAPIServer) routerHndleFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := fmt.Sprintf("%s%s", r.Method, r.RequestURI)
		log.Println(key)
		if mockdata, ok := m.router[key]; ok {
			mockdata.SetResponseWriter(w)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{
				"message": "500 - request key not found: %s"
			}`, key)))
		}
	}
}

func (m *testMockAPIServer) Run() error {
	if m.Error != nil {
		return m.Error
	}
	if m.running {
		return nil
	}
	http.HandleFunc("/", m.routerHndleFunc())
	go func() {
		m.running = true
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	return nil
}

func loadMockRequestList(router *testMockAPIServer, mockRequest ...MockRequestList) error {
	for _, mockdataList := range mockRequest {
		for _, mockdata := range mockdataList {
			if _, ok := router.router[mockdata.GetKey()]; !ok {
				router.router[mockdata.GetKey()] = mockdata
			} else {
				return fmt.Errorf("duplicate request key: %s", mockdata.GetKey())
			}
		}
	}

	return nil

}
