/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parser

import (
	"strings"
)

type ParserMode int

const (
	eofChar = 0x100

	ParserModeMysql = ParserMode(iota)
	ParserModePostgres
	ParserModeSQLite3
	ParserModeMssql
)

// The main parser function for sqldef.
func ParseDDL(sql string, mode ParserMode) (Statement, error) {
	_ = "STUB: not implemented"
	return *new(Statement), nil
}

// yyParse returns 0 when error recovery rules are triggered, so check LastError

// Tokenizer is the struct used to generate SQL
// tokens for the parser.
type Tokenizer struct {
	AllowComments        bool
	lastChar             rune
	Position             int
	lastToken            string
	LastError            error
	posVarIndex          int
	ParseTree            Statement
	partialDDL           *DDL
	multi                bool
	specialComment       *Tokenizer
	mode                 ParserMode
	peeking              bool // true when peeking ahead to avoid infinite recursion
	lastIdentifierQuoted bool // true if the last scanned identifier was quoted

	buf     string
	bufPos  int
	bufSize int
}

// NewTokenizer creates a new Tokenizer for a given SQL string.
func NewTokenizer(sql string, mode ParserMode) *Tokenizer { _ = "STUB: not implemented"; return nil }

// keywords maps keyword strings to their token IDs.
// Keywords marked as UNUSED are recognized but not actively used in the grammar.
//
// When adding new keywords, also add them to either the reserved_keyword or
// non_reserved_keyword grammar in parser.y. This allows the keyword to be used
// as an identifier in certain contexts.
var keywords = map[string]int{
	"accessible":             UNUSED,
	"action":                 ACTION,
	"add":                    ADD,
	"after":                  AFTER,
	"against":                AGAINST,
	"all":                    ALL,
	"alter":                  ALTER,
	"always":                 ALWAYS,
	"analyze":                ANALYZE,
	"and":                    AND,
	"array":                  ARRAY,
	"as":                     AS,
	"asc":                    ASC,
	"asensitive":             UNUSED,
	"auto_increment":         AUTO_INCREMENT,
	"auto_random":            AUTO_RANDOM,
	"autoincrement":          AUTOINCREMENT,
	"before":                 BEFORE,
	"begin":                  BEGIN,
	"between":                BETWEEN,
	"bigint":                 BIGINT,
	"bigserial":              BIGSERIAL,
	"binary":                 BINARY,
	"_binary":                UNDERSCORE_BINARY,
	"bit":                    BIT,
	"blob":                   BLOB,
	"bool":                   BOOL,
	"boolean":                BOOLEAN,
	"both":                   UNUSED,
	"bpchar":                 BPCHAR,
	"by":                     BY,
	"cache":                  CACHE,
	"call":                   UNUSED,
	"called":                 CALLED,
	"cascade":                CASCADE,
	"case":                   CASE,
	"cast":                   CAST,
	"change":                 UNUSED,
	"char":                   CHAR,
	"character":              CHARACTER,
	"charset":                CHARSET,
	"check":                  CHECK,
	"citext":                 CITEXT,
	"close":                  CLOSE,
	"clustered":              CLUSTERED,
	"nonclustered":           NONCLUSTERED,
	"collate":                COLLATE,
	"column":                 COLUMN,
	"columns":                COLUMNS,
	"columnstore":            COLUMNSTORE,
	"comment":                COMMENT_KEYWORD,
	"committed":              COMMITTED,
	"commit":                 COMMIT,
	"concurrently":           CONCURRENTLY,
	"async":                  ASYNC,
	"condition":              UNUSED,
	"constraint":             CONSTRAINT,
	"continue":               CONTINUE,
	"create":                 CREATE,
	"convert":                CONVERT,
	"cosine":                 COSINE,
	"cost":                   COST,
	"substr":                 SUBSTR,
	"substring":              SUBSTRING,
	"cross":                  CROSS,
	"current_date":           CURRENT_DATE,
	"current_time":           CURRENT_TIME,
	"current_timestamp":      CURRENT_TIMESTAMP,
	"current_user":           CURRENT_USER,
	"cursor":                 CURSOR,
	"cycle":                  CYCLE,
	"database":               DATABASE,
	"databases":              DATABASES,
	"day_hour":               UNUSED,
	"day_microsecond":        UNUSED,
	"day_minute":             UNUSED,
	"day_second":             UNUSED,
	"date":                   DATE,
	"data":                   DATA,
	"daterange":              DATERANGE,
	"datetime":               DATETIME,
	"datetime2":              DATETIME2,
	"datetimeoffset":         DATETIMEOFFSET,
	"deallocate":             DEALLOCATE,
	"dec":                    UNUSED,
	"decimal":                DECIMAL,
	"declare":                DECLARE,
	"default":                DEFAULT,
	"deferrable":             DEFERRABLE,
	"deferred":               DEFERRED,
	"definer":                DEFINER,
	"delayed":                UNUSED,
	"delete":                 DELETE,
	"desc":                   DESC,
	"describe":               DESCRIBE,
	"deterministic":          UNUSED,
	"distance":               DISTANCE,
	"domain":                 DOMAIN,
	"distinct":               DISTINCT,
	"distinctrow":            UNUSED,
	"div":                    DIV,
	"double":                 DOUBLE,
	"drop":                   DROP,
	"duplicate":              DUPLICATE,
	"each":                   EACH,
	"else":                   ELSE,
	"elseif":                 UNUSED,
	"enclosed":               UNUSED,
	"end":                    END,
	"engine":                 ENGINE,
	"enum":                   ENUM,
	"euclidean":              EUCLIDEAN,
	"escape":                 ESCAPE,
	"escaped":                UNUSED,
	"exclude":                EXCLUDE,
	"exists":                 EXISTS,
	"exec":                   EXEC,
	"execute":                EXECUTE,
	"except":                 EXCEPT,
	"exit":                   EXIT,
	"explain":                EXPLAIN,
	"extension":              EXTENSION,
	"expansion":              EXPANSION,
	"extract":                EXTRACT,
	"extended":               EXTENDED,
	"false":                  FALSE,
	"fetch":                  FETCH,
	"first":                  FIRST,
	"float":                  FLOAT_TYPE,
	"float4":                 UNUSED,
	"float8":                 UNUSED,
	"for":                    FOR,
	"force":                  FORCE,
	"foreign":                FOREIGN,
	"found":                  FOUND,
	"from":                   FROM,
	"full":                   FULL,
	"fulltext":               FULLTEXT,
	"function":               FUNCTION,
	"generated":              GENERATED,
	"geometry":               GEOMETRY,
	"geometrycollection":     GEOMETRYCOLLECTION,
	"get":                    UNUSED,
	"getdate":                GETDATE,
	"global":                 GLOBAL,
	"grant":                  GRANT,
	"group":                  GROUP,
	"group_concat":           GROUP_CONCAT,
	"handler":                HANDLER,
	"hash":                   HASH,
	"having":                 HAVING,
	"high_priority":          UNUSED,
	"holdlock":               HOLDLOCK,
	"hour_microsecond":       UNUSED,
	"hour_minute":            UNUSED,
	"hour_second":            UNUSED,
	"identity":               IDENTITY,
	"if":                     IF,
	"ignore":                 IGNORE,
	"immediate":              IMMEDIATE,
	"immutable":              IMMUTABLE,
	"in":                     IN,
	"include":                INCLUDE,
	"increment":              INCREMENT,
	"index":                  INDEX,
	"infile":                 UNUSED,
	"input":                  INPUT,
	"inout":                  INOUT,
	"inner":                  INNER,
	"initially":              INITIALLY,
	"insensitive":            UNUSED,
	"insert":                 INSERT,
	"instead":                INSTEAD,
	"int":                    INT,
	"int1":                   UNUSED,
	"int2":                   UNUSED,
	"int3":                   UNUSED,
	"int4":                   UNUSED,
	"int4range":              INT4RANGE,
	"int8":                   UNUSED,
	"int8range":              INT8RANGE,
	"integer":                INTEGER,
	"intersect":              INTERSECT,
	"interval":               INTERVAL,
	"into":                   INTO,
	"invoker":                INVOKER,
	"io_after_gtids":         UNUSED,
	"is":                     IS,
	"isolation":              ISOLATION,
	"inherit":                INHERIT,
	"iterate":                UNUSED,
	"join":                   JOIN,
	"json":                   JSON,
	"jsonb":                  JSONB,
	"key":                    KEY,
	"keys":                   KEYS,
	"key_block_size":         KEY_BLOCK_SIZE,
	"kill":                   UNUSED,
	"language":               LANGUAGE,
	"leakproof":              LEAKPROOF,
	"last":                   LAST,
	"last_insert_id":         LAST_INSERT_ID,
	"leading":                UNUSED,
	"leave":                  UNUSED,
	"left":                   LEFT,
	"less":                   LESS,
	"level":                  LEVEL,
	"like":                   LIKE,
	"limit":                  LIMIT,
	"list":                   LIST,
	"linear":                 LINEAR,
	"lines":                  UNUSED,
	"linestring":             LINESTRING,
	"load":                   UNUSED,
	"localtime":              LOCALTIME,
	"localtimestamp":         LOCALTIMESTAMP,
	"lock":                   LOCK,
	"long":                   UNUSED,
	"longblob":               LONGBLOB,
	"longtext":               LONGTEXT,
	"loop":                   UNUSED,
	"low_priority":           UNUSED,
	"m":                      M,
	"master_bind":            UNUSED,
	"match":                  MATCH,
	"materialized":           MATERIALIZED,
	"maxvalue":               MAXVALUE,
	"mediumblob":             MEDIUMBLOB,
	"mediumint":              MEDIUMINT,
	"mediumtext":             MEDIUMTEXT,
	"middleint":              UNUSED,
	"minute_microsecond":     UNUSED,
	"minute_second":          UNUSED,
	"minvalue":               MINVALUE,
	"mod":                    MOD,
	"mode":                   MODE,
	"modifies":               UNUSED,
	"money":                  MONEY,
	"multilinestring":        MULTILINESTRING,
	"multipoint":             MULTIPOINT,
	"multipolygon":           MULTIPOLYGON,
	"names":                  NAMES,
	"natural":                NATURAL,
	"nchar":                  NCHAR,
	"new":                    NEW,
	"next":                   NEXT,
	"no":                     NO,
	"nolock":                 NOLOCK,
	"none":                   NONE,
	"not":                    NOT,
	"now":                    NOW,
	"nowait":                 NOWAIT,
	"no_write_to_binlog":     UNUSED,
	"ntext":                  NTEXT,
	"null":                   NULL,
	"nulls":                  NULLS,
	"numeric":                NUMERIC,
	"numrange":               NUMRANGE,
	"nvarchar":               NVARCHAR,
	"of":                     OF,
	"offset":                 OFFSET,
	"on":                     ON,
	"only":                   ONLY,
	"open":                   OPEN,
	"optimize":               OPTIMIZE,
	"optimizer_costs":        UNUSED,
	"option":                 OPTION,
	"optionally":             UNUSED,
	"or":                     OR,
	"order":                  ORDER,
	"out":                    OUT,
	"outer":                  OUTER,
	"outfile":                UNUSED,
	"output":                 OUTPUT,
	"over":                   OVER,
	"overlaps":               OVERLAPS,
	"owned":                  OWNED,
	"paglock":                PAGLOCK,
	"parallel":               PARALLEL,
	"parser":                 PARSER,
	"partial":                PARTIAL,
	"period":                 PERIOD,
	"partition":              PARTITION,
	"partitions":             PARTITIONS,
	"permissive":             PERMISSIVE,
	"point":                  POINT,
	"policy":                 POLICY,
	"polygon":                POLYGON,
	"precision":              PRECISION,
	"primary":                PRIMARY,
	"prior":                  PRIOR,
	"privileges":             PRIVILEGES,
	"processlist":            PROCESSLIST,
	"procedure":              PROCEDURE,
	"query":                  QUERY,
	"restrictive":            RESTRICTIVE,
	"range":                  RANGE,
	"read":                   READ,
	"reads":                  UNUSED,
	"read_write":             UNUSED,
	"readuncommitted":        READUNCOMMITTED,
	"real":                   REAL,
	"recursive":              RECURSIVE,
	"references":             REFERENCES,
	"regexp":                 REGEXP,
	"release":                UNUSED,
	"rename":                 RENAME,
	"reorganize":             REORGANIZE,
	"repair":                 REPAIR,
	"repeat":                 UNUSED,
	"repeatable":             REPEATABLE,
	"replace":                REPLACE,
	"replication":            REPLICATION,
	"require":                UNUSED,
	"resignal":               UNUSED,
	"restrict":               RESTRICT,
	"restricted":             RESTRICTED,
	"return":                 RETURN,
	"returns":                RETURNS,
	"revoke":                 REVOKE,
	"right":                  RIGHT,
	"rlike":                  REGEXP,
	"rollback":               ROLLBACK,
	"row":                    ROW,
	"rowid":                  ROWID,
	"rowlock":                ROWLOCK,
	"rows":                   ROWS,
	"safe":                   SAFE,
	"schema":                 SCHEMA,
	"schemas":                UNUSED,
	"scroll":                 SCROLL,
	"second_microsecond":     UNUSED,
	"security":               SECURITY,
	"select":                 SELECT,
	"sensitive":              UNUSED,
	"separator":              SEPARATOR,
	"sequence":               UNUSED,
	"serial":                 SERIAL,
	"serializable":           SERIALIZABLE,
	"session":                SESSION,
	"set":                    SET,
	"setof":                  SETOF,
	"share":                  SHARE,
	"show":                   SHOW,
	"signal":                 UNUSED,
	"signed":                 SIGNED,
	"simple":                 SIMPLE,
	"smalldatetime":          SMALLDATETIME,
	"smallint":               SMALLINT,
	"smallmoney":             SMALLMONEY,
	"smallserial":            SMALLSERIAL,
	"spatial":                SPATIAL,
	"specific":               UNUSED,
	"sql":                    SQL,
	"sqlexception":           SQLEXCEPTION,
	"sqlstate":               SQLSTATE,
	"sqlwarning":             SQLWARNING,
	"sql_big_result":         UNUSED,
	"sql_cache":              SQL_CACHE,
	"sql_calc_found_rows":    UNUSED,
	"sql_no_cache":           SQL_NO_CACHE,
	"sql_small_result":       UNUSED,
	"srid":                   SRID,
	"ssl":                    UNUSED,
	"stable":                 STABLE,
	"start":                  START,
	"starting":               UNUSED,
	"status":                 STATUS,
	"stored":                 STORED,
	"straight_join":          STRAIGHT_JOIN,
	"stream":                 STREAM,
	"strict":                 STRICT,
	"table":                  TABLE,
	"tables":                 TABLES,
	"tablock":                TABLOCK,
	"terminated":             UNUSED,
	"text":                   TEXT,
	"text_pattern_ops":       TEXT_PATTERN_OPS,
	"than":                   THAN,
	"then":                   THEN,
	"time":                   TIME,
	"timestamp":              TIMESTAMP,
	"tinyblob":               TINYBLOB,
	"tinyint":                TINYINT,
	"tinytext":               TINYTEXT,
	"to":                     TO,
	"top":                    TOP,
	"trailing":               UNUSED,
	"transaction":            TRANSACTION,
	"trigger":                TRIGGER,
	"true":                   TRUE,
	"tsrange":                TSRANGE,
	"tstzrange":              TSTZRANGE,
	"truncate":               TRUNCATE,
	"type":                   TYPE,
	"uncommitted":            UNCOMMITTED,
	"undo":                   UNUSED,
	"union":                  UNION,
	"unique":                 UNIQUE,
	"unlock":                 UNUSED,
	"unsigned":               UNSIGNED,
	"unsafe":                 UNSAFE,
	"update":                 UPDATE,
	"updlock":                UPDLOCK,
	"usage":                  UNUSED,
	"use":                    USE,
	"using":                  USING,
	"utc_date":               UTC_DATE,
	"utc_time":               UTC_TIME,
	"utc_timestamp":          UTC_TIMESTAMP,
	"uniqueidentifier":       UNIQUEIDENTIFIER,
	"uuid":                   UUID,
	"value":                  VALUE,
	"values":                 VALUES,
	"variables":              VARIABLES,
	"variadic":               VARIADIC,
	"varbinary":              VARBINARY,
	"varchar":                VARCHAR,
	"varcharacter":           UNUSED,
	"varying":                VARYING,
	"vector":                 VECTOR,
	"virtual":                VIRTUAL,
	"view":                   VIEW,
	"volatile":               VOLATILE,
	"vschema_tables":         VSCHEMA_TABLES,
	"when":                   WHEN,
	"where":                  WHERE,
	"while":                  WHILE,
	"with":                   WITH,
	"without":                WITHOUT,
	"write":                  WRITE,
	"xor":                    UNUSED,
	"year":                   YEAR,
	"year_month":             UNUSED,
	"zerofill":               ZEROFILL,
	"zone":                   ZONE,
	"::":                     TYPECAST,
	"allow_page_locks":       ALLOW_PAGE_LOCKS,
	"allow_row_locks":        ALLOW_ROW_LOCKS,
	"fillfactor":             FILLFACTOR,
	"ignore_dup_key":         IGNORE_DUP_KEY,
	"off":                    OFF,
	"pad_index":              PAD_INDEX,
	"statistics_incremental": STATISTICS_INCREMENTAL,
	"statistics_norecompute": STATISTICS_NORECOMPUTE,
	"lead":                   LEAD,
	"lag":                    LAG,
	"getutcdate":             GETUTCDATE,
	"newid":                  NEWID,
	"sysutcdatetime":         SYSUTCDATETIME,
	"newsequentialid":        NEWSEQUENTIALID,
	"openjson":               OPENJSON,
	"string_split":           STRING_SPLIT,
	"apply":                  APPLY,
	"within":                 WITHIN,
	"trim":                   TRIM,
	"try_cast":               TRY_CAST,

	// SET options for SQL Server
	"concat_null_yields_null":  CONCAT_NULL_YIELDS_NULL,
	"cursor_close_on_commit":   CURSOR_CLOSE_ON_COMMIT,
	"quoted_identifier":        QUOTED_IDENTIFIER,
	"arithabort":               ARITHABORT,
	"fmtonly":                  FMTONLY,
	"nocount":                  NOCOUNT,
	"noexec":                   NOEXEC,
	"numeric_roundabort":       NUMERIC_ROUNDABORT,
	"ansi_defaults":            ANSI_DEFAULTS,
	"ansi_null_dflt_off":       ANSI_NULL_DFLT_OFF,
	"ansi_null_dflt_on":        ANSI_NULL_DFLT_ON,
	"ansi_nulls":               ANSI_NULLS,
	"ansi_padding":             ANSI_PADDING,
	"ansi_warnings":            ANSI_WARNINGS,
	"forceplan":                FORCEPLAN,
	"showplan_all":             SHOWPLAN_ALL,
	"showplan_text":            SHOWPLAN_TEXT,
	"showplan_xml":             SHOWPLAN_XML,
	"implicit_transactions":    IMPLICIT_TRANSACTIONS,
	"remote_proc_transactions": REMOTE_PROC_TRANSACTIONS,
	"xact_abort":               XACT_ABORT,
}

