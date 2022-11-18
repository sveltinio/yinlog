package yinlog

import (
	"fmt"
	"os"
	"strings"
)

// LogEntry represents a single log entry.
type LogEntry struct {
	Logger        *Logger
	Level         Level
	Message       string
	Entries       []LogEntry
	prefix        string
	indentChar    string
	indentSize    int
	initialPrefix string
}

// NewLogEntry returns a new entry.
func NewLogEntry(logger *Logger) *LogEntry {
	return &LogEntry{
		Logger: logger,
	}
}

// NewListLogger returns a new entry with `entries` set.
func NewListLogger() *LogEntry {
	l := &Logger{
		Printer: iconAndColorOnlyTextPrinter(),
		Level:   DebugLevel,
	}
	return NewLogEntry(l).withList()
}

// Title set the test to be used as title for ListLogger.
func (entry *LogEntry) Title(title string) {
	entry.Message = title
}

// WithList returns a new entry with `entries` set.
func (entry *LogEntry) withList() *LogEntry {
	return &LogEntry{
		Logger:        entry.Logger,
		Entries:       []LogEntry{},
		prefix:        "",
		initialPrefix: "",
		indentChar:    "\u0020", //space
		indentSize:    2,
	}
}

// IsListLogger returns true when ListLogger.
func (entry *LogEntry) IsListLogger() bool {
	return len(entry.Entries) > 0
}

// String returns the entry as string.
func (entry *LogEntry) String() string {
	return entry.Logger.Printer.Format(entry)
}

// SetPrefix sets the char used as prefix char for entries. Whitespace as default.
func (entry *LogEntry) SetPrefix(c string) {
	entry.prefix = c
	entry.initialPrefix = c
}

func (entry *LogEntry) resetPrefix() string {
	return entry.initialPrefix
}

// SetIndentChar sets the char used as prefix when indent list entries. Whitespace as default.
func (entry *LogEntry) SetIndentChar(c string) {
	entry.indentChar = c
}

// SetIndentSize sets how many time the indent char has to be repeated when indent list entries.
func (entry *LogEntry) SetIndentSize(size int) {
	entry.indentSize = size
}

// Append add an entry to the `entries`.
func (entry *LogEntry) Append(level Level, msg string) {
	elem := LogEntry{
		Logger:  entry.Logger,
		Level:   level,
		Message: msg,
		prefix:  entry.prefix,
	}
	entry.Entries = append(entry.Entries, elem)
}

// Render prints log entry when ListLogger.
func (entry *LogEntry) Render() {
	if entry.IsListLogger() {
		entry.Logger.log(entry)
	}
}

// Indent add a string as prefix when indent list entries.
func (entry *LogEntry) Indent() {
	entry.prefix = fmt.Sprintf("%s%s%s", entry.prefix, strings.Repeat("\u0020", entry.indentSize), entry.indentChar)
}

// Unindent resets the indent size.
func (entry *LogEntry) Unindent() {
	entry.indentSize = 0
	entry.Indent()
	entry.prefix = entry.resetPrefix()
}

// Plain level message.
func (entry *LogEntry) Plain(msg string) {
	entry.Level = DefaultLevel
	entry.Message = msg
	entry.Logger.log(entry)
}

// Plainln prints a new empty line.
func (entry *LogEntry) Plainln() {
	entry.Level = DefaultLevel
	entry.Message = "\n"
	entry.Logger.log(entry)
}

// Debug level message.
func (entry *LogEntry) Debug(msg string) {
	entry.Level = DebugLevel
	entry.Message = msg
	entry.Logger.log(entry)
}

// Info level message.
func (entry *LogEntry) Info(msg string) {
	entry.Level = InfoLevel
	entry.Message = msg
	entry.Logger.log(entry)
}

// Error level message.
func (entry *LogEntry) Error(msg string) {
	entry.Level = ErrorLevel
	entry.Message = msg
	entry.Logger.log(entry)
}

// Warning level message.
func (entry *LogEntry) Warning(msg string) {
	entry.Level = WarningLevel
	entry.Message = msg
	entry.Logger.log(entry)
}

// Success level message.
func (entry *LogEntry) Success(msg string) {
	entry.Level = SuccessLevel
	entry.Message = msg
	entry.Logger.log(entry)
}

// Important level message.
func (entry *LogEntry) Important(msg string) {
	entry.Level = ImportantLevel
	entry.Message = msg
	entry.Logger.log(entry)
}

// Fatal level message.
func (entry *LogEntry) Fatal(msg string) {
	entry.Level = FatalLevel
	entry.Message = msg
	entry.Logger.log(entry)
	os.Exit(1)
}

// Plainf level formatted message.
func (entry *LogEntry) Plainf(msg string, v ...interface{}) {
	entry.Plain(fmt.Sprintf(msg, v...))
}

// Debugf level formatted message.
func (entry *LogEntry) Debugf(msg string, v ...interface{}) {
	entry.Debug(fmt.Sprintf(msg, v...))
}

// Infof level formatted message.
func (entry *LogEntry) Infof(msg string, v ...interface{}) {

	entry.Info(fmt.Sprintf(msg, v...))
}

// Errorf level formatted message.
func (entry *LogEntry) Errorf(msg string, v ...interface{}) {
	entry.Error(fmt.Sprintf(msg, v...))
}

// Successf level formatted message.
func (entry *LogEntry) Successf(msg string, v ...interface{}) {
	entry.Success(fmt.Sprintf(msg, v...))
}

// Warningf level formatted message.
func (entry *LogEntry) Warningf(msg string, v ...interface{}) {
	entry.Warning(fmt.Sprintf(msg, v...))
}

// Importantf level formatted message.
func (entry *LogEntry) Importantf(msg string, v ...interface{}) {
	entry.Important(fmt.Sprintf(msg, v...))
}

// Fatalf level formatted message.
func (entry *LogEntry) Fatalf(msg string, v ...interface{}) {
	entry.Fatal(fmt.Sprintf(msg, v...))
}
