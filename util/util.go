package util

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// FunctionName - should return name of calling function
func FunctionName() string {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in FunctionName(), Error Info: ", errD)
		}
	}()
	pc, _, _, _ := runtime.Caller(1)
	funcName := strings.TrimSuffix(runtime.FuncForPC(pc).Name(), ".func1") // This is for defer function
	funcName = strings.TrimSuffix(funcName, ".1")                          // This is for go runtine function
	return funcName
}

// RecoverExceptionDetails will take one formal parameter as function name - should return exeception detail formated as: packageName.functionName:lineNumber
// Each format detail appended  with "<<" if it is multiple stack frames(LIFO). For example: packageName.functionName:lineNumber << packageName.functionName:lineNumber
func RecoverExceptionDetails(strfuncName string) string {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in RecoverExceptionDetails(), Error Info: ", errD)
		}
	}()
	var output string
	flag := false
	for skip := 1; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		strfunctionName := runtime.FuncForPC(pc).Name()
		if strings.Contains(file, "/runtime/") && strings.Contains(strfunctionName, "runtime.") {
			flag = true
			continue
		}
		if flag && strings.HasSuffix(file, ".go") {
			output += strfunctionName + ":" + strconv.Itoa(line) + " << "
			if strfuncName == strfunctionName {
				output = strings.TrimSuffix(output, " << ")
				break
			}
		}
	}
	return output
}
