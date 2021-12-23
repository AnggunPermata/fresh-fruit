package main

import (
	"fmt"
	"github.com/anggunpermata/fresh-fruit/config"
	"github.com/anggunpermata/fresh-fruit/helper/logger"
	"github.com/anggunpermata/fresh-fruit/server/routes"
	"github.com/labstack/echo"
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
	e := echo.New()

	routes.New(e)
	config.InitPort()
	config.CORSWithConfig(e)

	e.Use(logger.MiddlewareLogging)
	e.HTTPErrorHandler = logger.ErrorHandler

	port := fmt.Sprintf(":%d", config.PORT)
	logger.MakeLogEntry(nil).Infof("starting fresh-fruit version: %s", config.LoadEnv("VERSION"))
	if err := e.Start(port); err != nil{
		e.Logger.Fatal(err)
	}

}
