package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CSVHandlerTestSuite struct {
	suite.Suite
	csv CSVHandler
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
}

/*-------------------Runner-----------------------------*/

func TestCSVHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CSVHandlerTestSuite))
}
