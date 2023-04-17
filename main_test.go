package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/psbernardo/dockertest/infra/testingservice"
	"github.com/psbernardo/dockertest/infra/testingservice/mockapi"
	"github.com/psbernardo/dockertest/infra/testingservice/testsetup/loadtestdata"
	internal "github.com/psbernardo/dockertest/internal"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	userJSON = "{\"id\":4,\"name\":\"Bryan\",\"lastName\":\"Bernardo\",\"age\":23}\n"
)

type MainTestSuite struct {
	suite.Suite
	testingservice.SuiteTest
	require require.Assertions
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

func (tu *MainTestSuite) SetupTest() {
	tu.require = *tu.Require()
	tu.require.Nil(mockapi.NewMockAPIServer().LoadDefaultMockDataTest().Run())
	tu.require.Nil(tu.SetupTestServices())

}

func (tu *MainTestSuite) TearDownTest() {
	tu.require.Nil(tu.TearDownTestServices())
}

func (tu *MainTestSuite) TestConsumeRestAPIFromDocker() {

	// if we need to load some data to database
	// before calling the test usecase
	mariaDBTest, err := tu.NewMariaDBTestClient(
		loadtestdata.WithNewPerson(),  // load test data 1
		loadtestdata.WithNewPerson2(), // load test data 2
		loadtestdata.WithNewPerson(),
	)
	tu.require.Nil(err)

	usecase := internal.NewUseCase(tu.NewThirdPartyAPITestClient(), mariaDBTest)

	// Setup test
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	handler := NewHanlder(usecase)

	err = handler.CreatePerson(c)
	tu.require.Nil(err)
	tu.require.Equal(http.StatusCreated, rec.Code)
	tu.require.Equal(userJSON, rec.Body.String())

}