// keywordStrings contains the reverse mapping of token to keyword strings
var keywordStrings = map[int]string{}

// IsKeyword returns true if the given string is a SQL keyword.
// The check is case-insensitive.
func IsKeyword(s string) bool { _ = "STUB: not implemented"; return false }

var encodeRef = map[byte]byte{
	'\x00': '0',
	'\'':   '\'',
	'"':    '"',
	'\b':   'b',
	'\n':   'n',
	'\r':   'r',
	'\t':   't',
	26:     'Z', // ctl-Z
	'\\':   '\\',
}

// sqlEncodeMap specifies how to escape binary data with '\'.
// Complies to http://dev.mysql.com/doc/refman/5.1/en/string-syntax.html
var sqlEncodeMap [256]byte

// sqlDecodeMap is the reverse of sqlEncodeMap
var sqlDecodeMap [256]byte

// dontEscape tells you if a character should not be escaped.
var dontEscape = byte(255)

func init() {
	// Convert keywords to keywordStrings
	for str, id := range keywords {
		if id == UNUSED {
			continue
		}
		keywordStrings[id] = str
	}

	// Convert encodeRef to sqlEncodeMap and sqlDecodeMap
	for i := range sqlEncodeMap {
		sqlEncodeMap[i] = dontEscape
		sqlDecodeMap[i] = dontEscape
	}
	for i := range sqlEncodeMap {
		b := byte(i)
		if to, ok := encodeRef[b]; ok {
			sqlEncodeMap[b] = to
			sqlDecodeMap[to] = b
		}
	}
}

