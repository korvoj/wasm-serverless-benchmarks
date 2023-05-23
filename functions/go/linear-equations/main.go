package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateRandomMatrix(randomInstance *rand.Rand, rows int, columns int) *mat.Dense {

	data := make([]float64, rows*columns)
	for i := range data {
		data[i] = randomInstance.NormFloat64()
	}
	return mat.NewDense(rows, columns, data)
}

func main() {
	if len(os.Args) > 1 {
		seed := rand.NewSource(time.Now().UnixNano())
		randomInstance := rand.New(seed)
		unknowns, _ := strconv.Atoi(os.Args[1])
		matrixResult := mat.NewDense(unknowns, 1, nil)
		a := generateRandomMatrix(randomInstance, unknowns, unknowns)
		b := generateRandomMatrix(randomInstance, unknowns, 1)
		matrixResult.Solve(a, b)
		formatted := mat.Formatted(matrixResult, mat.Prefix(""), mat.Squeeze())
		fmt.Println(formatted)
	} else {
		log.Fatalln("Please pass a number as the first argument...")
	}
}
