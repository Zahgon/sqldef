package database

import (
	"github.com/sqldef/sqldef/v3/parser"
)

// A tuple of an original DDL and a Statement
type DDLStatement struct {
	DDL       string
	Statement parser.Statement
}

type Parser interface {
	Parse(sql string) ([]DDLStatement, error)
}

type GenericParser struct {
	mode parser.ParserMode
}

func NewParser(mode parser.ParserMode) GenericParser {
	_ = "STUB: not implemented"
	return *new(GenericParser)
}

func (p GenericParser) Parse(sql string) ([]DDLStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p GenericParser) splitDDLs(str string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Right now, the parser isn't capable of splitting statements by itself.
// So we just attempt parsing until it succeeds. I'll let the parser do it in the future.

// remove scanned tokens

// trimMarginComments pulls out any leading or trailing comments from a raw sql query.
// This function also trims leading (if there's a comment) and trailing whitespace.
func trimMarginComments(sql string) string { _ = "STUB: not implemented"; return "" }

// trailingCommentStart returns the first index of trailing comments.
// If there are no trailing comments, returns the length of the input string.
// NOTE: MySQL version comments (/*!NNNNN ... */) are NOT treated as comments
// because they contain SQL code that should be executed.
func trailingCommentStart(text string) (start int) { _ = "STUB: not implemented"; return 0 }

// Eat up any whitespace. Leading whitespace will be considered part of
// the trailing comments.

// Find the beginning of the comment

// Badly formatted sql :/

// Check if this is a MySQL version comment (/*!NNNNN ... */) or
// a TiDB extension comment (/*T! ... */ or /*T![feature] ... */).
// These are NOT actual comments - they contain SQL code that should be executed.

// leadingCommentEnd returns the first index after all leading comments, or
// 0 if there are no leading comments.
func leadingCommentEnd(text string) (end int) { _ = "STUB: not implemented"; return 0 }

// Eat up any whitespace. Trailing whitespace will be considered part of
// the leading comments.

// Found visible characters. Look for '/*' at the beginning
// and '*/' somewhere after that.

// Missing end comment :/

func isNonSpace(r rune) bool { _ = "STUB: not implemented"; return false }
