package main

import (
	"testing"

	"github.com/psbernardo/dockertest/internal/testsetup"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
	require     require.Assertions
	testService *testsetup.TestService
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

func (tu *MainTestSuite) SetupTest() {
	tu.require = *tu.Require()
	testservice, err := testsetup.NewTestService(
		testsetup.WithThirdPartyAPITest(),
	)
	tu.require.Nil(err)
	tu.testService = testservice

}

func (tu *MainTestSuite) TearDownTest() {
	err := tu.testService.TearDownServices()
	tu.require.Nil(err)
}

func (tu *MainTestSuite) TestConsumeRestAPIFromDocker() {
	r := Rest{
		baseURL: tu.testService.ThirdPartyAPIHost,
	}

	statusCode, err := r.Get()
	tu.require.Nil(err)
	tu.require.Equal(200, statusCode)
}
