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

/*-------------------Runner-----------------------------*/

func TestCSVHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CSVHandlerTestSuite))
}
