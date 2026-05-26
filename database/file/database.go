package file

import (
	"database/sql"

	"github.com/sqldef/sqldef/v3/database"
)

// Pseudo database for comparison between files
type FileDatabase struct {
	file            string
	generatorConfig database.GeneratorConfig
}

func NewDatabase(file string) *FileDatabase { _ = "STUB: not implemented"; return nil }

func (f FileDatabase) ExportDDLs() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f FileDatabase) DB() *sql.DB { _ = "STUB: not implemented"; return nil }

func (f FileDatabase) Close() error { _ = "STUB: not implemented"; return nil }

func (f FileDatabase) GetDefaultSchema() string { _ = "STUB: not implemented"; return "" }

func (d *FileDatabase) SetGeneratorConfig(config database.GeneratorConfig) {
	_ = "STUB: not implemented"
	return
}

func (d *FileDatabase) GetGeneratorConfig() database.GeneratorConfig {
	_ = "STUB: not implemented"
	return *new(database.GeneratorConfig)
}

func (d *FileDatabase) GetTransactionQueries() database.TransactionQueries {
	_ = "STUB: not implemented"
	return *new(database.TransactionQueries)
}

func (d *FileDatabase) GetConfig() database.Config {
	_ = "STUB: not implemented"
	return *new(database.Config)
}
