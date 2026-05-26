package sqldef

import (
	_ "embed"

	"github.com/sqldef/sqldef/v3/database"
	"github.com/sqldef/sqldef/v3/schema"
)

//go:embed VERSION
var version string

// GetVersion returns the version of sqldef read from VERSION file.
func GetVersion() string { _ = "STUB: not implemented"; return "" }

// GetRevision returns the git revision of sqldef.
// It is automatically populated from Go's embedded VCS info.
// Returns empty string if vcs.revision is not available (e.g., when built without .git directory).
func GetRevision() string { _ = "STUB: not implemented"; return "" }

// GetFullVersion returns the version string for --version output.
// Returns "X.Y.Z (abc1234)" if revision is available, otherwise just "X.Y.Z".
func GetFullVersion() string { _ = "STUB: not implemented"; return "" }

type Options struct {
	DesiredDDLs string
	CurrentFile string
	DryRun      bool
	Export      bool
	Check       bool // Like --dry-run, but exit 2 when DDL would be applied (CI gate).
	BeforeApply string
	Config      database.GeneratorConfig
}

// CheckExitCode is the exit code returned by --check when the database schema
// diverges from the desired schema (i.e. one or more DDL statements would have
// been applied). Matches the convention of diff(1) and `terraform plan -detailed-exitcode`.
const CheckExitCode = 2

// Main function shared by all commands
func Run(generatorMode schema.GeneratorMode, db database.Database, sqlParser database.Parser, options *Options) {
	_ = "STUB: not implemented"
	// Set the generator config on the database for privilege filtering
	// Note: MySQL will populate MysqlLowerCaseTableNames from the server
	return
}

// Schema is out of sync: signal to CI without erroring.

func ParseFiles(files []string) []string { _ = "STUB: not implemented"; return nil }

// assume default:"-"

func ReadFiles(filepaths []string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ReadFile(filepath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParseSkipTables(skipFile string) []string { _ = "STUB: not implemented"; return nil }
