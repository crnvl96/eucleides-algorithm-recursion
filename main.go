package main

import (
	"fmt"

	"github.com/crnvl96/eucleidesAlgorithm/helpers"
)

func main() {
	size := helpers.CalculateMaxTerrainSize(1680, 640)
	fmt.Println(size)
}
