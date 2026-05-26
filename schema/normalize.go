package schema

import (
	"github.com/sqldef/sqldef/v3/parser"
)

var (
	dataTypeAliases = map[string]string{
		"bool":    "boolean",
		"int":     "integer",
		"char":    "character",
		"numeric": "decimal",
		"varchar": "character varying",
	}
	postgresDataTypeAliases = map[string]string{
		"int2":   "smallint",
		"int4":   "integer",
		"int8":   "bigint",
		"float4": "real",
		"float8": "double precision",
		"float":  "double precision",
		"bpchar": "character",

		// Timezone type aliases (the timezone flag is stored separately in Column.timezone)
		"timestamptz":                 "timestamp",
		"timestamp with time zone":    "timestamp",
		"timestamp without time zone": "timestamp",
		"timetz":                      "time",
		"time with time zone":         "time",
		"time without time zone":      "time",
	}
	// PostgreSQL serial types to their underlying integer types
	postgresSerialTypes = map[string]string{
		"smallserial": "smallint",
		"serial":      "integer",
		"bigserial":   "bigint",
	}
	mssqlDataTypeAliases = map[string]string{}
	mysqlDataTypeAliases = map[string]string{
		"boolean": "tinyint",
	}
)

// effectiveTypeName returns the comparison-canonical type of a column, after applying
// equivalences that depend on the column's full definition (charset/collate/CHECK), not
// just its declared typeName.
//
// MariaDB does not have a native JSON storage type: `CREATE TABLE t (j JSON)` is rewritten
// internally to `j longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin CHECK (json_valid(j))`,
// and that is exactly what SHOW CREATE TABLE returns. Without this equivalence, every
// `--dry-run` and `--export` against a MariaDB database that contains JSON columns reports a
// spurious `CHANGE COLUMN ... json` for each one, because the desired side parses as `json`
// while the readback parses as `longtext`.
//
// The normalization is applied symmetrically: only when the *other* side is also a JSON
// column (either declared as `json` or already matching the MariaDB storage pattern). This
// preserves comparison parity on MySQL, where the same `longtext + inline CHECK` text in a
// user schema is stored with the CHECK lifted to a table-level constraint at readback —
// asymmetric inline-vs-table CHECK placement that would otherwise break this fix on MySQL.
func effectiveTypeName(col Column, other Column, mode GeneratorMode) string {
	_ = "STUB: not implemented"
	return ""
}

// isMariaDBJSONColumn returns true if col matches the longtext+utf8mb4_bin+json_valid()
// shape that MariaDB emits for JSON columns.
func isMariaDBJSONColumn(col Column) bool { _ = "STUB: not implemented"; return false }

// normalizeTypeName normalizes a type name using dataTypeAliases and mode-specific aliases.
// This is the central function for all type name normalization in the generator.
func normalizeTypeName(typeName string, mode GeneratorMode) string {
	_ = "STUB: not implemented"
	// Normalize to lowercase for case-insensitive comparison
	return ""
}

// Apply common aliases

// Apply database-specific aliases

// normalizeConvertType normalizes a ConvertType's type name.
// This handles type aliases like int -> integer and properly handles array types like int[] -> integer[]
func normalizeConvertType(convertType *parser.ConvertType, mode GeneratorMode) *parser.ConvertType {
	_ = "STUB: not implemented"
	return nil
}

// Check if the type is an array type (ends with [])

// For array types, normalize the base type and then re-append the []

// BuildPostgresConstraintName generates a constraint name following PostgreSQL's naming convention.
// It automatically truncates names to 63 characters (NAMEDATALEN - 1) using PostgreSQL's algorithm:
// - If column > 28 chars: reduce column to 28 first, then apply remaining overflow to table
// - If column == 28 chars and table <= 29 chars: truncate table
// - If column == 28 chars and table > 29 chars: truncate table
// - If column < 28 chars: truncate table
// In summary: when column <= 28, always truncate the table first
func buildPostgresConstraintName(tableName, columnName, suffix string) string {
	_ = "STUB: not implemented"
	return ""
}

// Column exceeds 28: reduce to 28 first, then put remaining overflow on table

// Column can only be reduced to 28, put the rest on table

// Column <= 28: always truncate table

// buildPostgresConstraintNameIdent builds a PostgreSQL auto-generated constraint name
// and returns it as an Ident with quote information inferred from case.
func buildPostgresConstraintNameIdent(tableName, columnName, suffix string) Ident {
	_ = "STUB: not implemented"
	return *new(Ident)
}

