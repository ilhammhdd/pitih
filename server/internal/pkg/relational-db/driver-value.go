package relationalDB

import "database/sql/driver"

func ConvertToDriverValue(args ...any) []driver.Value {
	result := make([]driver.Value, len(args))

	for idx := range args {
		result[idx] = driver.Value(args[idx])
	}

	return result
}
