package logs

import (
   "fmt"
   "time"
   "os"
   "strings"
   "github.com/sirupsen/logrus"
   //"runtime"
   // "path"

)

//Custom log format definition
type MyFormatter struct{}

func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}


func initLog() {

        logrus.SetReportCaller(true)
        logrus.SetFormatter(new(MyFormatter))
        logrus.SetOutput(os.Stdout)

}

func Info( msg string) {

      initLog()
      logrus.Info(msg)

}

func Debug( msg string) {
     initLog()
     logrus.Debug(msg)
}

func Error( msg string) {
     initLog()
     logrus.Error(msg)
}

func Fatal( msg string) {
     initLog()
     logrus.Fatal(msg) //log之后会调用os.Exit(1)
}

func Panic( msg string) {
     initLog()
     logrus.Panic(msg) //log之后会panic()
}