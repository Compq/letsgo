package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 4
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	suite.Equal(4, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

// type Browsers struct {
// 	Chrome  string
// 	FireFox string
// 	IE      string
// }

// var browsers Browsers

// func TestHomeSuite(t *testing.T) {
// 	testCases := []struct {
// 		Browsers string
// 	}{
// 		{Browsers: "browsers.FireFox"}, //as for demo fr showig in logs 'Failed on Browser.Type
// 		{Browsers: "browsers.Chrome"},
// 		{Browsers: "browsers.IE"},
// 	}

// 	for _, tC := range testCases {
// 		t.Run(tC.Browsers, func(t *testing.T) {
// 			TestHome200(t)
// 			TestHome404(t)

// 		})
// 	}
// }

// func TestAllSuites(t *testing.T) {
// 	testCases := []struct {
// 		Suites string
// 	}{
// 		{Suites: "Descript"},
// 	}

// 	for _, tC := range testCases {

// 		t.Run(tC.Suites, func(t *testing.T) {
// 			TestHomeSuite(t)

// 		})
// 	}
// }
