package schema

import (
	"github.com/sqldef/sqldef/v3/database"
	"github.com/sqldef/sqldef/v3/parser"
)

type GeneratorMode int

const (
	GeneratorModeMysql = GeneratorMode(iota)
	GeneratorModePostgres
	GeneratorModeSQLite3
	GeneratorModeMssql
)

// tidbTableOption defines a TiDB-specific table option that the generator
// should compare between current and desired schemas.
type tidbTableOption struct {
	key          string // option key as it appears in the options map (e.g. "SHARD_ROW_ID_BITS")
	defaultValue string // value to use when the option is removed from desired
}

// tidbTableOptions is the single source of truth for TiDB table options
// that the generator manages. When adding a new TiDB table option:
//  1. Add an entry here
//  2. If the option uses /*T![feature] ... */ comment syntax,
//     add the feature name to extractTiDBComment in parser/token.go
var tidbTableOptions = []tidbTableOption{
	{key: "SHARD_ROW_ID_BITS", defaultValue: "0"},
	{key: "PRE_SPLIT_REGIONS", defaultValue: "0"},
	{key: "AUTO_ID_CACHE", defaultValue: "0"},
}

// This struct holds simulated schema states during GenerateIdempotentDDLs().
type Generator struct {
	mode          GeneratorMode
	desiredTables []*Table
	currentTables []*Table

	desiredViews []*View
	currentViews []*View

	desiredTriggers []*Trigger
	currentTriggers []*Trigger

	desiredFunctions []*Function
	currentFunctions []*Function

	desiredTypes []*Type
	currentTypes []*Type

	desiredDomains []*Domain
	currentDomains []*Domain

	desiredPartitionOfs []*CreatePartitionOf
	currentPartitionOfs []*CreatePartitionOf

	// Track FKs that have been handled during primary key changes
	handledForeignKeys map[string]bool

	// Track tables that have been dropped to skip COMMENT cleanup for them
	droppedTables map[string]bool

	// Track columns that have been dropped to skip COMMENT cleanup for them
	// Key is "schema.table.column"
	droppedColumns map[string]bool

	// Track indexes that have been dropped to skip COMMENT cleanup for them
	// Key is "schema.index_name"
	droppedIndexes map[string]bool

	// Map index names to their owning tables (for comment cleanup after table drops)
	// Key is "schema.index_name", value is the table's QualifiedName
	indexToTable map[string]QualifiedName

	desiredComments []*Comment
	currentComments []*Comment

	desiredExtensions []*Extension
	currentExtensions []*Extension

	desiredSchemas []*Schema
	currentSchemas []*Schema

	desiredPrivileges []*GrantPrivilege
	currentPrivileges []*GrantPrivilege

	defaultSchema string

	algorithm string
	lock      string

	config database.GeneratorConfig
}

// Parse argument DDLs and call `generateDDLs()`
func GenerateIdempotentDDLs(mode GeneratorMode, sqlParser database.Parser, desiredSQL string, currentSQL string, config database.GeneratorConfig, defaultSchema string) ([]string, error) {
	_ = "STUB: not implemented"
	// TODO: invalidate duplicated tables, columns
	return nil, nil
}

// Build index-to-table mapping before any tables are dropped

// sortIndexesByName returns indexes sorted by name.
// This ensures deterministic DDL ordering when processing indexes from the database,
// which may return indexes in different orders (e.g., MySQL vs TiDB).
func sortIndexesByName(indexes []Index) []Index { _ = "STUB: not implemented"; return nil }

// Main part of DDL generation
func (g *Generator) generateDDLs(desiredDDLs []DDL) ([]string, error) {
	_ = "STUB: not implemented"
	// These variables are used to control the output order of the DDL.
	// `CREATE SCHEMA` should execute first, and DDLs that add indexes and foreign keys should execute last.
	// Other DDLs are stored in interDDLs.
	return nil, nil
}

// Comments on indexes must come after CREATE INDEX

// Incrementally examine desiredDDLs

// Table already exists, guess required DDLs.

// Table not found. Check if it's a rename from another table.

// Found the old table, generate rename DDL

// PostgreSQL automatically transfers comments when renaming tables

// Update the old table's name to the new name

// Now generate DDLs for any column/index changes

// Old table not found, create as new table

// copy table

// Table not found and no rename, create table.

// copy table

// Only add to desiredTables if it doesn't already exist (it may have been pre-populated from aggregation)

// copy table

// Check if this partition child table already exists

// Partition child table doesn't exist, create it

// copy

// For now, we don't support modifying partition bounds - only create or keep as-is

// Index comments must come after CREATE INDEX statements

// Index comments must come after CREATE INDEX

// Clean up obsoleted triggers BEFORE dropping tables
// Triggers must be dropped before their associated tables to avoid "relation does not exist" errors

// Clean up obsoleted views BEFORE dropping columns
// Views must be dropped before columns they depend on to avoid "other objects depend on it" errors

// Sort tables to be dropped by dependencies (dependent tables first)

// Remove dropped tables from currentTables and track them for later
// (to skip generating COMMENT cleanup DDLs for dropped tables)

// Drop partition child tables that no longer exist in desired schema

// Clean up obsoleted indexes, columns in remaining tables

// Already handled in drop tables above

// Table is expected to exist. Drop foreign keys prior to index deletion

// Skip foreign keys without constraint names - they're likely from column-level REFERENCES
// that haven't been fully processed yet

// Check if FK exists in desired state by name

// Also check if there's a matching FK by columns (for unnamed FKs in desired schema)
// This applies to all databases that support named FK constraints (not SQLite3)

// Foreign key is expected to exist.

// Skip if referenced table is being dropped - already handled in generateDropTableDDLsWithDependencies

// The foreign key seems obsoleted. Check and drop it as needed.

// TODO: simulate to remove foreign key from `currentTable.foreignKeys`?

// Table is expected to exist. Drop exclusion constraints.

// Exclusion constraint is expected to exist.

// Check indexes
// Sort current indexes by name for deterministic DDL ordering (DB may return indexes in different order)

// Alter statement for primary key index should be generated above.

// Also check foreign key index names (these are plain strings)

// For MySQL, also check if this index supports an unnamed FK (the FK index name might be the column name)

// Index is expected to exist.

// Check if this index was renamed (don't drop if it was renamed)

// Index was renamed, don't drop it

// The index seems obsoleted. Check and drop it as needed.

// TODO: simulate to remove index from `currentTable.indexes`?

// Check CHECK constraints BEFORE dropping columns.
// This is important because CHECK constraints may reference columns that are about to be dropped.
// Databases require CHECK constraints to be dropped before the columns they reference.

// First try to find by name

// Also check if this constraint matches any CHECK by definition
// This handles auto-generated constraint names for column-level CHECKs (MySQL/MSSQL)
// and unnamed CHECK constraints in the desired schema (PostgreSQL)

// Check columns.
// Use sorted columns to ensure deterministic DDL ordering
// Drop columns in reverse order (last column first) to be more intuitive

// Column is expected to exist.

// Check if this column is being renamed (not dropped)

// Column is obsoleted. Drop column.

// Track dropped column for later (to skip generating COMMENT cleanup DDLs)

// Check policies.

// Clean up obsoleted domains

// Clean up obsoleted extensions

// Clean up obsoleted functions

// Clean up obsoleted types

// Clean up obsoleted comments

// Skip comments for tables that have been dropped
// PostgreSQL automatically removes comments when the table is dropped

// Skip comments for columns that have been dropped
// PostgreSQL automatically removes column comments when the column is dropped

// Skip comments for indexes that have been dropped
// PostgreSQL automatically removes index comments when the index is dropped

// Check if this comment still exists in desired comments

// Only generate NULL statement if the comment is completely absent from desired schema
// If desiredComment exists but is empty, it means the desired schema has "COMMENT ... IS NULL",
// which will be handled by generateDDLsForComment

// Comment was completely removed, generate COMMENT ... IS NULL

// Check each grantee individually for orphaned privileges

// Skip if managed roles is empty (means "manage no roles") or grantee is not in managed roles

// Check if this grantee exists in any desired privilege for the same table

// Comment out DROP/REVOKE statements when enable_drop is false

// commentOutDropStatements converts DROP/REVOKE statements to SQL comments.
// This makes the output testable and visible in --dry-run output.
func commentOutDropStatements(ddls []string) []string { _ = "STUB: not implemented"; return nil }

// isDropStatement checks if a DDL statement is a DROP or REVOKE statement.
// Note: DROP CONSTRAINT and DROP CHECK are NOT included because they are
// required for non-destructive schema changes (e.g., changing defaults).
func isDropStatement(ddl string) bool { _ = "STUB: not implemented"; return false }

func (g *Generator) generateDDLsForAbsentColumn(currentTable *Table, desiredTable *Table, column *Column) []string {
	_ = "STUB: not implemented"

	// Only MSSQL has column default constraints. They need to be deleted before dropping the column.
	return nil
}

