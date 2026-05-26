package schema

import (
	"github.com/sqldef/sqldef/v3/database"
	"github.com/sqldef/sqldef/v3/parser"
)

type (
	Ident         = database.Ident
	QualifiedName = database.QualifiedName
)

var (
	NewIdentWithQuoteDetected = database.NewIdentWithQuoteDetected
	NewNormalizedIdent        = database.NewNormalizedIdent
)

// identsEqual compares two Idents with quote-awareness based on database mode.
// This is for general identifiers (columns, indexes, constraints) - NOT table names.
//
// For MySQL: Column/index/constraint names are always case-insensitive.
// For PostgreSQL in quote-aware mode: unquoted identifiers are normalized to lowercase.
// For other databases: always uses case-insensitive comparison.
func identsEqual(a, b Ident, mode GeneratorMode, legacyIgnoreQuotes bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Quote-aware comparison: normalize unquoted identifiers to lowercase

// MySQL/MSSQL/SQLite3: always case-insensitive for non-table identifiers

// tableIdentsEqual compares two table/schema name Idents with quote-awareness.
// This respects MySQL's lower_case_table_names setting which only affects table names.
//
// For MySQL: respects mysqlLowerCaseTableNames setting:
//   - 0 (Linux default): Case-sensitive comparison
//   - 1 or 2 (Windows/macOS): Case-insensitive comparison
//
// For PostgreSQL in quote-aware mode: unquoted identifiers are normalized to lowercase.
// For other databases: always uses case-insensitive comparison.
func tableIdentsEqual(a, b Ident, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) bool {
	_ = "STUB: not implemented"
	return false
}

// Case-sensitive: exact match required

// Case-insensitive (1 or 2)

// Quote-aware comparison: normalize unquoted identifiers to lowercase

// MSSQL/SQLite3: always case-insensitive

// qualifiedNamesEqual compares two QualifiedNames (table names) with quote-awareness.
// Uses tableIdentsEqual since this is specifically for table name comparison.
func qualifiedNamesEqual(a, b QualifiedName, defaultSchema string, mode GeneratorMode, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) bool {
	_ = "STUB: not implemented"
	return false
}

type DDL interface {
	Statement() string
}

type CreateTable struct {
	statement string
	table     Table
}

type CreateIndex struct {
	statement string
	tableName QualifiedName
	index     Index
}

type AddIndex struct {
	statement  string
	tableName  QualifiedName
	constraint bool
	index      Index
}

type AddPrimaryKey struct {
	statement string
	tableName QualifiedName
	index     Index
}

type AddForeignKey struct {
	statement  string
	tableName  QualifiedName
	foreignKey ForeignKey
}

type AddExclusion struct {
	statement string
	tableName QualifiedName
	exclusion Exclusion
}

type AddPolicy struct {
	statement string
	tableName QualifiedName
	policy    Policy
}

type GrantPrivilege struct {
	statement  string
	tableName  QualifiedName
	grantees   []string
	privileges []string
}

type RevokePrivilege struct {
	statement     string
	tableName     QualifiedName
	grantees      []string
	privileges    []string
	cascadeOption bool // CASCADE option for REVOKE
}

// CreatePartitionOf represents a PostgreSQL CREATE TABLE ... PARTITION OF statement
type CreatePartitionOf struct {
	statement   string
	tableName   QualifiedName
	parentTable QualifiedName
	boundSpec   PartitionBound
}

// PartitionBound represents the partition bound specification
type PartitionBound struct {
	// For RANGE partitions: FROM (values) TO (values)
	From parser.Exprs
	To   parser.Exprs
	// For LIST partitions: IN (values)
	In parser.Exprs
	// For DEFAULT partition
	IsDefault bool
}

type Table struct {
	name        QualifiedName
	columns     map[string]*Column
	indexes     []Index
	checks      []CheckDefinition
	foreignKeys []ForeignKey
	exclusions  []Exclusion
	policies    []Policy
	privileges  []TablePrivilege
	options     map[string]string
	renamedFrom Ident           // Previous table name if renamed via @renamed annotation
	partition   *TablePartition // Partition definition (MySQL/MariaDB)
}

// TablePartition represents partition information for a table
type TablePartition struct {
	Type        string                // RANGE, RANGE COLUMNS, LIST, LIST COLUMNS, HASH, LINEAR HASH, KEY, LINEAR KEY
	Definitions []PartitionDefinition // Individual partition definitions
}

// PartitionDefinition represents a single partition
type PartitionDefinition struct {
	Name     Ident
	LessThan parser.Exprs // For RANGE: VALUES LESS THAN
	In       parser.Exprs // For LIST: VALUES IN
	Maxvalue bool         // For VALUES LESS THAN MAXVALUE
}

type Column struct {
	name                       Ident
	position                   int
	typeName                   string
	typeIdent                  Ident // Type name with quote information (for custom types like domains)
	unsigned                   bool
	notNull                    *bool
	autoIncrement              bool
	autoRandom                 bool
	autoRandomShardBits        int
	autoRandomRange            int
	array                      bool
	defaultDef                 *DefaultDefinition
	sridDef                    *SridDefinition
	length                     *Value
	scale                      *Value
	displayWidth               *Value
	check                      *CheckDefinition
	charset                    string
	collate                    string
	timezone                   bool // for Postgres `with time zone`
	keyOption                  ColumnKeyOption
	onUpdate                   *Value
	comment                    *Value
	enumValues                 []string
	references                 Ident
	referenceDeferrable        *bool // for Postgres: DEFERRABLE, NOT DEFERRABLE, or nil
	referenceInitiallyDeferred *bool // for Postgres: INITIALLY DEFERRED, INITIALLY IMMEDIATE, or nil
	identity                   *Identity
	sequence                   *Sequence
	generated                  *Generated
	renamedFrom                Ident // Previous column name if renamed via @renamed annotation
	// TODO: keyopt
	// XXX: zerofill?
}

type Index struct {
	name              Ident
	indexType         string // Parsed only in "create table" but not parsed in "add index". Only used inside `generateDDLsForCreateTable`.
	columns           []IndexColumn
	primary           bool
	unique            bool
	vector            bool // for MariaDB vector indexes
	constraint        bool // for Postgres/MSSQL `ADD CONSTRAINT UNIQUE`
	async             bool // for Aurora DSQL
	concurrently      bool // for PostgreSQL
	constraintOptions *ConstraintOptions
	where             parser.Expr    // for Postgres `Partial Indexes`
	included          []string       // for MSSQL
	clustered         bool           // for MSSQL
	partition         IndexPartition // for MSSQL
	options           []IndexOption
	renamedFrom       Ident // Previous index name if renamed via @renamed annotation
}

type IndexColumn struct {
	columnExpr    parser.Expr // never nil as it's always initialized in the parser
	length        *int
	direction     string
	operatorClass string

	withoutOverlaps bool
}

// ColumnName returns the column name if this is a simple column reference.
// For functional indexes or expressions, it returns the string representation.
// FIXME: parser.String(ic.columnExpr) is not actually a correct column name.
func (ic IndexColumn) ColumnName() string {
	_ = "STUB: not implemented"
	// Check if it's a simple column reference (ColName)
	return ""
}

// For expressions, return the full expression string

// IndexColumn.direction
const (
	AscScr  = "asc"
	DescScr = "desc"
)

type IndexOption struct {
	optionName string
	value      *Value
}

type IndexPartition struct {
	partitionName string
	column        string
}

type ConstraintOptions struct {
	deferrable        bool
	initiallyDeferred bool
}

type ForeignKey struct {
	constraintName     Ident
	indexName          Ident
	indexColumns       []Ident
	referenceTableName QualifiedName
	referenceColumns   []Ident
	onDelete           string
	onUpdate           string
	notForReplication  bool
	constraintOptions  *ConstraintOptions
	period             bool
}

type Exclusion struct {
	constraintName Ident
	indexType      string
	where          parser.Expr
	exclusions     []ExclusionPair
}

type ExclusionPair struct {
	expression string
	operator   string
}

type Policy struct {
	name          Ident
	referenceName string
	permissive    string
	scope         string
	roles         []string
	using         parser.Expr
	withCheck     parser.Expr
}

type TablePrivilege struct {
	tableName       string
	grantee         string
	privileges      []string
	withGrantOption bool
}

type View struct {
	statement    string
	viewType     string
	securityType string
	name         QualifiedName
	definition   parser.SelectStatement // never nil
	indexes      []Index
	columns      []string
	withData     bool // true for "WITH DATA"
	withNoData   bool // true for "WITH NO DATA"
}

// TriggerEvent represents a single trigger event (INSERT, UPDATE, DELETE, or UPDATE OF columns)
type TriggerEvent struct {
	eventType string  // "INSERT", "UPDATE", "DELETE"
	columns   []Ident // For UPDATE OF col1, col2 - nil for INSERT/DELETE/plain UPDATE
}

type Trigger struct {
	statement     string
	name          QualifiedName
	tableName     QualifiedName
	time          string
	event         []TriggerEvent
	whenCondition string
	body          []string
}

// Function represents a PostgreSQL CREATE FUNCTION statement
type Function struct {
	statement  string
	name       QualifiedName
	args       []FunctionArg
	returnType string
	body       string
	language   string
	orReplace  bool
	options    []string // Additional options like IMMUTABLE, SECURITY DEFINER, SET timezone = 'UTC', etc.
}

// FunctionArg is the schema-level representation of a function argument.
// Used by dependency analysis to detect Function -> Type/Domain edges.
type FunctionArg struct {
	name Ident
	typ  string
}

type Value struct {
	valueType ValueType
	raw       string

	// ValueType-specific (behaves like a union)
	strVal   string  // ValueTypeStr
	intVal   int     // ValueTypeInt
	floatVal float64 // ValueTypeFloat
	bitVal   bool    // ValueTypeBit, ValueTypeBool
}

type ValueType int

const (
	ValueTypeStr = ValueType(iota)
	ValueTypeInt
	ValueTypeFloat
	ValueTypeHexNum
	ValueTypeHex
	ValueTypeValArg
	ValueTypeBit
	ValueTypeBool
)

type ColumnKeyOption int

const (
	ColumnKeyNone = ColumnKeyOption(iota)
	ColumnKeyPrimary
	ColumnKeySpatialKey
	ColumnKeyUnique
	ColumnKeyUniqueKey
	ColumnKey
)

type Identity struct {
	behavior          string
	notForReplication bool
}

type Sequence struct {
	Name        string
	IfNotExists bool
	Type        string
	IncrementBy *int
	MinValue    *int
	NoMinValue  bool
	MaxValue    *int
	NoMaxValue  bool
	StartWith   *int
	Cache       *int
	Cycle       bool
	NoCycle     bool
	OwnedBy     string
}

type DefaultDefinition struct {
	expression     parser.Expr // never nil
	constraintName Ident       // only for MSSQL
}

type SridDefinition struct {
	value *Value
}

type CheckDefinition struct {
	definition        parser.Expr // never nil
	constraintName    Ident
	notForReplication bool
	noInherit         bool
}

// EnumValue represents a single enum value with optional rename information
type EnumValue struct {
	value       string
	renamedFrom Ident // Previous enum value if renamed via @renamed annotation
}

// TODO: include type information
type Type struct {
	name       QualifiedName
	statement  string
	enumValues []EnumValue
}

type Domain struct {
	name         QualifiedName
	statement    string
	dataType     string
	defaultValue *DefaultDefinition
	notNull      bool
	collation    string
	constraints  []DomainConstraint
}

type DomainConstraint struct {
	name       string
	expression parser.Expr
}

type Generated struct {
	expr          string
	exprAST       parser.Expr // preserved for dependency analysis (function calls in the expression)
	generatedType GeneratedType
}

type GeneratedType int

const (
	GeneratedTypeVirtual = GeneratedType(iota)
	GeneratedTypeStored
)

type Comment struct {
	statement string
	comment   parser.Comment
}

type Extension struct {
	statement string
	extension parser.Extension
}

type Schema struct {
	statement string
	schema    parser.Schema
}

func (c *CreateTable) Statement() string { _ = "STUB: not implemented"; return "" }

func (c *CreateIndex) Statement() string { _ = "STUB: not implemented"; return "" }

func (a *AddIndex) Statement() string { _ = "STUB: not implemented"; return "" }

func (a *AddPrimaryKey) Statement() string { _ = "STUB: not implemented"; return "" }

func (a *AddForeignKey) Statement() string { _ = "STUB: not implemented"; return "" }

func (a *AddExclusion) Statement() string { _ = "STUB: not implemented"; return "" }

func (a *AddPolicy) Statement() string { _ = "STUB: not implemented"; return "" }

func (g *GrantPrivilege) Statement() string { _ = "STUB: not implemented"; return "" }

func (r *RevokePrivilege) Statement() string { _ = "STUB: not implemented"; return "" }

func (c *CreatePartitionOf) Statement() string { _ = "STUB: not implemented"; return "" }

func (v *View) Statement() string { _ = "STUB: not implemented"; return "" }

func (t *Trigger) Statement() string { _ = "STUB: not implemented"; return "" }

func (f *Function) Statement() string { _ = "STUB: not implemented"; return "" }

func (t *Type) Statement() string { _ = "STUB: not implemented"; return "" }

func (d *Domain) Statement() string { _ = "STUB: not implemented"; return "" }

func (t *Comment) Statement() string { _ = "STUB: not implemented"; return "" }

func (t *Extension) Statement() string { _ = "STUB: not implemented"; return "" }

func (t *Schema) Statement() string { _ = "STUB: not implemented"; return "" }

func (t *Table) PrimaryKey() *Index { _ = "STUB: not implemented"; return nil }

func (keyOption ColumnKeyOption) isUnique() bool { _ = "STUB: not implemented"; return false }