// buildMysqlForeignKeyName builds a MySQL auto-generated foreign key constraint name
// using the format {table}_{column}_fk (similar to PostgreSQL's {table}_{column}_fkey).
// MySQL's actual auto-generated names are {table}_ibfk_{N} but since we can't predict N,
// we use a column-based deterministic name instead.
func buildMysqlForeignKeyName(tableName, columnName string) string {
	_ = "STUB: not implemented"
	return ""
}

// buildMysqlForeignKeyNameIdent builds a MySQL auto-generated foreign key constraint name
// and returns it as an Ident.
func buildMysqlForeignKeyNameIdent(tableName, columnName string) Ident {
	_ = "STUB: not implemented"
	return *new(Ident)
}

// buildMssqlForeignKeyName builds a MSSQL auto-generated foreign key constraint name
// using the format FK_{table}_{column}.
// Note: This is NOT the same as MSSQL's native auto-generated names which use
// the format FK__{table}__{column}__{hash} (e.g., "FK__posts__user_id__4CA06362").
// We use a deterministic column-based name for consistency and predictability.
func buildMssqlForeignKeyName(tableName, columnName string) string {
	_ = "STUB: not implemented"
	return ""
}

// buildMssqlForeignKeyNameIdent builds a MSSQL auto-generated foreign key constraint name
// and returns it as an Ident.
// Note: This is NOT the same as MSSQL's native auto-generated names which include
// a hash suffix. We use a deterministic column-based name for consistency.
func buildMssqlForeignKeyNameIdent(tableName, columnName string) Ident {
	_ = "STUB: not implemented"
	return *new(Ident)
}

// normalizeCheckExpr normalizes a CHECK constraint expression AST for comparison
// mode parameter controls PostgreSQL-specific normalization (IN to ANY conversion)
func normalizeCheckExpr(expr parser.Expr, mode GeneratorMode) parser.Expr {
	_ = "STUB: not implemented"
	return *new(parser.Expr)
}

// Remove certain casts that PostgreSQL simplifies
// - text, character varying: Always removed
// - date, timestamp without time zone: Removed when cast from string literals
// - time, time without time zone: Kept but normalized (PostgreSQL preserves these for precision)

// Always remove text casts

// Remove date/timestamp casts from string literals
// PostgreSQL simplifies '2020-01-01'::date to '2020-01-01' in CHECK constraints

// Remove redundant array typecasts on ARRAY constructors
// PostgreSQL normalizes ARRAY['a'::varchar]::text[] to ARRAY['a'::varchar::text]
// by pushing down the array typecast to each element. Since we strip ::text casts
// on elements, we also need to strip ::text[] on the ARRAY constructor itself.
// e.g., ARRAY['a'::varchar]::text[] -> ARRAY['a'::varchar] (after stripping ::text[])
//       ARRAY['a'::varchar::text]   -> ARRAY['a'::varchar] (after stripping ::text on element)
// Empty arrays (ARRAY[]) need the typecast or PostgreSQL can't determine the type.

// Non-empty array with array typecast: strip the redundant typecast

// For time types, keep the cast but use normalized type name

// Unwrap parentheses around simple expressions (literals, column names, etc.)
// MSSQL/PostgreSQL may add unnecessary parens like (1) instead of 1 or (name) instead of name

// Normalize operands and unwrap unnecessary parentheses around them

// MySQL adds parentheses around each operand in AND chains, so unwrap them

// Normalize operands and unwrap unnecessary parentheses around them

// MySQL adds parentheses around each operand in OR chains, so unwrap them
// Always safe to unwrap in OR chains since OR has the lowest precedence

// Try to convert OR chain of equality comparisons to IN expression
// MSSQL transforms IN (a, b, c) to col=a OR col=b OR col=c
// We normalize back to IN for comparison

// The generic parser may parse "= ANY(ARRAY[...])" as a FuncExpr on the right side
// We need to normalize this to set the Any/All flags properly

// Convert "column = ANY(array)" to ComparisonExpr with Any=true

// Convert "column = ALL(array)" to ComparisonExpr with All=true

// Unwrap ParenExpr from right side for ANY/ALL to ensure consistent formatting
// The parser may create ParenExpr(ArrayConstructor) which formats as ANY(ARRAY
// We want to normalize to ArrayConstructor directly which formats as ANY (ARRAY