// In the caller, `mergeTable` manages `g.currentTables`.
func (g *Generator) generateDDLsForCreateTable(currentTable Table, desired CreateTable) ([]string, error) {
	_ = "STUB: not implemented"

	// Track foreign keys that need to be recreated after primary key changes
	return nil, nil
}

// Examine each column

// deep copy to avoid modifying the original

// Check for conflict: can't rename a column if the old name still exists

// We may not be able to add AUTO_INCREMENT yet. It will be added after adding keys (primary or not) at the "Add new AUTO_INCREMENT" place.
// prevent to

// Check if this is a renamed column

// Generate RENAME COLUMN DDL
// Use quote info from renamedFrom annotation, not from the found column
// (database exports always quote identifiers)

// PostgreSQL automatically transfers comments when renaming columns

// After renaming, check if type/constraints need to be changed

// MySQL uses CHANGE COLUMN for rename

// SQL Server uses sp_rename
// For sp_rename, we need to handle schema prefixes properly

// Only include schema if it's not the default

// After renaming, check if type/constraints need to be changed
// Skip if the column is part of the current primary key - the primary key handling logic
// will properly handle foreign key dependencies

// Use consistent table name format (without default schema prefix)

// For SQLite, when type needs to change:
// 1. Add new column with new name and type
// 2. Copy data from old column to new column
// 3. Drop old column

// 1. Add new column with desired name and definition

// 2. Copy data from old column to new column

// 3. Drop the old column

// Simple rename without type change

// Fallback to regular ADD for unsupported databases

// Regular column addition (not a rename)

// Column not found, add column.

// Change column data type or order as needed.

// Change column type and orders, *except* AUTO_INCREMENT and UNIQUE KEY.

// MySQL has limitations (Error 3106) with generated columns that require using
// DROP COLUMN + ADD COLUMN instead of CHANGE COLUMN in these cases:
// 1. Changing storage type (VIRTUAL <-> STORED)
// 2. Converting regular column to generated column
// 3. Converting generated column to regular column
// For all other generated column modifications (expression changes, attribute changes, etc.),
// CHANGE COLUMN works correctly and is more efficient than DROP+ADD.

// Both columns are generated - check if storage type is changing

// One column is generated and the other is not - MySQL requires DROP+ADD

// Add UNIQUE KEY. TODO: Probably it should be just normalized to an index after the parser phase.

// TODO: deal with a case that the index is not a UNIQUE KEY.

// enumTypeChange is true when the type change involves an enum on
// either side. In that case the column default must be re-applied
// around the ALTER (see classifyEnumTypeChange's docstring).

// Serial types require changing both column type and sequence type

// Change type - use desiredColumn for escaping to match user's quote style

// Handle IDENTITY and NOT NULL in the correct order for PostgreSQL:
// - When adding IDENTITY: must SET NOT NULL first (IDENTITY requires NOT NULL)
// - When removing IDENTITY: must DROP IDENTITY first (can't drop NOT NULL from IDENTITY column)

// Step 1: When removing IDENTITY, drop it first

// Step 2: When adding IDENTITY, set NOT NULL first if needed

// Step 3: Add or modify IDENTITY

// Modify existing IDENTITY (not adding or removing)

// Step 4: Handle NOT NULL changes unrelated to IDENTITY

// Step 5: After removing IDENTITY, drop NOT NULL if needed

// default

// The default was already dropped (if present) before the ALTER
// COLUMN TYPE above; re-apply the desired default in the new type.

// drop - use desiredColumn for escaping to match user's quote style

// set - use desiredColumn for escaping to match user's quote style

// First, check if the current column has a CHECK constraint
// If so, use it directly for comparison instead of searching by the desired name

// Current column has a CHECK - check if its name matches the desired name

// Names match (accounting for quoting), use the current column's check

// Names don't match - this is a rename scenario
// The current constraint should be dropped and the new one added

// Current column has no CHECK, search in table-level constraints

// Check if current column's CHECK matches a table-level CHECK in desired.
// PostgreSQL exports single-column CHECKs as column-level, but user may define them as table-level.

// Determine if we need to drop the current column's constraint
// This handles the case where names are different (quoted vs unquoted)
// We need to drop if: current has a check AND (definition differs OR constraint names differ)

// Drop the current constraint if it exists

// Current column has a CHECK with a different name that needs to be dropped

// TODO: support adding a column's `references`

// Skip if the column is part of the current primary key - the primary key handling logic
// will properly handle foreign key dependencies when the PK changes

// Change column definition

// For MSSQL, column-level CHECKs might actually be table-level CHECKs that MSSQL converted
// Check if the current column-level CHECK matches a table-level CHECK in desired

// Current has column-level CHECK, desired doesn't
// Check if it matches a table-level CHECK in desired

// This column-level CHECK is actually a table-level CHECK
// It will be handled in the table-level CHECK processing

// IDENTITY

// remove

// DEFAULT

// drop

// set

// Remove old AUTO_INCREMENT/AUTO_RANDOM from deleted column before deleting key (primary or not)
// and if primary key changed

// Examine primary key

// Check if there are foreign keys referencing this table's primary key

// If there are foreign keys referencing this table,
// we need to drop them first before modifying the primary key

// Track dropped FKs to avoid duplicates using normalized names for case-insensitive comparison

// Create a unique key for this FK using normalized names

// Already processed this FK

// Update the current state to reflect that we've dropped this FK
// This prevents duplicate FK creation when processing the referencing table

// Remove the FK from the current table's FK list

// Also drop the index if it exists (MySQL creates implicit indexes for FKs)
// PostgreSQL and SQL Server don't create implicit indexes for FKs

// Look for the corresponding desired foreign key to get updated columns
// We need to find the desired table that references our table

// Only recreate the foreign key if:
// 1. The referencing table exists in the desired schema, AND
// 2. The foreign key exists in the desired schema

// Mark this FK as globally handled so we don't add it again in normal FK processing
// Use normalized names for case-insensitive deduplication

// If the table doesn't exist in desired schema or the FK doesn't exist,
// we don't recreate it (it will be dropped with the table)

// Add the DROP FK statements before we modify the primary key

// When dropping PRIMARY KEY, also drop implicit NOT NULL if the column should be nullable
// PRIMARY KEY implicitly adds NOT NULL, so we need to remove it when reverting

// Extract column name from the index column expression

// Find the column in the desired table to check if it should be nullable

// Column should be nullable, remove the implicit NOT NULL

// MySQL doesn't support ALTER COLUMN ... DROP NOT NULL
// Instead, we use CHANGE COLUMN with the full column definition

// MSSQL doesn't support ALTER COLUMN ... DROP NOT NULL either
// Instead, we use ALTER COLUMN with the full column definition

// Store the FK recreation DDLs to be added at the end

// Examine each index

// Drop and add index as needed.

// Check if this is a renamed index

// Generate RENAME INDEX DDL

// Index not found and not a rename, add index.

// Add new AUTO_INCREMENT/AUTO_RANDOM after adding index and primary key

// Examine each foreign key

// Auto-generate constraint name if not specified

// When desired FK has no explicit constraint name, first try to find
// a matching FK in current state by columns. This handles the case where
// databases auto-generate names (e.g., MySQL's "books_ibfk_1", PostgreSQL's
// "table_column_fkey", or MSSQL's "FK__posts__user_id__...") but the user's
// schema doesn't specify a name.

// Use the existing constraint name from the current state

// If no matching FK found in current state, generate a deterministic name

// Use the first column name for the constraint name

// Create a modified ForeignKey with the generated constraint name if needed

// Use the matched FK from earlier if we found one, otherwise look up by name

// Drop and add foreign key as needed.

// Foreign key not found, add foreign key.
// But first check if we've already handled this FK during primary key changes

// Examine each exclusion

// Drop and add exclusion as needed.

// Exclusion not found, add exclusion.

// Examine each check

// First try to find by name

// Also try to find by definition if not found by name
// This handles auto-generated constraint names for MySQL/MSSQL, and
// for PostgreSQL when the desired constraint has no explicit name

// Constraint exists but has different definition, need to replace it

// SQLite does not support ALTER TABLE for CHECK constraints
// Modifying CHECK constraints requires recreating the table, which is not supported

// Constraint exists with same definition but different name
// Don't generate DDL for renaming - constraint names don't matter if the definition is the same
// This handles cases where MSSQL auto-generates names like CK__table__column__hash
// and sqldef auto-generates names like table_column_check

// Constraint doesn't exist, add it

// Examine table comment

// Examine TiDB table options

// Compare partitions (MySQL/MariaDB only)

// Add FK recreation DDLs at the end (they will be executed after all table modifications)

// generatePartitionDDLs compares partitions between current and desired tables
// and generates ADD PARTITION / DROP PARTITION statements
func (g *Generator) generatePartitionDDLs(currentTable Table, desiredTable Table) []string {
	_ = "STUB: not implemented"

	// If neither has partitions, nothing to do
	return nil
}

// If only current has partitions (desired removes all partitioning), we don't handle
// removing partitioning entirely - that would require REMOVE PARTITIONING

// If only desired has partitions, the table creation should handle it