// Lex returns the next token form the Tokenizer.
// This function is used by go yacc.
func (tkn *Tokenizer) Lex(lval *yySymType) int { _ = "STUB: not implemented"; return 0 }

// For ID tokens, create an Ident with value and quoted flag

// For other tokens, use the str field as before

func (tkn *Tokenizer) getLineInfo(position int) (lineNum int, lineContent string, columnNum int) {
	_ = "STUB: not implemented"
	return 0, "", 0
}

// Find the end of the current line

// Extract the line content

// Calculate column number (position within the line)

// Error is called by go yacc if there's a parsing error.
func (tkn *Tokenizer) Error(err string) { _ = "STUB: not implemented"; return }

// Add a pointer to show the exact position

// Try and re-sync to the next statement

// Scan scans the tokenizer for the next token and returns
// the token type and an optional value.
func (tkn *Tokenizer) Scan() (int, string) { _ = "STUB: not implemented"; return 0, "" }

// Enter specialComment scan mode.
// for scanning such kind of comment: /*! MySQL-specific code */

// return the specialComment scan result as the result

// leave specialComment scan mode after all stream consumed.

// Check for ~~ (LIKE) and ~~* (ILIKE) pattern operators

// Check for ~* (case-insensitive regex) or ~ (regex)

// PostgreSQL user-defined operator starting with '<='
// e.g., pgvector cosine distance: <=>

