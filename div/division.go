package div

import (
	"errors"
	"fmt"

	util "calculator-testify/util"
)

func DivisionOfTwoRationalNumber(dividentFloat32, divisorFloat32 float32) (quotient float32, err error) {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", util.RecoverExceptionDetails(util.FunctionName()), " and recovered in DivisionOfTwoRationalNumber function, Error Info: ", errD)
		}
	}()
	if divisorFloat32 == 0 {
		err := errors.New("division by zero error")
		return quotient, err
	}
	quotient = dividentFloat32 / divisorFloat32
	return quotient, nil
}