// Both have partitions - compare partition definitions
// Find partitions to add (in desired but not in current)

// Find partitions to drop (in current but not in desired)

// generateAddPartitionDDL generates ALTER TABLE ADD PARTITION statement
func (g *Generator) generateAddPartitionDDL(table Table, part PartitionDefinition) string {
	_ = "STUB: not implemented"
	return ""
}

// Quote partition name only if it needs quoting (contains special chars/spaces)
// Don't preserve quotes from source since MariaDB quotes differently than MySQL

// LIST partition: VALUES IN (...)

// RANGE partition: VALUES LESS THAN MAXVALUE

// RANGE partition: VALUES LESS THAN (...)

// Fallback (shouldn't happen with valid partition definitions)

// generateDropPartitionDDL generates ALTER TABLE DROP PARTITION statement
func (g *Generator) generateDropPartitionDDL(table Table, part PartitionDefinition) string {
	_ = "STUB: not implemented"
	return ""
}

// Quote partition name only if it needs quoting (contains special chars/spaces)
// Don't preserve quotes from source since MariaDB quotes differently than MySQL

// escapePartitionName quotes a partition name only if it needs quoting.
// MySQL/MariaDB partition names need quoting if they contain spaces or special characters.
func (g *Generator) escapePartitionName(name string) string { _ = "STUB: not implemented"; return "" }

// formatExprs formats parser.Exprs for use in DDL statements
func (g *Generator) formatExprs(exprs parser.Exprs) string { _ = "STUB: not implemented"; return "" }

// Shared by `CREATE INDEX` and `ALTER TABLE ADD INDEX`.
// This manages `g.currentTables` unlike `generateDDLsForCreateTable`...
func (g *Generator) generateDDLsForCreateIndex(tableName QualifiedName, desiredIndex Index, action string, statement string) ([]string, error) {
	_ = "STUB: not implemented"
	// For CREATE INDEX, handle statement regeneration based on mode
	return nil, nil
}

// Quote-aware mode: always regenerate to ensure proper schema qualification and quoting

// Legacy mode with CONCURRENTLY config: insert CONCURRENTLY into original statement
// This preserves the original formatting while adding the keyword

// Otherwise: use the original statement as-is

// Views or non-existent tables

// Index not found, add index.

// Check if the view exists in desired views (might be created in the same migration)

// View will be created, add the index

// Check if it's a desired table that hasn't been created yet

// Table will be created, add the index

// Creating index on non-existent table/view, just add the statement

// Check if this is a renamed index

// Generate RENAME INDEX DDL

// PostgreSQL automatically transfers comments when renaming indexes

// Update the current table's indexes to reflect the rename

// Replace with the renamed index

// Index not found and not a rename, add index.

// Index found. If it's different, drop and add index.

// simulate index change. TODO: use []*Index in table and destructively modify it

func (g *Generator) generateDDLsForAddForeignKey(tableName QualifiedName, desiredForeignKey ForeignKey, action string, statement string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Foreign Key not found, add foreign key

// Foreign key found, If it's different, drop and add or alter foreign key.

// Examine indexes in desiredTable to delete obsoleted indexes later

// Only add to desiredTable.foreignKeys if it doesn't already exist (it may have been pre-populated from aggregation)

func (g *Generator) generateDDLsForAddExclusion(tableName QualifiedName, desiredExclusion Exclusion, action string, statement string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exclusion not found, add exclusion

// Exclusion key found, If it's different, drop and add or alter exclusion.

// Examine indexes in desiredTable to delete obsoleted indexes later

// Only add to desiredTable.exclusions if it doesn't already exist (it may have been pre-populated from aggregation)

func (g *Generator) generateDDLsForCreatePolicy(tableName QualifiedName, desiredPolicy Policy, action string, statement string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Policy not found, add policy.

// policy found. If it's different, drop and add or alter policy.

// Examine policies in desiredTable to delete obsoleted policies later

// Only add to desiredTable.policies if it doesn't already exist (it may have been pre-populated from aggregation)

func (g *Generator) shouldDropAndCreateView(currentView *View, desiredView *View) bool {
	_ = "STUB: not implemented"
	return false
}

// In the case of PostgreSQL, if there are any deletions or changes to columns,
// you cannot use REPLACE VIEW, so you need to DROP and CREATE VIEW.
//
// ref: https://www.postgresql.org/docs/current/sql-createview.html
//
// > CREATE OR REPLACE VIEW is similar, but if a view of the same name already exists, it is replaced.
// > The new query must generate the same columns that were generated by the existing view query
// > (that is, the same column names in the same order and with the same data types), but it may add additional
// > columns to the end of the list. The calculations giving rise to the output columns may be completely different.

// If we couldn't extract columns from the definitions, fall back to DROP and CREATE

// If columns are removed, we need to DROP and CREATE.

// If all existing columns are identical and only a new column is added, use REPLACE; otherwise, execute DROP and CREATE.

func (g *Generator) generateDDLsForCreateView(desiredView *View) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// View not found, add view.

// copy view
// Don't copy indexes from desired to current - they'll be added when the CREATE INDEX is processed

// TODO: Fix the definition comparison for materialized views and enable this
// View found. If it's different, create or replace view.
// Use AST-based comparison with table lookup for SELECT * expansion

// Post-normalization fix for generic parser: strip remaining table qualifiers
// The generic parser's ColName.Name may not respect the empty qualifier we set

// Build the WITH [NO] DATA clause for materialized views

// When dropping views that may have dependents, we need to handle them
// For PostgreSQL, find dependent views and recreate them after

// Find all views that depend on this view

// Drop them first (in reverse dependency order)

// Store DDLs to recreate dependent views after the base view

// Recreate dependent views

// VIEW with the specified security type found. If it's different, create or replace view.

// Examine policies in desiredTable to delete obsoleted policies later
// Only add to desiredViews if it doesn't already exist (it may have been pre-populated from aggregation)

// findDependentViews finds all views that reference the given view name in their definitions.
// Returns views in topological order (views with no dependents first, most dependent last).
// Uses the proper dependency extraction from ddl_ordering.go.
func (g *Generator) findDependentViews(viewName QualifiedName) []*View {
	_ = "STUB: not implemented"
	// Normalize the target view name for comparison
	return nil
}

// Build dependency graph for all views

// Extract dependencies using the proper AST-based extraction

// Find views that directly or indirectly depend on the target view

// Skip the target view itself

// Check if this view depends on the target view

// Sort dependents in topological order using the dependency graph

// stripTableQualifiers removes table qualifiers from column references in SQL
// E.g., "users.name" -> "name", "t.id" -> "id"
// This is needed for PostgreSQL 13-15 where table qualifiers are included in column references.
func stripTableQualifiers(sql string) string {
	_ = "STUB: not implemented"
	// Match table.column patterns where:
	// - table name is [a-z_][a-z0-9_]* (identifier)
	// - followed by a dot
	// - followed by column name [a-z_][a-z0-9_]* (identifier)
	// We use word boundaries to avoid matching within quoted strings
	return ""
}

// Replace "table.column" with just "column" (keeping capture group 1)

// createTableLookup returns a TableLookupFunc that looks up tables from both
// desired and current table lists.
func (g *Generator) createTableLookup() TableLookupFunc {
	_ = "STUB: not implemented"
	return *new(TableLookupFunc)
}

func (g *Generator) formatTriggerEvent(event TriggerEvent) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) formatTriggerEvents(events []TriggerEvent, sep string) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) generateDDLsForCreateTrigger(triggerName QualifiedName, desiredTrigger *Trigger) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Trigger not found, add trigger.

// Trigger found. If it's different, drop and recreate (or alter for MSSQL).

// Only add to desiredTriggers if it doesn't already exist (it may have been pre-populated from aggregation)

// generateDDLsForCreateFunction generates DDLs for CREATE FUNCTION statements
func (g *Generator) generateDDLsForCreateFunction(desired *Function) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Function does not exist, create it

// Function exists but is different, use CREATE OR REPLACE
// Modify the statement to include OR REPLACE if not already present

// Track the function as processed

func (g *Generator) findFunctionByName(functions []*Function, name QualifiedName) *Function {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) areSameFunctionDefinition(a, b *Function) bool {
	_ = "STUB: not implemented"
	// Compare function properties
	return false
}

// Compare options (order-independent, case-insensitive for option names)

// areSameFunctionOptions compares two sets of function options.
// Options are compared in a normalized way to handle differences in formatting.
func (g *Generator) areSameFunctionOptions(a, b []string) bool {
	_ = "STUB: not implemented"
	// Normalize and filter out default options
	return false
}

// normalizeFunctionOptions normalizes a list of function options for comparison.
// It removes default options that PostgreSQL doesn't export.
func normalizeFunctionOptions(options []string) []string {
	_ = "STUB: not implemented"
	// PostgreSQL default options (not exported by pg_get_functiondef)
	return nil
}

// default volatility
// default null behavior
// default security
// default parallel safety

// Skip default options

// normalizeFunctionOption normalizes a function option for comparison.
// Handles differences like:
// - "TimeZone" vs timezone (quoted vs unquoted identifiers)
// - SET x TO y vs SET x = y
// - 'value' vs value (quoted vs unquoted values)
// - RETURNS NULL ON NULL INPUT vs STRICT
// - Case differences
func normalizeFunctionOption(opt string) string { _ = "STUB: not implemented"; return "" }

// Remove double quotes from identifiers (e.g., "timezone" -> timezone)

// Remove single quotes from values (e.g., 'public' -> public)
// This normalizes SET search_path = 'public', 'pg_temp' to SET search_path = public, pg_temp

// Normalize SET clause: replace " to " with " = " for consistency
// Match patterns like "set x to y" and convert to "set x = y"

// Replace " to " with " = " but be careful not to replace within values
// Simple approach: replace the first occurrence of " to " after "set "

// Normalize whitespace

// Normalize equivalent options
// RETURNS NULL ON NULL INPUT is equivalent to STRICT

func (g *Generator) generateDDLsForCreateType(desired *Type) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle RENAME VALUE for values with @renamed annotation in desired

// Handle ADD VALUE for new values (not renamed)

// Type not found, add type.

// Only add to desiredTypes if it doesn't already exist (it may have been pre-populated from aggregation)

// containsEnumValue checks if the given value exists in the enum values.
func containsEnumValue(enumValues []EnumValue, value string) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Generator) generateDDLsForCreateDomain(desired *Domain) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only add to desiredDomains if it doesn't already exist (it may have been pre-populated from aggregation)

