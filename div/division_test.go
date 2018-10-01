package div

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

//testCase is type struct used for each test case
type testCase struct {
	dividentFloat32, divisorFloat32, expectedResultFloat32 float32
}

//CalculatorTestSuite is a basic testing suite with methods for storing and retrieving the current *testing.T context.
type CalculatorTestSuite struct {
	suite.Suite
	testSuccessCases []testCase
	testFailureCases []testCase
}

//Run takes a testing suite and runs all of the tests attached to it.
func Test_divide_suite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

//SetupSuite will run before the tests in the suite are run
func (suite *CalculatorTestSuite) SetupTest() {
	suite.testSuccessCases = []testCase{
		{10, 20, 0.5},
		{30, 20, 1.5},
	}
	suite.testFailureCases = []testCase{
		{10, 0, 10},
		{30, 30, 1.5},
	}
}

/*
* DivisionOfTwoRationalNumber will take two formal parameters - should return quotient value
* 1. If divisor number is not zero - should return actual quotient value and nil.
* 2. If divisor number is zero - should return actual quotient value and err.
 */
func (suite *CalculatorTestSuite) Test_given_two_rational_number_function_DivisionOfTwoRationalNumber_should_return_quotient_value_on_success_case() {
	for _, tc := range suite.testSuccessCases {
		actualResultFloat32, err := DivisionOfTwoRationalNumber(tc.dividentFloat32, tc.divisorFloat32)
		require.Nil(suite.T(), err, "Error should be nil")
		require.Equal(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should be equal")
	}
}

func (suite *CalculatorTestSuite) Test_given_two_rational_number_function_DivisionOfTwoRationalNumber_should_return_quotient_value_on_failure_case() {
	for _, tc := range suite.testFailureCases {
		actualResultFloat32, err := DivisionOfTwoRationalNumber(tc.dividentFloat32, tc.divisorFloat32)
		if tc.divisorFloat32 == 0 {
			require.NotNil(suite.T(), err, "It should be error")
		} else {
			require.Nil(suite.T(), err, "Error should be nil")
		}
		require.NotEqual(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should not be equal")
	}
}

func (suite *CalculatorTestSuite) TearDownTest() {
	suite.testSuccessCases = nil
	suite.testFailureCases = nil
}
