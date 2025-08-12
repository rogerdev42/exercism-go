package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f := cb[file]
    total := 0
    for _, occupied := range f {
        if occupied {
            total++
        }
    }
    return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    total := 0
	for _, file := range cb {
        if rank < 1 || rank > len(file) {
            return 0
        }
        if file[rank - 1] {
            total++
        } 
    }
    return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	for _, file := range cb {
		return len(file) * len(cb)
	}
    return 0
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	total := 0
    for _, file := range cb {
		for _, occupied := range file {
            if occupied {
            	total++
            }
        }
  	}
    return total
}
