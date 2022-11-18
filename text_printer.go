package yinlog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// TextPrinter sets stdout as printer.
type TextPrinter struct {
	Writer  io.Writer
	Options *PrinterOptions
}

// SetPrinterOptions sets the printer options for TextPrinter.
func (tp *TextPrinter) SetPrinterOptions(options *PrinterOptions) {
	tp.Options = options
}

func fullTextPrinter() *TextPrinter {
	return &TextPrinter{
		Writer: os.Stdout,
		Options: &PrinterOptions{
			Timestamp:       true,
			TimestampFormat: time.RFC3339,
			Colors:          true,
			Icons:           true,
			Labels:          true,
		},
	}
}

func iconAndColorOnlyTextPrinter() *TextPrinter {
	return &TextPrinter{
		Writer: os.Stdout,
		Options: &PrinterOptions{
			Timestamp: false,
			Colors:    true,
			Icons:     true,
			Labels:    false,
		},
	}
}

// Print send the message string to the stdout.
func (tp *TextPrinter) Print(msg string) {
	stdOut := bufio.NewWriter(tp.Writer)
	stdOut.WriteString(msg)
	stdOut.WriteString("\n")
	stdOut.Flush()
}

// Format defines how the TextPrinter will format a log entry.
func (tp *TextPrinter) Format(entry *LogEntry) string {
	if entry.IsListLogger() {
		return formatList(entry, tp.Options)
	}
	return formatSingle(entry, tp.Options)
}

//=============================================================================

// Format defines how a single log entry will be formatted by the TextLogger printer.
func formatSingle(entry *LogEntry, options *PrinterOptions) string {
	return fmt.Sprintf("%s%s %s", entry.prefix, formatMarkup(entry.Level, options), formatText(entry.Message))
}

// FormatList defines how a log entry with List will be formatted by the TextLogger printer.
func formatList(entry *LogEntry, options *PrinterOptions) string {
	var buffer bytes.Buffer
	if entry.Message != "" {
		fmt.Fprintf(&buffer, "%s\n", formatListTitle(entry.Message))
	}
	for _, line := range entry.Entries {
		fmt.Fprintf(&buffer, "%s\n", formatListItem(&line, options))
	}
	return buffer.String()
}

func formatListTitle(msg string) string {
	return listTitleStyle(msg)
}

func formatListItem(entry *LogEntry, options *PrinterOptions) string {
	text := fmt.Sprintf("%s %s %s", entry.prefix, formatMarkup(entry.Level, options), formatText(entry.Message))
	logRow := lipgloss.JoinHorizontal(lipgloss.Center, listItemStyle(text, entry.indentSize))
	return fmt.Sprint(logRow)
}

func formatText(msg string) string {
	textColor := getColor(DefaultLevel)
	return fmt.Sprint(textStyle(textColor, msg))
}

func formatMarkup(level Level, options *PrinterOptions) string {
	timestamp := ""
	if options.Timestamp {
		timestamp = fmt.Sprintf("%s ", time.Now().Format(options.TimestampFormat))
	}
	color := lipgloss.AdaptiveColor{}
	if options.Colors {
		color = getColor(level)
	}
	icon := ""
	if options.Icons {
		icon = getIcon(level)
	}
	label := ""
	if options.Labels {
		label = fmt.Sprintf(" %s", getLevelLabel(level))
	}

	logRow := lipgloss.JoinHorizontal(lipgloss.Center, timestampStyle(timestamp), iconStyle(color, icon), levelStyle(color, label))
	return fmt.Sprint(logRow)
}
