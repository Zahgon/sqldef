package postgres

import (
	pgquery "github.com/pganalyze/pg_query_go/v6"
	"github.com/sqldef/sqldef/v3/database"
	"github.com/sqldef/sqldef/v3/parser"
)

// validationError is an error that should not trigger fallback to the generic parser
type validationError struct {
	msg string
}

func (e validationError) Error() string {
	_ = "STUB: not implemented"

	// PsqldefParserMode defines the parsing strategy for psqldef
	return ""
}

type PsqldefParserMode int

const (
	// PsqldefParserModeAuto is the default migration mode that prefers the generic parser
	// but falls back to pgquery when needed. This mode helps with gradual migration:
	// 1. First tries generic parser on the full SQL
	// 2. If that fails, uses pgquery to split SQL into statements
	// 3. For each statement, tries to convert pgquery AST to generic AST
	// 4. If conversion fails for a statement, tries generic parser on that statement
	PsqldefParserModeAuto PsqldefParserMode = iota
	// PsqldefParserModePgquery uses only pgquery without any fallback to generic parser
	PsqldefParserModePgquery
	// PsqldefParserModeGeneric uses only the generic parser without any fallback to pgquery
	PsqldefParserModeGeneric
)

type PostgresParser struct {
	parser database.GenericParser
	mode   PsqldefParserMode
}

func NewParser() PostgresParser { _ = "STUB: not implemented"; return *new(PostgresParser) }

func NewParserWithMode(mode PsqldefParserMode) PostgresParser {
	_ = "STUB: not implemented"
	return *new(PostgresParser)
}

func (p PostgresParser) Parse(sql string) ([]database.DDLStatement, error) {
	_ = "STUB: not implemented"
	// Workaround for comments (not needed?)
	//re := regexp.MustCompilePOSIX("^ *--.*")
	//sql = re.ReplaceAllString(sql, "")
	return nil, nil
}

// If generic parser only mode is enabled, skip pgquery entirely

// If pgquery only mode is enabled, skip generic parser entirely

// Auto mode: Try generic parser first for faster path and better error messages.
// If generic parser succeeds, we're done. If it fails, fall back to pgquery
// which can handle more PostgreSQL-specific syntax and provides statement splitting.

// Generic parser couldn't handle this SQL, use pgquery as fallback.
// This is expected during the migration period as the generic parser
// may not support all PostgreSQL features yet.

