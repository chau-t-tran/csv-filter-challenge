package main

import (
	"bufio"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CSVHandlerTestSuite struct {
	suite.Suite
	csv CSVHandler
}

/*-------------------Helpers----------------------------*/

func resetScanner(c *CSVHandler) {
	/*
		A new scanner needs to be made with every new test
		because I can't think of a better way to reset the file
		pointer.
	*/
	c.file.Seek(0, io.SeekStart)
	c.scanner = bufio.NewScanner(c.file)
}

/*-------------------Setups/Teardowns-------------------*/

func (suite *CSVHandlerTestSuite) SetupSuite() {
	csv, err := NewCSVHandler("data.csv")
	assert.NoError(suite.T(), err)
	suite.csv = csv
}

/*-------------------Tests------------------------------*/

func (suite *CSVHandlerTestSuite) TestReadLines() {
	rows := [][]string{
		{"first_name", "last_name", "dob"},
		{"Bobby", "Tables", "19700101"},
		{"Ken", "Thompson", "19430204"},
		{"Rob", "Pike", "19560101"},
		{"Robert", "Griesemer", "19640609"},
	}
	for _, expectedRow := range rows {
		actualRow, err := suite.csv.NextRow()
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), expectedRow, actualRow)
	}
	_, err := suite.csv.NextRow()
	assert.Error(suite.T(), err) // eof error
	resetScanner(&suite.csv)
}

func (suite *CSVHandlerTestSuite) TestReadLineWithFilter() {
	filter := []string{"Ken", "Thompson", "19430204"}
	expect := []string{"Ken", "Thompson", "19430204"} // technically the same data as filter, but they both have different purposes
	_, err := suite.csv.NextRowWithFilter(filter)
	assert.Error(suite.T(), err)

	_, err = suite.csv.NextRowWithFilter(filter)
	assert.Error(suite.T(), err)

	row, err := suite.csv.NextRowWithFilter(filter)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, row)

	_, err = suite.csv.NextRowWithFilter(filter)
	assert.Error(suite.T(), err)

	_, err = suite.csv.NextRowWithFilter(filter)
	assert.Error(suite.T(), err)

	_, err = suite.csv.NextRowWithFilter(filter)
	assert.Error(suite.T(), err)
	resetScanner(&suite.csv)
}

/*-------------------Runner-----------------------------*/

func TestCSVHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CSVHandlerTestSuite))
}