// Handle IN clauses based on mode

// PostgreSQL normalizes IN (values) to = ANY (ARRAY[values])

// Change operator and set ANY flag

// "not in"

// For other databases, keep IN but sort the tuple for consistent comparison

// For ANY/ALL expressions, normalize the array elements

// This means we just set the flag above from IN conversion
// Already handled

// Normalize existing ANY/ALL expressions (strip casts, preserve order)

// Normalize function name to lowercase (PostgreSQL convention)

// PostgreSQL normalizes BETWEEN to >= AND <=
// e.g., "score BETWEEN 0 AND 100" becomes "score >= 0 AND score <= 100"

// x BETWEEN a AND b -> x >= a AND x <= b

// x NOT BETWEEN a AND b -> x < a OR x > b

// Normalize column names while preserving quoting information:
// - Quoted identifiers that are NOT all lowercase preserve their case and remain quoted
// - Quoted identifiers that ARE all lowercase are normalized to unquoted (since "id" = id)
// - Unquoted identifiers are normalized to lowercase

// Strip the type prefix for all temporal literals.
// PostgreSQL's pg_get_constraintdef emits typed literals for time/timestamp/date
// when comparing to typed columns, but user SQL often writes bare literals.
// Dropping the type prefix on both sides makes them compare equal.

// For all other expression types (literals, etc.), return as-is

// normalizeExpr normalizes an expression.
// This is similar to normalizeCheckExpr but tailored for other contexts.
func normalizeExpr(expr parser.Expr, mode GeneratorMode) parser.Expr {
	_ = "STUB: not implemented"
	return *new(parser.Expr)
}

// Normalize column name and qualifier while preserving quoting:
// - Quoted identifiers that are NOT all lowercase preserve their case and remain quoted
// - Quoted identifiers that ARE all lowercase are normalized to unquoted (since "id" = id)
// - Unquoted identifiers are normalized to lowercase
// For Postgres and MySQL, remove table qualifiers (e.g., "users.name" -> "name")

// For Postgres and MySQL, remove table qualifiers

// For PostgreSQL, normalize date/time function calls to keywords
// The generic parser parses CURRENT_DATE in parentheses as a function call,
// but without parentheses as a keyword (SQLVal with ValArg type)
// e.g., (CURRENT_DATE) -> current_date(), but CURRENT_DATE -> current_date

// For Postgres, check for ARRAY constructors BEFORE normalizing
// PostgreSQL standardizes function arguments to use ARRAY['a', 'b'] notation
// but users may write them expanded as 'a', 'b', so we expand for comparison
// e.g., jsonb_extract_path_text(payload, ARRAY['amount']) -> jsonb_extract_path_text(payload, 'amount')
// e.g., jsonb_extract_path_text(payload, ARRAY['a', 'b']) -> jsonb_extract_path_text(payload, 'a', 'b')
// However, do NOT expand for ANY/ALL/SOME functions as they require the ARRAY constructor

// Expand ARRAY elements into separate normalized arguments

// Not an ARRAY, normalize normally

// Normalize function name to lowercase for PostgreSQL (PostgreSQL stores functions in lowercase)
// For MySQL, preserve the original case as MySQL preserves case for function names

// For PostgreSQL, unwrap unnecessary parentheses around simple expressions in casts
// PostgreSQL adds parentheses like (amount)::numeric, but we want to normalize to amount::numeric

// Remove redundant casts that PostgreSQL adds for typed literals
// PostgreSQL adds ::type casts when storing typed literals like DATE '2024-01-01'
// We strip these redundant casts to generate cleaner DDL
// However, we preserve necessary casts like ::interval, ::bpchar, ::json, ::jsonb, and numeric casts

// Only strip casts on simple string literals (not in expressions)

// Handle string literals

// PostgreSQL stores negative numbers as string literals with casts like '-20'::integer
// We convert these back to plain numeric literals

// Convert numeric string to actual numeric literal
// This unwraps '-20'::integer -> -20

// Strip redundant text casts on string literals

// Strip date/time casts on literals (PostgreSQL adds these for typed literals)

// Handle numeric literals that already have explicit types
// PostgreSQL adds redundant casts like 100::numeric or 3.14::double precision
// Strip these redundant casts when casting to numeric types

// The value is already a numeric literal, no cast needed

// Strip redundant text/character varying casts on column names
// PostgreSQL adds these implicitly when comparing varchar columns
// e.g., (status)::text = 'active'::text → status = 'active'

