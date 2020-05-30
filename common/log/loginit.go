package log

import (
	log4go "github.com/alecthomas/log4go"
)

func Init() {
	log4go.LoadConfiguration("log_conf.xml")
}
