package testsetup_test

import (
	"testing"

	"github.com/psbernardo/dockertest/internal/testsetup"
	"github.com/stretchr/testify/assert"
)

func TestSetupServer(t *testing.T) {
	_, err := testsetup.NewTestService(testsetup.WithMariaDBTest())
	assert.Nil(t, err)

}
