package ddl

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GenerateDDL(db *sqlx.DB, schemaName string) ([]string, error) {
	tableNames := fmt.Sprintf(`select table_name from information_schema.tables where table_schema = '%s';`, schemaName)

	tables := make([]string, 0)
	err := db.Select(&tables, tableNames)
	if err != nil {
		return nil, err
	}

	return tables, nil
}
