package main

import (
	"log"
	"os"

	"github.com/sqldef/sqldef/v3"
	"github.com/sqldef/sqldef/v3/database"
	"github.com/sqldef/sqldef/v3/database/file"
	"github.com/sqldef/sqldef/v3/database/mssql"
	"github.com/sqldef/sqldef/v3/schema"
	"github.com/sqldef/sqldef/v3/util"
)

// Return parsed options and schema filename
// TODO: Support `sqldef schema.sql -opt val...`
func parseOptions(args []string) (database.Config, *sqldef.Options) {
	_ = "STUB: not implemented"
	// MSSQL default: legacy_ignore_quotes is true (legacy mode)
	return *new(database.Config), nil
}

// Custom handlers for config flags to preserve order

// merge --config and --config-inline in order

func main() {
	util.InitSlog()

	config, options := parseOptions(os.Args[1:])

	var db database.Database
	if len(options.CurrentFile) > 0 {
		db = file.NewDatabase(options.CurrentFile)
	} else {
		var err error
		db, err = mssql.NewDatabase(config)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	}

	sqlParser := mssql.NewParser()
	sqldef.Run(schema.GeneratorModeMssql, db, sqlParser, options)
}