// PostgreSQL user-defined operator starting with '<'
// e.g., pgvector: <->, <#>, <+>

// Check for !~~* (NOT ILIKE) and !~~ (NOT LIKE)

// Check for !~* (NOT case-insensitive regex) or !~ (NOT regex)

// PostgreSQL dollar-quoted strings: $$...$$ or $tag$...$tag$

// skipStatement scans until the EOF, or end of statement is encountered.
func (tkn *Tokenizer) skipStatement() { _ = "STUB: not implemented"; return }

func (tkn *Tokenizer) skipBlank() { _ = "STUB: not implemented"; return }

func (tkn *Tokenizer) scanIdentifier(firstChar rune, isDbSystemVariable bool) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// Context-aware handling for "with" keyword
// Only peek if we're not already in a peek operation (prevents infinite recursion)

// If next token is DATA or NO (for "WITH NO DATA"), this is a data option

// PostgreSQL treats KEY as a non-reserved keyword usable as an unquoted
// column name. Surface a distinct PG_KEY token so the grammar can accept
// `key text NOT NULL` without colliding with MySQL's inline `KEY idx_name (col)`.

// keyword is case-insensitive

// dual must always be case-insensitive

// others are case-sensitive (unquoted identifiers)

func (tkn *Tokenizer) scanHex() (int, string) { _ = "STUB: not implemented"; return 0, "" }

