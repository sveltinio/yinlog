package main

import (
	logger "github.com/sveltinio/yinlog"
)

func main() {
	listLog := logger.NewListLogger()
	listLog.Title("\nLogging in a list")
	listLog.Append(logger.DebugLevel, "debug message")
	listLog.Append(logger.ErrorLevel, "error message")
	listLog.Append(logger.ImportantLevel, "important message")
	listLog.Append(logger.InfoLevel, "info message")
	listLog.Append(logger.SuccessLevel, "success message")
	listLog.Append(logger.WarningLevel, "warning message")

	//log.Fatal("fatal message")
	listLog.Render()

	listLog2 := logger.NewListLogger()
	listLog2.SetPrefix("->")
	listLog2.SetIndentChar("*")
	listLog2.SetIndentSize(2)
	listLog2.Title("\nLogging in a list with prefix, indent char, indent size and unindent")
	listLog2.Append(logger.DebugLevel, "debug message")
	listLog2.Append(logger.ErrorLevel, "error message")
	listLog2.Indent()
	listLog2.Append(logger.ImportantLevel, "important message")
	listLog2.Append(logger.InfoLevel, "info message")
	listLog2.Unindent()
	listLog2.Append(logger.SuccessLevel, "success message")
	listLog2.Append(logger.WarningLevel, "warning message")

	//log.Fatal("fatal message")
	listLog2.Render()

}
