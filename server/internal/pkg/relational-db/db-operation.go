package relationalDB

import (
	"context"
	"database/sql"

	"ilhammhdd.com/pitih/internal/pkg/errors"
)

type QueryAndArgs struct {
	Query string
	Args  []any
}

type DBOperation struct {
	DB *sql.DB
}

func (dbOps DBOperation) ExecTxContext(ctx context.Context, queryAndArgs []*QueryAndArgs) ([]*sql.Result, error) {
	if err := dbOps.DB.PingContext(ctx); err != nil {
		return nil, errors.WrapTrace(err)
	}

	tx, err := dbOps.DB.BeginTx(ctx, nil)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return nil, errors.WrapTrace(err)
		}
		return nil, errors.WrapTrace(err)
	}

	var results []*sql.Result
	for i := range queryAndArgs {
		preparedStmt, err := tx.PrepareContext(ctx, queryAndArgs[i].Query)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return nil, errors.WrapTrace(rollbackErr)
			}
			return nil, errors.WrapTrace(err)
		}

		result, err := preparedStmt.ExecContext(ctx, queryAndArgs[i].Args...)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return nil, errors.WrapTrace(rollbackErr)
			}
			return nil, errors.WrapTrace(err)
		}
		defer func() { errors.LogTraceIfErr(preparedStmt.Close()) }()
		results = append(results, &result)
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.WrapTrace(err)
	}

	return results, nil
}

func (dbOps DBOperation) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	stmt, err := dbOps.pingAndPrepare(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { errors.LogTraceIfErr(stmt.Close()) }()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, errors.WrapTrace(err)
	}

	return result, nil
}

func (dbOps DBOperation) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	stmt, err := dbOps.pingAndPrepare(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { errors.LogTraceIfErr(stmt.Close()) }()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, errors.WrapTrace(err)
	}

	return rows, nil
}

func (dbOps DBOperation) QueryRowContext(ctx context.Context, query string, args ...any) (*sql.Row, error) {
	stmt, err := dbOps.pingAndPrepare(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { errors.LogTraceIfErr(stmt.Close()) }()

	row := stmt.QueryRowContext(ctx, args...)

	return row, nil
}

func (dbOps DBOperation) QueryRowsToMapContext(ctx context.Context, query string, args ...any) ([]*map[string]any, error) {
	stmt, err := dbOps.pingAndPrepare(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { errors.LogTraceIfErr(stmt.Close()) }()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, errors.WrapTrace(err)
	}
	defer func() { errors.LogTraceIfErr(rows.Close()) }()

	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.WrapTrace(err)
	}

	columns := make([]any, len(cols))
	columnPointers := make([]any, len(cols))

	var resultRows []*map[string]any

	for rows.Next() {
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		err = rows.Scan(columnPointers...)
		if err != nil {
			return nil, errors.WrapTrace(err)
		}

		columnTypes, _ := rows.ColumnTypes()
		resultRow := make(map[string]any)

		for i, colName := range cols {
			val := columnPointers[i].(*any)
			columnTypeName := (*columnTypes[i]).DatabaseTypeName()
			if columnTypeName == "TINYINT" {
				if (*val).(int64) == 1 {
					resultRow[colName] = true
				} else if (*val).(int64) == 0 {
					resultRow[colName] = false
				}
			} else {
				resultRow[colName] = *val
			}
		}

		resultRows = append(resultRows, &resultRow)
	}

	return resultRows, nil
}

func (dbOps DBOperation) pingAndPrepare(ctx context.Context, query string) (*sql.Stmt, error) {
	if err := dbOps.DB.PingContext(ctx); err != nil {
		return nil, errors.WrapTrace(err)
	}

	stmt, err := dbOps.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.WrapTrace(err)
	}

	return stmt, nil
}
