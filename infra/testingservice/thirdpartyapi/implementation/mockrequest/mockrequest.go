package mockrequest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type MockRequest struct {
	httpMethod string
	request    string
	statusCode int

	response     string
	responseData interface{}
}

func NewMockRequest(requestURL string) *MockRequest {
	return &MockRequest{
		request:    requestURL,
		statusCode: http.StatusOK,
	}
}

func (m *MockRequest) GetKey() string {
	return fmt.Sprintf("%s%s", m.httpMethod, m.request)
}

func (m *MockRequest) Request(requestURL string) *MockRequest {
	m.request = requestURL
	return m
}

func (m *MockRequest) HttpMethod(method string) *MockRequest {
	m.httpMethod = method
	return m
}

func (m *MockRequest) ResponseString(response string) *MockRequest {
	m.response = response
	return m
}

func (m *MockRequest) ResponseCode(statusCode int) *MockRequest {
	m.statusCode = statusCode
	return m
}

func (m *MockRequest) Response(data interface{}) *MockRequest {
	m.responseData = data
	return m
}

func (m *MockRequest) SetResponseWriter(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(m.statusCode)
	if len(strings.TrimSpace(m.response)) > 0 {
		w.Write([]byte(m.response))
		return
	}

	json.NewEncoder(w).Encode(m.responseData)
}

type MockRequestList []*MockRequest
