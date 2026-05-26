// This package has SQL parser, its abstraction and SQL generator.
// Never touch database.
package schema

import (
	"github.com/sqldef/sqldef/v3/database"
	"github.com/sqldef/sqldef/v3/parser"
)

// Parse `ddls`, which is expected to `;`-concatenated DDLs
// and not to include destructive DDL.
func ParseDDLs(mode GeneratorMode, sqlParser database.Parser, sql string, defaultSchema string) ([]DDL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if this is a MultiStatement (e.g., from multi-table GRANT)

// Expand MultiStatement into individual DDLs

// Parse DDL like `CREATE TABLE` or `ALTER TABLE`.
// This doesn't support destructive DDL like `DROP TABLE`.
func parseDDL(mode GeneratorMode, ddl string, stmt parser.Statement, defaultSchema string) (DDL, error) {
	_ = "STUB: not implemented"
	return *new(DDL), nil
}

// Handle PARTITION OF tables differently

// TODO: handle other create DDL as error?

// Store raw AST and DDL; normalization is handled by the generator

// Normalize privilege names to uppercase for consistency

// For now, return the first grantee as a single statement

// Normalize privilege names to uppercase for consistency

// SET & USE statements are parsed but ignored - they're session-level settings, not schema objects

// parsePartitionOf handles CREATE TABLE ... PARTITION OF statements
func parsePartitionOf(mode GeneratorMode, stmt *parser.DDL, defaultSchema string, rawDDL string) (*CreatePartitionOf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseTable(mode GeneratorMode, stmt *parser.DDL, defaultSchema string, rawDDL string) (Table, error) {
	_ = "STUB: not implemented"
	return *new(Table), nil
}

// Normalize PostgreSQL type aliases from generic parser

// references is used for:
// 1. Schema-qualified type names (e.g., "public." for public.mytype) - stored with trailing dot
// 2. Simple REFERENCES clause without column names (e.g., "REFERENCES table_name")

// For simple REFERENCES (without column names), store the table name in references
// This is separate from the ForeignKey logic which handles REFERENCES with explicit columns

// Handle short timezone forms: timestamptz -> timestamp, timetz -> time
// The generic parser parses these as custom identifiers without setting Timezone flag.
// Clear typeIdent so the quote-aware comparison path doesn't see a stale name.

// Handle schema-qualified types from generic parser
// Generic parser stores "schema.type" in typeName field
// pgquery parser stores "schema." in references and "type" in typeName
// Normalize to the pgquery format for consistent comparison

// Store schema with trailing dot to match pgquery format

// FIXME: tight coupling in enum order

// Parse @renamed annotation for each column

// Convert inline foreign key references to ForeignKey objects
// This handles syntax like: column_name TYPE REFERENCES table_name(column_name)
// Note: We only convert when reference columns are explicitly specified.
// If not specified (e.g., "REFERENCES table_name"), we leave it as-is
// for the database-specific parser to handle (it will infer the primary key).

// Skip if no inline foreign key reference or if it's missing column names
// (empty ReferenceNames means it will use the primary key, which is database-specific)

// Build the foreign key object

// Leave constraint name empty for inline FK references.
// The generator will handle matching by columns against existing FKs
// (which may have auto-generated names like MySQL's "table_ibfk_N",
// PostgreSQL's "table_column_fkey", or MSSQL's "FK__table__column__...")
// or generate an appropriate name if creating a new FK.

// Only create constraintOptions if DEFERRABLE or INITIALLY DEFERRED is explicitly set to true
// This ensures we don't create an empty constraintOptions struct that would differ from
// database-parsed FKs (which have nil constraintOptions when not deferrable)

// Clear the references field from the column since it's now represented as a foreign key
// This prevents it from being used for type qualification

// MSSQL and MySQL: all columns participating in a PRIMARY KEY constraint have their nullability set to NOT NULL
// MSSQL: https://learn.microsoft.com/en-us/sql/relational-databases/tables/create-primary-keys#limitations
// MySQL: https://dev.mysql.com/doc/refman/8.4/en/create-table.html

// Auto-generate index/constraint name based on database conventions

// For MySQL or multi-column constraints, use just the column name
// Auto-generated names are unquoted

// Determine if this is a constraint
// Constraints have constraintOptions (set when CONSTRAINT keyword is used)
// For PostgreSQL: Constraints have constraintOptions

// For MSSQL, PRIMARY KEY is always a constraint

// Mark as constraint based on database-specific logic

// Parse @renamed annotation for this index

// Parse partition information

func parseIndex(stmt *parser.DDL, rawDDL string, mode GeneratorMode) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

// remove root paren expression

// Use PostgreSQL naming convention for UNIQUE constraints

// Auto-generated names are unquoted

// Extract index comments and look for @renamed annotation

// not supported in parser yet

func mustConvertToInt(val string) int { _ = "STUB: not implemented"; return 0 }

func mustConvertToFloat(val string) float64 { _ = "STUB: not implemented"; return 0 }

func parseValue(val *parser.SQLVal) *Value { _ = "STUB: not implemented"; return nil }

// Assume an integer length. Maybe useful only for index lengths.
// TODO: Change IndexColumn.Length in parser.y to integer in the first place
func parseLength(val *parser.SQLVal) (*int, error) { _ = "STUB: not implemented"; return nil, nil }

func parseIdentity(opt *parser.IdentityOpt) *Identity { _ = "STUB: not implemented"; return nil }

func parseDefaultDefinition(opt *parser.DefaultDefinition) *DefaultDefinition {
	_ = "STUB: not implemented"
	return nil
}

func parseSridDefinition(opt *parser.SridDefinition) *SridDefinition {
	_ = "STUB: not implemented"
	return nil
}

func parseIdentitySequence(opt *parser.IdentityOpt) *Sequence {
	_ = "STUB: not implemented"
	return nil
}

func parseGenerated(genc *parser.GeneratedColumn) *Generated { _ = "STUB: not implemented"; return nil }

func parseExclusion(exclusion *parser.ExclusionDefinition) Exclusion {
	_ = "STUB: not implemented"
	return *new(Exclusion)
}

// PostgreSQL defaults to btree when no index method is specified
// Normalize to lowercase to match PostgreSQL's pg_get_constraintdef output

// normalizeQualifiedName creates a QualifiedName from a parser.TableName
func normalizeQualifiedName(mode GeneratorMode, tableName parser.TableName, defaultSchema string) QualifiedName {
	_ = "STUB: not implemented"
	return *new(QualifiedName)
}

// normalizeQualifiedObjectName creates a QualifiedName from a parser.ObjectName
func normalizeQualifiedObjectName(mode GeneratorMode, objectName parser.ObjectName, defaultSchema string) QualifiedName {
	_ = "STUB: not implemented"
	return *new(QualifiedName)
}

// normalizeColNameToQualifiedName creates a QualifiedName from a parser.ColName.
// This is used for trigger names which can be schema-qualified like [dbo].[trigger_name].
// Unlike table names, trigger names do not get a default schema if none was specified.
func normalizeColNameToQualifiedName(mode GeneratorMode, colName *parser.ColName, defaultSchema string) QualifiedName {
	_ = "STUB: not implemented"
	return *new(QualifiedName)
}

// ColName.Qualifier is a TableName; for trigger schema, the schema is in Qualifier.Schema
// For [dbo].[insert_log], Qualifier.Schema = "dbo", Qualifier.Name = ""

// Fallback: if Schema is empty but Name is set, use Name as schema

// Note: unlike tables, triggers don't get a default schema - if no schema was specified,
// we leave it empty to preserve the original behavior

func normalizedTable(mode GeneratorMode, tableName Ident, defaultSchema string) Ident {
	_ = "STUB: not implemented"
	return *new(Ident)
}

// avoid qualifying empty references (e.g., built-in types)

// Replace pseudo collation "binary" with "{charset}_bin"
func normalizeCollate(collate string, table parser.TableSpec) string {
	_ = "STUB: not implemented"
	return ""
}

// Convert back `type BoolVal bool`
func castBool(val parser.BoolVal) bool { _ = "STUB: not implemented"; return false }

func castBoolPtr(val *parser.BoolVal) *bool { _ = "STUB: not implemented"; return nil }

// extractRenameFrom extracts the old name from a @renamed annotation.
// Returns an Ident with both the name and whether it was quoted.
// e.g., `@renamed from="OldName"` -> Ident{Name: "OldName", Quoted: true}
// e.g., `@renamed from=oldname` -> Ident{Name: "oldname", Quoted: false}
func extractRenameFrom(comment string) Ident {
	_ = "STUB: not implemented"
	// First try to match @renamed (preferred)
	return *new(Ident)
}

// If @renamed not found, try @rename (deprecated) for backward compatibility

// If @rename is found, issue a deprecation warning

// The regex has 2 capture groups (double quotes or unquoted)
// matches[0] = full match, matches[1] = quoted, matches[2] = unquoted

// double-quoted identifier

// unquoted identifier

// parseEnumValuesWithRename parses enum values from DDL and extracts @renamed annotations.
// It looks for patterns like: 'value' /* @renamed from=old_value */ or 'value' -- @renamed from=old_value
func parseEnumValuesWithRename(ddl string, rawValues []string, mode GeneratorMode) []EnumValue {
	_ = "STUB: not implemented"
	return nil
}

// Extract comments using tokenizer-based approach

// Match comments to enum values

// extractEnumValueComments extracts inline comments from a CREATE TYPE ... AS ENUM statement
// and maps them to enum values
func extractEnumValueComments(rawDDL string, mode GeneratorMode) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// EOF

// Track CREATE TYPE ... AS ENUM sequence

// Reset if CREATE is not followed by TYPE

// Wait for opening parenthesis

// Reset AS if not followed by ENUM

// Track parentheses for enum values

// Skip table name parentheses, wait for ENUM keyword

// After a comma, expect a new enum value
// Don't clear currentEnumValue yet - the comment might come after the comma

// Strip quotes from the value for the map key

// generatorModeToParserMode converts GeneratorMode to ParserMode
func generatorModeToParserMode(mode GeneratorMode) parser.ParserMode {
	_ = "STUB: not implemented"
	return *new(parser.ParserMode)
}

func extractTableComment(rawDDL string, mode GeneratorMode) string {
	_ = "STUB: not implemented"
	return ""
}

// Store the first comment after CREATE TABLE

// EOF

// Look for CREATE keyword

// Look for TABLE keyword after CREATE

// After CREATE TABLE, capture the first comment we encounter
// This could be before or after the opening parenthesis

// Continue scanning to handle all cases

// Reset if we found CREATE but next token is not TABLE

// extractColumnComments extracts inline comments (-- comments) from a CREATE TABLE statement
// and maps them to column names
func extractColumnComments(rawDDL string, mode GeneratorMode) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// EOF

// Track CREATE TABLE statements

// Reset if we found CREATE but next token is not TABLE

// Track parentheses depth to know when we're inside column definitions

// After a comma inside the table definition, expect a new column

// Don't clear currentColumnName yet - the comment might come after the comma

// Capture potential column name at the start of a column definition

// Associate comment with the current column name
// Comments can appear after the column definition but before the next column

// Only store if we haven't already stored a comment for this column

func extractIndexComments(rawDDL string, mode GeneratorMode) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// EOF

// Track CREATE TABLE statements

// Scan ahead to see if it's CREATE TABLE

// Track parentheses depth to know when we're inside table definition

// After a comma inside the table definition, reset index tracking

// Don't clear currentIndexName yet - the comment might come after the comma

// Found an INDEX or KEY keyword inside CREATE TABLE

// Found UNIQUE keyword which might be followed by INDEX or KEY

// This is a CONSTRAINT ... UNIQUE definition
// Use the constraint name as the index name

// CONSTRAINT can be followed by a name and then UNIQUE, which creates an index

// These indicate other types of constraints, not regular indexes

// Capture potential index name or constraint name

// This is the constraint name

// Keep afterConstraintKeyword true to catch UNIQUE keyword next

// This is the index name

// Check if this ID is "KEY" or "INDEX"

// Next ID will be the index name

// This is the index name for UNIQUE without KEY/INDEX keyword

// Associate comment with the current index name

// Only store if we haven't already stored a comment for this index

// Now handle standalone CREATE INDEX statements

// EOF

// UNIQUE can appear after CREATE

// Continue looking for INDEX

// Part of CREATE INDEX IF NOT EXISTS
// Next tokens will be NOT and EXISTS

// Part of IF NOT EXISTS

// This is the index name

// Associate comment with the index from CREATE INDEX statement

// Only store if we haven't already stored a comment for this index

// Reset for next potential index

// After ON keyword, we're past the index name
