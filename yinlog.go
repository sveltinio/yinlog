package yinlog

// Logger represents a logger with configurable Printer and Level.
type Logger struct {
	Printer Printer
	Level   Level
}

// New returns a new logger.
func New() *Logger {
	return &Logger{
		Printer: fullTextPrinter(),
		Level:   DebugLevel,
	}
}

// SetPrinter set the logger Printer.
func (l *Logger) SetPrinter(printer Printer) {
	l.Printer = printer
}

// SetLevel sets the logger Level.
func (l *Logger) SetLevel(level Level) {
	l.Level = level
}

// Plain level message.
func (l *Logger) Plain(msg string) {
	NewLogEntry(l).Plain(msg)
}

// Debug level message.
func (l *Logger) Debug(msg string) {
	NewLogEntry(l).Debug(msg)
}

// Info level message.
func (l *Logger) Info(msg string) {
	NewLogEntry(l).Info(msg)
}

// Error level message.
func (l *Logger) Error(msg string) {
	NewLogEntry(l).Error(msg)
}

// Success level message.
func (l *Logger) Success(msg string) {
	NewLogEntry(l).Success(msg)
}

// Warning level message.
func (l *Logger) Warning(msg string) {
	NewLogEntry(l).Warning(msg)
}

// Important level message.
func (l *Logger) Important(msg string) {
	NewLogEntry(l).Important(msg)
}

// Fatal level message.
func (l *Logger) Fatal(msg string) {
	NewLogEntry(l).Fatal(msg)
}

// Debugf level formatted message.
func (l *Logger) Debugf(msg string, v ...interface{}) {
	NewLogEntry(l).Debugf(msg, v...)
}

// Infof level formatted message.
func (l *Logger) Infof(msg string, v ...interface{}) {
	NewLogEntry(l).Infof(msg, v...)
}

// Errorf level formatted message.
func (l *Logger) Errorf(msg string, v ...interface{}) {
	NewLogEntry(l).Errorf(msg, v...)
}

// Successf level formatted message.
func (l *Logger) Successf(msg string, v ...interface{}) {
	NewLogEntry(l).Successf(msg, v...)
}

// Warningf level formatted message.
func (l *Logger) Warningf(msg string, v ...interface{}) {
	NewLogEntry(l).Warningf(msg, v...)
}

// Importantf level formatted message.
func (l *Logger) Importantf(msg string, v ...interface{}) {
	NewLogEntry(l).Importantf(msg, v...)
}

// Fatalf level formatted message.
func (l *Logger) Fatalf(msg string, v ...interface{}) {
	NewLogEntry(l).Fatalf(msg, v...)
}

//=============================================================================

// log will send a message at the level given to the Printer.
func (l *Logger) log(item *LogEntry) {
	l.Printer.Print(item.String())
}
