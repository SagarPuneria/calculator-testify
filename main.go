package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	add "calculator-testify/add"
	div "calculator-testify/div"
	mul "calculator-testify/mul"
	rdm "calculator-testify/random"
	sub "calculator-testify/sub"
	util "calculator-testify/util"
)

func main() {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", util.RecoverExceptionDetails(util.FunctionName()), " and recovered in main function, Error Info: ", errD)
		}
	}()
	fmt.Println("Go Calculator")
	fmt.Println("=============")
	greeting := "Enter your choice\n\t1. Addition of two numbers\n\t2. Subtraction of two numbers\n\t3. Multiplication of two numbers\n\t4. Division of two numbers\n\t5. Random of two numbers\n\t6.Exit"
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(greeting)
		scanner.Scan()
		choiceInteger, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error in choice input and Error Info:", err)
			fmt.Println("Enter correct integer number")
			continue
		}
		var value1Float32, value2Float32 float32
		if choiceInteger > 0 && choiceInteger < 6 {
			fmt.Println("Enter the numbers for calculation:")
			value1Float32 = readInputValue(scanner)
			value2Float32 = readInputValue(scanner)
		}
		executeSwitchStatementBasedOnGivenChoice(choiceInteger, value1Float32, value2Float32)
	}
}

// readInputValue which Scans and read a line from Stdin(Console) - return Console input value.
func readInputValue(scanner *bufio.Scanner) (inputFloat32 float32) {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", util.RecoverExceptionDetails(util.FunctionName()), " and recovered in readInputValue function, Error Info: ", errD)
		}
	}()
	for scanner.Scan() {
		inputFloat64, err := strconv.ParseFloat(scanner.Text(), 32)
		if err != nil {
			fmt.Println("Error in console input and Error Info:", err)
			fmt.Println("Enter correct rational number")
			continue
		}
		inputFloat32 = float32(inputFloat64)
		break
	}
	return
}

func executeSwitchStatementBasedOnGivenChoice(choice int, value1Float32, value2Float32 float32) {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", util.RecoverExceptionDetails(util.FunctionName()), " and recovered in executeSwitchStatementBasedOnGivenChoice function, Error Info: ", errD)
		}
	}()
	switch choice {
	case 1:
		resultFloat32 := add.AddtionOfTwoRationalNumber(value1Float32, value2Float32)
		fmt.Printf("Addition:  %.2f + %.2f = %.2f\n", value1Float32, value2Float32, resultFloat32)
	case 2:
		resultFloat32 := sub.SubtractionOfTwoRationalNumber(value1Float32, value2Float32)
		fmt.Printf("Subtraction:  %.2f - %.2f = %.2f\n", value1Float32, value2Float32, resultFloat32)
	case 3:
		resultFloat32 := mul.MultiplicationOfTwoRationalNumber(value1Float32, value2Float32)
		fmt.Printf("Multiplication:  %.2f x %.2f = %.2f\n", value1Float32, value2Float32, resultFloat32)
	case 4:
		quotientFloat32, err := div.DivisionOfTwoRationalNumber(value1Float32, value2Float32)
		if err != nil {
			fmt.Println("Error Info:", err)
		} else {
			fmt.Printf("Division:  %.2f / %.2f = %.2f\n", value1Float32, value2Float32, quotientFloat32)
		}
	case 5:
		r := &rdm.Calc{Rnd: &rdm.Calc{}}
		min, max := int(value1Float32), int(value2Float32)
		randomIntNumber, err := r.Random(&min, &max)
		if err != nil {
			fmt.Println("Error Info:", err)
		} else {
			fmt.Printf("Random of  %d and  %d = %d\n", min, max, randomIntNumber)
		}
	case 6:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}
