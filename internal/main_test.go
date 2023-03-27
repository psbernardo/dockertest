package main_test

import (
	"testing"

	internal "github.com/psbernardo/dockertest/internal"
	"github.com/psbernardo/dockertest/test/testsetup"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
	testsetup.SuiteTest
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
	r := internal.Rest{
		BaseURL: tu.TestService.ThirdPartyAPIHost,
	}

	statusCode, err := r.Get()
	tu.require.Nil(err)
	tu.require.Equal(200, statusCode)
}
