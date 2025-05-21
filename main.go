package main

import (
	"github.com/OfflineBot/gomat/matrix2"
)

func main() {
	x := matrix2.NewMatrix2([][]float32 {
		{1.0},
		{2.0},
	});
	
	y := matrix2.NewMatrix2([][]float32 {
		{1.0, 2.0},
	})
	z := x.Dot(y)
	z.Println()
}
