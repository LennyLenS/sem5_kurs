package generatelib

import (
	"fmt"
	mt "lib/table"
	"math/rand"
)

func GenerateRandTable(m, k int) mt.Table {
	newTable := mt.Table{
		Info: mt.TableInfo{
			ColumnName: []string{"id", "sex"},
			Rows:       m,
			Cols:       2,
		},
		Grid: make([][]string, m),
	}

	for i := range m {
		newTable.Grid[i] = make([]string, 2)
		newTable.Grid[i][0] = fmt.Sprint(i)
		if rand.Intn(2)%2 == 0 {
			newTable.Grid[i][1] = "f"
		} else {
			newTable.Grid[i][1] = "m"
		}
	}
	return newTable
}

func GenerateRandTable2(m, k int) mt.Table {
	newTable := mt.Table{
		Info: mt.TableInfo{
			ColumnName: []string{"name", "age"},
			Rows:       m,
			Cols:       2,
		},
		Grid: make([][]string, m),
	}

	for i := range m {
		newTable.Grid[i] = make([]string, 2)
		newTable.Grid[i][0] = GetRandString(1)
		newTable.Grid[i][1] = fmt.Sprint((rand.Intn(k)))
	}
	return newTable
}
