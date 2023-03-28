package main_test

import (
	"testing"

	"github.com/psbernardo/dockertest/infra/testingservice"
	internal "github.com/psbernardo/dockertest/internal"
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
	tu.require.Nil(tu.SetupTestServices())

}

func (tu *MainTestSuite) TearDownTest() {
	tu.require.Nil(tu.TearDownTestServices())
}

func (tu *MainTestSuite) TestConsumeRestAPIFromDocker() {
	businessLogic := internal.NewTestRest(tu.NewThirdPartyAPIClient())
	statusCode, err := businessLogic.Get()
	tu.require.Nil(err)
	tu.require.Equal(200, statusCode)
}
