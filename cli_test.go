package main

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CLITestSuite struct {
	suite.Suite
	orderedArgs []string
	stdout      string
}

/*-------------------Setups/Teardowns-------------------*/

func (suite *CLITestSuite) SetupSuite() {
	suite.orderedArgs = []string{"Ken", "Thompson", "19430204"}
	suite.stdout = "Ken,Thompson,19430204\n"
}

/*-------------------Tests------------------------------*/

func (suite *CLITestSuite) TestImplicitArgs() {
	args := []string{"cmd", "-f=data.csv", "-i", "Ken", "Thompson", "19430204"}
	os.Args = args
	cli, err := NewCLI()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "data.csv", cli.path)
	assert.Equal(suite.T(), IMPLICIT, cli.mode)
	assert.Equal(suite.T(), suite.orderedArgs, cli.args)

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func (suite *CLITestSuite) TestIncompleteImplicitArgs() {
	args := []string{"cmd", "-f=data.csv", "-i", "Ken", "Thompson"}
	os.Args = args
	cli, err := NewCLI()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "data.csv", cli.path)
	assert.Equal(suite.T(), IMPLICIT, cli.mode)
	assert.Equal(suite.T(), []string{"Ken", "Thompson", "*"}, cli.args)

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
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
	cli, err := NewCLI()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "data.csv", cli.path)
	assert.Equal(suite.T(), EXPLICIT, cli.mode)
	assert.Equal(suite.T(), suite.orderedArgs, cli.args)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func (suite *CLITestSuite) TestIncompleteExplicitArgs() {
	args := []string{
		"cmd",
		"-f=data.csv",
		"-e",
		"dob=19430204",
		"first_name=Ken",
	}
	os.Args = args
	cli, err := NewCLI()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "data.csv", cli.path)
	assert.Equal(suite.T(), EXPLICIT, cli.mode)
	assert.Equal(suite.T(), []string{"Ken", "*", "19430204"}, cli.args)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func (suite *CLITestSuite) TestPromptArgs() {
	/*
		Prompt mode will interactive prompt for user input for each field
		in the .csv file.
	*/
	args := []string{
		"cmd",
		"-f=data.csv",
		"-p",
	}
	os.Args = args

	input := []byte("Ken\nThompson\n19430204\n")
	mockStdin, err := ioutil.TempFile("", "nil")
	assert.NoError(suite.T(), err)

	defer os.Remove(mockStdin.Name())

	_, err = mockStdin.Write(input)
	assert.NoError(suite.T(), err)

	_, err = mockStdin.Seek(0, 0)
	assert.NoError(suite.T(), err)

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = mockStdin

	cli, err := NewCLI()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "data.csv", cli.path)
	assert.Equal(suite.T(), PROMPT, cli.mode)
	assert.Equal(suite.T(), suite.orderedArgs, cli.args)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func (suite *CLITestSuite) TestRun() {
	args := []string{
		"cmd",
		"-f=data.csv",
		"-e",
		"last_name=Thompson",
		"dob=19430204",
		"first_name=Ken",
	}
	os.Args = args

	mockStdoutReader, mockStdout, err := os.Pipe()
	assert.NoError(suite.T(), err)

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	os.Stdout = mockStdout

	cli, err := NewCLI()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "data.csv", cli.path)
	assert.Equal(suite.T(), EXPLICIT, cli.mode)
	assert.Equal(suite.T(), suite.orderedArgs, cli.args)

	cli.Run()

	buf := make([]byte, 1024)
	n, err := mockStdoutReader.Read(buf)
	assert.NoError(suite.T(), err)

	result := string(buf[:n])
	assert.Equal(suite.T(), suite.stdout, result)
}

/*-------------------Runner-----------------------------*/

func TestCLITestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}
