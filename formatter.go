package yinlog

// Formatter is the interface defining the methods to be implemented by a printer.
type Formatter interface {
	Format(*LogEntry) string
}
