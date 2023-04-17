package main

import (
	"net/http"
	"testing"

	"github.com/psbernardo/dockertest/infra/testingservice"
	"github.com/psbernardo/dockertest/infra/testingservice/mockapi"
	"github.com/psbernardo/dockertest/infra/testingservice/testsetup/loadtestdata"
	internal "github.com/psbernardo/dockertest/internal"
	"github.com/psbernardo/dockertest/internal/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
		loadtestdata.WithNewPerson(), // load test data 1
	)
	tu.require.Nil(err)
	usecase := internal.NewUseCase(tu.NewThirdPartyAPIClient(), mariaDBTest)
	handler := NewHanlder(usecase)

	RunAllHTTPTest(tu.T(),

		NewHttpTest("Get person id 4").
			withHTTPMethod(http.MethodPost).
			withHandler(handler.CreatePerson).
			withExpectedStatusCode(http.StatusCreated).
			withPathParameters(map[string]string{"id": "4"}).
			withExpectedResponse(model.Person{
				ID:       4,
				Name:     "Bryan",
				LastName: "Bernardo",
				Age:      23,
			}),

		NewHttpTest("Get person id 3").
			withHTTPMethod(http.MethodPost).
			withHandler(handler.CreatePerson).
			withExpectedStatusCode(http.StatusCreated).
			withPathParameters(map[string]string{"id": "3"}).
			withExpectedResponse(model.Person{
				ID:       3,
				Name:     "Patrick",
				LastName: "Bernardo",
				Age:      28,
			}),

		NewHttpTest("Get person id 5").
			withHTTPMethod(http.MethodPost).
			withHandler(handler.CreatePerson).
			withExpectedStatusCode(http.StatusCreated).
			withPathParameters(map[string]string{"id": "5"}).
			withExpectedResponse(model.Person{
				ID:       5,
				Name:     "Pearson",
				LastName: "Specter",
				Age:      30,
			}),
	)

}
