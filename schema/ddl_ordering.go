package schema

import (
	"github.com/sqldef/sqldef/v3/parser"
)

// normalizeIdentKey returns a normalized string representation of an Ident
// for use as a map key in dependency graphs. This ensures that identifiers
// which refer to the same database object produce the same key.
//
// When legacyIgnoreQuotes is true, all identifiers are normalized to lowercase
// for case-insensitive matching (backward compatible behavior).
//
// When legacyIgnoreQuotes is false, behavior per database:
//   - PostgreSQL: Unquoted fold to lowercase, quoted preserve case
//   - MySQL: Respects mysqlLowerCaseTableNames (0=case-sensitive, 1/2=case-insensitive)
//   - MSSQL/SQLite3: Always case-insensitive (fold all to lowercase)
func normalizeIdentKey(ident Ident, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) string {
	_ = "STUB: not implemented"
	// Legacy mode: case-insensitive matching for all databases
	return ""
}

// Case-sensitive: preserve case

// Case-insensitive (1 or 2): fold to lowercase

// MSSQL/SQLite3: always case-insensitive

// normalizeNameKey returns a normalized string representation of a
// QualifiedName for use as a map key in dependency graphs.
func normalizeNameKey(name QualifiedName, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) string {
	_ = "STUB: not implemented"
	return ""
}

// topologicalSort performs a stable topological sort on items based on their dependencies
// using Kahn's algorithm (BFS-based). It returns the sorted items in dependency order,
// or an empty slice if a circular dependency is detected.
//
// The algorithm is stable: when multiple items have no dependencies between them
// (independent items), they are output in their original input order. This ensures
// deterministic and predictable output.
//
// Time complexity: O(V + E) where V is the number of items and E is the number of dependencies.
func topologicalSort[T any](items []T, dependencies map[string][]string, getID func(T) string) []T {
	_ = "STUB: not implemented"
	// Build item map and track original indices for stable sorting
	return nil
}

// Calculate in-degrees (number of dependencies each item has)
// and build reverse dependency map (dependents) for efficiency
// Use items order (not map iteration) for deterministic behavior

// Priority queue: nodes with zero in-degree, maintained in sorted order by original index
// Using a simple slice here; items are kept sorted by original index for stable output

// Initialize queue with all nodes that have no dependencies

// Sort initial queue by original index to ensure stable output

// Process node with smallest original index (maintains input order for independent items)

// Reduce in-degree for all items that depend on curr

// This item is now ready (all its dependencies have been processed)
// Insert into queue maintaining sorted order by original index

// Binary search to find insertion position

// Insert at position while maintaining order

// Check if all nodes were processed (if not, there's a circular dependency)

// Circular dependency detected

// SortTablesByDependencies sorts CREATE TYPE/DOMAIN/FUNCTION/TABLE/VIEW DDLs
// by a unified dependency graph, ensuring objects are created in the correct
// order (dependencies before dependents). CREATE EXTENSION/SCHEMA stay at the
// front (no inbound edges from sorted kinds) and other DDLs stay at the tail.
//
// Edges harvested:
//   - Table -> Table (foreign keys)
//   - Table -> Function (column DEFAULT / CHECK / generated expressions)
//   - Table -> Type/Domain (column type references)
//   - Domain -> Function (Domain CHECK expressions)
//   - Domain -> Type/Domain (Domain underlying data type)
//   - Function -> Type/Domain (argument and return types)
//   - View -> Table/View (SELECT body)
func SortTablesByDependencies(ddls []DDL, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) []DDL {
	_ = "STUB: not implemented"
	return nil
}

// resolvableObjects gates the expensive column/expression walks: when no
// types/domains/functions exist (typical for MySQL/SQLite and many plain
// PostgreSQL schemas), every edge harvested below would resolve to nothing.

// sortItems is arranged in fixed-bucket order (types -> domains -> functions
// -> tables -> views) so the topological sort's stable tie-breaking falls
// back to that order for independent items, and so does the cycle fallback.

// resolveTypeReference resolves a type-name string (possibly schema-qualified)
// to a normalized dependency-graph key, returning (key, true) only if the
// reference matches an entry in knownTypes or knownDomains. PostgreSQL only:
// the other engines do not surface user-defined TYPE/DOMAIN.
func resolveTypeReference(typeName string, knownTypes, knownDomains map[string]bool, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// extractFunctionCallNames walks an expression and returns the normalized
// names of every function call that matches an entry in knownFunctions.
func extractFunctionCallNames(expr parser.Expr, knownFunctions map[string]bool, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) []string {
	_ = "STUB: not implemented"
	return nil
}

func walkExprForFunctionCalls(expr parser.Expr, knownFunctions map[string]bool, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int, deps map[string]bool) {
	_ = "STUB: not implemented"
	return
}

// extractViewDependencies extracts all table/view names that a view depends on
// by walking the SelectStatement AST and collecting TableName references.
// Returns normalized names suitable for use as dependency graph keys.
func extractViewDependencies(stmt parser.SelectStatement, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) []string {
	_ = "STUB: not implemented"
	return nil
}

// Convert map to slice in deterministic order

// extractDependenciesFromSelectStatement recursively extracts table/view dependencies from a SelectStatement
func extractDependenciesFromSelectStatement(stmt parser.SelectStatement, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int, deps map[string]bool) {
	_ = "STUB: not implemented"
	return
}

// Extract from the FROM clause

// Extract from WITH clause (CTE references)

// Recursively extract from both sides of the UNION

// Unwrap parenthesized SELECT

// extractDependenciesFromWith extracts dependencies from WITH clause (CTEs)
func extractDependenciesFromWith(with *parser.With, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int, deps map[string]bool) {
	_ = "STUB: not implemented"
	return
}

// extractDependenciesFromTableExprs extracts table/view names from TableExprs
func extractDependenciesFromTableExprs(exprs parser.TableExprs, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int, deps map[string]bool) {
	_ = "STUB: not implemented"
	return
}

// extractDependenciesFromTableExpr extracts table/view names from a single TableExpr
func extractDependenciesFromTableExpr(expr parser.TableExpr, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int, deps map[string]bool) {
	_ = "STUB: not implemented"
	return
}

// Extract from the actual table expression

// Recursively extract from both sides of the JOIN

// Recursively extract from parenthesized table expressions

// extractDependenciesFromSimpleTableExpr extracts table/view names from SimpleTableExpr
func extractDependenciesFromSimpleTableExpr(expr parser.SimpleTableExpr, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int, deps map[string]bool) {
	_ = "STUB: not implemented"
	return
}

// This is an actual table/view reference

// Always use schema.tableName format for consistency

// Recursively extract from subquery
