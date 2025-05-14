package main

import (
	"github.com/OfflineBot/gomat/matrix1"
	"github.com/OfflineBot/gomat/matrix2"
)

func main() {
	x := matrix2.NewMatrix2([][]float32 {
		[]float32 {1.0, 2.0},
		[]float32 {3.0, 4.0},
		[]float32 {1.0, 2.0},
	});
	y := matrix1.NewMatrix1([]float32 {1.0, 1.0});
	x.AddMatrix1(y)
	x.Println()
}
