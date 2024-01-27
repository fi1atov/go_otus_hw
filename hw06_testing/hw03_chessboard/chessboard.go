package chessboard

func GetChessBoard(size int) (result string) {
	// fmt.Scanf("%d", &size)
	for x := 1; x <= size; x++ {
		if x%2 == 0 {
			result += " "
		}
		for y := 1; y <= size-1; y++ {
			if y%2 != 0 {
				result += "#"
			} else {
				result += " "
			}
		}
		if x%2 != 0 {
			result += " "
		}
		result += "\n"
	}

	return result
}
