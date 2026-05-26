// This package has database database layer. Never deal with DDL construction.
package database

import (
	"database/sql"

	"github.com/sqldef/sqldef/v3/parser"
)

type Config struct {
	DbName        string
	User          string
	Password      string
	Host          string
	Port          int
	Socket        string
	SkipView      bool
	SkipExtension bool
	SkipPartition bool

	// Only MySQL
	MySQLEnableCleartextPlugin bool

	// Only MySQL and PostgreSQL
	SslMode string

	// Only MySQL
	SslCa string

	// Only PostgreSQL
	TargetSchema []string

	// Only MySQL and PostgreSQL
	DumpConcurrency int

	// Only PostgreSQL, especially for Aurora DSQL limitation
	DisableDdlTransaction bool

	// Only MSSQL
	TrustedConnection bool   // Use Windows authentication
	Instance          string // Instance name
	TrustServerCert   bool   // Trust server certificate
}

type GeneratorConfig struct {
	TargetTables            []string
	SkipTables              []string
	SkipViews               []string
	TargetSchema            []string
	Algorithm               string
	Lock                    string
	DumpConcurrency         int
	ManagedRoles            []string // Roles whose privileges are managed by sqldef (empty means no privileges are managed)
	EnableDrop              bool     // Whether to enable DROP/REVOKE operations
	CreateIndexConcurrently bool     // Whether to add CONCURRENTLY to CREATE INDEX statements
	DisableDdlTransaction   bool     // Do not use a transaction for DDL statements
	LegacyIgnoreQuotes      bool     // true = ignore quotes (legacy), false = preserve quotes

	// MySQL-specific: value of lower_case_table_names server variable.
	// 0 = case-sensitive (Linux default), 1 or 2 = case-insensitive (Windows/macOS).
	// Default is 0 (case-sensitive) for offline mode compatibility.
	MysqlLowerCaseTableNames int
}

type TransactionQueries struct {
	Begin    string
	Commit   string
	Rollback string
}

// Ident is an alias for parser.Ident.
// Represents an identifier with quote information for quote-aware identifier handling.
type Ident = parser.Ident

// NewIdent is an alias for parser.NewIdent.
var NewIdent = parser.NewIdent

// NewIdentWithQuoteDetected creates an Ident with the Quoted flag inferred from content:
//   - If the name contains uppercase letters, it must have been quoted
//     (PostgreSQL folds unquoted identifiers to lowercase)
//   - If the name contains special characters (dots, spaces, etc.), it requires quoting
//   - If the name is all lowercase without special chars, it's treated as unquoted.
//     This is correct because in PostgreSQL, "users" (quoted lowercase) and users
//     (unquoted) are semantically equivalent and can be referenced interchangeably.
//
// Note: This does NOT check for reserved keywords. Keywords are handled separately
// at DDL output time because:
//   - The Quoted flag represents whether quoting is needed to preserve the identifier's form
//   - Keyword escaping is a SQL syntax requirement, not an identifier property
//
// Use this for identifiers from the database or auto-generated constraint names.
func NewIdentWithQuoteDetected(name string) Ident { _ = "STUB: not implemented"; return *new(Ident) }

// NeedsQuoting returns true if an identifier needs to be quoted in SQL output.
// This is the complete check for DDL generation, combining:
//   - Non-standard characters (uppercase, special chars, invalid start)
//   - Reserved keywords
//
// Use this when generating SQL output to determine if quoting is required.
func NeedsQuoting(name string) bool { _ = "STUB: not implemented"; return false }

// hasNonStandardChars returns true if an identifier contains characters
// that require quoting to preserve the identifier's form. This checks:
//   - Uppercase letters (PostgreSQL folds unquoted to lowercase)
//   - Special characters that aren't allowed in unquoted identifiers
//   - Invalid first character (must be letter or underscore)
//
// This does NOT check for reserved keywords - use NeedsQuoting for that.
func hasNonStandardChars(name string) bool { _ = "STUB: not implemented"; return false }

// Check if it has uppercase letters

// First character: must be letter or underscore

// Remaining characters: letters, digits, underscores, or $ are allowed

// NewNormalizedIdent normalizes an Ident for comparison:
//   - Quoted identifiers: preserve case, set Quoted based on whether name has uppercase
//   - Unquoted identifiers: normalize to lowercase, set Quoted=false
func NewNormalizedIdent(ident Ident) Ident { _ = "STUB: not implemented"; return *new(Ident) }

// QualifiedName represents a schema-qualified table name with quote information.
type QualifiedName struct {
	Schema Ident // empty if not specified (will use default schema)
	Name   Ident
}

// IsEmpty returns true if the qualified name has no name set.
func (q QualifiedName) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// RawString returns the raw qualified name as "schema.name" or just "name" if no schema.
// This is NOT escaped for SQL output and NOT normalized for comparison.
// Use this for logging, debugging, or map keys.
func (q QualifiedName) RawString() string { _ = "STUB: not implemented"; return "" }

// Abstraction layer for multiple kinds of databases
type Database interface {
	ExportDDLs() (string, error)
	DB() *sql.DB
	Close() error
	GetDefaultSchema() string
	SetGeneratorConfig(config GeneratorConfig)
	GetGeneratorConfig() GeneratorConfig
	GetTransactionQueries() TransactionQueries
	GetConfig() Config
}

func isDryRun(d Database) bool { _ = "STUB: not implemented"; return false }

func isSingleLineComment(s string) bool { _ = "STUB: not implemented"; return false }

func RunDDLs(d Database, ddls []string, beforeApply string, ddlSuffix string, logger Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// beforeApply is executed in transaction

// DDLs in transaction

// Skip commented DDLs (e.g., "-- Skipped: ...")

// Only commit if we started a transaction

// DDLs not in transaction

// Skip ddlSuffix and execution for commented DDLs (e.g., "-- Skipped: ...")

func TransactionSupported(ddl string) bool { _ = "STUB: not implemented"; return false }

func MergeGeneratorConfigs(configs []GeneratorConfig) GeneratorConfig {
	_ = "STUB: not implemented"
	return *new(GeneratorConfig)
}

func ParseGeneratorConfigString(yamlString string, defaults GeneratorConfig) GeneratorConfig {
	_ = "STUB: not implemented"
	return *new(GeneratorConfig)
}

func ParseGeneratorConfig(configFile string, defaults GeneratorConfig) GeneratorConfig {
	_ = "STUB: not implemented"
	return *new(GeneratorConfig)
}

// MergeGeneratorConfig merges two configs, with the second one taking precedence
func MergeGeneratorConfig(base, override GeneratorConfig) GeneratorConfig {
	_ = "STUB: not implemented"

	// Override fields if they are set in the override config
	return *new(GeneratorConfig)
}

// LegacyIgnoreQuotes: override always takes precedence (set by first config with database-specific default)

func parseGeneratorConfigFromBytes(buf []byte, defaults GeneratorConfig) GeneratorConfig {
	_ = "STUB: not implemented"
	return *new(GeneratorConfig)
}

// Use the provided default, override if explicitly set in config
