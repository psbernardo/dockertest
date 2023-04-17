package mockapi

import "net/http"

var healthCheck = []*MockRequest{
	NewMockRequest("/health").
		HttpMethod("GET").
		ResponseCode(http.StatusOK).
		ResponseString(""),
}

var MockRequestPersonList = []*MockRequest{
	NewMockRequest("/person/3").
		HttpMethod("GET").
		ResponseCode(http.StatusOK).
		ResponseString(`{
		"id":3,
		"name":"Patrick",
		"lastName":"Bernardo",
		"age":28
		}`),

	NewMockRequest("/person/4").
		HttpMethod("GET").
		ResponseCode(http.StatusOK).
		ResponseString(`{
		"id":4,
		"name":"Bryan",
		"lastName":"Bernardo",
		"age":23
		}`),

	NewMockRequest("/person/5").
		HttpMethod("GET").
		ResponseCode(http.StatusOK).
		ResponseString(`{
		"id":5,
		"name":"Pearson",
		"lastName":"Specter",
		"age":30
		}`),
}
