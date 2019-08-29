package syslogLDS

import (
	"fmt"
	"log/syslog"
)

var instance *syslog.Writer

func getBitMask() syslog.Priority {
	var bitmask syslog.Priority
	bitmask = syslog.LOG_LOCAL7 | syslog.LOG_NOTICE
	return bitmask
}

func GetInstanceLog() *syslog.Writer {

	if instance == nil {
		var bitmask = getBitMask()
		var err error
		instance, err = syslog.New(bitmask, "LDS")
		if err != nil {
			fmt.Println("Erro to instance syslog Class")
			return nil //TODO: Inserir tratamento de exceção
		}
	}
	return instance
}