func (g *Generator) generateAlterDomainDDLs(current, desired *Domain) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: We don't handle collation changes as PostgreSQL doesn't support changing collation via ALTER DOMAIN
// The user would need to drop and recreate the domain to change collation

func (g *Generator) findDomainConstraintByExpression(constraints []DomainConstraint, expression parser.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Generator) generateDDLsForComment(desired *Comment) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If both current and desired comments are NULL/empty, no change is needed.

// Comment not found or different. Generate statement from normalized AST.

// escapeCommentObjectParts returns the escaped object parts for a COMMENT statement.
// Object is []Ident representing: [schema, table] for TABLE, [schema, table, column] for COLUMN.
// The object is first normalized (schema prepended if missing), then the schema part
// is normalized if it matches the default schema.
func (g *Generator) escapeCommentObjectParts(comment *Comment) []string {
	_ = "STUB: not implemented"
	return nil
}

// Determine schema position based on object type
// CONSTRAINT and TRIGGER have schema at position 1: [name, schema, table]
// All others have schema at position 0: [schema, name, ...]

// This element is the schema - normalize if it's the default schema

// generateNormalizedCommentStatement creates a COMMENT statement with normalized identifiers.
func (g *Generator) generateNormalizedCommentStatement(comment *Comment) string {
	_ = "STUB: not implemented"
	return ""
}

// Build the SQL statement based on object type

// These have syntax: COMMENT ON CONSTRAINT/TRIGGER name ON [schema.]table IS ...
// parts is [name, schema, table] or [name, table]

// FUNCTION has syntax: COMMENT ON FUNCTION [schema.]name(args) IS ...

// Standard syntax: COMMENT ON TYPE [schema.]name IS ...

// Escape the comment value (single quotes need to be doubled)

// buildFunctionSignature builds a function signature string for COMMENT ON FUNCTION.
func (g *Generator) buildFunctionSignature(comment *Comment) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) generateDDLsForExtension(desired *Extension) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extension not found, add extension.

// copy extension

// Only add to desiredExtensions if it doesn't already exist (it may have been pre-populated from aggregation)

func (g *Generator) generateDDLsForSchema(desired *Schema) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Schema not found, add schema.

// copy schema

// Only add to desiredSchemas if it doesn't already exist (it may have been pre-populated from aggregation)

// Even though simulated table doesn't have a foreign key, references could exist in column definitions.
// This carefully generates DROP CONSTRAINT for such situations.
func (g *Generator) generateDDLsForAbsentForeignKey(currentForeignKey ForeignKey, currentTable Table, desiredTable Table) []string {
	_ = "STUB: not implemented"
	return nil
}

// Even though simulated table doesn't have an index, primary or unique could exist in column definitions.
// This carefully generates DROP INDEX for such situations.
func (g *Generator) generateDDLsForAbsentIndex(currentIndex Index, currentTable Table, desiredTable Table) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If nil, it will be `DROP COLUMN`-ed and we can usually ignore it.
// However, it seems like you need to explicitly drop it first for MSSQL.

// TODO: check length of currentIndex.columns
// TODO: handle this. Rename primary key column...?

// Columns become empty if the index is a PostgreSQL's expression index.

// Match by current name or by renamed-from name
// (PostgreSQL automatically updates constraint column references on RENAME COLUMN)

// No unique column. Drop unique key index.

func (g *Generator) generateDataType(column Column) string { _ = "STUB: not implemented"; return "" }

// Determine the full type name including schema qualification

// Normalize PostgreSQL shortcuts to their canonical forms for output
// Note: We DON'T normalize general aliases like varchar->character varying or numeric->decimal
// Those are preserved as-is in the output. We only normalize PostgreSQL-specific shortcuts.

// Note: timestamptz and timetz are already normalized by the parsers
// (they set the timezone flag and convert the type name to timestamp/time)

// Only qualify type names with schema for PostgreSQL when:
// 1. references is not empty (including "public." for enum types)
// 2. the type name doesn't already contain a dot
// 3. it's not a built-in type (built-in types shouldn't have references set to non-empty schema)

// Preserve quoting for case-sensitive types like domains.

func (g *Generator) generateColumnDefinition(column Column, enableUnique bool) (string, error) {
	_ = "STUB: not implemented"
	// TODO: make string concatenation faster?
	return "", nil
}

// [CHARACTER SET] and [COLLATE] should be placed before [NOT NULL | NULL] on MySQL

// Generated column definitions have this syntax on MySQL
// col_name data_type [GENERATED ALWAYS] AS (expr)
//  [VIRTUAL | STORED] [NOT NULL | NULL]
//  [UNIQUE [KEY]] [[PRIMARY] KEY]
//  [COMMENT 'string']

// TODO: Should this use StringConstant?

// Normalize CHECK expression to match PostgreSQL's output
// This ensures typed literals are properly converted (e.g., time '...' -> '...'::time)

// noop

// noop

type AggregatedSchema struct {
	Tables       []*Table
	PartitionOfs []*CreatePartitionOf // PostgreSQL partition child tables
	Views        []*View
	Triggers     []*Trigger
	Functions    []*Function
	Types        []*Type
	Domains      []*Domain
	Comments     []*Comment
	Extensions   []*Extension
	Schemas      []*Schema
	Privileges   []*GrantPrivilege
}

// generateCreateIndexStatement generates a CREATE INDEX statement from an Index struct.
// This is used to regenerate CREATE INDEX statements with proper schema-qualified table names.
func (g *Generator) generateCreateIndexStatement(table QualifiedName, index Index) string {
	_ = "STUB: not implemented"
	// Build column list with proper quoting
	return ""
}

// For simple column references (ColName), use escapeSQLIdent to preserve quoting

// For expressions (functional indexes), format with quote awareness

// Legacy mode: use parser.String for backward compatibility

// Start building the statement
// PostgreSQL syntax: CREATE [UNIQUE] INDEX [CONCURRENTLY] name ON table

// Add index method if specified (e.g., USING btree)

// Add index options (WITH clause must come before WHERE in PostgreSQL)

// Add WHERE clause for partial indexes

// insertConcurrentlyIntoCreateIndex inserts the CONCURRENTLY keyword into a CREATE INDEX statement.
// It handles both "CREATE INDEX" and "CREATE UNIQUE INDEX" statements.
func insertConcurrentlyIntoCreateIndex(statement string) string {
	_ = "STUB: not implemented"
	return ""
}

// Fallback: return original statement

// generateAddIndex generates DDL to add an index.
func (g *Generator) generateAddIndex(table QualifiedName, index Index) string {
	_ = "STUB: not implemented"
	return ""
}

// For simple column references (ColName), use escapeSQLIdent to preserve quoting

// For expressions (functional indexes), format with quote awareness

// Legacy mode: use parser.String for backward compatibility

// definition of partition is valid only in the syntax `CREATE INDEX ...`

// If the current name is just the column name (common with generic parser),
// replace it with the PostgreSQL convention

// Construct index type with optional VECTOR keyword for MariaDB vector indexes

func (g *Generator) generateIndexOptionDefinition(indexOptions []IndexOption) string {
	_ = "STUB: not implemented"
	return ""
}

// Handle multiple vector index options (M and DISTANCE)

func (g *Generator) generateConstraintOptions(ConstraintOptions *ConstraintOptions) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) generateForeignKeyDefinition(foreignKey ForeignKey) string {
	_ = "STUB: not implemented"
	// TODO: make string concatenation faster?
	return ""
}

// Empty constraint name is already invalidated in generateDDLsForCreateIndex

func (g *Generator) generateExclusionDefinition(exclusion Exclusion) string {
	_ = "STUB: not implemented"
	return ""
}

