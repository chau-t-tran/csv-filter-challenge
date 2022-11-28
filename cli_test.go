package main

import (
	"flag"
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
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cli, err := NewCLI()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "data.csv", cli.path)
	assert.Equal(suite.T(), IMPLICIT, cli.mode)
	assert.Equal(suite.T(), suite.orderedArgs, cli.args)
}

func (suite *CLITestSuite) TestExplicitArgs() {
	/*
		Explicit is similar to implicit, except the CLI constructor
		re-arranges all filter args in the order shown in the .csv
		file.
	*/
	args := []string{
		"cmd",
		"-f=data.csv",
		"-e",
		"last_name=Thompson",
		"dob=19430204",
		"first_name=Ken",
	}
	os.Args = args
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cli, err := NewCLI()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "data.csv", cli.path)
	assert.Equal(suite.T(), EXPLICIT, cli.mode)
	assert.Equal(suite.T(), suite.orderedArgs, cli.args)
}

/*-------------------Runner-----------------------------*/

func TestCLITestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}
