package logger

import (
	"log"

	"github.com/err-him/gozap"
)

//Default Intialize
func init() {

	config := gozap.Configuration{
		EnableConsole:     false,            //Whether to print the outpul on the console, Good for debugging purpose in local
		ConsoleLevel:      gozap.Debug,      //Debug level log
		ConsoleJSONFormat: false,            //Console log in JSON format, false will print in raw format on console
		EnableFile:        true,             // Logging in File
		FileLevel:         gozap.Info,       // File log leve\
		FileJSONFormat:    false,            // File JSON Format, False will print in file in raw Format
		FileLocation:      "./logs/app.log", //File location where log needs to be appended
	}

	err := gozap.NewLogger(config)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}
}

func Debugf(format string, args ...interface{}) {
	gozap.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	gozap.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	gozap.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	gozap.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	gozap.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	gozap.Panicf(format, args...)
}
