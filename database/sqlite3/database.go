package sqlite3

import (
	"database/sql"

	"github.com/sqldef/sqldef/v3/database"
	_ "modernc.org/sqlite"
)

type Sqlite3Database struct {
	config          database.Config
	db              *sql.DB
	generatorConfig database.GeneratorConfig
}

func NewDatabase(config database.Config) (database.Database, error) {
	_ = "STUB: not implemented"
	return *new(database.Database), nil
}

func (d *Sqlite3Database) ExportDDLs() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *Sqlite3Database) tableNames() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Sqlite3Database) exportTableDDL(table string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *Sqlite3Database) views() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Sqlite3Database) indexes() ([]string, error) {
	_ = "STUB: not implemented"

	// Exclude automatically generated indexes for unique constraint
	return nil, nil
}

func (d *Sqlite3Database) triggers() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Sqlite3Database) DB() *sql.DB { _ = "STUB: not implemented"; return nil }

func (d *Sqlite3Database) Close() error { _ = "STUB: not implemented"; return nil }

func (d *Sqlite3Database) GetDefaultSchema() string { _ = "STUB: not implemented"; return "" }

func (d *Sqlite3Database) SetGeneratorConfig(config database.GeneratorConfig) {
	_ = "STUB: not implemented"
	return
}

func (d *Sqlite3Database) GetGeneratorConfig() database.GeneratorConfig {
	_ = "STUB: not implemented"
	return *new(database.GeneratorConfig)
}

func (d *Sqlite3Database) GetTransactionQueries() database.TransactionQueries {
	_ = "STUB: not implemented"
	return *new(database.TransactionQueries)
}

func (d *Sqlite3Database) GetConfig() database.Config {
	_ = "STUB: not implemented"
	return *new(database.Config)
}