// generateRenameIndex generates DDL statements to rename an index.
func (g *Generator) generateRenameIndex(tableName QualifiedName, oldIndexName Ident, newIndexName Ident, desiredIndex *Index) []string {
	_ = "STUB: not implemented"
	return nil
}

// MySQL uses ALTER TABLE ... RENAME INDEX

// PostgreSQL uses ALTER INDEX ... RENAME TO
// Qualify the old index name with schema

// SQL Server uses sp_rename - use raw names without escaping

// SQLite doesn't support renaming indexes directly - drop and recreate

// generateDropIndex generates a DDL statement to drop an index.
func (g *Generator) generateDropIndex(tableName QualifiedName, indexName Ident, constraint bool) string {
	_ = "STUB: not implemented"
	return ""
}

// For DROP INDEX, we need schema.indexname

// escapeQualifiedName escapes a QualifiedName using quote-aware logic.
// Both schema and table names use quote-aware logic when legacy_ignore_quotes is false.
func (g *Generator) escapeQualifiedName(name QualifiedName) string {
	_ = "STUB: not implemented"
	return ""
}

// If schema is empty, don't add schema prefix

// escapeTableName escapes a table name using quote-aware logic.
// Both schema and table names use quote-aware logic when legacy_ignore_quotes is false.
func (g *Generator) escapeTableName(table *Table) string { _ = "STUB: not implemented"; return "" }

// escapeColumnName escapes a column name using quote-aware logic.
func (g *Generator) escapeColumnName(column *Column) string { _ = "STUB: not implemented"; return "" }

// escapeViewName escapes a view name using quote-aware logic.
func (g *Generator) escapeViewName(view *View) string { _ = "STUB: not implemented"; return "" }

// escapeTypeName escapes a type name using quote-aware logic.
func (g *Generator) escapeTypeName(t *Type) string { _ = "STUB: not implemented"; return "" }

// escapeDomainName escapes a domain name using quote-aware logic.
func (g *Generator) escapeDomainName(d *Domain) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) forceEscapeSQLName(name string) string { _ = "STUB: not implemented"; return "" }

// standard SQL (PostgreSQL, SQLite)

// escapeSQLIdent escapes an Ident for SQL output, respecting the quote-aware normalization setting.
// When legacy_ignore_quotes is false:
//   - Quoted identifiers preserve their case and are always quoted in output
//   - Unquoted identifiers are normalized to lowercase and are NOT quoted in output
func (g *Generator) escapeSQLIdent(ident Ident) string { _ = "STUB: not implemented"; return "" }

// escapeSQLNameQuoteAware escapes an identifier name for SQL output,
// taking into account whether it was originally quoted and the legacy_ignore_quotes setting.
func (g *Generator) escapeSQLNameQuoteAware(name string, wasQuoted bool) string {
	_ = "STUB: not implemented"
	// Legacy mode: always quote everything (backward compatible behavior)
	return ""
}

// Quote-aware mode

// Originally quoted: preserve case and quote in output

// Originally unquoted: normalize to lowercase and don't quote

// Originally quoted: preserve case and quote in output

// Originally unquoted: do nothing since the RDBMS here is case-insensitive

// identsEqual compares two Ident values (columns, indexes, constraints) for equality.
// This does NOT use MysqlLowerCaseTableNames because that only affects table names.
//
// When legacy_ignore_quotes is false (quote-aware mode for PostgreSQL):
//   - Unquoted identifiers are normalized to lowercase before comparison
//   - Quoted identifiers preserve their case
//   - `users` (unquoted) == `"users"` (quoted lowercase) because unquoted normalizes to lowercase
//   - `"Users"` (quoted) != `"users"` (quoted) because case differs
//
// When legacy_ignore_quotes is true or nil (legacy mode):
//   - Compare case-insensitively (backward compatible behavior)
func (g *Generator) identsEqual(a, b Ident) bool { _ = "STUB: not implemented"; return false }

// qualifiedNamesEqual compares two QualifiedName values for equality.
// An empty schema is treated as equivalent to the default schema.
// When legacy_ignore_quotes is true (or nil), schema names are compared case-insensitively.
// When legacy_ignore_quotes is false, schema names use quote-aware comparison
// (quoted "MySchema" is different from unquoted myschema).
func (g *Generator) qualifiedNamesEqual(a, b QualifiedName) bool {
	_ = "STUB: not implemented"
	return false
}

// normalizeDefaultSchema returns an Ident for a schema, treating the default schema
// (e.g., "public") as unquoted when it's lowercase. This ensures consistent output where
// the default schema appears without quotes. For non-default schemas, the original
// quote status is preserved.
func (g *Generator) normalizeDefaultSchema(schema Ident) Ident {
	_ = "STUB: not implemented"
	return *new(Ident)
}

// escapeAndJoinNames escapes a list of names with comma separation
func (g *Generator) escapeAndJoinNames(names []Ident) string { _ = "STUB: not implemented"; return "" }

// validateAndEscapeGrantee validates and escapes a grantee name to prevent SQL injection
func (g *Generator) validateAndEscapeGrantee(grantee string) (string, error) {
	_ = "STUB: not implemented"
	// PUBLIC is a special keyword and should not be quoted
	return "", nil
}

// Check for potentially dangerous characters that shouldn't be in role names
// PostgreSQL role names can contain letters, digits, underscores, and some special chars
// but we'll be conservative to prevent injection
// Note: quotes, backticks, and brackets are allowed as escapeSQLName handles them

// Use escapeSQLName which handles proper escaping including quotes/brackets/backticks

// normalizeOldTableName creates a QualifiedName from a renamedFrom Ident,
// using the schema from the new table name if not specified.
func (g *Generator) normalizeOldTableName(oldName Ident, newTable QualifiedName) QualifiedName {
	_ = "STUB: not implemented"
	// Use the schema from the new table name
	return *new(QualifiedName)
}

// generateRenameTableDDL generates a DDL statement to rename a table.
// Uses quote-aware escaping for both old and new table names.
func (g *Generator) generateRenameTableDDL(oldTable QualifiedName, newTable QualifiedName) string {
	_ = "STUB: not implemented"
	return ""
}

// For PostgreSQL, RENAME TO should only include the table name without schema

// must be qualified
// must not be qualified

// MSSQL uses sp_rename for renaming tables

// must be qualified
// must not be qualified

func (g *Generator) notNull(column Column) bool { _ = "STUB: not implemented"; return false }

func isAddConstraintForeignKey(ddl string) bool { _ = "STUB: not implemented"; return false }

// isPrimaryKey checks if a column is part of the table's primary key.
//
// TODO: This function uses direct string comparison (indexColumn.ColumnName() == column.name.Name)
// instead of quote-aware comparison via identsEqual(). This could cause incorrect results
// when the column name and index column name have different case representations in the parsed SQL
// (e.g., column "UserId" vs PRIMARY KEY (userid)). In practice, this doesn't manifest for
// PostgreSQL because database exports normalize identifiers consistently. However, it could
// cause issues in offline mode where SQL is parsed directly without database normalization.
// Consider using g.identsEqual() or normalizeIdentKey() for correctness.
func (g *Generator) isPrimaryKey(column Column, table Table) bool {
	_ = "STUB: not implemented"
	return false
}

// Destructively modify table1 to have table2 columns/indexes
func mergeTable(table1 *Table, table2 Table) {
	_ = "STUB: not implemented"
	// Update/add all columns from table2
	return
}

// Add indexes from table2 that don't exist in table1

func aggregateDDLsToSchema(ddls []DDL, mode GeneratorMode, defaultSchema string, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) (*AggregatedSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// copy table

// Copy the partition of statement

// TODO: check duplicated creation

// TODO: check duplicated creation

// TODO: check duplicated creation

// TODO: multi-column primary key?

// Note: REVOKE statements in desired schemas are not recommended
// The desired schema should describe the target state with GRANTs only
// This case is kept for backwards compatibility but may be removed

func formatPrivilegesForGrant(privileges []string) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) generateDDLsForGrantPrivilege(desired *GrantPrivilege) ([]string, error) {
	_ = "STUB: not implemented"
	// Grantees should already be filtered by FilterPrivileges
	// If multiple grantees made it here, they all have the same privileges to grant
	return nil, nil
}

// Track REVOKE operations per grantee

// Track GRANT operations grouped by privileges to grant

// privileges key -> grant group

// Before revoking, check if this privilege is granted by any other
// desired GRANT statement for the same grantee and table

// DO NOT update current privileges here - this breaks idempotency
// The state should only be updated after DDLs are successfully applied

func (g *Generator) generateDDLsForRevokePrivilege(desired *RevokePrivilege) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DO NOT update current privileges here - this breaks idempotency
// The state should only be updated after DDLs are successfully applied

func equalPrivileges(a, b []string) bool { _ = "STUB: not implemented"; return false }

func (g *Generator) findTableByName(tables []*Table, name QualifiedName) *Table {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) findPartitionOfByName(partitionOfs []*CreatePartitionOf, name QualifiedName) *CreatePartitionOf {
	_ = "STUB: not implemented"
	return nil
}

