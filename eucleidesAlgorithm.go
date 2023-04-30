package eucleidesAlgorithm

func getGreaterSize(height int, width int) (int, int) {
	var greaterSize int
	var minorSize int

	if height > width {
		greaterSize = height
		minorSize = width
	} else {
		greaterSize = width
		minorSize = height
	}

	return greaterSize, minorSize
}

func CalculateMaxTerrainSize(height int, width int) int {
	greaterSize, minorSize := getGreaterSize(height, width)

	if greaterSize%minorSize == 0 {
		return minorSize
	}

	return CalculateMaxTerrainSize(minorSize, greaterSize%minorSize)
}
