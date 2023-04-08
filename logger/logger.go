package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Loger logrus.Logger

func init() {
	Loger.SetLevel(logrus.DebugLevel)
	writer1 := io.Writer(os.Stdout)
	writer2, _ := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	Loger.SetFormatter(&logrus.JSONFormatter{})
	Loger.SetOutput(io.MultiWriter(writer1, writer2))
}
