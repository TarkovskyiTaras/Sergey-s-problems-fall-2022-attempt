package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 3. Матрицы генерировать случайными числами и размер матрицы вводит юзер.
// 3.1. Заменить все числа, которые делятся на 5 без остачи в двухмерной матрице на 8

func main() {

	sliceLength := 5 //horizontal dimension
	sliceHeight := 5 //vertical dimension

	timeNow := time.Now()
	seconds := timeNow.Unix()
	rand.Seed(seconds)
	resultSliceOfSlices := make([][]int, sliceHeight)
	for i, _ := range resultSliceOfSlices {
		resultSliceOfSlices[i] = make([]int, sliceLength)
	}
	for _, y := range resultSliceOfSlices {

		for i, _ := range y {

			randNumb := rand.Intn(100)
			y[i] = randNumb
		}
	}
	fmt.Println(resultSliceOfSlices)
}
