package main

import (
	logger "github.com/sveltinio/yinlog"
)

func main() {
	log := logger.New()
	log.Printer.SetPrinterOptions(&logger.PrinterOptions{
		Timestamp: false,
		Colors:    true,
		Labels:    false,
		Icons:     true,
	})

	log.Plain("plain message")
	log.Debug("debug message")
	log.Error("error message")
	log.Important("important message")
	log.Info("info message")
	log.Success("success message")
	log.Warning("warning message")
	//log.Fatal("fatal message")

	log.Plainln()
	value := "yinlog"
	log.Plainf("plain message with %s", value)
	log.Debugf("debug message with %s", value)
	log.Errorf("error message with %s", value)
	log.Importantf("important message with %s", value)
	log.Infof("info message with %s", value)
	log.Successf("success message with %s", value)
	log.Warningf("warning message with %s", value)
	// log.Fatalf("fatal message with value %s", value)

}
