package main_test

import (
	"testing"

	"github.com/psbernardo/dockertest/infra/testingservice"
	"github.com/psbernardo/dockertest/infra/testingservice/mockapi"
	"github.com/psbernardo/dockertest/infra/testingservice/testsetup/loadtestdata"
	internal "github.com/psbernardo/dockertest/internal"
	"github.com/psbernardo/dockertest/internal/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// func TestMain(m *testing.M) {
// 	t := testsetup.SuiteTest{}
// 	t.SetupTestServices()
// 	log.Println("Do stuff BEFORE the tests!")
// 	exitVal := m.Run()
// 	log.Println("Do stuff AFTER the tests!")
// 	t.TearDownTestServices()
// 	os.Exit(exitVal)
// }

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
	)
	tu.require.Nil(err)

	usecase := internal.NewTestRest(tu.NewThirdPartyAPITestClient(), mariaDBTest)
	person, err := usecase.FetchAndCreate(3)
	tu.require.Nil(err)
	tu.require.Equal(&model.Person{
		ID:       3,
		Name:     "Patrick",
		LastName: "Bernardo",
		Age:      28,
	}, person)
}
