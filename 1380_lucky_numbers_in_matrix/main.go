package luckynumbersinmatrix

func luckyNumbers (matrix [][]int) []int {
    rows, columns := len(matrix),len(matrix[0])
    lucky := make([]int, 0)
    for i := 0; i < rows; i++ {
        minValue := matrix[i][0]
        minColumn := 0
        for j := 1; j < columns; j++ {
            if minValue > matrix[i][j] {
                minValue = matrix[i][j]
                minColumn = j
            }
        }
        for j := 0; j < rows; j++ {
            if matrix[j][minColumn] > minValue {
                minValue = 0
                break
            }
        }
        if minValue != 0 {
            lucky = append(lucky, minValue)
        }
    }
    return lucky
}