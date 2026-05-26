package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sqldef/sqldef/v3/database/file"

	"github.com/sqldef/sqldef/v3"
	"github.com/sqldef/sqldef/v3/database"
	"github.com/sqldef/sqldef/v3/database/postgres"
	"github.com/sqldef/sqldef/v3/schema"
	"github.com/sqldef/sqldef/v3/util"
)

// Return parsed options and schema filename
// TODO: Support `sqldef schema.sql -opt val...`
func parseOptions(args []string) (database.Config, *sqldef.Options) {
	_ = "STUB: not implemented"
	// PostgreSQL-specific default: legacy_ignore_quotes is true for backward compatibility
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
		db, err = postgres.NewDatabase(config)

		// Emulate the default behavior (sslmode=prefer) of psql when PGSSLMODE is not set,
		// which is not supported by Go's lib/pq.
		if _, ok := os.LookupEnv("PGSSLMODE"); !ok && err == nil {
			e := db.DB().Ping()
			if e != nil && strings.Contains(fmt.Sprintf("%s", e), "SSL is not enabled") {
				db.Close()
				os.Setenv("PGSSLMODE", "disable")
				db, err = postgres.NewDatabase(config)
			}
		}

		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	}

	sqlParser := postgres.NewParser()
	sqldef.Run(schema.GeneratorModePostgres, db, sqlParser, options)
}
