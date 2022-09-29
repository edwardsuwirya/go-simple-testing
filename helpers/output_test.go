package helpers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OutputTestSuite struct {
	suite.Suite
	outputService Output
}

//Secara otomatis SetupTest dijalankan
func (suite *OutputTestSuite) SetupTest() {
	suite.outputService = Output{}
}

func (suite *OutputTestSuite) TestOutput_Console() {
	resultDummy := 10
	err := suite.outputService.Console(resultDummy)
	assert.Nil(suite.T(), err)
}

func (suite *OutputTestSuite) TestOutput_Error() {
	err := suite.outputService.Error(errors.New("error"))
	assert.Nil(suite.T(), err)
}

func TestOutputTestSuite(t *testing.T) {
	suite.Run(t, new(OutputTestSuite))
}
