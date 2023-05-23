package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		floatNumber, _ := strconv.ParseFloat(os.Args[1], 64)
		resultSin := math.Sin(floatNumber)
		resultCos := math.Cos(floatNumber)
		resultSqrt := math.Sqrt(floatNumber)
		fmt.Printf("Sin: %f\nCos: %f\nSqrt: %f\n", resultSin, resultCos, resultSqrt)
	} else {
		log.Fatalln("Please pass a number as the first argument...")
	}
}
