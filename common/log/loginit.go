package log

import (
	log4go "github.com/alecthomas/log4go"
)

func init() {
	log4go.LoadConfiguration("log_conf.xml")
}
