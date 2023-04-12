package mockrequest

import "net/http"

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
}