// findTableQuoteAware finds a table using quote-aware comparison without requiring a Generator.
func findTableQuoteAware(tables []*Table, name QualifiedName, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) *Table {
	_ = "STUB: not implemented"
	return nil
}

// findViewQuoteAware finds a view using quote-aware comparison without requiring a Generator.
func findViewQuoteAware(views []*View, name QualifiedName, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) *View {
	_ = "STUB: not implemented"
	return nil
}

// findColumnByName finds a column using quote-aware comparison
func (g *Generator) findColumnByName(columns map[string]*Column, name Ident) *Column {
	_ = "STUB: not implemented"
	return nil
}

// findIndexByName finds an index by its identifier using quote-aware comparison.
func (g *Generator) findIndexByName(indexes []Index, name Ident) *Index {
	_ = "STUB: not implemented"
	return nil
}

func findIndexOptionByName(options []IndexOption, name string) *IndexOption {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) findCheckConstraintInTable(table *Table, constraintName Ident) *CheckDefinition {
	_ = "STUB: not implemented"
	// First, look for table-level check constraints
	return nil
}

// Then, look for column-level check constraints

// findCheckConstraintByName finds a CHECK constraint in a list by name
func (g *Generator) findCheckConstraintByName(checks []CheckDefinition, constraintName Ident) *CheckDefinition {
	_ = "STUB: not implemented"
	return nil
}

// findCheckConstraintByDefinitionInList finds a CHECK constraint in a list by comparing definitions
func (g *Generator) findCheckConstraintByDefinitionInList(checks []CheckDefinition, check *CheckDefinition) *CheckDefinition {
	_ = "STUB: not implemented"
	return nil
}

// findCheckConstraintByDefinition finds a CHECK constraint in a table by comparing definitions.
// This is used for MySQL when column-level CHECKs are converted to table-level CONSTRAINTs
// with auto-generated names.
func (g *Generator) findCheckConstraintByDefinition(table *Table, check *CheckDefinition) *CheckDefinition {
	_ = "STUB: not implemented"
	return nil
}

// Search table-level checks

// Search column-level checks

func (g *Generator) findForeignKeyByName(foreignKeys []ForeignKey, constraintName Ident) *ForeignKey {
	_ = "STUB: not implemented"
	return nil
}

// findForeignKeyByColumns finds a foreign key by matching its columns and reference table/columns.
// This is used when the desired FK has no explicit constraint name, to match against existing FKs
// that may have auto-generated names (e.g., MySQL's "table_ibfk_N" pattern).
func (g *Generator) findForeignKeyByColumns(foreignKeys []ForeignKey, desired ForeignKey) *ForeignKey {
	_ = "STUB: not implemented"
	return nil
}

// findMatchingDesiredForeignKey finds a desired foreign key that matches the current FK by source columns.
// This is used to check if a current FK (with an auto-generated name) should be kept or dropped.
// We only match by source columns (not reference table) because the reference table might change
// when modifying an FK - in that case, we want to keep the FK for modification rather than dropping it.
func (g *Generator) findMatchingDesiredForeignKey(desiredForeignKeys []ForeignKey, current ForeignKey) *ForeignKey {
	_ = "STUB: not implemented"
	return nil
}

// Only match against desired FKs that don't have explicit names

// foreignKeysMatchBySourceColumns checks if two foreign keys have the same source columns.
// This is used to determine if an FK should be kept for modification (even if reference table changes).
func (g *Generator) foreignKeysMatchBySourceColumns(fk1, fk2 ForeignKey) bool {
	_ = "STUB: not implemented"
	return false
}

// foreignKeysMatchByColumns checks if two foreign keys match by their columns and reference table/columns.
func (g *Generator) foreignKeysMatchByColumns(fk1, fk2 ForeignKey) bool {
	_ = "STUB: not implemented"
	// Match by index columns
	return false
}

// Match by reference table

// Match by reference columns

// indexSupportsUnnamedForeignKey checks if an index supports an unnamed foreign key.
// MySQL creates implicit indexes for foreign keys, named after the first column.
// TiDB creates implicit indexes with names like fk_1, fk_2, etc.
// This function checks if the index supports any unnamed FK by:
// 1. Matching index name to the FK's first column (MySQL behavior)
// 2. Matching index columns to the FK's columns (TiDB and general case)
func (g *Generator) indexSupportsUnnamedForeignKey(index Index, foreignKeys []ForeignKey) bool {
	_ = "STUB: not implemented"
	return false
}

// Only check unnamed FKs

// Check if index name matches the first column of the FK (MySQL behavior)

// Check if index columns match the FK columns (TiDB behavior: fk_N naming)

// indexColumnsMatchForeignKey checks if the index columns match the foreign key columns.
func (g *Generator) indexColumnsMatchForeignKey(index Index, fk ForeignKey) bool {
	_ = "STUB: not implemented"
	return false
}

// findForeignKeysReferencingTable finds all foreign keys from all tables that reference the given table
func (g *Generator) findForeignKeysReferencingTable(referencedTable QualifiedName) []struct {
	tableName  QualifiedName
	foreignKey ForeignKey
} {
	_ = "STUB: not implemented"
	return nil
}

// Check all current tables for foreign keys that reference this table

func (g *Generator) findExclusionByName(exclusions []Exclusion, constraintName Ident) *Exclusion {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) findPolicyByName(policies []Policy, name Ident) *Policy {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) findViewByName(views []*View, name QualifiedName) *View {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) findTriggerByName(triggers []*Trigger, name QualifiedName) *Trigger {
	_ = "STUB: not implemented"
	return nil
}

// findType finds a type by matching both schema and name with quote-aware comparison.
// This handles both exact matches and case-insensitive matches for unquoted names.
func (g *Generator) findType(types []*Type, desiredType *Type) *Type {
	_ = "STUB: not implemented"
	return nil
}

// isEnumCastableStringType reports whether typeName is a PostgreSQL string-like
// type that can be safely cast to/from an enum via "USING col::enum" for both
// empty and populated tables.
//
// Fixed-length character(N) / bpchar is intentionally excluded: its values are
// space-padded on the right, so "USING col::enum" would compare e.g.
// 'active  ' against enum labels and fail at runtime. Supporting it would
// require emitting rtrim() inside USING, which silently corrupts enum labels
// that legitimately end with whitespace.
func isEnumCastableStringType(typeName string) bool { _ = "STUB: not implemented"; return false }

// classifyEnumTypeChange inspects an ALTER COLUMN ... TYPE change that may
// involve a PostgreSQL enum.
//
// using is the body of a USING expression when the cast has no implicit form:
//   - string -> enum:                "col::enum_t"
//   - enum_a -> enum_b (different):  "col::text::enum_b" (no direct cast)
//
// PostgreSQL has an assignment cast from enum to string types (text, varchar,
// citext, name), so enum -> string returns no USING.
//
// enumInvolved is true whenever either side resolves to an enum (regardless of
// whether USING is needed). Callers use it to drop and re-apply the column
// default around the ALTER, since the existing default still references the
// old type after the cast and would otherwise block DROP TYPE.
func (g *Generator) classifyEnumTypeChange(currentColumn, desiredColumn Column) (using string, enumInvolved bool) {
	_ = "STUB: not implemented"
	return "", false
}

// findEnumTypeForColumn looks up the enum type that the column refers to in the given
// types list, or returns nil if the column does not refer to an enum type.
func (g *Generator) findEnumTypeForColumn(column Column, types []*Type) *Type {
	_ = "STUB: not implemented"
	return nil
}

// typeIdent carries the quote info and is set for user-defined types
// (domains, enums, ...). The generic parser leaves it empty when the
// source SQL writes a schema-qualified custom type like "myschema.item",
// because parseTable splits "myschema.item" into references="myschema."
// + typeName="item" before typeIdent gets populated. Fall back to
// typeName so quoted schemas still find their unquoted enum.

// findDomainByName finds a domain using quote-aware comparison including schema
func (g *Generator) findDomainByName(domains []*Domain, name QualifiedName) *Domain {
	_ = "STUB: not implemented"
	return nil
}

// findCommentByObject finds a comment by its object path using quote-aware comparison.
// Objects are normalized (schema prepended if missing) before comparison.
func (g *Generator) findCommentByObject(comments []*Comment, targetComment *parser.Comment) *Comment {
	_ = "STUB: not implemented"
	return nil
}

// isCommentOnDroppedTable checks if a comment belongs to a table that has been dropped.
// This is used to skip generating COMMENT ... IS NULL for dropped tables,
// since PostgreSQL automatically removes comments when a table is dropped.
func (g *Generator) isCommentOnDroppedTable(comment *Comment) bool {
	_ = "STUB: not implemented"
	return false
}

// Extract table qualified name based on object type
// Different object types have different structures:
// - OBJECT_TABLE: [schema, table]
// - OBJECT_COLUMN: [schema, table, column]
// - OBJECT_CONSTRAINT: [constraint, schema, table]
// - OBJECT_TRIGGER: [trigger, schema, table]
// - OBJECT_INDEX: [schema, index] - need to look up table from index

