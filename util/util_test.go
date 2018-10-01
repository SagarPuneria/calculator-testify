package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

//testCase is type struct used for each test case
type testCase struct {
	expectedFunctionNameString               string
	expectedExceptionDetailExcludeLineNumber string
}

//utilTestSuite is a basic testing suite with methods for storing and retrieving the current *testing.T context.
type utilTestSuite struct {
	suite.Suite
	testFuncNameSuccessCases testCase
	testFuncNameFailureCases testCase
}

//Run takes a testing suite and runs all of the tests attached to it.
func Test_util_suite(t *testing.T) {
	suite.Run(t, new(utilTestSuite))
}

//SetupSuite will run before the tests in the suite are run
func (suite *utilTestSuite) SetupTest() {
	suite.testFuncNameSuccessCases.expectedFunctionNameString =
		"calculator/util.(*utilTestSuite).Test_function_FunctionName_should_return_name_of_calling_function_on_success_case"
	suite.testFuncNameFailureCases.expectedFunctionNameString =
		"calculator/util.(*utilTestSuite).Test_function_FunctionName_should_return_name_of_calling_function_on_success_case"
	suite.testFuncNameSuccessCases.expectedExceptionDetailExcludeLineNumber = "calculator/util.generateException"
	suite.testFuncNameFailureCases.expectedExceptionDetailExcludeLineNumber = "calculator/util"
}

// FunctionName - should return name of calling function
func (suite *utilTestSuite) Test_function_FunctionName_should_return_name_of_calling_function_on_success_case() {
	actualFuncNameString := FunctionName()
	if returnBooleanValue := assert.Equal(suite.T(), suite.testFuncNameSuccessCases.expectedFunctionNameString, actualFuncNameString, "Expected and Calculated value should be equal"); !returnBooleanValue {
		require.Equal(suite.T(), suite.testFuncNameSuccessCases.expectedFunctionNameString, actualFuncNameString, "Expected and Calculated value should be equal")
	}
}

func (suite *utilTestSuite) Test_function_FunctionName_should_return_name_of_calling_function_on_failure_case() {
	actualFuncNameString := FunctionName()
	if returnBooleanValue := assert.NotEqual(suite.T(), suite.testFuncNameFailureCases.expectedFunctionNameString, actualFuncNameString, "Expected and Calculated value should not be equal"); !returnBooleanValue {
		require.NotEqual(suite.T(), suite.testFuncNameFailureCases.expectedFunctionNameString, actualFuncNameString, "Expected and Calculated value should not be equal")
	}
}

// RecoverExceptionDetails will take one formal parameter as function name - should return exeception detail formated as: packageName.functionName:lineNumber
// Each format detail appended  with "<<" if it is multiple stack frames(LIFO). For example: packageName.functionName:lineNumber << packageName.functionName:lineNumber
func (suite *utilTestSuite) Test_given_function_name_function_RecoverExceptionDetails_should_return_exception_detail_of_calling_function_on_success_case() {
	actualExceptionDetailString := generateException()
	idx := strings.LastIndex(actualExceptionDetailString, ":")
	actualExceptionDetailString = actualExceptionDetailString[0:idx]
	if returnBooleanValue := assert.Equal(suite.T(), suite.testFuncNameSuccessCases.expectedExceptionDetailExcludeLineNumber, actualExceptionDetailString, "Expected and Calculated value should be equal"); !returnBooleanValue {
		require.Equal(suite.T(), suite.testFuncNameSuccessCases.expectedExceptionDetailExcludeLineNumber, actualExceptionDetailString, "Expected and Calculated value should be equal")
	}
}

func (suite *utilTestSuite) Test_given_function_name_function_RecoverExceptionDetails_should_return_exception_detail_of_calling_function_on_failure_case() {
	actualExceptionDetailString := generateException()
	idx := strings.LastIndex(actualExceptionDetailString, ":")
	actualExceptionDetailString = actualExceptionDetailString[0:idx]
	if returnBooleanValue := assert.NotEqual(suite.T(), suite.testFuncNameFailureCases.expectedExceptionDetailExcludeLineNumber, actualExceptionDetailString, "Expected and Calculated value should not be equal"); !returnBooleanValue {
		require.NotEqual(suite.T(), suite.testFuncNameFailureCases.expectedExceptionDetailExcludeLineNumber, actualExceptionDetailString, "Expected and Calculated value should not be equal")
	}
}

// Function generateException intentionally generate panic/exception - should return exception detail
func generateException() (actualExceptionDetailString string) {
	defer func() {
		if errD := recover(); errD != nil {
			actualExceptionDetailString = RecoverExceptionDetails(FunctionName())
		}
	}()
	panic("Generating panic/exception intentionally for testing")
}

//TearDownTest will run after all the tests in the suite have been run
func (suite *utilTestSuite) TearDownTest() {
	suite.testFuncNameSuccessCases = testCase{}
	suite.testFuncNameFailureCases = testCase{}
}
