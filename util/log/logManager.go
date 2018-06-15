package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kataras/golog"
)

var fileLog os.File

var logMode string

var MyLogger *golog.Logger

// NewLogFile ...
func NewLogFile(mode string) {
	logMode = mode
	go func() {
		var hour int
		for {
			nowHour := time.Now().Hour()
			if nowHour != hour {
				t := time.Now()
				hour = t.Hour()
				fileLog.Close()
				createFile(t)
			}
			time.Sleep(time.Millisecond * 1)
		}
	}()
}

func createFile(t time.Time) {
	sYear := strconv.Itoa(t.Year())
	sMonth := strconv.Itoa(int(t.Month()))
	sDay := strconv.Itoa(t.Day())
	sHour := strconv.Itoa(t.Hour())

	if len(sMonth) == 1 {
		sMonth = fmt.Sprintf("0%s", sMonth)
	}

	if len(sDay) == 1 {
		sDay = fmt.Sprintf("0%s", sDay)
	}

	if len(sHour) == 1 {
		sHour = fmt.Sprintf("0%s", sHour)
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := fmt.Sprintf("%s/logs/%s%s/%s/", pwd, sYear, sMonth, sDay)
	err = os.MkdirAll(filePath, 0777)
	if err != nil {
		panic(err)
	}

	fileName := fmt.Sprintf("%s.log", t.Format("2006010215"))
	fileLog, err := os.OpenFile(filePath+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(io.MultiWriter(fileLog, os.Stdout))

	MyLogger = golog.New()
	MyLogger.SetOutput(io.MultiWriter(fileLog, os.Stdout))
	MyLogger.SetLevel(logMode)

	MyLogger.Println("MyLogger ~~~~")

}
