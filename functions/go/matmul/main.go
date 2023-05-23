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

func generateRandomSquareMatrix(randomInstance *rand.Rand, dimensions int) *mat.Dense {
	data := make([]float64, dimensions*dimensions)
	for i := range data {
		data[i] = randomInstance.NormFloat64()
	}
	return mat.NewDense(dimensions, dimensions, data)
}

func main() {
	if len(os.Args) > 1 {
		seed := rand.NewSource(time.Now().UnixNano())
		randomInstance := rand.New(seed)
		dimensions, _ := strconv.Atoi(os.Args[1])
		matrix := mat.NewDense(dimensions, dimensions, nil)
		a := generateRandomSquareMatrix(randomInstance, dimensions)
		b := generateRandomSquareMatrix(randomInstance, dimensions)
		matrix.Mul(a, b)
		formatted := mat.Formatted(matrix, mat.Prefix(""), mat.Squeeze())
		fmt.Println(formatted)
	} else {
		log.Fatalln("Please pass a number as the first argument...")
	}
}