// parsePgquery parses SQL using the pgquery parser
func (p PostgresParser) parsePgquery(sql string) ([]database.DDLStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attempt to convert pgquery AST to our generic AST format

// Check if this is a validation error (should not fallback)

// In Auto mode, if we can't convert the pgquery AST to generic AST,
// try parsing this individual statement with the generic parser directly.
// This handles cases where the generic parser can parse the statement
// but we haven't implemented the AST conversion from pgquery yet.

func (p PostgresParser) parseStmt(node *pgquery.Node) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

// In pgquery parser, CreateFunctionStmt is ignored.

func (p PostgresParser) parseCreateStmt(stmt *pgquery.CreateStmt) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

func (p PostgresParser) parseIndexStmt(stmt *pgquery.IndexStmt) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

// go_pgquery doesn't support ASYNC, will be set by generic parser

func (p PostgresParser) parseViewStmt(stmt *pgquery.ViewStmt) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

func (p PostgresParser) parseSelectStmt(stmt *pgquery.SelectStmt) (parser.SelectStatement, error) {
	_ = "STUB: not implemented"
	return *new(parser.SelectStatement), nil
}

// pgquery's DistinctClause contains the DISTINCT ON expressions, if any

func (p PostgresParser) parseResTarget(stmt *pgquery.ResTarget) (parser.SelectExpr, error) {
	_ = "STUB: not implemented"
	return *new(parser.SelectExpr), nil
}

func (p PostgresParser) parseExpr(stmt *pgquery.Node) (parser.Expr, error) {
	_ = "STUB: not implemented"
	return *new(parser.Expr), nil
}

// normalize

// Ignore table name for easy comparison

// For casts, use the raw type name from pgquery to preserve "bpchar" instead of "character"
// This matches what the generic parser produces from ::bpchar syntax

// Fallback to normalized type if raw extraction fails

// Convert lower case for compatibility with legacy parser

// IN operator is internally converted to = ANY (ARRAY[...]) in PostgreSQL

// For IN expressions, convert ValTuple to ArrayConstructor

// Convert ValTuple to ArrayConstructor

// Handle list of values (used in IN expressions)

func (p PostgresParser) parseIndexColumn(stmt *pgquery.Node) (parser.IndexColumn, error) {
	_ = "STUB: not implemented"
	return *new(parser.IndexColumn), nil
}

func (p PostgresParser) parseArrayElement(node parser.Expr) (parser.Expr, error) {
	_ = "STUB: not implemented"
	return *new(parser.Expr), nil
}

func (p PostgresParser) parseCommentStmt(stmt *pgquery.CommentStmt) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

// parseIdentList converts a pgquery list of strings to []parser.Ident.
// pgquery doesn't preserve quoting information, so we assume unquoted (false).
func (p PostgresParser) parseIdentList(list *pgquery.List) ([]parser.Ident, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pgquery doesn't preserve quoting info, assume unquoted

func (p PostgresParser) parseTableName(relation *pgquery.RangeVar) (parser.TableName, error) {
	_ = "STUB: not implemented"
	return *new(parser.TableName), nil
}

func (p PostgresParser) parseExtensionStmt(stmt *pgquery.CreateExtensionStmt) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

func (p PostgresParser) parseAlterTableStmt(stmt *pgquery.AlterTableStmt) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

func (p PostgresParser) parseConstraint(constraint *pgquery.Constraint, tableName parser.TableName) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

func (p PostgresParser) parseExclusion(constraint *pgquery.Constraint) (*parser.ExclusionDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If there's no expression, just use the column name as an expression

func (p PostgresParser) parseForeignKey(constraint *pgquery.Constraint) (*parser.ForeignKeyDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PostgresParser) parseFkAction(action string) parser.Ident {
	_ = "STUB: not implemented"
	// https://github.com/pganalyze/pg_query_go/blob/v2.2.0/parser/include/nodes/parsenodes.h#L2145-L2149C23
	return *new(parser.Ident)
}

// pgquery cannot distinguish between unspecified action and no action.
// Empty for no action to match existing behavior.

func (p PostgresParser) parseColumnDef(columnDef *pgquery.ColumnDef, tableName parser.TableName) (*parser.ColumnDefinition, *parser.ForeignKeyDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Postgres only supports stored generated column

func (p PostgresParser) parseDefaultValue(rawExpr *pgquery.Node) (*parser.DefaultDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Preserve CastExpr nodes in defaults (fix for incorrect normalization)
// PostgreSQL stores defaults with explicit casts, so we should preserve them
// Special case: Convert NullVal to SQLVal for consistency

// Mark as used
// Handle NULL::type casts by converting NullVal to SQLVal
// PostgreSQL represents NULL as NullVal in pg_query AST but we use SQLVal
// Preserve the cast by wrapping the SQLVal in a CastExpr
// Use lowercase null to match the lexer's keyword normalization

// For other CastExpr cases, check if the cast is semantically meaningful
// Strip redundant casts like ::text, ::character varying on string literals
// But preserve important casts like ::interval, ::bpchar, ::json, ::jsonb, ::integer[], and numeric casts

// Check if this is a redundant cast that should be stripped

// Preserve all other casts (interval, bpchar, json, jsonb, integer[], timestamp, numeric casts on strings, etc.)

// getRawTypeName extracts the raw type name from a TypeName node with minimal normalization.
// This preserves type names like "bpchar" instead of normalizing them to "character",
// but still normalizes PostgreSQL internal names like "int4" to "integer" for consistency.
func (p PostgresParser) getRawTypeName(node *pgquery.TypeName) string {
	_ = "STUB: not implemented"
	return ""
}

// Get the last name, skipping schema prefix like "pg_catalog"

// Normalize PostgreSQL internal type names to SQL standard names
// but preserve types like "bpchar" that should not be normalized

// Keep the type name as-is (including "bpchar", "timetz", "timestamptz", etc.)

func (p PostgresParser) parseTypeName(node *pgquery.TypeName) (parser.ColumnType, error) {
	_ = "STUB: not implemented"
	return *new(parser.ColumnType), nil
}

// For test compatibility, keep bool as bool.
// TODO: Delete this exception.

// TODO: use this pattern more, fixing failed tests as well

// TODO: Whitelist types explicitly. We're missing 'json' and 'text' at least.

// Schema-qualified type: store schema prefix with trailing dot in References.Name
// This is for backward compatibility with how the pgquery parser stores schema prefixes

func (p PostgresParser) parseTypmods(typmods []*pgquery.Node) ([]*parser.SQLVal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PostgresParser) parseStringList(list *pgquery.List) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p PostgresParser) parseCheckConstraint(constraint *pgquery.Constraint) (*parser.CheckDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PostgresParser) parseCreateSchemaStmt(stmt *pgquery.CreateSchemaStmt) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

func (p PostgresParser) parseCreatePolicyStmt(stmt *pgquery.CreatePolicyStmt) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

// This is a workaround to handle cases where PostgreSQL automatically adds or removes type casting.
//
// Example:
//
// ```
// $ cat schema.sql
// CREATE TABLE test (
// t text CHECK (t ~ '[0-9]'),
// i integer CHECK (i = ANY (ARRAY[1,2,3]::integer[]))
// );
//
// $ psql sandbox < schema.sql
// $ psqldef sandbox --export
// CREATE TABLE "public"."test" (
// "t" text CONSTRAINT test_t_check CHECK (t ~ '[0-9]'::text),
// "i" integer CONSTRAINT test_i_check CHECK (i = ANY (ARRAY[1, 2, 3]))
// );
// ```
//
// Looking at the export result, PostgreSQL automatically adds `::text` type casting to '[0-9]',
// and removes `::integer[]` from `ARRAY[1,2,3]`. In such cases, if you don't remove the type casting,
// the generator will fail to calculate the diff.
//
// Ideally, the generator should be smart enough to handle the calculation of diff while keeping the type casting.
// However, as a workaround, it is handled by the parser.
//
// Since this function's support is not complete, updates will be necessary in the future.
func shouldDeleteTypeCast(sourceNode *pgquery.Node, targetType parser.ColumnType) bool {
	_ = "STUB: not implemented"
	return false
}

// Do not delete type cast from '{1,2,3}'::integer[]

// Delete type cast from '[0-9]'::text

// Delete type cast from '2022-01-01'::date

// Do not delete type cast from '1 day'::interval

// Delete type cast from ARRAY[1,2,3]::integer[]

func (p PostgresParser) parseGrantStmt(stmt *pgquery.GrantStmt) (parser.Statement, error) {
	_ = "STUB: not implemented"
	return *new(parser.Statement), nil
}

// For now, only support table grants

// Check for unsupported WITH GRANT OPTION

// Check for unsupported CASCADE/RESTRICT (for REVOKE)
// Note: DROP_RESTRICT is the default behavior and is allowed

// Handle multiple tables - return multiple DDL statements

// ALL PRIVILEGES case

// Always false since we error on WITH GRANT OPTION above
// Always false since we error on CASCADE/RESTRICT above

// If we have multiple statements, return a composite statement

// Return a MultiStatement wrapper for multiple tables
