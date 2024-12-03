package relationalDB

import "strings"

type ArgsResolver[T any] struct {
	mapScanDest MapScanDestFunc[T]
	allColumns  []string
}

func NewArgsResolver[T any](mapScanDest MapScanDestFunc[T], allColumns []string) *ArgsResolver[T] {
	return &ArgsResolver[T]{mapScanDest: mapScanDest, allColumns: allColumns}
}

func (argsResolver *ArgsResolver[T]) ResolveInsertArgs(query string, models ...*T) []any {
	insertIntoClauseIdx, insertIntoClauseIdxOK := getClauseIndex(query, InsertClause)
	if !insertIntoClauseIdxOK {
		return nil
	}

	valuesClauseIdx, valuesClauseIdxOK := getClauseIndex(query, ValuesClause)
	if !valuesClauseIdxOK {
		return nil
	}

	tableColumns := query[insertIntoClauseIdx:valuesClauseIdx]
	columnOpenBracketIdx := strings.Index(tableColumns, "(")
	columnCloseBracketIdx := strings.Index(tableColumns, ")")

	var columns []string
	if columnOpenBracketIdx == -1 && columnCloseBracketIdx == -1 {
		columns = argsResolver.allColumns
	} else {
		columns = strings.Split(tableColumns[columnOpenBracketIdx+1:columnCloseBracketIdx], ",")
	}

	insertArgs := make([]any, 0, len(models)*len(columns))
	for idxModel := range models {
		for idxColumn := range columns {
			insertArgs = append(insertArgs, argsResolver.mapScanDest(strings.TrimSpace(columns[idxColumn]), models[idxModel])...)
		}
	}

	return insertArgs
}

func (argsResolver *ArgsResolver[T]) ResolveUpdateArgs(query string, model *T, whereArgs ...any) []any {
	setClauseIdx, setClauseIdxOK := getClauseIndex(query, SetClause)
	if !setClauseIdxOK {
		return nil
	}

	whereClauseIdx, whereClauseIdxOK := getClauseIndex(query, WhereClause)
	if !whereClauseIdxOK {
		return nil
	}

	columnParamsRaw := query[setClauseIdx+len(SetClause) : whereClauseIdx]
	columnParams := strings.Split(columnParamsRaw, ",")

	args := make([]any, 0, len(columnParams)+len(whereArgs))

	for idx := range columnParams {
		columnAndParam := strings.Split(columnParams[idx], "=")
		if len(columnAndParam) == 2 {
			args = append(args, argsResolver.mapScanDest(strings.TrimSpace(columnAndParam[0]), model)...)
		}
	}

	args = append(args, whereArgs...)

	return args
}
