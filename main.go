package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

func init(){
	lumberjackLogger:= &lumberjack.Logger{
		Filename:   "temp/logfiles.log",
		MaxSize:    5, //megabytes
		MaxAge:     30, //days
		MaxBackups: 3, //max file backups based on MaxSize
		Compress:   true,
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(io.MultiWriter(os.Stdout, lumberjackLogger))
}

func main() {

}