func (tkn *Tokenizer) scanBitLiteral() (int, string) { _ = "STUB: not implemented"; return 0, "" }

func (tkn *Tokenizer) scanLiteralIdentifier(sepChar rune) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// The previous char was not a backtick.

// Premature EOF.

// Literal identifiers are quoted

func (tkn *Tokenizer) scanBindVar() (int, string) { _ = "STUB: not implemented"; return 0, "" }

func (tkn *Tokenizer) scanMantissa(base int, buffer *strings.Builder) {
	_ = "STUB: not implemented"
	return
}

func (tkn *Tokenizer) scanNumber(seenDecimalPoint bool) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// 0x construct.

// A letter cannot immediately follow a number.

func (tkn *Tokenizer) scanString(delim rune, typ int) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// Unterminated string.

// Reached the end of the buffer without finding a delim or
// escape character.

// String terminates mid escape character.

// Correctly terminated string, which is not a double delim.

// scanDollarQuotedString scans a PostgreSQL dollar-quoted string.
// Supports both $$...$$ and $tag$...$tag$ styles.
func (tkn *Tokenizer) scanDollarQuotedString() (int, string) {
	_ = "STUB: not implemented"
	return 0,

		// Build the opening delimiter (already consumed the first $)
		""
}

// Check if it's $$ or $tag$

