import "slices"

func rotate(matrix [][]int) {
    n := len(matrix)

    // транспонирование
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    // разворот строк
    for _, row := range matrix {
        slices.Reverse(row)
    }
}
