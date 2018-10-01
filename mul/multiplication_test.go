package mul

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

//testCase is type struct used for each test case
type testCase struct {
	value1Float32, value2Float32, expectedResultFloat32 float32
}

//CalculatorTestSuite is a basic testing suite with methods for storing and retrieving the current *testing.T context.
type CalculatorTestSuite struct {
	suite.Suite
	testSuccessCases []testCase
	testFailureCases []testCase
}

//SetupSuite will run before the tests in the suite are run
func (suite *CalculatorTestSuite) SetupTest() {
	suite.testSuccessCases = []testCase{
		{10, 20, 200},
		{20, 30, 600},
	}
	suite.testFailureCases = []testCase{
		{10, 20, 100},
		{20, 30, 200},
	}
}

//Run takes a testing suite and runs all of the tests attached to it.
func Test_multiply_suite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

//MultiplicationOfTwoRationalNumber will take two formal parameters - should return multiply value
func (suite *CalculatorTestSuite) Test_given_two_rational_number_function_MultiplicationOfTwoRationalNumber_should_return_multiply_value_on_success_case() {
	for _, tc := range suite.testSuccessCases {
		actualResultFloat32 := MultiplicationOfTwoRationalNumber(tc.value1Float32, tc.value2Float32)
		if returnBooleanValue := assert.Equal(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should be equal"); !returnBooleanValue {
			require.Equal(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should be equal")
		}
	}
}

func (suite *CalculatorTestSuite) Test_given_two_rational_number_function_MultiplicationOfTwoRationalNumber_should_return_multiply_value_on_failure_case() {
	for _, tc := range suite.testFailureCases {
		actualResultFloat32 := MultiplicationOfTwoRationalNumber(tc.value1Float32, tc.value2Float32)
		if returnBooleanValue := assert.NotEqual(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should not be equal"); !returnBooleanValue {
			require.NotEqual(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should not be equal")
		}
	}
}

//TearDownTest will run after all the tests in the suite have been run
func (suite *CalculatorTestSuite) TearDownTest() {
	suite.testSuccessCases = nil
	suite.testFailureCases = nil
}
