package database

type Logger interface {
	Print(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
}

type StdoutLogger struct{}

func (s StdoutLogger) Print(v ...any) { _ = "STUB: not implemented"; return }

func (s StdoutLogger) Printf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (s StdoutLogger) Println(v ...any) { _ = "STUB: not implemented"; return }

type NullLogger struct{}

func (n NullLogger) Print(v ...any)                 { _ = "STUB: not implemented"; return }
func (n NullLogger) Printf(format string, v ...any) { _ = "STUB: not implemented"; return }
func (n NullLogger) Println(v ...any)               { _ = "STUB: not implemented"; return }
