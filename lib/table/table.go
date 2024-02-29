package tables

type Table struct {
	Info TableInfo
	Grid [][]string
}

type TableInfo struct {
	ColumnName []string
	Rows       int
	Cols       int
}

func (m Table) Copy() Table {
	newM := Table{
		Info: TableInfo{
			ColumnName: m.Info.ColumnName,
			Rows:       m.Info.Rows,
			Cols:       m.Info.Cols,
		},
		Grid: make([][]string, m.Info.Rows),
	}

	for i := range newM.Grid {
		newM.Grid[i] = make([]string, m.Info.Cols)
		copy(newM.Grid[i], m.Grid[i])
	}

	return newM
}
