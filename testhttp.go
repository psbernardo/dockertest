package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type HttpTest struct {
	TestName   string
	HTTPMethod string
	Handler    func(c echo.Context) error

	ExpectedStatusCode int
	ExpectedResponse   interface{}

	RequestBody     interface{}
	PathParameters  map[string]string
	QueryParameters map[string]string
}

func NewHttpTest(testName string) *HttpTest {
	return &HttpTest{
		TestName:        testName,
		HTTPMethod:      http.MethodGet,
		PathParameters:  make(map[string]string),
		QueryParameters: make(map[string]string),
	}
}

func (h *HttpTest) withHTTPMethod(httpMethod string) *HttpTest {
	h.HTTPMethod = httpMethod
	return h
}

func (h *HttpTest) withHandler(hndler func(c echo.Context) error) *HttpTest {
	h.Handler = hndler
	return h
}

func (h *HttpTest) withExpectedStatusCode(statusCode int) *HttpTest {
	h.ExpectedStatusCode = statusCode
	return h
}

func (h *HttpTest) shoulResponse(response interface{}) *HttpTest {
	h.ExpectedResponse = response
	return h
}

func (h *HttpTest) withRequestBody(requestBody interface{}) *HttpTest {
	h.RequestBody = requestBody
	return h
}

func (h *HttpTest) withPathParameters(mapPath map[string]string) *HttpTest {
	for key, value := range mapPath {
		h.PathParameters[key] = value
	}
	return h
}

func (h *HttpTest) withPathQueryParameter(queryParam map[string]string) *HttpTest {
	for key, value := range queryParam {
		h.QueryParameters[key] = value
	}
	return h
}

func (h *HttpTest) requestBody() (io.Reader, error) {
	if h.RequestBody == nil {
		return nil, nil
	}
	jsonStr, err := json.Marshal(h.RequestBody)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonStr)
	return body, nil
}

func (h *HttpTest) getExpectedResponse() (string, error) {
	jsonStr, err := json.Marshal(h.ExpectedResponse)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

func (h *HttpTest) Test(t *testing.T) {
	t.Run(h.TestName, func(t *testing.T) {
		assert := assert.New(t)
		e := echo.New()
		requestBody, err := h.requestBody()
		assert.Nil(err)
		req := httptest.NewRequest(h.HTTPMethod, "/", requestBody)
		if len(h.QueryParameters) > 0 {
			q := req.URL.Query()
			for key, value := range h.QueryParameters {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()
		}

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		for key, value := range h.PathParameters {
			c.SetPath(fmt.Sprintf("/:%s", key))
			c.SetParamNames(key)
			c.SetParamValues(value)
		}
		assert.NotNil(h.Handler)
		err = h.Handler(c)
		assert.Nil(err)
		assert.Equal(h.ExpectedStatusCode, rec.Code)
		expectedJSON, err := h.getExpectedResponse()
		assert.Nil(err)
		expectedJSON = clearJson(expectedJSON)
		resBody := rec.Body.String()
		resBody = clearJson(resBody)
		assert.Equal(expectedJSON, resBody)

	})

}

func RunAllHTTPTest(t *testing.T, httpTestList ...*HttpTest) {
	for _, testCase := range httpTestList {
		testCase.Test(t)
	}
}

func clearJson(data string) string {
	data = strings.Trim(string(data), "\"")
	data = strings.TrimSuffix(data, "\n")
	return data
}
