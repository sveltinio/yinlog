package yinlog

// Printer is the interface defining the methods to be implemented by a printer.
type Printer interface {
	Print(string)
	SetPrinterOptions(options *PrinterOptions)
	Formatter
}

// PrinterOptions sets stdout as printer.
type PrinterOptions struct {
	Timestamp       bool
	TimestampFormat string
	Colors          bool
	Labels          bool
	Icons           bool
}