// Structure: [schema, table, ...] or [table, ...]

// Structure: [name, schema, table] or [name, table]

// Structure: [schema, index] or [index]
// Need to find which table this index belongs to

// trackDroppedColumn records a column as dropped for later use in comment cleanup.
func (g *Generator) trackDroppedColumn(table *Table, column *Column) {
	_ = "STUB: not implemented"
	return
}

// droppedColumnKey returns a normalized key for tracking dropped columns.
// Uses normalizeIdentKey to handle case-insensitive matching for unquoted identifiers.
func (g *Generator) droppedColumnKey(schema, table, column Ident) string {
	_ = "STUB: not implemented"
	return ""
}

// isCommentOnDroppedColumn checks if a comment belongs to a column that has been dropped.
// This is used to skip generating COMMENT ... IS NULL for dropped columns,
// since PostgreSQL automatically removes column comments when the column is dropped.
func (g *Generator) isCommentOnDroppedColumn(comment *Comment) bool {
	_ = "STUB: not implemented"
	return false
}

// Column comment structure: [schema, table, column] or [table, column]

// trackDroppedIndex records an index as dropped for later use in comment cleanup.
func (g *Generator) trackDroppedIndex(table *Table, index Index) { _ = "STUB: not implemented"; return }

// droppedIndexKey returns a normalized key for tracking dropped indexes.
func (g *Generator) droppedIndexKey(schema, indexName Ident) string {
	_ = "STUB: not implemented"
	return ""
}

// isCommentOnDroppedIndex checks if a comment belongs to an index that has been dropped.
// This is used to skip generating COMMENT ... IS NULL for dropped indexes,
// since PostgreSQL automatically removes index comments when the index is dropped.
func (g *Generator) isCommentOnDroppedIndex(comment *Comment) bool {
	_ = "STUB: not implemented"
	return false
}

// buildIndexToTableMap builds a mapping from index names to their owning tables.
// This must be called before any tables are dropped, as it uses currentTables.
func (g *Generator) buildIndexToTableMap() { _ = "STUB: not implemented"; return }

// findTableForIndex finds the table that owns the given index.
// object is [schema, index] or [index] from an OBJECT_INDEX comment.
// Returns empty QualifiedName if the index is not found.
func (g *Generator) findTableForIndex(object []Ident) QualifiedName {
	_ = "STUB: not implemented"
	return *new(QualifiedName)
}

// indexMapKey generates a map key for index lookup using quote-aware normalization.
func (g *Generator) indexMapKey(schema, name Ident) string { _ = "STUB: not implemented"; return "" }

// identsSliceEqual compares two []Ident slices for equality using quote-aware comparison.
// For PostgreSQL in quote-aware mode (legacy_ignore_quotes: false):
//   - Quoted identifiers preserve case and are compared case-sensitively
//   - Unquoted identifiers are normalized to lowercase before comparison
//
// For legacy mode or non-PostgreSQL databases, comparison is case-insensitive.
func (g *Generator) identsSliceEqual(a, b []Ident) bool { _ = "STUB: not implemented"; return false }

// generateCommentNullStatement creates a COMMENT ... IS NULL statement.
func (g *Generator) generateCommentNullStatement(comment *Comment) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) findExtensionByName(extensions []*Extension, name Ident) *Extension {
	_ = "STUB: not implemented"
	return nil
}

func findSchemaByName(schemas []*Schema, name string) *Schema {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) haveSameColumnDefinition(current Column, desired Column) bool {
	_ = "STUB: not implemented"
	// Not examining AUTO_INCREMENT, AUTO_RANDOM, and UNIQUE KEY because it'll be added in a later stage
	return false
}

// `PRIMARY KEY` implies `NOT NULL`

// (current.check == desired.check) && /* workaround. CHECK handling in general should be improved later */
// detect change column only when set explicitly. TODO: can we calculate implicit charset?
// detect change column only when set explicitly. TODO: can we calculate implicit collate?

func (g *Generator) areSameGenerated(generatedA, generatedB *Generated) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: Difference between bracketed and unbracketed, as Expr values are not fully comparable.

func (g *Generator) haveSameDataType(current Column, desired Column) bool {
	_ = "STUB: not implemented"
	return false
}

// Quote-aware comparison for custom types (domains, etc.)
//
// PostgreSQL's format_type returns the internal type name without quotes.
// For a domain created with CREATE DOMAIN "PositiveInt", it returns "PositiveInt".
// For a domain created with CREATE DOMAIN PositiveInt, it returns "positiveint".
//
// The canonical form for comparison is:
// - Quoted identifiers: use the exact name (case-sensitive)
// - Unquoted identifiers: lowercase the name

// Legacy behavior: case-insensitive normalized comparison

// Normalize length for MySQL YEAR type.
// MySQL's YEAR(4) is deprecated and equivalent to YEAR. MySQL's SHOW CREATE TABLE
// returns just "year" without a length, so we ignore the length for comparison.

// Normalize default precision/scale for numeric/decimal types.

// Length is typically an integer value, but MSSQL can use "max".
// For example, VARCHAR(MAX)

func (g *Generator) areSameCheckDefinition(checkA *CheckDefinition, checkB *CheckDefinition) bool {
	_ = "STUB: not implemented"
	return false
}

// Unwrap outermost parentheses if present (MySQL adds extra parens)

// unwrapOutermostParenExpr removes the outermost ParenExpr if the expression is wrapped in one.
// This is needed because some databases (like MySQL) add extra parentheses around CHECK expressions.
// It preserves parentheses around OR expressions to maintain correct operator precedence.
func unwrapOutermostParenExpr(expr parser.Expr) parser.Expr {
	_ = "STUB: not implemented"
	return *new(parser.Expr)
}

func (g *Generator) buildForeignKeyDDL(tableName QualifiedName, fk *ForeignKey) string {
	_ = "STUB: not implemented"
	return ""
}

// normalizeCheckExprString returns a normalized string representation of a CHECK constraint expression
// For PostgreSQL, this converts IN (a,b,c) to = ANY (ARRAY[a,b,c])
func (g *Generator) normalizeCheckExprString(expr parser.Expr) string {
	_ = "STUB: not implemented"
	return ""
}

// Unwrap outermost parentheses for consistent output (comparison does this too)

// In quote-aware mode, use formatExprQuoteAware to preserve quoting in column names
// In legacy mode, use parser.String for backward compatibility (no quoting in expressions)

// formatExprQuoteAware formats an expression with quote-aware column name handling.
// This walks the AST and uses escapeSQLIdent for column names to preserve quoting.
func (g *Generator) formatExprQuoteAware(expr parser.Expr) string {
	_ = "STUB: not implemented"
	return ""
}

// IsExpr has Operator (e.g., "is null", "is not null") and Expr

// For function expressions, format arguments with quote awareness
// Normalize function name to lowercase (PostgreSQL convention)

// For other expression types, fall back to parser.String

// formatSelectExprQuoteAware formats a SelectExpr (used in function arguments) with quote awareness.
func (g *Generator) formatSelectExprQuoteAware(expr parser.SelectExpr) string {
	_ = "STUB: not implemented"
	return ""
}

func areSameIdentityDefinition(identityA *Identity, identityB *Identity) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Generator) areSameDefaultValue(currentDefault *DefaultDefinition, desiredDefault *DefaultDefinition, columnType string) bool {
	_ = "STUB: not implemented"
	// Normalize: DEFAULT NULL is the same as no default
	return false
}

// Both null (or absent) - they're the same

// One is null, the other isn't - they're different

// Normalize expressions first to handle typed literals and other database-specific normalizations
// This ensures that TypedLiterals are converted to their underlying values before comparison

// Strip type casts remaining after normalizeExpr (e.g., custom types like ENUMs/domains).
// PostgreSQL stores defaults with explicit casts (e.g., 'pending'::order_status),
// but users write DEFAULT 'pending' without the cast. Both are semantically identical.

// Check if both are simple SQLVal (vs complex expressions) after normalization

// If both are simple values (SQLVal), use value comparison

// If one is SQLVal and the other is not, they're different

// Both are complex expressions - use string comparison on already-normalized expressions

// unwrapCast strips a type cast from an expression, returning the inner expression.
// This handles custom type casts (e.g., ENUM, domain) that normalizeExpr does not strip.
func unwrapCast(expr parser.Expr) parser.Expr { _ = "STUB: not implemented"; return *new(parser.Expr) }

// isNumericColumnType determines if a column type should be compared numerically.
// This is used to decide how to compare default values.
func (g *Generator) isNumericColumnType(typeName string) bool {
	_ = "STUB: not implemented"
	return false
}

// areSameValue compares two default values with knowledge of the column type.
func (g *Generator) areSameValue(current, desired *Value, columnType string) bool {
	_ = "STUB: not implemented"
	return false
}

// Special handling for MySQL boolean values (BOOLEAN is stored as TINYINT(1))
// MySQL converts: false → 0, true → 1

