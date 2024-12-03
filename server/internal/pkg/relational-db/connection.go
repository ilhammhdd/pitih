package relationalDB

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"ilhammhdd.com/pitih/config"
	"ilhammhdd.com/pitih/internal/pkg/errors"
)

const dsnFormat = "%s:%s@tcp(%s:%s)/%s?parseTime=true"

var ErrConfigNotSet = errors.NewTrace("configuration have not been set")

var DB *sql.DB

func Open() (*sql.DB, error) {
	if config.Config == nil {
		return nil, ErrConfigNotSet
	}

	dsn := parseDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.WrapTrace(err)
	}

	DB = db
	return DB, nil
}

func parseDSN() string {
	return fmt.Sprintf(
		dsnFormat,
		config.Config.RelationalDB.User,
		config.Config.RelationalDB.Password,
		config.Config.RelationalDB.Host,
		config.Config.RelationalDB.Port,
		config.Config.RelationalDB.Name)
}
