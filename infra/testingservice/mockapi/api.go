package mockapi

import (
	"fmt"
	"log"
	"net/http"
)

type testMockAPIServer struct {
	router   map[string]*MockRequest
	Error    error
	running  bool
	httpPort string
}

var (
	mockAPI *testMockAPIServer
)

func NewMockAPIServer() *testMockAPIServer {

	if mockAPI == nil {
		mockAPI = &testMockAPIServer{
			router:   make(map[string]*MockRequest),
			httpPort: "8000",
		}

	}

	return mockAPI
}

func (m *testMockAPIServer) LoadDefaultMockDataTest() *testMockAPIServer {
	if err := loadMockRequestList(m,
		// Add here all mock request
		MockRequestPersonList,
		healthCheck,
	); err != nil {
		m.Error = err
	}
	return m
}

func (m *testMockAPIServer) LoadMockData(mockRequest ...MockRequestList) *testMockAPIServer {
	if err := loadMockRequestList(m,
		// Add here all mock request
		MockRequestPersonList,
		healthCheck,
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

	for !m.HealthCheck() {
		log.Println("health check test")
		http.HandleFunc("/", mockAPI.routerHndleFunc())
		go func() {
			mockAPI.running = true
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", mockAPI.httpPort), nil))
		}()
	}
	return nil
}

func (m *testMockAPIServer) HealthCheck() bool {
	response, err := http.Get(fmt.Sprintf("%s/health", fmt.Sprintf("http://localhost:%s", mockAPI.httpPort)))

	if err != nil {
		return false
	}

	return response.StatusCode == http.StatusOK
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
