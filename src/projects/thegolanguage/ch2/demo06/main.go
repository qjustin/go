package main

import (
	"math"
)

func main() {
	for t := 0.0; t < cycles*2*math.Pi; t += res {
		x := math.Sin(t)
		y := math.Sin(t*freq + phase)
		img.SetColorIndex(
			size+int(x*size+0.5),
			size+int(y*size+0.5),
			// 最后插入的逗号不会导致编译错误，这是Go编译器的一个特性
			blackIndex,
		)
	}
}
