package usecase_test

import (
	"context"
	"testing"

	"github.com/psbernardo/dockertest/infra/testingservice"
	"github.com/psbernardo/dockertest/infra/testingservice/mockapi"
	"github.com/psbernardo/dockertest/infra/testingservice/testsetup/loadtestdata"
	internal "github.com/psbernardo/dockertest/internal"
	"github.com/psbernardo/dockertest/internal/mocking"
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
	require     require.Assertions
	mockAPIPort int
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

func (tu *MainTestSuite) SetupTest() {
	tu.require = *tu.Require()
	mockAPIPort, err := mockapi.NewMockAPIServer().LoadDefaultMockDataTest().Run()
	tu.require.Nil(err)
	tu.mockAPIPort = mockAPIPort
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

	// test catch up

	usecase := internal.NewUseCase(tu.NewThirdPartyAPIClient(tu.mockAPIPort), mariaDBTest)
	person, err := usecase.FetchAndCreate(context.Background(), 3)
	tu.require.Nil(err)
	tu.require.Equal(&model.Person{
		ID:       3,
		Name:     "Patrick",
		LastName: "Bernardo",
		Age:      28,
	}, person)
}

func TestUseCase(t *testing.T) {
	mockingRepo := &mocking.MockRepository{}
	mockThirdParty := &mocking.MockThirdPartyAPI{}

	mockThirdParty.GetPersonFn = func(ctx context.Context, personId int) (*model.Person, error) {
		return &model.Person{
			ID:       3,
			Name:     "Patrick",
			LastName: "Bernardo",
			Age:      28,
		}, nil
	}

	mockingRepo.CreatePersonFn = func(ctx context.Context, person *model.Person) (*model.Person, error) {
		return &model.Person{
			ID:       3,
			Name:     "Patrick",
			LastName: "Bernardo",
			Age:      28,
		}, nil
	}

	usecase := internal.NewUseCase(mockThirdParty, mockingRepo)

	person, _ := usecase.FetchAndCreate(context.Background(), 3)

	expected := &model.Person{
		ID:       3,
		Name:     "Patrick",
		LastName: "Bernardo",
		Age:      28,
	}
	if *person != *expected {
		t.Error("not equal")
	}

}
