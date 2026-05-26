package testutil

import (
	"log/slog"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/sqldef/sqldef/v3/database"
	"github.com/sqldef/sqldef/v3/schema"
	"github.com/sqldef/sqldef/v3/util"
)

var stripHeredocRegex = regexp.MustCompilePOSIX("^\t*")

type TestCase struct {
	Current            string  // default: empty schema
	Desired            string  // default: empty schema
	Up                 *string // expected DDL for current → desired migration
	Down               *string // expected DDL for desired → current migration
	Output             *string `yaml:"output,omitempty"` // DEPRECATED: use 'up' and 'down' instead
	Error              *string // default: nil
	MinVersion         string  `yaml:"min_version"`
	MaxVersion         string  `yaml:"max_version"`
	User               string
	Flavor             string   // database flavor (e.g., "mariadb", "mysql")
	ManagedRoles       []string `yaml:"managed_roles"`        // Roles whose privileges are managed by sqldef (empty means no privileges are managed)
	EnableDrop         *bool    `yaml:"enable_drop"`          // Whether to enable DROP/REVOKE operations
	LegacyIgnoreQuotes *bool    `yaml:"legacy_ignore_quotes"` // nil or true = ignore quotes (legacy default), false = preserve quotes
	Offline            bool     `yaml:"offline"`
	Config             struct { // Optional config settings for the test
		CreateIndexConcurrently bool `yaml:"create_index_concurrently"`
		DisableDdlTransaction   bool `yaml:"disable_ddl_transaction"`
	} `yaml:"config"`
}

func init() {
	util.InitSlog()

	// In test environments, suppress INFO-level logs to prevent them from contaminating test output comparisons.
	// Users can still see DEBUG/INFO logs by setting LOG_LEVEL=debug or LOG_LEVEL=info environment variable,
	// which will override this default. Warnings and errors will still appear by default.
	if os.Getenv("LOG_LEVEL") == "" {
		// Set default test log level to WARN to hide INFO messages like "Using generic parser only mode"
		opts := &slog.HandlerOptions{
			Level: slog.LevelWarn,
		}
		handler := slog.NewTextHandler(os.Stderr, opts)
		slog.SetDefault(slog.New(handler))
	}
}

// CreateTestDatabaseName generates a unique database name for a test case.
// The name is sanitized to be a valid database name (lowercase, alphanumeric + underscore)
// and uses FNV hash to ensure uniqueness.
//
// Parameters:
//   - testName: The test name to sanitize
//   - dbLimit: Database name length limit. For example:
//   - PostgreSQL: 63 characters
//   - SQL Server: 128 characters
//
// The resulting format is: sqldef_test_{sanitized}_{hash}
// where hash is the first 8 characters of the FNV-1a hash (in hex) of the original test name.
func CreateTestDatabaseName(testName string, dbLimit int) string {
	_ = "STUB: not implemented"
	return ""
}

// Calculate maximum length for the sanitized portion
// dbLimit = len(prefix) + len(sanitized) + len("_") + len(hash)
// sanitized = dbLimit - len(prefix) - 1 - hashLen

// Sanitize the test name: lowercase, replace non-alphanumeric with underscore

// Truncate to maxSanitizedLen to ensure the full name stays within database limits

// Create a short hash from the full test name for uniqueness

