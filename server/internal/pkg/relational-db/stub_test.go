package relationalDB_test

import relationalDB "ilhammhdd.com/pitih/internal/pkg/relational-db"

const stubTableName = "stub_table"

const (
	stubID          = "id"
	stubName        = "name"
	stubage         = "age"
	stubForeigner   = "foreigner"
	stubStocksOwned = "stocksOwned"
)

type stubModel struct {
	id          uint64
	name        string
	age         uint16
	foreigner   bool
	stocksOwned float32
}

var stubAllColumns = []string{
	stubID,
	stubName,
	stubage,
	stubForeigner,
	stubStocksOwned,
}

var stubMapScanDestFunc = relationalDB.MapScanDestFunc[stubModel](func(column string, model *stubModel) []any {
	switch column {
	case stubID:
		return []any{&model.id}
	case stubName:
		return []any{&model.name}
	case stubage:
		return []any{&model.age}
	case stubForeigner:
		return []any{&model.foreigner}
	case stubStocksOwned:
		return []any{&model.stocksOwned}
	case "*":
		return []any{
			&model.id,
			&model.name,
			&model.age,
			&model.foreigner,
			&model.stocksOwned,
		}
	default:
		return nil
	}
})