// Simple $$ delimiter

// Tagged $tag$ delimiter

// Scan until we find the closing delimiter

// Check if this is the start of the closing delimiter

// Found the closing delimiter

// Not the closing delimiter, restore and include in content

func (tkn *Tokenizer) scanCommentType1(prefix string) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (tkn *Tokenizer) scanCommentType2() (int, string) { _ = "STUB: not implemented"; return 0, "" }

// scanCommentType2OrTiDBComment reads a block comment and checks if it's a
// TiDB-specific comment like /*T![auto_rand] AUTO_RANDOM(5) */.
// If so, the inner SQL is expanded into the token stream via specialComment.
func (tkn *Tokenizer) scanCommentType2OrTiDBComment() (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// extractTiDBComment extracts the SQL from a TiDB-specific comment.
// Two formats are supported:
//   - /*T![feature_name] SQL */ — feature-gated (e.g. auto_rand, clustered_index)
//   - /*T! SQL */               — ungated (e.g. SHARD_ROW_ID_BITS, PRE_SPLIT_REGIONS)
func extractTiDBComment(comment string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// strip /* and */

// Ungated form: /*T! SQL */

// Feature-gated form: /*T![feature_name] SQL */

// Supported features: expand the inner SQL into the token stream

// Unsupported features (e.g. clustered_index): ignore

func (tkn *Tokenizer) consumeNext(buffer *strings.Builder) { _ = "STUB: not implemented"; return }

// This should never happen.

func (tkn *Tokenizer) next() { _ = "STUB: not implemented"; return }

// peekToken peeks ahead to determine the next token
// without consuming any characters. This is used for context-aware tokenization.
func (tkn *Tokenizer) peekToken() (int, string) {
	_ = "STUB: not implemented"
	// Save current state and restore on exit
	return 0, ""
}

// Set peeking flag to prevent infinite recursion

// extractMysqlComment extracts the version and SQL from a comment-only query
// such as /*!50708 sql here */
func extractMysqlComment(sql string) (version string, innerSQL string) {
	_ = "STUB: not implemented"
	return "", ""
}

func isIdentifierFirstChar(ch rune) bool { _ = "STUB: not implemented"; return false }

func isIdentifierMetaChar(ch rune) bool { _ = "STUB: not implemented"; return false }

func isAsciiDigit(ch rune) bool { _ = "STUB: not implemented"; return false }

func digitVal(ch rune) int { _ = "STUB: not implemented"; return 0 }

// larger than any legal digit val

// isOperatorChar returns true if ch is a valid PostgreSQL operator character.
// PostgreSQL allows: + - * / < > = ~ ! @ # % ^ & | ` ?
func isOperatorChar(ch rune) bool { _ = "STUB: not implemented"; return false }

// isSpecialOperatorChar returns true if ch is a "special" operator character.
// Multi-char operators ending in + or - must contain at least one of these.
func isSpecialOperatorChar(ch rune) bool { _ = "STUB: not implemented"; return false }

// scanCustomOp scans a PostgreSQL user-defined operator starting with the given character.
// Called after the first character has been consumed and we've confirmed the next char is a valid operator char.
//
// PostgreSQL operator rules:
// - Cannot contain -- or /* (comment sequences)
// - Multi-char operators ending in + or - must contain at least one of: ~ ! @ # % ^ & | ` ?
func (tkn *Tokenizer) scanCustomOp(firstChar rune) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// Consume operator characters

// Check for prohibited comment sequences: -- and /*

// Put back this character and return what we have so far (minus the prev char)
// This is tricky - for now, return a lex error for invalid operators

// Multi-char operators ending in + or - must contain a special char
