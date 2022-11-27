package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CLITestSuite struct {
	suite.Suite
	orderedArgs []string
}

/*-------------------Setups/Teardowns-------------------*/

func (suite *CLITestSuite) SetupSuite() {
	suite.orderedArgs = []string{"Ken", "Thompson", "19430204"}
}

/*-------------------Tests------------------------------*/

func (suite *CLITestSuite) TestImplicitArgs() {
	args := []string{"cmd", "-f=data.csv", "-i", "Ken", "Thompson", "19430204"}
	os.Args = args
	cli := NewCLI()
	assert.Equal(suite.T(), "data.csv", cli.path)
	assert.Equal(suite.T(), IMPLICIT, cli.mode)
	assert.Equal(suite.T(), suite.orderedArgs, cli.args)
}

/*-------------------Runner-----------------------------*/

func TestCLITestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}
