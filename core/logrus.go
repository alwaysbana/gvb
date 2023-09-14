package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb/global"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestmap := entry.Time.Format("2006-01-02 15:04:05")

	log := global.Config.Logger
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestmap, levelColor, entry.Level, fileVal, funcVal, entry.Message)

	} else {
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestmap, levelColor, entry.Level, entry.Message)

	}
	return b.Bytes(), nil
}

//对于
func InitLogLevel() logrus.Level {
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	return level
}
func InitLogger() *logrus.Logger {
	mlog := logrus.New() //新建一个实例
	mlog.SetOutput(os.Stdout)
	mlog.SetReportCaller(global.Config.Logger.ShowLine) //开启返回行号和函数名
	mlog.SetFormatter(&LogFormatter{})                  //设置自己定义的format
	mlog.SetLevel(InitLogLevel())                       //设置最低的level//设置输出类型
	//fmt.Println("橙子")
	//mlog.SetReportCaller(true) //开启返回函数名和行号
	return mlog
}

func InitDefaultLogger() {
	//全局log
	logrus.SetOutput(os.Stdout)
	//fmt.Println("橙子")
	//设置输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) //开启返回行号和函数名
	logrus.SetFormatter(&LogFormatter{})                  //设置自己定义的format
	logrus.SetLevel(InitLogLevel())                       //设置最低的level
}
