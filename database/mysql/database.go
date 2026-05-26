package mysql

import (
	"database/sql"

	"github.com/sqldef/sqldef/v3/database"
)

type MysqlDatabase struct {
	config              database.Config
	db                  *sql.DB
	lowerCaseTableNames int // 0 = case-sensitive, 1 or 2 = case-insensitive
	generatorConfig     database.GeneratorConfig
}

func NewDatabase(config database.Config) (database.Database, error) {
	_ = "STUB: not implemented"
	return *new(database.Database), nil
}

// Query MySQL version and lower_case_table_names for case sensitivity handling

// queryMySQLServerInfo logs the MySQL version and returns the lower_case_table_names setting.
// This helps debug case sensitivity issues since MySQL behavior differs:
// - lower_case_table_names=0 (Linux default): Case-sensitive table names
// - lower_case_table_names=1 or 2 (macOS/Windows): Case-insensitive table names
func queryMySQLServerInfo(db *sql.DB) int { _ = "STUB: not implemented"; return 0 }

// Default to case-sensitive

func (d *MysqlDatabase) ExportDDLs() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *MysqlDatabase) tableNames() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *MysqlDatabase) exportTableDDL(table string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TODO: escape table name

func (d *MysqlDatabase) views() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *MysqlDatabase) triggers() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// can be NULL when the trigger is migrated from MySQL 5.6 to 5.7

func (d *MysqlDatabase) DB() *sql.DB { _ = "STUB: not implemented"; return nil }

func (d *MysqlDatabase) Close() error { _ = "STUB: not implemented"; return nil }

func (d *MysqlDatabase) GetDefaultSchema() string { _ = "STUB: not implemented"; return "" }

func mysqlBuildDSN(config database.Config) string { _ = "STUB: not implemented"; return "" }

func registerTLSConfig(pemPath string) error { _ = "STUB: not implemented"; return nil }

func (d *MysqlDatabase) SetGeneratorConfig(config database.GeneratorConfig) {
	_ = "STUB: not implemented"
	return
}

func (d *MysqlDatabase) GetGeneratorConfig() database.GeneratorConfig {
	_ = "STUB: not implemented"
	return *new(database.GeneratorConfig)
}

func (d *MysqlDatabase) GetTransactionQueries() database.TransactionQueries {
	_ = "STUB: not implemented"
	return *new(database.TransactionQueries)
}

func (d *MysqlDatabase) GetConfig() database.Config {
	_ = "STUB: not implemented"
	return *new(database.Config)
}
