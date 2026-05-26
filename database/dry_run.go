package database

import (
	"database/sql"
	"database/sql/driver"
)

type DryRunDatabase struct {
	wrapped         Database
	dryRunDB        *sql.DB
	generatorConfig GeneratorConfig
}

func NewDryRunDatabase(db Database) (*DryRunDatabase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unique name per database instance

func (d *DryRunDatabase) ExportDDLs() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *DryRunDatabase) DB() *sql.DB { _ = "STUB: not implemented"; return nil }

func (d *DryRunDatabase) Close() error { _ = "STUB: not implemented"; return nil }

func (d *DryRunDatabase) GetDefaultSchema() string { _ = "STUB: not implemented"; return "" }

func (d *DryRunDatabase) SetGeneratorConfig(config GeneratorConfig) {
	_ = "STUB: not implemented"
	return
}

// Get the config back from wrapped in case it was modified (e.g., MySQL adds lowerCaseTableNames)

func (d *DryRunDatabase) GetGeneratorConfig() GeneratorConfig {
	_ = "STUB: not implemented"
	return *new(GeneratorConfig)
}

func (d *DryRunDatabase) GetTransactionQueries() TransactionQueries {
	_ = "STUB: not implemented"
	return *new(TransactionQueries)
}

func (d *DryRunDatabase) GetConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

type dryRunDriver struct {
	txQueries TransactionQueries
}

func (d *dryRunDriver) Open(name string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

type dryRunConn struct {
	txQueries TransactionQueries
}

func (c *dryRunConn) Prepare(query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

func (c *dryRunConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *dryRunConn) Begin() (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

type dryRunTx struct {
	txQueries TransactionQueries
}

func (tx *dryRunTx) Commit() error { _ = "STUB: not implemented"; return nil }

func (tx *dryRunTx) Rollback() error { _ = "STUB: not implemented"; return nil }

type dryRunStmt struct {
	query string
}

func (s *dryRunStmt) Close() error { _ = "STUB: not implemented"; return nil }

func (s *dryRunStmt) NumInput() int { _ = "STUB: not implemented"; return 0 }

func (s *dryRunStmt) Exec(args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (s *dryRunStmt) Query(args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

type dryRunResult struct{}

func (r *dryRunResult) LastInsertId() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *dryRunResult) RowsAffected() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

type dryRunRows struct {
	closed bool
}

func (r *dryRunRows) Columns() []string { _ = "STUB: not implemented"; return nil }

func (r *dryRunRows) Close() error { _ = "STUB: not implemented"; return nil }

func (r *dryRunRows) Next(dest []driver.Value) error { _ = "STUB: not implemented"; return nil }
