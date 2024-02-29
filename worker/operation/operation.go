package operation

import (
	"go/token"
	is "lib/infostructs"
	tb "lib/table"
	tr "lib/trees"
	"math"
	"sync"
)

type BinaryOp1 struct {
	Op    token.Token
	Left  *tr.TableLeaf
	Right *tr.TableLeaf
}

type Input struct {
	Root   BinaryOp1
	Tables map[string]tb.Table
}

const MinInterval = 1000

type TableLimit struct {
	Start int
	End   int
}

func SUM(op Input, workerInfo *is.WorkerInfo) tb.Table {
	table1 := op.Tables[op.Root.Left.TableName]
	table2 := op.Tables[op.Root.Right.TableName]

	var result tb.Table
	result.Grid = append(result.Grid, table1.Grid...)
	result.Grid = append(result.Grid, table2.Grid...)
	result.Info.Cols = table1.Info.Cols
	result.Info.ColumnName = table1.Info.ColumnName
	result.Info.Rows = table1.Info.Rows + table2.Info.Rows

	return result
}

func MULGor(info TableLimit, t1 tb.Table, t2 tb.Table) tb.Table {
	var result tb.Table
	for i := info.Start; i <= info.End; i++ {
		for k := 0; k < t2.Info.Rows; k++ {
			s := make([]string, 0)
			s = append(s, t1.Grid[i]...)
			s = append(s, t2.Grid[k]...)
			result.Grid = append(result.Grid, [][]string{s}...)
			result.Info.ColumnName = append(t1.Info.ColumnName, t2.Info.ColumnName...)
			result.Info.Cols = t1.Info.Cols + t2.Info.Cols
			result.Info.Rows++
		}
	}
	return result
}

func MUL(op Input, workerInfo *is.WorkerInfo) tb.Table {
	table1 := op.Tables[op.Root.Left.TableName]
	table2 := op.Tables[op.Root.Right.TableName]
	resulting := make(map[int]tb.Table)
	var result tb.Table
	var quantityInterval int
	currentInterval := table1.Info.Rows / workerInfo.Cores
	currentInterval = max(MinInterval, currentInterval)
	quantityInterval = int(math.Ceil(float64(table1.Info.Rows) / float64(currentInterval)))
	var sliceTableInfo []TableLimit = make([]TableLimit, quantityInterval)

	for i := 0; i < quantityInterval; i++ {
		sliceTableInfo[i] = TableLimit{
			Start: i * currentInterval,
			End:   min((i+1)*currentInterval-1, table1.Info.Rows-1),
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < quantityInterval; i++ {
		wg.Add(1)
		go func() {
			resulting[i] = MULGor(sliceTableInfo[i], table1, table2)
			wg.Done()
		}()
	}
	wg.Wait()

	result.Info.ColumnName = resulting[0].Info.ColumnName
	for i := 0; i < quantityInterval; i++ {
		result.Grid = append(result.Grid, resulting[i].Grid...)
		result.Info.Cols = resulting[i].Info.Cols
		result.Info.Rows += resulting[i].Info.Rows
	}

	return result
}

func QUOGor(info TableLimit, t1 tb.Table, t2 tb.Table) tb.Table {
	var result tb.Table
	for i := info.Start; i <= info.End; i++ {
		flag := true
		for k := 0; k < t2.Info.Rows; k++ {
			flag = true
			for j := 0; j < t1.Info.Cols; j++ {
				if t1.Grid[i][j] != t2.Grid[k][j] {
					flag = false
					break
				}
			}
			if flag {
				break
			}
		}
		if flag {
			result.Grid = append(result.Grid, [][]string{t1.Grid[i]}...)
			result.Info.ColumnName = t1.Info.ColumnName
			result.Info.Cols = t1.Info.Cols
			result.Info.Rows++
		}
	}
	return result
}

func QUO(op Input, workerInfo *is.WorkerInfo) tb.Table {
	table1 := op.Tables[op.Root.Left.TableName]
	table2 := op.Tables[op.Root.Right.TableName]
	resulting := make(map[int]tb.Table)
	var result tb.Table
	var quantityInterval int
	currentInterval := table1.Info.Rows / workerInfo.Cores
	currentInterval = max(MinInterval, currentInterval)
	quantityInterval = int(math.Ceil(float64(table1.Info.Rows) / float64(currentInterval)))
	var sliceTableInfo []TableLimit = make([]TableLimit, quantityInterval)

	for i := 0; i < quantityInterval; i++ {
		sliceTableInfo[i] = TableLimit{
			Start: i * currentInterval,
			End:   min((i+1)*currentInterval-1, table1.Info.Rows-1),
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < quantityInterval; i++ {
		wg.Add(1)
		go func() {
			resulting[i] = QUOGor(sliceTableInfo[i], table1, table2)
			wg.Done()
		}()
	}
	wg.Wait()

	result.Info.ColumnName = table1.Info.ColumnName
	for i := 0; i < quantityInterval; i++ {
		result.Grid = append(result.Grid, resulting[i].Grid...)
		result.Info.Cols = resulting[i].Info.Cols
		result.Info.Rows += resulting[i].Info.Rows
	}

	return result
}

func SUBGor(info TableLimit, t1 tb.Table, t2 tb.Table) tb.Table {
	var result tb.Table
	for i := info.Start; i <= info.End; i++ {
		flag := true
		for k := 0; k < t2.Info.Rows; k++ {
			flag = true
			for j := 0; j < t1.Info.Cols; j++ {
				if t1.Grid[i][j] != t2.Grid[k][j] {
					flag = false
					break
				}
			}
			if flag {
				break
			}
		}
		if !flag {
			result.Grid = append(result.Grid, [][]string{t1.Grid[i]}...)
			result.Info.ColumnName = t1.Info.ColumnName
			result.Info.Cols = t1.Info.Cols
			result.Info.Rows++
		}
	}
	return result
}

func SUB(op Input, workerInfo *is.WorkerInfo) tb.Table {
	table1 := op.Tables[op.Root.Left.TableName]
	table2 := op.Tables[op.Root.Right.TableName]
	resulting := make(map[int]tb.Table)
	var result tb.Table
	var quantityInterval int
	currentInterval := table1.Info.Rows / workerInfo.Cores
	currentInterval = max(MinInterval, currentInterval)
	quantityInterval = int(math.Ceil(float64(table1.Info.Rows) / float64(currentInterval)))
	var sliceTableInfo []TableLimit = make([]TableLimit, quantityInterval)

	for i := 0; i < quantityInterval; i++ {
		sliceTableInfo[i] = TableLimit{
			Start: i * currentInterval,
			End:   min((i+1)*currentInterval-1, table1.Info.Rows-1),
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < quantityInterval; i++ {
		wg.Add(1)
		go func() {
			resulting[i] = SUBGor(sliceTableInfo[i], table1, table2)
			wg.Done()
		}()
	}
	wg.Wait()

	result.Info.ColumnName = table1.Info.ColumnName
	for i := 0; i < quantityInterval; i++ {
		result.Grid = append(result.Grid, resulting[i].Grid...)
		result.Info.Cols = resulting[i].Info.Cols
		result.Info.Rows += resulting[i].Info.Rows
	}

	return result
}