func ReadTests(pattern string) (map[string]TestCase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Track which file each test case came from for better error messages

// Check for deprecated 'output' field

// Validate up/down dependency: both must be present or both must be absent

// defaults to true

func RunTest(t *testing.T, db database.Database, test TestCase, mode schema.GeneratorMode, sqlParser database.Parser, version string, allowedFlavor string) {
	_ = "STUB: not implemented"
	return
}

// Determine LegacyIgnoreQuotes: use test value if specified, otherwise default to true (legacy mode)
// default: true for backward compatibility

// Set config first to populate database-specific settings (e.g., MysqlLowerCaseTableNames)

// If test requires a specific flavor, check if it matches the current environment
// Supports both positive (flavor: mariadb) and negative (flavor: !tidb) matching
// Instead of skipping mismatched tests, we run them and expect them to fail.
// This validates that flavor annotations are correct - if a test passes on a
// non-matching flavor, the annotation should be removed or corrected.

// If no flavor is explicitly set, default to "mysql" for MySQL tests

// Negative match: test excludes this flavor, expect failure if running on excluded flavor

// Positive match: test requires specific flavor, expect failure on other flavors

// Create error collector for expected failure mode

// Check results for expected failure mode

// Test failed as expected - flavor annotation is correct

// Test passed when it should have failed - flavor annotation is wrong

// left < right: compareVersion() < 0
// left = right: compareVersion() = 0
// left > right: compareVersion() > 0
func compareVersion(t *testing.T, leftVersion string, rightVersion string) int {
	_ = "STUB: not implemented"
	return 0
}

// Compare only specified segments (e.g., "10.0" vs "10" -> compare "10" and "10")

// parseVersionSegment extracts the leading numeric part from a version segment.
// This handles formats like "8" -> 8, "4-TiDB-v8" -> 4, "11-MariaDB" -> 11
func parseVersionSegment(segment string) (int, error) {
	_ = "STUB: not implemented"
	// Find the first non-digit character
	return 0, nil
}

func splitDDLs(mode schema.GeneratorMode, sqlParser database.Parser, str string, defaultSchema string, legacyIgnoreQuotes bool, mysqlLowerCaseTableNames int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func runDDLs(db database.Database, ddls []string) error { _ = "STUB: not implemented"; return nil }

/* beforeApply */ /* ddlSuffix */

func joinDDLs(ddls []string) string { _ = "STUB: not implemented"; return "" }

// filterSkippedDDLs removes DDLs that were commented out (skipped) from the list.
// This is used for idempotency checks because skipped DDLs represent intentionally
// unapplied changes that will naturally reappear on subsequent comparisons.
func filterSkippedDDLs(ddls []string) []string { _ = "STUB: not implemented"; return nil }

// MustExecute executes a command within a test and fails the test if it errors.
func MustExecute(t *testing.T, command string, args ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// MustExecuteNoTest executes a command and terminates the program if it errors.
// Use this in TestMain or other setup code where *testing.T is not available.
func MustExecuteNoTest(command string, args ...string) string { _ = "STUB: not implemented"; return "" }

// BuildForTest builds the current package, adding -cover flag if GOCOVERDIR is set.
// Use this in TestMain to build binaries that support coverage collection.
func BuildForTest() { _ = "STUB: not implemented"; return }

func Execute(command string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type stringLogger struct {
	buf strings.Builder
}

func (l *stringLogger) Print(v ...any) { _ = "STUB: not implemented"; return }

func (l *stringLogger) Printf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (l *stringLogger) Println(v ...any) { _ = "STUB: not implemented"; return }

func (l *stringLogger) String() string { _ = "STUB: not implemented"; return "" }

// ApplyWithOutput applies desired DDLs to a database and returns the CLI output format
// This mimics the behavior of running psqldef/mysqldef/etc from the command line
func ApplyWithOutput(db database.Database, mode schema.GeneratorMode, sqlParser database.Parser, desiredDDLs string, config database.GeneratorConfig) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

/* beforeApply */

// QueryRows executes a query and returns the results as a tab-separated string.
// This is a common helper for all *Query functions in *def_test.go files.
func QueryRows(db database.Database, query string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func WriteFile(path string, content string) { _ = "STUB: not implemented"; return }

func StripHeredoc(heredoc string) string { _ = "STUB: not implemented"; return "" }

// errorCollector is used to intercept test failures when running in "expect failure" mode.
// When a test's flavor doesn't match the current environment, we expect the test to fail.
// This collector captures errors instead of failing the test immediately, allowing us to
// verify the failure behavior at the end.
type errorCollector struct {
	t             *testing.T
	expectFailure bool
	hasErrors     bool
	errors        []string
}

func (c *errorCollector) Helper() { _ = "STUB: not implemented"; return }

func (c *errorCollector) Error(args ...any) { _ = "STUB: not implemented"; return }

func (c *errorCollector) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (c *errorCollector) Fatal(args ...any) { _ = "STUB: not implemented"; return }

// In expectFailure mode, panic to stop test execution (will be recovered)

func (c *errorCollector) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

// In expectFailure mode, panic to stop test execution (will be recovered)

// errorCollectorStop is a special panic value used to stop test execution
// in expectFailure mode without actually failing the test.
type errorCollectorStop struct{}

// runTestImplWithCollector runs the test implementation with error collection support.
// When expectFailure is true, errors are collected instead of failing the test immediately.
func runTestImplWithCollector(collector *errorCollector, db database.Database, test TestCase, mode schema.GeneratorMode, sqlParser database.Parser, config database.GeneratorConfig, version string) {
	_ = "STUB: not implemented"
	return
}

// Recover from panics caused by Fatal/Fatalf in expectFailure mode

// Re-panic if it's not our stop signal

// Otherwise, the test stopped due to an expected failure

// testReporter is an interface for reporting test failures, allowing us to
// intercept failures when testing flavor mismatches.
type testReporter interface {
	Helper()
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

func runTestImplWithReporter(r testReporter, db database.Database, test TestCase, mode schema.GeneratorMode, sqlParser database.Parser, config database.GeneratorConfig) {
	_ = "STUB: not implemented"
	return
}

// Test idempotency of current schema

// Main test

// Bidirectional migration test
// PHASE 1: Test forward migration (current → desired) should produce Up

// Handle expected errors

// PHASE 2: Test idempotency of desired schema

// Filter out skipped DDLs for idempotency check because they weren't applied

// PHASE 3: Test reverse migration (desired → current) should produce Down

// PHASE 4: Test idempotency of current schema after reverse migration

// Filter out skipped DDLs for idempotency check because they weren't applied

// Idempotency-only test (neither up nor down specified)

// Handle expected errors

// For idempotency-only tests, we expect no DDLs (or just apply and test idempotency)

// Test idempotency of desired schema

func runOfflineTestWithReporter(r testReporter, test TestCase, mode schema.GeneratorMode, sqlParser database.Parser, config database.GeneratorConfig, defaultSchema string) {
	_ = "STUB: not implemented"
	return
}

// Bidirectional migration test (offline mode)
// PHASE 1: Test forward migration (current → desired) should produce Up

// PHASE 2: Test idempotency of desired schema

// PHASE 3: Test reverse migration (desired → current) should produce Down

// PHASE 4: Test idempotency of current schema

// Idempotency-only test (neither up nor down specified)

// Test idempotency of desired schema