// Strip redundant casts on NULL values and normalize to lowercase
// PostgreSQL stores DEFAULT NULL as NULL::type, but we normalize to just null
// (matching the lexer's keyword normalization to lowercase)

// Strip all type casts on NULL and return lowercase null (matching lexer)

// Preserve all other casts (interval, bpchar, json, jsonb, etc.)

// Remove redundant implicit casts that PostgreSQL adds for function arguments
// e.g., expr::bigint::double precision -> expr::bigint
// PostgreSQL adds these when a function expects double precision but gets bigint

// If the inner cast is to an integer type, remove the outer double precision cast

// Remove redundant array typecasts on ARRAY constructors
// PostgreSQL normalizes ARRAY[expr::type]::type[] to ARRAY[(expr)::type]
// The array typecast is redundant since the ARRAY constructor already produces the right type
// e.g., ARRAY[current_date::text]::text[] -> ARRAY[(CURRENT_DATE)::text]
// HOWEVER: Empty arrays (ARRAY[]) NEED the typecast or PostgreSQL can't determine the type

// Check if this is an array type cast (type string ends with [])

// This is an array type (e.g., text[], int[])
// Only strip the typecast if the array is NOT empty

// Non-empty array: strip the redundant typecast

// Empty array: preserve the typecast (ARRAY[]::int[] is required)

// Normalize the type name in the cast expression to handle type aliases
// e.g., int[]::int[] should become int[]::integer[]

// For PostgreSQL and MySQL, unwrap parentheses around most expressions to normalize
// Both databases add parentheses around many expressions, but we want a canonical form
// We always unwrap ParenExpr during normalization to get a canonical form
// The only exception is when parentheses are around complex nested expressions
// where they're needed for precedence (like CASE inside a larger expression)

// Preserve parentheses around COLLATE expressions, as they're semantically significant

// Always unwrap single-layer parentheses for normalization
// This handles cases like (NOT deleted), (a = 1), ((col)::type), etc.

// The generic parser may parse "= ANY(ARRAY[...])" as a FuncExpr on the right side.
// Normalize this to set the Any/All flags properly.

// Unwrap ParenExpr from right side for ANY/ALL to ensure consistent formatting.

// Handle IN clauses based on mode.

// PostgreSQL normalizes IN (values) to = ANY (ARRAY[values]) and NOT IN to <> ALL (ARRAY[values]).

// PostgreSQL normalizes NOT IN (values) to <> ALL (ARRAY[values]).

// For other databases, keep IN but sort the tuple for consistent comparison.

// For existing ANY/ALL expressions, normalize and sort the array elements.

// Collapse UnaryExpr with minus/plus on numeric literals to SQLVal
// This ensures "-20" and "- 20" (unary minus on 20) are treated the same

// Create negative integer: -N

// Double negative: --N → N

// Create negative float: -N.M

// Double negative: --N.M → N.M

// Unary plus has no effect on numeric values

// Recurse into subquery SELECT so that nested FROM clauses receive the
// same normalization as the outer SELECT (notably database-prefix
// stripping on MySQL — MariaDB's SHOW CREATE VIEW emits `db.table`
// everywhere including subquery FROM clauses, but user-written DDL
// rarely does, causing spurious view re-creations on every diff).

// Normalize CAST(expr AS type) to expr::type (CastExpr) for consistency
// PostgreSQL represents both forms identically in its internal representation

// PostgreSQL adds ELSE NULL to CASE expressions, normalize it away

// Also handle ELSE NULL::type (cast to null)

// PostgreSQL normalizes typed literals by removing the type prefix
// e.g., DATE '2024-01-01' -> '2024-01-01'
// e.g., TIME '12:00:00' -> '12:00:00'
// e.g., TIMESTAMP '2024-01-01 12:00:00' -> '2024-01-01 12:00:00'
// Only normalize for PostgreSQL mode

// For literals and other types, return as-is

// normalizeSelectExprs normalizes SELECT expressions for comparison
func normalizeSelectExprs(exprs parser.SelectExprs, mode GeneratorMode) parser.SelectExprs {
	_ = "STUB: not implemented"
	return *new(parser.SelectExprs)
}

// normalizeSelectExpr normalizes a single SELECT expression for views
func normalizeSelectExpr(expr parser.SelectExpr, mode GeneratorMode) parser.SelectExpr {
	_ = "STUB: not implemented"
	return *new(parser.SelectExpr)
}

