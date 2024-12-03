package relationalDB

import (
	"strings"
)

type MapScanDestFunc[T any] func(column string, model *T) []any

type SelectModeler[T any] struct {
	ScanDests []any
	Model     *T
}

func NewSelectModeler[T any](query, alias string, mapScanDest MapScanDestFunc[T]) *SelectModeler[T] {
	aliasPrefix := alias + "."

	selectClauseIdx, selectClauseIdxOK := getClauseIndex(query, SelectClause)
	if !selectClauseIdxOK {
		return nil
	}

	fromClauseIdx, fromClauseIdxOK := getClauseIndex(query, FromClause)
	if !fromClauseIdxOK {
		return nil
	}

	columns := query[selectClauseIdx+len(SelectClause) : fromClauseIdx-1]
	columnsSplit := strings.Split(columns, ",")

	var model T
	scanDests := make([]any, 0, len(columnsSplit))

	for idx := range columnsSplit {
		column := strings.TrimSpace(columnsSplit[idx])
		if strings.HasPrefix(column, aliasPrefix) {
			columnSplit := strings.Split(column, ".")
			if len(columnSplit) == 2 {
				scanDests = append(scanDests, mapScanDest(columnSplit[1], &model)...)
			}
		}
	}

	return &SelectModeler[T]{ScanDests: scanDests, Model: &model}
}
