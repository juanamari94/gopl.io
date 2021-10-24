//!+
// Cf2 converts its numeric argument to Celsius and Fahrenheit. If none are provided, it gets them from stdin.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch2/lengthconv"
	"gopl.io/ch2/tempconv"
	"gopl.io/ch2/weightconv"
)

func convert(input float64) {
	fahrenheit := tempconv.Fahrenheit(input)
	celsius := tempconv.Celsius(input)
	meters := lengthconv.Meters(input)
	feet := lengthconv.Feet(input)
	kg := weightconv.Kilograms(input)
	lb := weightconv.Pounds(input)

	fmt.Println("Temperature")
	fmt.Printf("%s = %s, %s = %s\n\n", fahrenheit, tempconv.FToC(fahrenheit), celsius, tempconv.CToF(celsius))
	fmt.Println("Length")
	fmt.Printf("%s = %s, %s = %s\n\n", meters, lengthconv.MToF(meters), feet, lengthconv.FToM(feet))
	fmt.Println("Weight")
	fmt.Printf("%s = %s, %s = %s\n\n", kg, weightconv.KgToLb(kg), lb, weightconv.LbToKg(lb))
}

func processArguments(args []string) {
	for _, arg := range args {
		parsedArg, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf2: %v\n", err)
			os.Exit(1)
		}
		convert(parsedArg)
	}
}

func main() {
	if len(os.Args) < 2 {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		splitInput := strings.Split(input[:len(input)-1], " ")
		processArguments(splitInput)
	} else {
		processArguments(os.Args[1:])
	}
}

//!-
