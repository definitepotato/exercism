package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupiedSquares := 0

	for idx, cbFile := range cb {
		if idx == file {
			for _, occupied := range cbFile {
				if occupied {
					occupiedSquares += 1
				}
			}
		}
	}

	return occupiedSquares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	occupiedRanks := 0

	for _, cbFile := range cb {
		if cbFile[rank-1] {
			occupiedRanks += 1
		}
	}

	return occupiedRanks
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	squares := 0

	for _, file := range cb {
		squares += len(file)
	}

	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedSquares := 0

	for i := 0; i < len(cb)+1; i++ {
		occupiedSquares += CountInRank(cb, i)
	}

	return occupiedSquares
}
