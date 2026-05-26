package mssql

import (
	"github.com/sqldef/sqldef/v3/database"
)

type MssqlParser struct {
	parser database.GenericParser
}

var _ database.Parser = (*MssqlParser)(nil)

func NewParser() MssqlParser { _ = "STUB: not implemented"; return *new(MssqlParser) }

func (p MssqlParser) Parse(sql string) ([]database.DDLStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// normalizeMssqlSyntax rewrites a few SQL Server export patterns into forms the
// generic parser already understands. Each pass must preserve comments, quoted
// strings, and bracketed identifiers so only real SQL tokens are rewritten.
func normalizeMssqlSyntax(sql string) string { _ = "STUB: not implemented"; return "" }

// SQL Server exports DEFAULT constraints as nested parenthesized expressions
// like DEFAULT ((0)) or DEFAULT (newid()). The generic parser only needs the
// underlying expression, so collapse the redundant outer parentheses here.
func normalizeDefaultExpressions(sql string) string { _ = "STUB: not implemented"; return "" }

func normalizeQualifiedReservedIdentifiers(sql string) string { _ = "STUB: not implemented"; return "" }

func normalizeBareReservedIdentifiers(sql string) string { _ = "STUB: not implemented"; return "" }

// SQL Server can export bracketed built-in cast targets and reserved identifiers
// in positions where the generic grammar only accepts plain type names or normal
// identifiers. Normalize those edge cases here instead of teaching the grammar
// every SQL Server-only spelling variant.
func normalizeBracketedCastTypes(sql string) string { _ = "STUB: not implemented"; return "" }

func scanBracketedCastType(sql string, start int) (string, int, bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func scanReservedQualifiedIdentifier(sql string, start int) (string, int, bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func stripOuterSQLParens(expr string) string { _ = "STUB: not implemented"; return "" }

func scanBalancedParenthesizedSQL(sql string, start int) (string, int, bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func hasKeywordAt(sql string, pos int, keyword string) bool {
	_ = "STUB: not implemented"
	return false
}

func isIdentChar(ch byte) bool { _ = "STUB: not implemented"; return false }

func isSpace(ch byte) bool { _ = "STUB: not implemented"; return false }

func hasLineCommentPrefix(sql string, pos int) bool { _ = "STUB: not implemented"; return false }

func hasBlockCommentPrefix(sql string, pos int) bool { _ = "STUB: not implemented"; return false }

func scanLineComment(sql string, start int) int { _ = "STUB: not implemented"; return 0 }

func scanBlockComment(sql string, start int) int { _ = "STUB: not implemented"; return 0 }

func scanSingleQuotedString(sql string, start int) int { _ = "STUB: not implemented"; return 0 }

func scanDoubleQuotedString(sql string, start int) int { _ = "STUB: not implemented"; return 0 }

func scanBracketIdentifier(sql string, start int) int { _ = "STUB: not implemented"; return 0 }