// For PostgreSQL, strip automatic aliases like ?column?

// For MySQL, strip redundant aliases where the alias matches the column name
// MySQL adds "column_name as column_name" which is redundant

// The alias is the same as the column name, strip it

// normalizeTableExprs normalizes FROM clause table expressions
func normalizeTableExprs(exprs parser.TableExprs, mode GeneratorMode) parser.TableExprs {
	_ = "STUB: not implemented"
	return *new(parser.TableExprs)
}

// normalizeTableExpr normalizes a single TableExpr
func normalizeTableExpr(expr parser.TableExpr, mode GeneratorMode) parser.TableExpr {
	_ = "STUB: not implemented"
	return *new(parser.TableExpr)
}

// For MySQL, normalize table names to remove database prefix
// MySQL stores views with database.table references, but we want just table names

// Remove the database/schema part, keep only the table name

// Remove schema/database

// PostgreSQL and MariaDB add parentheses around JOINs when storing views.
// Unwrap these to get a canonical form.

// Single expression in parentheses - unwrap it

// Multiple expressions - normalize but keep parens

// normalizeJoinCondition normalizes the JOIN ON/USING condition
func normalizeJoinCondition(cond parser.JoinCondition, mode GeneratorMode) parser.JoinCondition {
	_ = "STUB: not implemented"
	return *

	// For PostgreSQL and MySQL, preserve table qualifiers in JOIN ON clauses
	// They're needed for disambiguation (e.g., "u.id = o.user_id")
	// We only normalize the expression structure (parentheses, etc.), not column qualifiers
	new(parser.JoinCondition)
}

// normalizeExprPreservingQualifiers is like normalizeExpr but preserves table qualifiers in column references
// This is used for JOIN ON clauses where qualifiers are semantically important
func normalizeExprPreservingQualifiers(expr parser.Expr, mode GeneratorMode) parser.Expr {
	_ = "STUB: not implemented"
	return *new(parser.Expr)
}

// Keep the qualifier but normalize the names to lowercase

// For PostgreSQL and MySQL, unwrap unnecessary parentheses

// For other expressions, use the regular normalizeExpr

// normalizeWhere normalizes WHERE clause
func normalizeWhere(where *parser.Where, mode GeneratorMode) *parser.Where {
	_ = "STUB: not implemented"
	return nil
}

// normalizeGroupBy normalizes GROUP BY clause
func normalizeGroupBy(groupBy parser.GroupBy, mode GeneratorMode) parser.GroupBy {
	_ = "STUB: not implemented"
	return *new(parser.GroupBy)
}

// normalizeOrderBy normalizes ORDER BY clause
func normalizeOrderBy(orderBy parser.OrderBy, mode GeneratorMode) parser.OrderBy {
	_ = "STUB: not implemented"
	return *new(parser.OrderBy)
}

// normalizeWith normalizes a WITH clause (Common Table Expressions) for comparison.
func normalizeWith(with *parser.With, mode GeneratorMode) *parser.With {
	_ = "STUB: not implemented"
	return nil
}

// TableLookupFunc is a function type for looking up a table by name.
// Used by normalizeViewDefinition to expand SELECT * expressions.
type TableLookupFunc func(name QualifiedName) *Table

// normalizeViewDefinition normalizes a view definition AST for comparison.
// This function removes database-specific formatting differences that don't affect the logical meaning.
// If tableLookup is provided, SELECT * expressions are expanded to explicit column names
// (PostgreSQL expands * when storing view definitions).
func normalizeViewDefinition(stmt parser.SelectStatement, mode GeneratorMode, tableLookup TableLookupFunc) parser.SelectStatement {
	_ = "STUB: not implemented"
	return *new(parser.SelectStatement)
}

// Expand SELECT * if we have table lookup capability

// Remove comments for view comparison - they don't affect semantic meaning

// hasStarExpr checks if there's a StarExpr in the SELECT list.
func hasStarExpr(exprs parser.SelectExprs) bool { _ = "STUB: not implemented"; return false }

// extractTableNameFromFrom extracts the table name from a simple FROM clause.
// Returns empty QualifiedName if the FROM clause is complex (joins, subqueries, etc.)
func extractTableNameFromFrom(from parser.TableExprs) QualifiedName {
	_ = "STUB: not implemented"
	return *new(QualifiedName)
}

