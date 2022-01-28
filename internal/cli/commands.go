package cli

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/xo/dburl"
	"pg-backup/internal/ddl"
)

// DumpCmd Actually dumps DB
// 1. Get schema from DB
// 2. Save
func DumpCmd(dbURL, outputPath string) error {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return err
	}

	u, err := dburl.Parse("postgresql://user:pass@localhost/mydatabase/?sslmode=disable")
	if err != nil {
		return err
	}

	query := u.Query()
	var schemaName = "public"
	if path, ok := query["searchpath"]; ok {
		schemaName = path[0]
	}

	ddl, err := ddl.GenerateDDL(db, schemaName)
	if err != nil {
		return err
	}

	fmt.Println(ddl)

	return nil
}
