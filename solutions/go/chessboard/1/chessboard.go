package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    var count int = 0
    occupied := cb[file]
    for _, files := range occupied {
        if files {
            count += 1
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    occupiedRank := 0
    if rank >= 1 && rank <= 8 {
    	for _, file := range cb {
            if file[rank-1] {
                occupiedRank += 1
            }
        }
    }
    return occupiedRank
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    squares := 0
	for _, files := range cb {
        squares += len(files)
    }
    return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var totalOccupied int = 0
    for _, files := range cb {
        for _, file := range files {
            if file {
                totalOccupied += 1
            }
        }
    }
    return totalOccupied
}
