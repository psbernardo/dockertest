package mockrequest

import "net/http"

var MockRequestLocusList = []*MockRequest{
	NewMockRequest("/locus/task/3").
		HttpMethod("GET").
		ResponseCode(http.StatusOK).
		ResponseString(`{
		"taskId":3,
		"deliveryDate":"2023-05-03",
		}`),

	NewMockRequest("/locus/task/4").
		HttpMethod("GET").
		ResponseCode(http.StatusOK).
		ResponseString(`{
		"taskId":4,
		"deliveryDate":"2023-05-03",
		}`),
}
