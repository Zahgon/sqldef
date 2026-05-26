package mssql

import (
	"database/sql"
	"regexp"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/sqldef/sqldef/v3/database"
)

const indent = "    "

type databaseInfo struct {
	tableName   []string
	columns     map[string][]column
	indexDefs   map[string][]*indexDef
	foreignDefs map[string][]string
	checkDefs   map[string][]string
}

type MssqlDatabase struct {
	config          database.Config
	db              *sql.DB
	defaultSchema   *string
	info            databaseInfo
	generatorConfig database.GeneratorConfig
}

func NewDatabase(config database.Config) (database.Database, error) {
	_ = "STUB: not implemented"
	return *new(database.Database), nil
}

func (d *MssqlDatabase) ExportDDLs() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *MssqlDatabase) updateDatabaesInfo() error { _ = "STUB: not implemented"; return nil }

func (d *MssqlDatabase) updateTableNames() error { _ = "STUB: not implemented"; return nil }

func (d *MssqlDatabase) tableNames() []string { _ = "STUB: not implemented"; return nil }

func (d *MssqlDatabase) exportTableDDL(table string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *MssqlDatabase) buildExportTableDDL(table string, columns []column, indexDefs []*indexDef, foreignDefs []string, checkDefs []string) string {
	_ = "STUB: not implemented"
	return ""
}

// PRIMARY KEY

// UNIQUE CONSTRAINTS (that were created as constraints, not indexes)

// Skip unique constraints - they're already handled inside the table definition

type column struct {
	Name        string
	dataType    string
	MaxLength   string
	Scale       string
	Nullable    bool
	Identity    *identity
	DefaultName string
	DefaultVal  string
	Check       *check
}

func (c column) getLength() (string, bool) { _ = "STUB: not implemented"; return "", false }

// The default precision is 18.

type identity struct {
	SeedValue         string
	IncrementValue    string
	NotForReplication bool
}

type check struct {
	Name              string
	Definition        string
	NotForReplication bool
}

func (d *MssqlDatabase) updateColumns() error { _ = "STUB: not implemented"; return nil }

func (d *MssqlDatabase) getColumns(table string) []column { _ = "STUB: not implemented"; return nil }

type indexDef struct {
	name       string
	columns    []string
	primary    bool
	unique     bool
	constraint bool
	indexType  string
	filter     *string
	included   []string
	options    []indexOption
}

type indexOption struct {
	name  string
	value string
}

func (d *MssqlDatabase) updateIndexDefs() error { _ = "STUB: not implemented"; return nil }

// `sys.stats.is_incremental` only exists SQL Server 2014 (12.x) and above.
// https://learn.microsoft.com/en-us/sql/relational-databases/system-catalog-views/sys-stats-transact-sql?view=sql-server-ver16

func (d *MssqlDatabase) getIndexDefs(table string) []*indexDef {
	_ = "STUB: not implemented"
	return nil
}

// foreignKeyInfo holds the aggregated information for a composite foreign key
type foreignKeyInfo struct {
	constraintName    string
	columns           []string
	foreignTableName  string
	foreignColumns    []string
	foreignUpdateRule string
	foreignDeleteRule string
	notForReplication bool
}

func (d *MssqlDatabase) updateForeignDefs() error {
	_ = "STUB: not implemented"
	// Order by constraint_column_id to get columns in the correct order for composite keys
	return nil
}

// First pass: aggregate columns by constraint name
// Key: schema.table:constraint_name

// schema.table -> list of constraint names (in order)

// Add column to existing FK

// New FK

// Second pass: generate FK definitions

// Build column list: [col1],[col2],...

func (d *MssqlDatabase) getForeignDefs(table string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (d *MssqlDatabase) updateCheckDefs() error {
	_ = "STUB: not implemented"
	// Query for table-level CHECK constraints only (those without a parent_column_id)
	// Column-level CHECKs are handled in updateColumns()
	return nil
}

func (d *MssqlDatabase) getCheckDefs(table string) []string { _ = "STUB: not implemented"; return nil }

func boolToOnOff(in bool) string { _ = "STUB: not implemented"; return "" }

var (
	suffixSemicolon = regexp.MustCompile(`;$`)
	spaces          = regexp.MustCompile(`[ ]+`)
	lineComment     = regexp.MustCompile(`(?m)--.*$`)
)

func (d *MssqlDatabase) views() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

//XXX - Line comments should be removed before removing newlines.

func (d *MssqlDatabase) triggers() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *MssqlDatabase) DB() *sql.DB { _ = "STUB: not implemented"; return nil }

func (d *MssqlDatabase) Close() error { _ = "STUB: not implemented"; return nil }

func (d *MssqlDatabase) GetDefaultSchema() string { _ = "STUB: not implemented"; return "" }

func mssqlBuildDSN(config database.Config) string { _ = "STUB: not implemented"; return "" }

func splitTableName(table string, defaultSchmea string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func forceQuoteName(name string) string { _ = "STUB: not implemented"; return "" }

func (d *MssqlDatabase) quoteIdentifier(name string) string {
	_ = "STUB: not implemented"
	// SQL Server exports are canonicalized with brackets so round-trips are stable
	// even when the original schema mixed quoted and unquoted identifier styles.
	return ""
}

func (d *MssqlDatabase) joinQuotedIdentifiers(names []string) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *MssqlDatabase) SetGeneratorConfig(config database.GeneratorConfig) {
	_ = "STUB: not implemented"
	return
}

func (d *MssqlDatabase) GetGeneratorConfig() database.GeneratorConfig {
	_ = "STUB: not implemented"
	return *new(database.GeneratorConfig)
}

func (d *MssqlDatabase) GetTransactionQueries() database.TransactionQueries {
	_ = "STUB: not implemented"
	return *new(database.TransactionQueries)
}

func (d *MssqlDatabase) GetConfig() database.Config {
	_ = "STUB: not implemented"
	return *new(database.Config)
}