// getSortedColumns converts a column map to a slice sorted by position.
// This is necessary because Go maps have non-deterministic iteration order,
// but column order matters for:
// - DDL generation (columns should appear in their original declaration order)
// - SELECT * expansion (PostgreSQL expands * to columns in ordinal position order)
func getSortedColumns(columns map[string]*Column) []*Column { _ = "STUB: not implemented"; return nil }

// expandStarExpr replaces StarExpr with explicit column references.
func expandStarExpr(exprs parser.SelectExprs, table *Table) parser.SelectExprs {
	_ = "STUB: not implemented"
	return *new(parser.SelectExprs)
}

// Get columns sorted by position to match PostgreSQL's expansion order

// normalizeViewColumnsFromDefinition extracts and normalizes column names from a view definition.
// This handles differences in how PostgreSQL versions format column names:
// - PostgreSQL 13-15: includes table qualifiers (e.g., "users.id")
// - PostgreSQL 16+: omits unnecessary qualifiers (e.g., "id")
func normalizeViewColumnsFromDefinition(def parser.SelectStatement, mode GeneratorMode) []string {
	_ = "STUB: not implemented"
	return nil
}

// For other statement types (e.g., UNION), we can't easily extract columns

// normalizeOperator converts operator to lowercase and applies PostgreSQL-specific mappings.
// PostgreSQL stores certain operators in a canonical form:
// - LIKE is stored as ~~
// - NOT LIKE is stored as !~~
// - != is stored as <>
func normalizeOperator(op string, mode GeneratorMode) string { _ = "STUB: not implemented"; return "" }

// normalizeName lowercases them for consistent comparison.
// TODO: Identifier case-sensitivity varies by RDBMS and settings:
//   - PostgreSQL: case-insensitive by default, case-sensitive when quoted
//   - MySQL: depends on settings, such as lower_case_table_names
//   - MSSQL: depends on collation settings
//     For now, we lowercase everything for normalization.
func normalizeName(name string) string { _ = "STUB: not implemented"; return "" }

var postgresTablePrivilegeList = []string{
	"DELETE",
	"INSERT",
	"REFERENCES",
	"SELECT",
	"TRIGGER",
	"TRUNCATE",
	"UPDATE",
}

func normalizePrivilegesForComparison(privileges []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Sort privileges in PostgreSQL canonical order
func sortPrivilegesByCanonicalOrder(privileges []string) { _ = "STUB: not implemented"; return }

// sortAndDeduplicateValues sorts and deduplicates a slice of expressions based on their string representation.
// This ensures that semantically equivalent lists are treated as identical regardless of order or duplicates.
// For example: [b, a, b] becomes [a, b]
func sortAndDeduplicateValues[T parser.Expr](values []T) []T { _ = "STUB: not implemented"; return nil }

// reuse underlying array

// tryConvertOrChainToIn attempts to convert an OR chain of equality comparisons
// (e.g., col=a OR col=b OR col=c) into an IN expression (e.g., col IN (a, b, c))
// Returns nil if the conversion is not applicable.
func tryConvertOrChainToIn(orExpr *parser.OrExpr) parser.Expr {
	_ = "STUB: not implemented"
	return *new(parser.Expr)
}

// Walk the OR chain and collect comparisons
// Also handle already-normalized IN expressions from nested ORs

// Handle IN expressions that were already normalized

// Extract values from IN clause

// normalizeCommentObject returns a normalized object path for a COMMENT statement.
// For PostgreSQL, this prepends the default schema if missing:
//   - OBJECT_TABLE: [table] -> [schema, table]
//   - OBJECT_COLUMN: [table, column] -> [schema, table, column]
//   - INDEX/VIEW/TYPE/DOMAIN/FUNCTION: [name] -> [schema, name]
//   - CONSTRAINT/TRIGGER: [name, table] -> [name, schema, table]
//
// For other databases or when schema is already present, returns the original object.
func normalizeCommentObject(comment *parser.Comment, mode GeneratorMode, defaultSchema string) []Ident {
	_ = "STUB: not implemented"
	return nil
}

// These types need [schema, name]

// COLUMN comments need [schema, table, column]

// These types need [name, schema, table] - schema is inserted at position 1

// Prepend default schema (unquoted)

func isPostgresSerialType(typeName string) bool { _ = "STUB: not implemented"; return false }

func getSerialUnderlyingType(typeName string) string { _ = "STUB: not implemented"; return "" }
