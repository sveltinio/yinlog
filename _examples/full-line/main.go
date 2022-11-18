package main

import (
	"fmt"

	logger "github.com/sveltinio/yinlog"
)

func main() {
	log := logger.New()
	log.Plain("plain message")
	log.Debug("debug message")
	log.Error("error message")
	log.Important("important message")
	log.Info("info message")
	log.Success("success message")
	log.Warning("warning message")
	//log.Fatal("fatal message")

	yinlogStr := "yinlog"
	charm := "lipgloss"
	fmt.Println()
	log.Plainf("plain message with %s and %s", yinlogStr, charm)
	log.Debugf("debug message with %s and %s", yinlogStr, charm)
	log.Errorf("error message with %s and %s", yinlogStr, charm)
	log.Importantf("important message with %s and %s", yinlogStr, charm)
	log.Infof("info message with %s and %s", yinlogStr, charm)
	log.Successf("success message with %s and %s", yinlogStr, charm)
	log.Warningf("warning message with %s and %s", yinlogStr, charm)
	// log.Fatalf("fatal message with value %s", value)
}
