package world

const WORLDSIZE = 10

type World struct {
    matrix [WORLDSIZE][WORLDSIZE]string
}

func NewWorld() *World {
    var matrix = [WORLDSIZE][WORLDSIZE]string{}
    for row, _ := range matrix {
        for column, _ := range matrix[row] {
            matrix[row][column] = "None"
        }
    }
    return &World{matrix: matrix}
}
