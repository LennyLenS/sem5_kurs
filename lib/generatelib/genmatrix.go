package generatelib

import (
	"fmt"
	mt "lib/table"
	"math/rand"
)

func GenerateRandTable(m, k int) mt.Table {
	newMatrix := mt.Table{
		Info: mt.TableInfo{
			ColumnName: []string{"id", "sex"},
			Rows:       m,
			Cols:       2,
		},
		Grid: make([][]string, m),
	}

	for i := range m {
		newMatrix.Grid[i] = make([]string, 2)
		newMatrix.Grid[i][0] = fmt.Sprint(i)
		if rand.Intn(2)%2 == 0 {
			newMatrix.Grid[i][1] = "f"
		} else {
			newMatrix.Grid[i][1] = "m"
		}
	}
	return newMatrix
}
