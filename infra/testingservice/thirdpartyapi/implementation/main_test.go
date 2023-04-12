package main_test

import (
	"testing"

	api "github.com/psbernardo/dockertest/infra/testingservice/thirdpartyapi/implementation"
)

func TestDuplicateRequest(t *testing.T) {
	if _, err := api.NewMockRequestRouter(); err != nil {
		t.Errorf("error %s", err.Error())
	}
}
