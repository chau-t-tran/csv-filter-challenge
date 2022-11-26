package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CSVHandlerTestSuite struct {
	suite.Suite
}

/*-------------------Tests------------------------------*/

func (suite *CSVHandlerTestSuite) TestCSVHandlerConstructor() {
	_, err := NewCSVHandler("data.csv")
	assert.NoError(suite.T(), err)
}

/*-------------------Runner-----------------------------*/

func TestCSVHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CSVHandlerTestSuite))
}