// For numeric types (DECIMAL, INT, FLOAT, etc.): use numeric comparison
// This handles DECIMAL precision normalization: '0.1' vs '0.10'
// Use 256 bits of precision to safely handle DECIMAL(65,30) (MySQL's max)
// 65 decimal digits * 3.32 bits/digit ≈ 216 bits, so 256 bits is safe

// Fallback to string comparison if parse fails

// For string types (VARCHAR, CHAR, TEXT, etc.): use exact string comparison
// This ensures '1.00' != '1.0000' for VARCHAR columns

// areSameIdentifiers compares two values for identifiers/keywords (case-insensitive).
func (g *Generator) areSameIdentifiers(current, desired *Value) bool {
	_ = "STUB: not implemented"
	return false
}

// areSameEvents compares two TriggerEvent slices for equality.
// Events are sorted before comparison because PostgreSQL reorders them alphabetically
// e.g., "INSERT OR UPDATE OR DELETE" becomes "INSERT OR DELETE OR UPDATE" in pg_get_triggerdef
func areSameEvents(eventsA, eventsB []TriggerEvent) bool { _ = "STUB: not implemented"; return false }

func (g *Generator) areSameTriggerDefinition(triggerA, triggerB *Trigger) bool {
	_ = "STUB: not implemented"
	return false
}

// Compare table names using quote-aware comparison

// Compare WHEN conditions
// Normalize: lowercase, remove spaces, strip matching outer parentheses
// pg_get_triggerdef() may wrap the condition in extra parentheses

// Normalize PROCEDURE to FUNCTION for PostgreSQL compatibility
// pg_get_triggerdef() always returns EXECUTE FUNCTION even if created with EXECUTE PROCEDURE

// stripMatchingOuterParens removes matching outer parentheses from a string.
// e.g., "((a>0))" → "a>0", but "(a)or(b)" is unchanged because the outer parens don't match.
func stripMatchingOuterParens(s string) string { _ = "STUB: not implemented"; return "" }

func isNullValue(value *Value) bool { _ = "STUB: not implemented"; return false }

func isNullDefault(def *DefaultDefinition) bool { _ = "STUB: not implemented"; return false }

// Check if it's a direct SQLVal

// Check if it's a CastExpr wrapping a NULL value (e.g., NULL::character varying)

func (g *Generator) areSamePrimaryKeys(primaryKeyA *Index, primaryKeyB *Index) bool {
	_ = "STUB: not implemented"
	return false
}

// For MSSQL, when comparing PRIMARY KEY constraints,
// ignore the name if one is auto-generated (PK__*) and the other is unnamed/synthetic ("PRIMARY")

// Check if one has an auto-generated name and the other is synthetic

// Compare everything except the name

// areSamePrimaryKeyColumns compares primary keys without checking the name
func (g *Generator) areSamePrimaryKeyColumns(indexA Index, indexB Index) bool {
	_ = "STUB: not implemented"
	return false
}

// For primary keys, we don't need to check other properties like where, included, options

func (g *Generator) areSameIndexes(indexA Index, indexB Index) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: check length?

// For MSSQL UNIQUE constraints (not regular indexes), don't compare options
// Inline constraints don't include WITH options in their DDL, but the database
// still has default options that show up in sys.indexes
// Only skip options comparison for constraints, not regular unique indexes

// Mysql: Default Index B-Tree (but not for vector indexes)

// For MSSQL and MySQL, constraint vs index distinction doesn't matter
// MSSQL: doesn't support PostgreSQL-style deferrable constraint options
// MySQL: CONSTRAINT name UNIQUE and UNIQUE KEY name are equivalent

func (g *Generator) formatIndexExprForComparison(expr parser.Expr) string {
	_ = "STUB: not implemented"
	return ""
}

// SQL Server metadata does not preserve whether a simple identifier was
// written as `name` or `[name]`, so compare index/PK column refs by their
// underlying identifier instead of their exported bracket style.

func (g *Generator) areSameWhereClause(whereA, whereB parser.Expr) bool {
	_ = "STUB: not implemented"
	// Both nil
	return false
}

// One is nil and the other is not

func (g *Generator) areSameForeignKeys(foreignKeyA ForeignKey, foreignKeyB ForeignKey) bool {
	_ = "STUB: not implemented"
	// Compare index columns (source columns of the FK)
	return false
}

// Compare ON UPDATE/DELETE actions

// Treat nil as equivalent to &ConstraintOptions{false, false} (the default).
// The pgquery parser always creates a non-nil ConstraintOptions even without
// a DEFERRABLE clause, while the generic parser creates nil.

func (g *Generator) areSameExclusions(exclusionA Exclusion, exclusionB Exclusion) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Generator) areSameExprs(exprA, exprB parser.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: case-insensitive comparison is not always correct

func (g *Generator) areSamePolicies(policyA, policyB Policy) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Generator) normalizeReferenceOption(action string) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO: Use interface to avoid defining following functions?

func convertIndexesToIndexNames(indexes []Index) []string { _ = "STUB: not implemented"; return nil }

func convertExclusionToConstraintNames(exclusions []Exclusion) []Ident {
	_ = "STUB: not implemented"
	return nil
}

func convertForeignKeysToIndexNames(foreignKeys []ForeignKey) []string {
	_ = "STUB: not implemented"
	return nil
}

// unexpected to reach else (really?)

func removeTableByName(tables []*Table, name string) []*Table {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) generateSerialSequenceAlterDDL(table *Table, column *Column, underlyingType string) string {
	_ = "STUB: not implemented"
	return ""
}

func generateSequenceClause(sequence *Sequence) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) generateDefaultDefinition(defaultDefinition DefaultDefinition) (string, error) {
	_ = "STUB: not implemented"
	// Type assertion: Check if it's a simple SQLVal
	return "", nil
}

// Simple value path - maintain existing formatting behavior

// NULL, CURRENT_TIMESTAMP, ...

// Complex expression path
// Normalize the expression to handle typed literals and other database-specific normalizations

// Enclose expression with parentheses to avoid syntax error
// https://dev.mysql.com/doc/refman/8.0/en/data-type-defaults.html#data-type-defaults-explicit
// https://www.sqlite.org/syntax/column-constraint.html

func generateSridDefinition(sridVal Value) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SRID option is only for MySQL 8.0.3 or later

func FilterTables(ddls []DDL, config database.GeneratorConfig) []DDL {
	_ = "STUB: not implemented"
	return nil
}

func skipTables(tables []string, config database.GeneratorConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func FilterViews(ddls []DDL, config database.GeneratorConfig) []DDL {
	_ = "STUB: not implemented"
	return nil
}

func FilterPrivileges(ddls []DDL, config database.GeneratorConfig) []DDL {
	_ = "STUB: not implemented"
	// If no roles specified, exclude all privileges
	return nil
}

// Skip all privilege-related DDLs

// Filter privileges to only include specified roles

// Map to consolidate grants by table and privileges

// Track order of insertion

// Track order of insertion

// Filter grantees to only include those in config

// Sort privileges for consistent key

// Add grantees to existing grant with same table and privileges

// Create new grant with filtered grantees

// Process each grantee separately and consolidate

// Merge privileges

// Create new revoke for this grantee

// Include all non-privilege DDLs

// Add all consolidated grants to the result in original order

func skipViews(views []string, config database.GeneratorConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func containsRegexpString(strs []string, str string) bool { _ = "STUB: not implemented"; return false }

// generateDropTableDDLsWithDependencies generates DROP TABLE statements in the correct order
// considering foreign key dependencies. Tables that reference other tables are dropped first.
// It also generates DROP CONSTRAINT statements for foreign keys from tables that will NOT be
// dropped but reference tables that WILL be dropped.
func (g *Generator) generateDropTableDDLsWithDependencies(tablesToDrop []*Table) []string {
	_ = "STUB: not implemented"
	return nil
}

// Build a set of tables to be dropped for quick lookup

// If there are no or only one table to drop, no sorting needed.

// Build reverse dependency graph for drops
// For drops: if table A references table B, then B depends on A (B can't be dropped until A is dropped)

// First, build a map of tables to be dropped for quick lookup

// Now build reverse dependencies
// For each table, find which other tables (in the drop list) reference it

// Skip self-referential FKs using quote-aware comparison

// If the referenced table is also being dropped

// The referenced table depends on this table being dropped first

// If circular dependency is detected, fall back to the original order.

// Drop foreign key constraints from tables that will NOT be dropped
// but reference tables that WILL be dropped

// Skip tables that will be dropped (their FKs will be dropped with the table)

// Skip foreign keys without constraint names

// Skip self-referential FKs

// If the referenced table is being dropped, we need to drop this FK first

func splitTableName(table string, defaultSchema string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func isValidAlgorithm(algorithm string) bool { _ = "STUB: not implemented"; return false }

func isValidLock(lock string) bool { _ = "STUB: not implemented"; return false }

// Escape a string and add quotes to form a legal SQL string constant.
func StringConstant(s string) string { _ = "STUB: not implemented"; return "" }
