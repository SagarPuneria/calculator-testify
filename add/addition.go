package add

import (
	util "calculator-testify/util"
	"fmt"
)

func AddtionOfTwoRationalNumber(value1Float32, value2Float32 float32) float32 {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", util.RecoverExceptionDetails(util.FunctionName()), " and recovered in AddTwoRationalNumber function, Error Info: ", errD)
		}
	}()
	resultFloat32 := value1Float32 + value2Float32
	return resultFloat32
}
