package rdm

import (
	"fmt"
	"math/rand"
	"time"

	util "calculator/util"
)

type random interface {
	getRandom(int, int) (int, error)
}

type Calc struct {
	Rnd random
}

func (c *Calc) Random(min, max *int) (randNumber int, err error) {
	defer func() {
		if errD := recover(); errD != nil {
			err = fmt.Errorf("%v", errD)
			fmt.Println("Exception occurred at ", util.RecoverExceptionDetails(util.FunctionName()), " and recovered in Random function, Error Info: ", errD)
		}
	}()
	if *max < *min {
		*min, *max = *max, *min
	}
	res := *max - *min
	if res <= 0 {
		panic(fmt.Sprintf("Difference of max=%d and min=%d should be greater than zero but actual difference is %d", *max, *min, res))
	}
	randNumber, err = c.Rnd.getRandom(res, *min)
	return
}

func (c *Calc) getRandom(res, min int) (randNumber int, err error) {
	defer func() {
		if errD := recover(); errD != nil {
			err = fmt.Errorf("%v", errD)
			fmt.Println("Exception occurred at ", util.RecoverExceptionDetails(util.FunctionName()), " and recovered in getRandom function, Error Info: ", errD)
		}
	}()
	rand.Seed(time.Now().Unix())
	randNumber = rand.Intn(res) + min
	return
}
